package hcli

import (
	"context"
	"fmt"
	"strings"

	"github.com/jtagcat/hcli/harg"
	"golang.org/x/exp/slog"
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

		// TODO:
		// Before Func
		Action Func
		// After  Func // ErrFunc

		// List of commands to execute
		SubCommands map[string]*Command
		parent      parentCommand

		// // Execute this function when an invalid flag is accessed from the context
		// InvalidFlagAccessHandler InvalidFlagAccessFunc

		// TODO:
		// default to log.Println to stderr
		Logger *slog.Logger
	}
	parentCommand struct {
		name    string
		command *Command
	}

	Func func(ctx context.Context,
		args []string, opts harg.Definitions,
		log *slog.Logger,
	) (exitCode int)
	// ErrFunc func(ctx Context, err error) error
)

//	func (ctx *Context) String() (string, bool) {
//		// oh no
//	}

// TODO: version and help
const (
	keyVersion = "version"
	keyHelp    = "help"
)

var DisallowedKeys = []string{keyVersion, keyHelp}

func (c *Command) ValidateTree() error {
	// TODO: uniqueness of flags and env
	// TODO: help and version is overwritten
	return fmt.Errorf("not implemented")
}

// For root command, name is usually os.Args[0], args os.Args[1:].
// See exit_codes.go for exitCode guidance.
func (c *Command) Run(ctx context.Context, name string, args []string) (exitCode int) {
	return c.run(ctx, name, args, nil)
}

func (c *Command) run(ctx context.Context, name string, args []string, parentDefs harg.Definitions) (exitCode int) {
	var subNames []string
	for name := range c.SubCommands {
		subNames = append(subNames, name)
	}

	var allDefs harg.Definitions
	for _, flag := range c.Flags {
		flag := flag.flag()
		def := flag.def()

		for _, name := range flag.Options {
			allDefs[name] = &def
		}
	}

	cleanArgs, choke, err := allDefs.Parse(args, subNames)
	if err != nil {
		c.Logger.Error("parsing arguments", err, slog.String("parse_type", "initial"), slog.String("command_name", name))
		return ExitUsage
	}

	defs := mergeDefs(parentDefs, allDefs) // previous + current

	// switch to subcommand
	if cleanArgs == nil && len(choke) != 0 {
		subName := strings.ToLower(choke[0])
		subCommand := c.SubCommands[subName]

		subCommand.parent = parentCommand{name: name, command: c}

		return subCommand.run(ctx, subName, choke[1:], defs)
	}

	// TODO: parse env (according to parent tree)

	// cleanArgs, _, err := defs.Parse()

	// TODO: validate all options (according to parent tree)
	// TODO: validate that defs is not nil, default logger?

	return c.Action(ctx, cleanArgs, defs, c.Logger)
}

// TODO:
func mergeDefs(previousLevel, currentLevel harg.Definitions) harg.Definitions {
	if len(previousLevel) == 0 {
		return currentLevel
	}

	panic("merge not implemented") // TODO:
}

// func (c *Command) normalize() error {
// 	// no duplicate env keys
// 	// no duplicate option keys within a tree
// 	// enforce uppercase for env and lowercase for long opts
// }

// func equalFoldsSlice(s string, target []string) bool {
// 	for _, t := range target {
// 		if strings.EqualFold(s, t) {
// 			return true
// 		}
// 	}

// 	return false
// }

// // all env will be uppercased, 1-letter env is forbidden, long options will be lowercased
// // does not use Default nor Condition
// func (c *Command) parseArgs(args []string) (_ harg.Definitions, parsed []string, _ error) {
// 	optDefs, envDefs, commonDefs := make(harg.Definitions), make(harg.Definitions), make(harg.Definitions)

// 	for _, flag := range c.Flags {
// 		f := flag.flag()
// 		def := harg.Definition{Type: f.Type, AlsoBool: f.AlsoBool, EnvCSV: f.EnvCSV}

// 		for _, opt := range f.Names {
// 			// lowercase long options
// 			if utf8.RuneCountInString(opt) > 1 {
// 				opt = strings.ToLower(opt)
// 			}

// 			if ok := optDefs.SetUnique(opt, &def); !ok {
// 				return nil, nil, fmt.Errorf("option name %s has duplicates (not unique): %w", opt, harg.ErrInvalidDefinition)
// 			}
// 			commonDefs[opt] = &def
// 		}

// 		opt := strings.ToUpper(f.Env)
// 		if utf8.RuneCountInString(opt) < 2 {
// 			// would conflict with short options
// 			return nil, nil, fmt.Errorf("environment name %q must be at least 2 characters: %w", opt, harg.ErrInvalidDefinition)
// 		}

// 		envDefs[opt] = &def
// 		if ok := commonDefs.SetUnique(opt, &def); !ok {
// 			return nil, nil, fmt.Errorf("environment name %s has duplicates (not unique): %w", opt, harg.ErrInvalidDefinition)
// 		}
// 	}

// 	if err := envDefs.ParseEnv(); err != nil {
// 		return nil, nil, fmt.Errorf("parsing environment: %w", err)
// 	}

// 	parsed, _, err := optDefs.Parse(args, nil)
// 	if err != nil {
// 		return nil, nil, fmt.Errorf("parsing options: %w", err)
// 	}

// 	return commonDefs, parsed, nil
// }

// TODO: continue working on run():
// - defs logic
// - traversing parent tree for globals
// - maybe allow harg to take in an arbritrary value (pointer)
// - if all else fails, build a reverse tree and start parsing from 0

// TODO: maybe global options for:
// - should unimplemented (yet specified by end user) be ignored or errored at
