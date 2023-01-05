package harg

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type (
	// must not start with a decimal digit (0,1,2,3,4,5,6,7,8,9) (for ergonomic negative values)
	Definitions map[string]*Definition // map[key]; 1-character: short option, >1: long option
	Definition  struct {
		Type Type

		// For short options (1-char key), this is ignored.
		// For long options:
		//   false: Allows spaces (`--key value`), in addition to `=` (`--key=value`).
		//   true: For defining boolean: `--key`, `---key`; for defining value: `--key=value`
		// Bools before a parsed Type are ignored. Any bools after Type are parsed as Type, and may result in ErrIncompatibleValue.
		AlsoBool bool

		// defs.ParseEnv(): If enabled, environment value will be split by commas (to slice).
		EnvCSV bool

		originalType Type // used in parsing AlsoBool
		parsed       option
	}
)

func (defs Definitions) Alias(name string, target string) error {
	defP, ok := defs[target]
	if !ok {
		return fmt.Errorf("definition name %s: %w", target, ErrOptionHasNoDefinition)
	}

	defs[name] = defP
	return nil
}

// TODO: hcli
// Does not overwrite existing (case-sensitive) definition names.
func (defs Definitions) SetUnique(name string, def *Definition) (ok bool) {
	if _, ok := defs[name]; ok {
		return false
	}

	defs[name] = def
	return true
}

func (defs Definitions) get(key string) (*Definition, error) {
	key = strings.ToLower(key)

	def, ok := defs[key]
	if ok {
		return def, nil
	}

	return nil, fmt.Errorf("%s: %w", optErrorName(key), ErrOptionHasNoDefinition)
}

func (defs Definitions) genericNormalize(transform func(key string, def *Definition) (newKey string, _ error)) error {
	for key, def := range defs {
		if def == nil || key == "" {
			delete(defs, key)

			continue
		}

		if def.Type > TypeMax {
			return fmt.Errorf("%s: %w", optErrorName(key), genericErr{
				Err: ErrInvalidDefinition, Wrapped: errors.New("Type does not exist"),
			})
		}

		if unicode.IsDigit(rune(key[0])) {
			return fmt.Errorf("%s: %w", optErrorName(key), genericErr{
				Err: ErrInvalidDefinition, Wrapped: errors.New("Definition key can't start with a digit"),
			})
		}

		if def.Type == Bool && def.AlsoBool {
			def.AlsoBool = false // for parseOptionContent()
		}

		new, err := transform(key, def)
		if err != nil {
			return fmt.Errorf("%s: %w", optErrorName(key), genericErr{
				Err: ErrInvalidDefinition, Wrapped: err,
			})
		}

		// alias, not delete, as opaque normalization might lead to unexpected key change (for retrival after parse)
		if key != new {
			defs[new] = def
		}
	}

	return nil
}

func (defs Definitions) normalizeOpts() error {
	return defs.genericNormalize(func(key string, def *Definition) (string, error) {
		// short args are case sensitive, skip
		if utf8.RuneCountInString(key) == 1 {
			def.AlsoBool = false
			return key, nil
		}

		// case insensitivize long args
		return strings.ToLower(key), nil
	})
}

func (defs Definitions) normalizeEnv() error {
	return defs.genericNormalize(func(key string, def *Definition) (string, error) {
		for _, r := range key {
			if r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r) {
				continue
			}

			return key, fmt.Errorf("must contain only underscores, letters and/or digits")
		}

		// capitalize all keys
		return strings.ToUpper(key), nil
	})
}

func optErrorName(key string) string {
	var keyType string
	if utf8.RuneCountInString(key) > 1 {
		keyType = "long"
	} else {
		keyType = "short"
	}

	return fmt.Sprintf("%s option %q", keyType, key)
}
