package harg

import (
	"fmt"
	"strings"

	"github.com/jtagcat/harg/internal"
)

// long option (--foo) (--foo=value) ?(--foo value)
// (--foo=ignored --foo=value) (--count --count) (--foo=elem1 --foo=elem2)
//
// caller should ensure len(args[i]) > 3; and defs.checkDefs()
func (defs *Definitions) parseLongOption(first int, args []string) (consumedNext bool, _ error) {
	argName := args[first][2:] // [2:]: remove suffix "--"
	if argName == "" {
		panic(fmt.Sprintf("parseLongOption caller did not ensure len(args[i]) > 2 for %d in %q", first, args))
	}

	key, value, valueFound := strings.Cut(argName, "=")

	def, err := defs.get(key)
	if err != nil {
		return false, err
	}

	// bool has no lookahead, default = true
	if value == "" && (def.Type == Bool || def.AlsoBool) {
		valueFound, value = true, "true"
	}

	// try to lookahead value (args: "--key", "value")
	if !valueFound && len(args)-1 > first {
		lookArg := args[first+1]

		consumedNext = argumentKind(&lookArg) == argument
		if consumedNext {
			value = lookArg
		}
	}

	return consumedNext, def.parseOptionContent(key, value)
}

// short option(s) (-f) (-fff) (-fb) (-fbvalue) (-fb value) (--n) (-y-ny)
//
// caller should ensure len(args[i]) >= 2; and defs.checkDefs()
func (defs *Definitions) parseShortOption(first int, args []string) (nextWasConsumed bool, _ error) {
	argRune := []rune(args[first][1:]) // [1:]: skip 0th "-"
	if len(argRune) == 0 {
		panic(fmt.Sprintf("parseShort caller did not ensure len(args[i]) > 1 for %d in %q", first, args))
	}

	var negateNext bool
	for optI, opt := range argRune {

		value := ""
		key := string(opt)

		if key == "-" {
			// new with harg: short option prefix "-" negates bools
			negateNext = true
			continue
		}

		def, err := defs.get(key)
		if err != nil {
			return false, err
		}

		if def.Type == Bool || def.AlsoBool {
			if negateNext {
				value = "false"
				negateNext = false
			} else {
				value = "true"
			}

			err = def.parseOptionContent(key, value)
			if err != nil {
				return false, err
			}

			continue
		}

		// valueful opt, break loop

		if len(argRune)-1 == optI {
			// value in same arg
			value = string(argRune[optI+1:])
		} else {
			// no value, space reached, try to lookahead (args: "-o", "value")
			if len(args)-1 > first { // there are more args
				lookArg := args[first+1]

				valueFound := argumentKind(&lookArg) == argument
				if valueFound {
					nextWasConsumed, value = true, lookArg
				}
			}
		}
		return true, def.parseOptionContent(key, value)

	}
	return false, nil
}

func (defs Definitions) get(key string) (*Definition, error) {
	key = strings.ToLower(key)

	def, ok := defs[key]
	if ok {
		return def, nil
	}

	return nil, fmt.Errorf("%s: %w", internal.KeyErrorName(key), ErrOptionHasNoDefinition)
}
