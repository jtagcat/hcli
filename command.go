package hcli

import (
	"fmt"
	"io"
	"strings"
	"unicode/utf8"

	"github.com/jtagcat/hcli/harg"
)

// TODO: keep track of aliases for help text
// categories
type (
	Command struct {
		// Name is moved outside to map, #TODO: unsure about root

		// // --help
		// // Full name of command for help, defaults to Name
		// HelpName string
		// // Description of the program.
		// Usage string
		// // Text to override the USAGE section of help
		// UsageText string
		// // Description of the program argument format.
		// ArgsUsage string
		// // Version of the program // use build args
		// Version string
		// // Description of the program
		// Description string
		// // Boolean to hide built-in help command and help flag
		// HideHelp bool
		// // Boolean to hide built-in help command but keep help flag.
		// // Ignored if HideHelp is true.
		// HideHelpCommand bool
		// // Boolean to hide built-in version flag and the VERSION section of help
		// HideVersion bool
		// // Execute this function if the proper command cannot be found
		// CommandNotFound CommandNotFoundFunc
		// // Execute this function if a usage error occurs
		// OnUsageError OnUsageErrorFunc
		// // List of all authors who contributed
		// Authors []*Author
		// // Copyright of the binary if any
		// Copyright string
		// // Boolean to hide this command from help or completion
		// Hidden bool
		// 		// CustomHelpTemplate the text template for the command help topic.
		// // cli.go uses text/template to render templates. You can
		// // render custom help text by setting this variable.
		// CustomHelpTemplate string
		// // Enable suggestions for commands and flags
		// Suggest bool
		// // Other custom info
		// Metadata map[string]interface{}
		// // Carries a function which returns app specific info.
		// ExtraInfo func() map[string]string
		// Error codes
		// TODO:

		// // The function to call when checking for bash command completions
		// BashComplete BashCompleteFunc
		// // Boolean to enable bash completion commands
		// EnableBashCompletion bool

		// List of flags to parse
		Flags []Flag

		Before Func
		Action Func
		After  Func // ErrFunc

		// List of commands to execute
		SubCommands map[string]*Command

		// // Execute this function when an invalid flag is accessed from the context
		// InvalidFlagAccessHandler InvalidFlagAccessFunc

		// default to log.Println to stderr
		Log io.Writer
	}

	Func func(ctx Context) (_ error, exitCode int)
	// ErrFunc func(ctx Context, err error) error

	Context struct {
		// ctx context.Context
	}
)

//	func (ctx *Context) String() (string, bool) {
//		// oh no
//	}
var (
	keyVersion = "version"
	keyHelp    = "help"
)

// For root command, name is usually os.Args[0]
func (c Command) Run(name string, args []string) (exitCode int) {
	versionOK := !flagNameUsed(c.Flags, keyVersion)
	if versionOK { // TODO: move to global options
		// this sounds like a horrible idea, not using the parser
		if len(args) > 1 && strings.EqualFold(args[1], "--version") {
			version()
			return 0
		}
		c.Flags = append(c.Flags, BoolFlag{Options: []string{keyVersion}})
	}

	// helpOK := !flagNameUsed(c.Flags, keyHelp)
	// if helpOK {
	// 	c.Flags = append(c.Flags, BoolFlag{Options: []string{keyHelp}})
	// }

	// global flags??

	// duplicate parsing: env could be only parsed once;
	// options chokereturn is kinda pointless, as all options are unordered anyway, instead command detection would be nice
	// say from first parsed argument
	// parse opts and env

	// merge opts and env based on c.Flags

	// if !equalFoldsSlice(parsed[0], mapKeys(c.SubCommands)) {
	// 	return c.run(name, args)
	// }

	// return c.SubCommands[chokeReturn[0]].run(chokeReturn[0], chokeReturn[1:])

	return 1
}

// this is called for any subcommands under Run(), difference being --version
func (c Command) run(name string, args []string) (exitCode int) {
	// handle (possible local) --help
}

func (c Command) normalize() error {
	// no duplicate env keys
	// no duplicate option keys within a tree
	// enforce uppercase for env and lowercase for long opts
}

func equalFoldsSlice(s string, target []string) bool {
	for _, t := range target {
		if strings.EqualFold(s, t) {
			return true
		}
	}

	return false
}

// all env will be uppercased, 1-letter env is forbidden, long options will be lowercased
// does not use Default nor Condition
func (c Command) parseArgs(args []string) (_ harg.Definitions, parsed []string, _ error) {
	optDefs, envDefs, commonDefs := make(harg.Definitions), make(harg.Definitions), make(harg.Definitions)

	for _, flag := range c.Flags {
		f := flag.flag()
		def := harg.Definition{Type: f.Type, AlsoBool: f.AlsoBool, EnvCSV: f.EnvCSV}

		for _, opt := range f.Options {
			// lowercase long options
			if utf8.RuneCountInString(opt) > 1 {
				opt = strings.ToLower(opt)
			}

			if ok := optDefs.SetUnique(opt, &def); !ok {
				return nil, nil, fmt.Errorf("option name %s has duplicates (not unique): %w", opt, harg.ErrInvalidDefinition)
			}
			commonDefs[opt] = &def
		}

		opt := strings.ToUpper(f.Env)
		if utf8.RuneCountInString(opt) < 2 {
			// would conflict with short options
			return nil, nil, fmt.Errorf("environment name %q must be at least 2 characters: %w", opt, harg.ErrInvalidDefinition)
		}

		envDefs[opt] = &def
		if ok := commonDefs.SetUnique(opt, &def); !ok {
			return nil, nil, fmt.Errorf("environment name %s has duplicates (not unique): %w", opt, harg.ErrInvalidDefinition)
		}
	}

	if err := envDefs.ParseEnv(); err != nil {
		return nil, nil, fmt.Errorf("parsing environment: %w", err)
	}

	parsed, _, err := optDefs.Parse(args, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("parsing options: %w", err)
	}

	return commonDefs, parsed, nil
}

// TODO: *Command doesn't want to fit in to any
func mapKeys[T comparable](m map[T]*Command) (keys []T) {
	for key := range m {
		keys = append(keys, key)
	}

	return
}

// uh so option parsing shall be recursive double-defined implicit globals:
// implicit globals: globals are defined once (in a subcommand tree) with flag property Global bool (name up to debate)
// this makes globals available before subcommand
// double-defined: it also makes the globals local
// (local options are unavailable before subcommand)
// recursive: subcommands may have subcommands, and globals are handled there as well
//
// conditions and defaults shall be applied after all parsing is done (replacing .Default() bool)
// acting on global variables is the responsibility of all subcommands

// TODO: maybe provide a convenience error wrapping to include an exit code
