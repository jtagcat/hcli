package harg

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// long option (--foo) (--foo=value) ?(--foo value)
// (--foo=ignored --foo=value) (--count --count) (--foo=elem1 --foo=elem2)
//
// caller should ensure len(args[i]) > 3; and defs.checkDefs()
func (defs *Definitions) parseLongOption(i *int, args *[]string) (consumedNext bool, _ error) {
	argName := (*args)[*i][2:] // [2:]: skip "--"
	if argName == "" {
		return false, fmt.Errorf("parseLongOption caller did not ensure len(args[i]) > 3 for %d in %q: %w", i, args, ErrInternalBug)
	}

	key, value, valueFound := strings.Cut(argName, "=")

	efKey, def, err := defs.find(key)
	if err != nil {
		return false, err
	}

	// bool has no lookahead, default = true
	if value == "" && (def.Type == e_bool || def.AlsoBool) {
		valueFound, value = true, "true"
	}

	// try to lookahead value (args: "--key", "value")
	if !valueFound && len(*args)-1 > *i {
		lookArg := (*args)[*i+1]

		consumedNext = argumentKind(&lookArg) == e_argument
		if consumedNext {
			value = lookArg
		}
	}

	return consumedNext, def.parseOptionContent(key, efKey, value)
}

// short option(s) (-f) (-fff) (-fb) (-fbvalue) (-fb value) (--n) (-y-ny)
//
// caller should ensure len(args[i]) >= 2; and defs.checkDefs()
func (defs *Definitions) parseShortOption(argI *int, args *[]string) (nextWasConsumed bool, _ error) {
	argRune := []rune((*args)[*argI][1:]) // [1:]: skip 0th "-"
	if len(argRune) == 0 {
		return false, fmt.Errorf("parseLongOption caller did not ensure len(args[i]) >= 2 for %d in %q: %w", argI, args, ErrInternalBug)
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

		efKey, def, err := defs.find(key)
		if err != nil {
			return false, err
		}

		if def.Type == e_bool || def.AlsoBool {
			if negateNext {
				value = "false"
				negateNext = false
			} else {
				value = "true"
			}

			err = def.parseOptionContent(key, efKey, value)
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
			if len(*args)-1 > *argI { // there are more args
				lookArg := (*args)[*argI+1]

				valueFound := argumentKind(&lookArg) == e_argument
				if valueFound {
					nextWasConsumed, value = true, lookArg
				}
			}
		}
		return true, def.parseOptionContent(key, efKey, value)

	}
	return false, nil
}

func (defs *Definitions) find(key string) (effectiveKey string, _ Definition, _ error) {
	var errPrelude string
	key = strings.ToLower(key)

	aliasKey, isAlias := defs.Aliases[key]
	if isAlias {
		key = *aliasKey
		errPrelude += fmt.Sprintf("alias %s: ", *aliasKey)
	}

	def, ok := defs.D[key]
	if ok {
		return key, def, nil
	}

	if utf8.RuneCountInString(key) > 1 {
		errPrelude += "long "
	} else {
		errPrelude += "short "
	}

	return "", Definition{}, fmt.Errorf(errPrelude+"option %s: %w", key, ErrOptionHasNoDefinition)
}
