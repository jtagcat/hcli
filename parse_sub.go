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
func (defs *Definitions) parseLongOption(args []string) (consumedNext bool, _ error) {
	argName := args[0][2:] // [2:]: remove suffix "--"
	if argName == "" {
		panic("parseLongOption caller did not ensure len(args[0]) > 2")
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
	if !valueFound && len(args) > 1 {
		consumedNext, value = lookAheadValue(args[1])
	}

	return consumedNext, def.parseOptionContent(key, value)
}

// short option(s) (-f) (-fff) (-fb) (-fbvalue) (-fb value) (--n) (-y-ny)
//
// caller should ensure len(args[i]) >= 2; and defs.checkDefs()
func (defs *Definitions) parseShortOption(args []string) (consumedNext bool, _ error) {
	argRune := []rune(args[0][1:]) // [1:]: skip 0th "-"
	if len(argRune) == 0 {
		panic("parseShortOption caller did not ensure len(args[0]) > 1")
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
			value = strings.TrimPrefix(value, "=")
		} else if len(args) > 1 { // try to lookahead value (args: "--key", "value")
			consumedNext, value = lookAheadValue(args[1])
		}

		return true, def.parseOptionContent(key, value)
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
