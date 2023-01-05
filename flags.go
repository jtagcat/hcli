package hcli

import (
	"strings"

	"github.com/jtagcat/hcli/harg"
)

type (
	Flag interface {
		flag() flag
	}
	flag struct {
		Level FlagLevel // Local/Global/Child/Parent

		Type    harg.Type
		Default any

		Options []string

		Env    string
		EnvCSV bool

		AlsoBool bool

		Condition FlagCondition

		Usage string
	}
	FlagCondition func(defaultValue any, def *harg.Definition) error
)

type FlagLevel uint8 // enum:
const (
	Local  FlagLevel = iota // only available in the defined command
	Global                  // available in the command, and all subcommands implementing it (with Child)
	Parent                  // available in all subcommands implementing it, not available in the same command
	Child
) //
var flagLevelMax = Child

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
