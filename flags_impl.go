package hcli

import (
	"fmt"

	"github.com/jtagcat/hcli/harg"
)

// implements FlagCondition //

// Requires the user to set the flag. Setting to default is valid.
func Defined(_ any, def *harg.Definition) error {
	if def.Default() {
		return fmt.Errorf("must be set")
	}

	return nil
}

// Requires the user to set the flag. Setting to default is valid.
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

// implements flag //

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

// Generatable:

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
