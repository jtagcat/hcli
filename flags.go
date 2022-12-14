package hcli

import (
	"fmt"
	"strings"

	"github.com/jtagcat/hcli/harg"
)

type (
	Flag interface {
		flag() flag
	}

	FlagCondition func(defaultValue any, def *harg.Definition) error
)

type FlagSource uint32 // enum:
const (
	OptEnv FlagSource = iota // Prefer Options (when set), fallback to Environment
	EnvOpt                   // Prefer Environment, fallback to Options
) //
var typeMax = EnvOpt

// Requires the user to set the flag. Setting to default is valid.
//
// Implements FlagCondition
func Defined(_ any, def *harg.Definition) error {
	if def.Default() {
		return fmt.Errorf("must be set")
	}

	return nil
}

// Requires the user to set the flag. Setting to default is valid.
//
// Implements FlagCondition
func NotDefault(defaultV any, def *harg.Definition) error {
	var set any

	switch def.Type {
	case harg.Bool:
		set, _ = def.Bool()
	case harg.String:
		set, _ = def.String()
	case harg.Int:
		set, _ = def.Int()
	case harg.Int64:
		set, _ = def.Int64()
	case harg.Uint:
		set, _ = def.Uint()
	case harg.Uint64:
		set, _ = def.Uint64()
	case harg.Float64:
		set, _ = def.Float64()
	case harg.Duration:
		set, _ = def.Duration()
	}

	if set != defaultV {
		return nil
	}

	return fmt.Errorf("must be non-default, default is %q", defaultV)
}

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

type flag struct {
	Type    harg.Type
	Default any

	Options []string

	Env    string
	EnvCSV bool

	AlsoBool bool

	Condition FlagCondition

	Usage string
}

// Generatable:

// bool

type BoolFlag struct {
	Options []string

	Env    string
	EnvCSV bool

	Default   bool
	Condition FlagCondition

	Usage string
}

func (b BoolFlag) flag() flag {
	return flag{
		Type:    harg.Bool,
		Default: b.Default,

		AlsoBool: false,

		Options:   b.Options,
		Env:       b.Env,
		EnvCSV:    b.EnvCSV,
		Condition: b.Condition,
		Usage:     b.Usage,
	}
}

// string

type StringFlag struct {
	Options []string

	Env    string
	EnvCSV bool

	AlsoBool bool

	Default   string
	Condition FlagCondition

	Usage string
}

func (b StringFlag) flag() flag {
	return flag{
		Type:    harg.String,
		Default: b.Default,

		Options:   b.Options,
		Env:       b.Env,
		EnvCSV:    b.EnvCSV,
		AlsoBool:  b.AlsoBool,
		Condition: b.Condition,
		Usage:     b.Usage,
	}
}

// TODO: ...
