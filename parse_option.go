package harg

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jtagcat/harg/internal"
)

// long option Bool (--foo) (---foo) or (--foo=value) (--foo) (--foo value)
//
// caller should ensure len(args[i]) > 3; and defs.checkDefs()
func (defs *Definitions) parseLongOption(args []string) (consumedNext bool, _ error) {
	argName := args[0][2:] // [2:]: remove prefix "--"
	if argName == "" {
		panic("parseLongOption caller did not ensure len(args[0]) > 2")
	}

	key, value, valueFound := strings.Cut(argName, "=")
	errContext := func() string { return internal.KeyErrorName(key) }

	key, negateBool := trimPrefix(key, "-") // ---foo (three dashes negate)

	def, err := defs.get(key)
	if err != nil {
		return false, err
	}

	if negateBool {
		if !(def.Type == Bool || def.AlsoBool) {
			return false, fmt.Errorf("parsing %s as %s: %w", errContext(), typeMetaM[def.Type].errName, internal.GenericErr{
				Err:     ErrIncompatibleValue,
				Wrapped: errors.New("only Bool option definitions can use negating prefix '---'"),
			})
		}

		if valueFound {
			return false, fmt.Errorf("parsing %s as %s: %w", errContext(), typeMetaM[def.Type].errName, internal.GenericErr{
				Err:     ErrIncompatibleValue,
				Wrapped: errors.New("negating prefix '---' can't have any value (---option=value)"),
			})
		}
	}

	// Bool has no lookahead, default = true
	if value == "" && (def.Type == Bool || def.AlsoBool) {
		return false, def.parseBoolValue(!negateBool, errContext)
	}

	if !valueFound && len(args) > 1 {
		consumedNext, value = lookAheadValue(args[1])
	}

	return consumedNext, def.parseValue(value, errContext)
}

// short option(s) (-f) (-fff) (-fb) (-fbvalue) (-fb value) (--n) (-y-ny)
//
// caller should ensure len(args[i]) >= 2; and defs.checkDefs()
func (defs *Definitions) parseShortOption(args []string) (consumedNext bool, _ error) {
	argRune := []rune(args[0][1:]) // [1:]: remove prefix "-"
	if len(argRune) == 0 {
		panic("parseShortOption caller did not ensure len(args[0]) > 1")
	}

	// loop through clustered (-abc = -a -b -c) options
	var negateNext bool
	for optI, opt := range argRune {

		value := ""
		key := string(opt)
		errContext := func() string { return internal.KeyErrorName(key) }

		if key == "-" {
			// short option prefix "-" negates
			negateNext = true
			continue
		}

		def, err := defs.get(key)
		if err != nil {
			return false, err
		}

		if def.Type == Bool || def.AlsoBool {
			err := def.parseBoolValue(!negateNext, errContext)
			if err != nil {
				return false, err
			}

			negateNext = false
			continue
		}

		if negateNext {
			return false, fmt.Errorf("parsing %s as %s: %w", errContext(), typeMetaM[def.Type].errName, internal.GenericErr{
				Err:     ErrIncompatibleValue,
				Wrapped: errors.New("only Bool option definitions can use negating prefix '-'"),
			})
		}
		// valueful opt, ending clustering loop

		if len(argRune)-1 == optI {
			consumedNext, value = lookAheadValue(args[1])
		} else {
			// value in same arg
			value = string(argRune[optI+1:])
			value = strings.TrimPrefix(value, "=")
		}

		return consumedNext, def.parseValue(value, errContext)
	}

	return false, nil
}

func lookAheadValue(nextArg string) (consumedNext bool, value string) {
	if argumentKind(nextArg) != argument {
		return false, ""
	}

	return true, nextArg
}

func (defs Definitions) get(key string) (*Definition, error) {
	key = strings.ToLower(key)

	def, ok := defs[key]
	if ok {
		return def, nil
	}

	return nil, fmt.Errorf("%s: %w", internal.KeyErrorName(key), ErrOptionHasNoDefinition)
}

// strings.TrimPrefix with ok
func trimPrefix(s, prefix string) (string, bool) {
	if !strings.HasPrefix(s, prefix) {
		return s, false
	}
	return strings.TrimPrefix(s, prefix), true
}
