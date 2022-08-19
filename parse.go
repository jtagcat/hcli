package harg

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	internal "github.com/jtagcat/harg/internal"
)

type (
	Definitions struct {
		D       DefinitionMap
		Aliases map[string]*string // map[alias slug]defSlug
	}
	DefinitionMap map[string]Definition // map[slug]; 1-character: short option, >1: long option
	Definition    struct {
		Type      Type
		Countable bool // where seen as boolean, see as countable (-vvv) instead.
		Slice     bool // where seen as value, append to slice instead of ignoring all but last

		// For short options (1-char length), true means it's always boolean
		// For long options:
		//   false: allows spaces (`--slug value` in addition to `--slug=value`)
		//   true: if "=" is not used, Type is changed to boolean (or countable)
		AlsoBoolean bool
	}
	Type int // enum
)

const ( // enum
	e_Boolean Type = iota
	// doesn't seem the best way, but let's try
	e_String
	e_Integer
	// TODO: ...?
)

var (
	// end user (runtime) error
	ErrOptionHasNoDefinition = errors.New("option has no definition")
	ErrLongOptionIsTooShort  = errors.New("long option (--foo) key name must have at least 2 characters")               // restriction for sanity
	ErrMixedSliceAlsoBoolean = errors.New("slice option with AlsoBoolean can't be given both boolean and slice inputs") // --foo=value --foo --foo=value

	// runtime error
	ErrShortOptionNoAlsoBoolean = errors.New("short option (-x) valueful definition mustn't be AlsoBoolean")
	ErrInternalBug              = errors.New("internal bug in harg") // anti-panic safetynet

	// depends on definitions (every exectime):
	ErrSlugConflict = errors.New("conflicting same-named alias")
)

func (defs *Definitions) Parse(
	// [^chokes]: Chokes allow for global-local-whatever argument definitions by using Parse() multiple times.
	// args: "--foo", "bar", "chokename", "--foo", "differentDef"
	//          ^ parsed ^    ^choke, chokeReturn: "chokename", "--foo", "differentDef"

	args []string, // usually os.Args
	chokes []string, //[^chokes]// [case insensitive] parse arguments until first choke (chokes after "--NotAlsoBoolean chokename" are not seen)
) (
	_ OptionsMap, // parsed options
	parsed []string, // non-options, arguments
	chokeReturn []string, //[^chokes]//  args[chokePos:], [0] is the found choke, [1:] are remaining unparsed args
	err error, // see above var(); errContext not provided: use fmt.Errorf("parsing arguments: %w", err)
) {
	chokeM := internal.SliceToLowercaseMap(chokes)
	optM := defs.D.toEmptyOptM()

	var skipNext bool
	for i, a := range args {

		if skipNext {
			// (current) i is "next", signal to skip
			// as i-1 already parsed i as it's value
			skipNext = false
			continue
		}

		switch argumentKind(&a) {
		case e_argument:
			if _, isChoke := chokeM[strings.ToLower(a)]; isChoke {
				return optM, parsed, args[i:], nil
			}
			parsed = append(parsed, a)

		case e_argumentDivider:
			if len(args)-1 == i { // no more args
				return optM, parsed, nil, nil
			}
			remainingArgs := args[i+1:]

			// look for chokes
			for lookI, lookA := range remainingArgs {
				if _, isChoke := chokeM[strings.ToLower(lookA)]; isChoke {

					parsed = append(parsed, remainingArgs[:lookI]...)
					return optM, parsed, remainingArgs[lookI:], nil
				}
			}

			return optM, append(parsed, remainingArgs...), nil, nil

		case e_shortOption:
			skipNext, err = defs.parseShortOption(&optM, &i, &args)
			if err != nil {
				return nil, nil, nil, err
			}

		case e_longOption:
			skipNext, err = defs.parseLongOption(&optM, &chokeM, &i, &args) // len(a) >= 3
			if err != nil {
				return nil, nil, nil, err
			}
		}

	}
	return optM, parsed, nil, nil
}

type argumentKindT int

const ( // enum
	e_argument        argumentKindT = iota
	e_argumentDivider               // "--"
	e_shortOption                   // "-something"
	e_longOption                    // "--something", len() >= 3
)

func argumentKind(arg *string) argumentKindT {
	if len(*arg) < 2 || !strings.HasPrefix(*arg, "-") {
		return e_argument // including "", "-"
	}

	// !HasPrefix "--"; len(2) checked above
	if !strings.HasPrefix((*arg)[1:], "-") {
		return e_shortOption
	}

	if len(*arg) == 2 {
		return e_argumentDivider
	}

	// len() >= 3
	return e_longOption
}

func (defs *Definitions) checkDefs() error {
	defs.D = internal.LowercaseLongMapNames(defs.D)
	defs.Aliases = internal.LowercaseLongMapNames(defs.Aliases)

	for n := range defs.D {
		if _, ok := defs.Aliases[n]; ok {
			return fmt.Errorf("option definition %s: %w", n, ErrSlugConflict)
		}
	}
	return nil
}

func (defM *DefinitionMap) toEmptyOptM() OptionsMap {
	return OptionsMap{} // TODO:
}

type (
	OptionsMap map[string]Option
	Option     struct {
		// TODO: ???
	}
)

// long option (--foo) (--foo=value) ?(--foo value)
// (--foo=ignored --foo=value) (--count --count) (--foo=elem1 --foo=elem2)
//
// caller should ensure len(args[i]) >= 3; and defs.checkDefs()
func (defs *Definitions) parseLongOption(optM *OptionsMap, chokeM *map[string]bool, i *int, args *[]string) (nextWasConsumed bool, _ error) {
	argName := (*args)[*i][2:] // [2:]: skip "--"
	if argName == "" {
		return false, fmt.Errorf("parseLongOption caller did not ensure len(args[i]) >= 3 for %d in %q: %w", i, args, ErrInternalBug)
	}

	key, value, valueFound := strings.Cut(argName, "=")
	if utf8.RuneCountInString(key) < 2 {
		return false, ErrLongOptionIsTooShort
	}

	def, err := defs.find(key)
	if err != nil {
		return false, err
	}

	if !valueFound && (def.Type == e_Boolean || def.AlsoBoolean) {
		valueFound, value = true, "true"
	}

	// if needed, try to lookahead value (args: "--key", "value")
	if !valueFound && len(*args)-1 > *i {
		lookArg := (*args)[*i+1]

		valueFound := def.lookaheadUsable(chokeM, lookArg)
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

func (def *Definition) lookaheadUsable(chokeM *map[string]bool, arg string) bool {
	if def.AlsoBoolean || def.Type == e_Boolean {
		return false
	}

	if _, isChoke := (*chokeM)[strings.ToLower(arg)]; isChoke {
		return false
	}

	return argumentKind(&arg) == e_argument
}

func (optM *OptionsMap) parseOptionContent(def *Definition, value *string, valueFound *bool) error {
	return fmt.Errorf("not implemented") // TODO:
}

// TODO:
// type (
// 	ArgBoolean struct {
// 		Defined bool // if found in args
// 		Value   bool // result
// 	}
// )
// type (
// 	RawArgs map[string]Arg[string] // string: identifier slug

// 	Arg[T any] struct {
// 		Defined bool
// 		HasData bool
// 		Data    []T
// 	}
