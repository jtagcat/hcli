package hcli

import (
	"fmt"

	"github.com/jtagcat/hcli/harg"
)

// implements Condition //

// Requires the user to set the flag. Setting to default is valid.
func Defined(def *harg.Definition) error {
	if def.Default() {
		return fmt.Errorf("must be set")
	}

	return nil
}

func getDefault[T any]() (defaultValue T) {
	return
}

// Requires the user to set the flag. Setting to default is valid.
func NotDefault[T comparable](def *harg.Definition) error {
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

	if got == defaultV {
		return fmt.Errorf("must be non-default, default is %q", defaultV)
	}

	return nil
}

// implements flag //

// bool

type (
	BoolFlag struct {
		Level FlagLevel // Local/Global/Child/Parent

		Type     harg.Type
		AlsoBool bool

		Options []string

		Env    string
		EnvCSV bool

		Default   bool // value to set when nothing is set
		Condition boolCondition

		Usage string
	}
	boolCondition func(got bool, def *harg.Definition) error
)

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
	baseFlag

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
