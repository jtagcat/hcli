package harg

import (
	"strings"

	internal "github.com/jtagcat/harg/internal"
)

type (
	Definitions struct {
		// string: identifier slug
		Short map[string]*Kind
		Long  map[string]*Kind
	}
	Kind int // enum
)

const ( // enum
	Boolean Kind = iota
	Count

	// doesn't seem the best way, but let's try
	// Value
	String
	Integer
	// Slice
	SliceString
	SliceInteger
	// TODO: ...?
)

func (defs *Definitions) Parse(
	// [^chokes]: Chokes allow for global-local-whatever argument definitions by using Parse() multiple times.
	// args: "--foo", "bar", "chokename", "--foo", "differentDef"
	//          ^ parsed ^    ^choke, chokeReturn: "chokename", "--foo", "differentDef"

	args []string, // usually os.Args
	chokes []string, //[^chokes]// [case insensitive] parse arguments until first choke (chokes after "--" aren't seen)
) (
	_ OptValues, // parsed options
	parsed []string, // non-options, arguments
	chokeReturn []string, //[^chokes]//  args[chokePos:], [0] is the found choke, [1:] are remaining unparsed args
	err error, // on undefined option (left of first choke)
) {
	chokeM := internal.SliceToLowercaseMap(chokes)
	opts := defs.toEmptyOpts()

	var nextWasConsumed bool
	for i, a := range args {

		if nextWasConsumed {
			// (current) i is "next",
			// skip, as i-1 already parsed i as it's value
			nextWasConsumed = false
			continue
		}

		if _, ok := chokeM[a]; ok {
			return opts, parsed, args[i:], nil
		}

		if a == "-" || !strings.HasPrefix(a, "-") {
			// normal argument
			parsed = append(parsed, a)
			continue
		}

		// !HasPrefix "--"
		if !strings.HasPrefix(a[1:], "-") {
			// short option
			nextWasConsumed, err = opts.parseShortOption(&i, &args)
			if err != nil {
				return OptValues{}, nil, nil, err
			}
			continue
		}

		// a == "--"
		if len(a) == 2 {
			if len(args) != i {
				// there are more args
				parsed = append(parsed, args[i+1:]...)
			}
			return opts, parsed, nil, nil
		}

		// long option
		nextWasConsumed, err = opts.parseLongOption(&i, &args) // len(a) >= 3
		if err != nil {
			return OptValues{}, nil, nil, err
		}
		continue

	}
	return opts, parsed, nil, nil
}

func (defs *Definitions) toEmptyOpts() OptValues {
	// v := make(map)

	// definitionsGroupToOpts(true, defs.Short, )
	// definitionsGroupToOpts(false, defs.Long)

	return OptValues{} // TODO:
}

func definitionsGroupToOpts(isShort bool, group map[string]*Kind) {
}

type OptValues struct {
	// TODO: ???
}

// long option (--foo) (--foo=value) ?(--foo value)
// (--foo=ignored --foo=value) (--count --count) (--foo=elem1 --foo=elem2)
//
// caller must ensure len(args[i]) >= 3
func (opts *OptValues) parseLongOption(i *int, args *[]string) (nextWasConsumed bool, _ error) {
	return false, nil // TODO:
}

// short option(s) (-f) (-fff) (-fb) (-fbvalue) (-fb value)
func (opts *OptValues) parseShortOption(i *int, args *[]string) (nextWasConsumed bool, _ error) {
	return false, nil // TODO:
	// ?implement negative boolean? _f
}

// TODO:
type (
	ArgBoolean struct {
		Defined bool // if found in args
		Value   bool // result
	}
)
