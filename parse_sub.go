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
func (defs *Definitions) parseLongOption(optM *OptionsMap, chokeM *map[string]bool, i *int, args *[]string) (nextWasConsumed bool, _ error) {
	argName := (*args)[*i][2:] // [2:]: skip "--"
	if argName == "" {
		return false, fmt.Errorf("parseLongOption caller did not ensure len(args[i]) > 3 for %d in %q: %w", i, args, ErrInternalBug)
	}

	key, value, valueFound := strings.Cut(argName, "=")

	def, err := defs.find(key)
	if err != nil {
		return false, err
	}

	// no-arg boolean
	if !valueFound && (def.Type == e_Boolean || def.AlsoBoolean) {
		valueFound, value = true, "true"
	}

	// if needed, try to lookahead value (args: "--key", "value")
	if !valueFound && len(*args)-1 > *i {
		lookArg := (*args)[*i+1]

		valueFound := def.lookaheadUsable(lookArg)
		if valueFound {
			nextWasConsumed, value = true, lookArg
		}
	}

	return nextWasConsumed, optM.parseOptionContent(&def, &value, &valueFound)
}

// short option(s) (-f) (-fff) (-fb) (-fbvalue) (-fb value)
//
// caller should ensure len(args[i]) >= 2; and defs.checkDefs()
func (defs *Definitions) parseShortOption(optM *OptionsMap, i *int, args *[]string) (nextWasConsumed bool, _ error) {
	argName := (*args)[*i][1:] // [1:]: skip "-"
	if argName == "" {
		return false, fmt.Errorf("parseLongOption caller did not ensure len(args[i]) >= 2 for %d in %q: %w", i, args, ErrInternalBug)
	}

	// 	var value string
	// 	var valueFound, negateNext bool
	//
	// 	for _, runeV := range argName {
	// 		char := string(runeV)
	//
	// 		if char == "_" {
	// 			// new with harg: short option prefix "_" negates it
	// 			negateNext = true
	// 			continue
	// 		}
	//
	// 		def, err := defs.find(char)
	// 		if err != nil {
	// 			return false, err
	// 		}
	//
	// 		// if !valueFound && len(*args)-1 > *i {
	// 		// 	lookArg := (*args)[*i+1]
	//
	// 		// 	valueFound := def.lookaheadUsable(chokeM, lookArg)
	// 		// 	if valueFound {
	// 		// 		nextWasConsumed, value = true, lookArg
	// 		// 	}
	// 		// }
	//
	// 		if def.Type == e_Boolean {
	// 			valueFound = true
	// 			if negateNext {
	// 				value = "false"
	// 			} else {
	// 				value = "true"
	// 			}
	// 		} else {
	// 			//
	// 			//
	// 		}
	//
	// 		if def.Type != e_Boolean && def.AlsoBoolean {
	// 			return false, fmt.Errorf("short definition %s: %w", char, ErrShortOptionNoAlsoBoolean)
	// 		}
	//
	// 		// if not bool
	//
	// 		optM.parseOptionContent(&def, b)
	// 		if negateNext {
	// 			b, negateNext = true, false
	// 		}
	//
	// 		// is it bool or not
	// 	}

	return false, nil // TODO:
	// ?implement negative boolean? _f
}

func (defs *Definitions) find(key string) (Definition, error) {
	var errPrelude string
	key = strings.ToLower(key)

	aliasKey, isAlias := defs.Aliases[key]
	if isAlias {
		key = *aliasKey
		errPrelude += fmt.Sprintf("alias %s: ", *aliasKey)
	}

	def, ok := defs.D[key]
	if ok {
		return def, nil
	}

	if utf8.RuneCountInString(key) > 1 {
		errPrelude += "long "
	} else {
		errPrelude += "short "
	}

	return Definition{}, fmt.Errorf(errPrelude+"option %s: %w", key, ErrOptionHasNoDefinition)
}

func (def *Definition) lookaheadUsable(arg string) bool {
	if def.Type == e_Boolean || def.AlsoBoolean {
		return false
	}

	return argumentKind(&arg) == e_argument
}
