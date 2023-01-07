package hcli

import (
	"github.com/jtagcat/hcli/harg"
)

type (
	Flag interface {
		flag() flag
		checkCondition(*harg.Definition) error
	}

	flag struct {
		Level FlagLevel // Local/Global/Child/Parent

		Type harg.Type

		AlsoBool bool
		Options  []string

		Env    string
		EnvCSV bool

		Default any // value to set when nothing is set

		Usage string

		child *ChildFlag
	}
)

type FlagLevel uint8 // enum:
const (
	Local  FlagLevel = iota // only available in the defined command
	Global                  // available in the command, and all subcommands implementing it (with Child)
	Parent                  // available in all subcommands implementing it, not available in the same command
) //
var flagLevelMax = Parent

func (f *flag) def() harg.Definition {
	return harg.Definition{
		Type:     f.Type,
		AlsoBool: f.AlsoBool,
		EnvCSV:   f.EnvCSV,
	}
}
