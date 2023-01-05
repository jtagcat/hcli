package hcli

import (
	"strings"

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
	}
)

type FlagLevel uint8 // enum:
const (
	Local  FlagLevel = iota // only available in the defined command
	Global                  // available in the command, and all subcommands implementing it (with Child)
	Parent                  // available in all subcommands implementing it, not available in the same command
	Child
) //
var flagLevelMax = Child

func (f *flag) def() harg.Definition {
	return harg.Definition{
		Type:     f.Type,
		AlsoBool: f.AlsoBool,
		EnvCSV:   f.EnvCSV,
	}
}

// TODO: remove?
func flagNameUsed(flags []Flag, name string) bool {
	for _, f := range flags {
		for _, o := range f.flag().Options {
			if strings.EqualFold(o, name) {
				return true
			}
		}
	}
	return false
}
