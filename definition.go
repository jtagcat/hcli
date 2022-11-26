package harg

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	internal "github.com/jtagcat/harg/internal"
)

type (
	// must not start with a decimal digit (0,1,2,3,4,5,6,7,8,9) (for ergonomic negative values)
	Definitions map[string]*Definition // map[slug]; 1-character: short option, >1: long option
	Definition  struct {
		Type Type

		// For short options (1-char length), this is ignored.
		// For long options:
		//   false: Allows spaces (`--slug value`), in addition to `=` (`--slug=value`).
		//   true: When "=" is not used, Type is Bool. Values (in `--slug=value`) are treated as bools, if strconv.ParseBool says so.
		// Bools before a parsed Type are ignored. Any bools after Type are parsed as Type, and may result in ErrIncompatibleValue.
		AlsoBool bool

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

func (defs Definitions) normalize() error {
	for name, def := range defs {
		if def == nil || name == "" {
			delete(defs, name)
			continue
		}

		if def.Type > typeMax {
			return fmt.Errorf("%s: %w", internal.KeyErrorName(name), internal.GenericErr{
				Err: ErrInvalidDefinition, Wrapped: errors.New("Type does not exist"),
			})
		}

		if unicode.IsDigit(rune(name[0])) {
			return fmt.Errorf("%s: %w", internal.KeyErrorName(name), internal.GenericErr{
				Err: ErrInvalidDefinition, Wrapped: errors.New("Definition name can't start with a digit"),
			})
		}

		if def.Type == Bool && def.AlsoBool {
			def.AlsoBool = false // for parseOptionContent()
		}

		// short args are case sensitive, skip
		if utf8.RuneCountInString(name) == 1 {
			def.AlsoBool = false
			continue
		}

		// case insensitivize long args
		lower := strings.ToLower(name)
		if name != lower {
			defs[lower] = def
			delete(defs, name)
		}
	}
	return nil
}
