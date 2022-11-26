package harg

import (
	"errors"
	"strings"
	"unicode/utf8"

	internal "github.com/jtagcat/harg/internal"
)

var (
	// end user (runtime) error
	ErrOptionHasNoDefinition = errors.New("option has no definition") // or invalid Alias() target
	ErrIncompatibleValue     = errors.New("incompatible value")       // eg strconv.Atoi("this is not a number")

	// library user error; always returned on Parse()
	ErrInvalidDefinition = errors.New("invalid definition")
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
	// parsed options get added to defs, see option_get.go (def.Touched(), .SlString(), .String(), ...)
	parsed []string, // non-options, arguments
	chokeReturn []string, // See above
	err error, // see above var() for possible errors
) {
	if len(args) == 0 {
		return nil, nil, nil
	}

	if err := defs.normalize(); err != nil {
		return nil, nil, err
	}
	chokeM := internal.SliceLowercaseIndex(chokes)

	for {
		var skipNext bool

		switch argumentKind(args[0]) {
		case argument:
			if _, isChoke := chokeM[strings.ToLower(args[0])]; isChoke {
				return parsed, args, nil
			}

			parsed = append(parsed, args[0])

		case argumentDivider:
			parsed = append(parsed, args[1:]...)

			return parsed, nil, nil

		case shortOption:
			skipNext, err = defs.parseShortOption(args) // len(a) > 1 or parseShortOption panics
			if err != nil {
				return nil, nil, err
			}

		case longOption:
			skipNext, err = defs.parseLongOption(args) // len(a) > 2 or parseLongOption panics
			if err != nil {
				return nil, nil, err
			}
		}

		if len(args) == 1 {
			break
		}

		if skipNext {
			if len(args) == 2 {
				break
			}

			args = args[2:]
			continue
		}

		args = args[1:]
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

func argumentKind(arg string) argumentKindT {
	if len(arg) < 2 || !strings.HasPrefix(arg, "-") {
		return argument // including "", "-"
	}

	// "-x"
	if !strings.HasPrefix(arg[1:], "-") {
		return shortOption // len(a) > 1 or parseShortOption panics
	}

	// begins with "--"
	switch utf8.RuneCountInString(arg) {
	case 2: // "--"
		return argumentDivider
	case 3: // "--x", single negating short
		return shortOption // len(a) > 1 or parseShortOption panics
	default: // >= 3 or parseLongOption panics
		return longOption
	}
}
