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
	chokes []string, //[^chokes]// [case insensitive] parse arguments until first choke
	// Chokes are not seen after "--", or in places of argument values ("--foo choke", "-f choke")
) (
	_ OptionsMap, // parsed options
	parsed []string, // non-options, arguments
	chokeReturn []string, //[^chokes]//  args[chokePos:], [0] is the found choke, [1:] are remaining unparsed args
	err error, // see above var(); errContext not provided: use fmt.Errorf("parsing arguments: %w", err)
) {
	if err := defs.checkDefs(); err != nil {
		return nil, nil, nil, err
	}
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
			if len(args)-1 != i { // there are more args
				parsed = append(parsed, args[i+1:]...)
			}
			return optM, parsed, nil, nil

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

	// "-x"
	if strings.HasPrefix((*arg)[1:], "-") {
		return e_shortOption
	}

	// begins with "--"
	switch utf8.RuneCountInString(*arg) {
	case 2: // "--"
		return e_argumentDivider
	case 3: // "--x", single negative short
		return e_shortOption
	default: // > 3
		return e_longOption
	}
}
