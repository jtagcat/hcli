package hcli

import (
	"io"
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
		Flags Flags

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

// func (ctx *Context) String() (string, bool) {
// 	// oh no
// }

// For root command, name is usually os.Args[0]
func (c Command) Run(name string, args []string) (exitCode int) {
	versionUsed := c.Flags.flagNameUsed("version")
	if !versionUsed {
		c.Flags[nonConflictingKey(c.Flags, "version")] = BoolFlag{Options: []string{"version"}}
	}
	helpUsed := c.Flags.flagNameUsed("help")
	if !helpUsed {
		c.Flags[nonConflictingKey(c.Flags, "help")] = BoolFlag{Options: []string{"help"}}
	}

	// parse

	// version

	// subcommand switch

	// subcommand help
	// local?? help

	// run()

	return 1
}

// bad function name
func nonConflictingKey(m Flags, name string) string { // m map[string]any
	original := name

	for i := 2; true; i++ {
		if _, ok := m[name]; !ok {
			return name
		}

		name = original + "-" + string(i)
	}

	return "" // unreachable (go compiler requires it, though)
}

// this is called for any subcommands under Run()
func (c Command) run(name string, args []string) (exitCode int) {
}
