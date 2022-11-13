package harg

import (
	"errors"
	"strings"
	"unicode/utf8"

	internal "github.com/jtagcat/harg/internal"
)

var (
	// end user (runtime) error
	ErrOptionHasNoDefinition = errors.New("option has no definition")
	ErrIncompatibleValue     = errors.New("incompatible value") // eg strconv.Atoi("this is not a number")

	// library user error (always returned on Parse())
	ErrInvalidDefinition = errors.New("invalid definition")

	// runtime error //TODO:
	ErrInternalBug = errors.New("internal bug in harg or undefined enum") // anti-panic safetynet
)

// Parse Definitions. See FORMAT.md for the spec. See parse_test.go for examples.
func (defs *Definitions) Parse(
	args []string, // usually os.Args[1:]
	// NB: Parse() does not remove program name (os.Args[0])

	chokes []string, // [case insensitive]
	// Chokes allow for global-local-whatever argument definitions by using Parse() multiple times:
	//
	// Parse() parses until first choke:
	// args: "--foo", "bar", "chokename", "--foo", "differentDef"
	//         ^ parsed ^     ^ choke ^
	//                        chokeReturn: "chokename", "--foo", "differentDef"
	//
	// Chokes are not seen after "--", or in argument values ("--foo choke", "-f choke")
) (
	// parsed options get added to defs (method parent)
	parsed []string, // non-options, arguments
	chokeReturn []string, //[^chokes]//  args[chokePos:], [0] is the found choke, [1:] are remaining unparsed args
	err error, // see above var(); errContext not provided: use fmt.Errorf("parsing arguments: %w", err)
) {
	if err := defs.normalize(); err != nil {
		return nil, nil, err
	}
	chokeM := internal.SliceLowercaseIndex(chokes)

	var skipNext bool
	for i, a := range args {

		if skipNext {
			// (current) i is "next", signal to skip
			// as i-1 already parsed i as it's value
			skipNext = false
			continue
		}

		switch argumentKind(&a) {
		case argument:
			if _, isChoke := chokeM[strings.ToLower(a)]; isChoke {
				return parsed, args[i:], nil
			}
			parsed = append(parsed, a)

		case argumentDivider:
			// append remaining args
			if len(args)-1 != i {
				parsed = append(parsed, args[i+1:]...)
			}

			return parsed, nil, nil

		case shortOption:
			skipNext, err = defs.parseShortOption(i, args)
			if err != nil {
				return nil, nil, err
			}

		case longOption:
			skipNext, err = defs.parseLongOption(i, args) // len(a) >= 3
			if err != nil {
				return nil, nil, err
			}
		}

	}

	return parsed, nil, nil
}

type argumentKindT uint32 // enum:
const (
	argument        argumentKindT = iota
	argumentDivider               // "--"
	shortOption                   // "-something"
	longOption                    // "--something", len() >= 3
)

func argumentKind(arg *string) argumentKindT {
	if len(*arg) < 2 || !strings.HasPrefix(*arg, "-") {
		return argument // including "", "-"
	}

	// "-x"
	if !strings.HasPrefix((*arg)[1:], "-") {
		return shortOption
	}

	// begins with "--"
	switch utf8.RuneCountInString(*arg) {
	case 2: // "--"
		return argumentDivider
	case 3: // "--x", single negative short
		return shortOption
	default: // > 3
		return longOption
	}
}
