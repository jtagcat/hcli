package hcli

import (
	"fmt"

	"github.com/jtagcat/hcli/harg"
)

// implements Condition //

// Requires the user to set the flag. Setting to default is valid.
func Defined[T comparable](_ T, def *harg.Definition) error {
	if def.Default() {
		return fmt.Errorf("must be set")
	}

	return nil
}

func getDefault[T any]() (defaultValue T) {
	return
}

// Requires the user to set the flag. Setting to default is valid.
func NotDefault[T comparable](defaultValue T, def *harg.Definition) error {
	got, _ := def.Any()

	if got != defaultValue {
		return fmt.Errorf("must be non-default, default is %q", defaultValue)
	}

	return nil
}

// implements flag //

// bool

type (
	BoolFlag struct {
		Level FlagLevel // Local/Global/Child/Parent

		Options []string

		Env    string
		EnvCSV bool

		Default   []bool // value to set when nothing is set
		Condition boolCondition

		Usage string
	}
	boolCondition func(flagDefault []bool, def *harg.Definition) error
)

func (b *BoolFlag) flag() flag {
	return flag{
		AlsoBool: false,

		Type:    harg.Bool,
		Default: b.Default,

		Level:   b.Level,
		Options: b.Options,
		Env:     b.Env,
		EnvCSV:  b.EnvCSV,
		Usage:   b.Usage,
	}
}

func (f *BoolFlag) checkCondition(def *harg.Definition) error {
	return f.Condition(f.Default, def)
}

// Generatable:

// string

type (
	StringFlag struct {
		Level FlagLevel // Local/Global/Child/Parent

		AlsoBool bool
		Options  []string

		Env    string
		EnvCSV bool

		Default   []string // value to set when nothing is set
		Condition stringCondition

		Usage string
	}
	stringCondition func(flagDefault []string, def *harg.Definition) error
)

func (b *StringFlag) flag() flag {
	return flag{
		Type:    harg.String,
		Default: b.Default,

		Level:    b.Level,
		Options:  b.Options,
		Env:      b.Env,
		EnvCSV:   b.EnvCSV,
		AlsoBool: b.AlsoBool,
		Usage:    b.Usage,
	}
}

func (f *StringFlag) checkCondition(def *harg.Definition) error {
	return f.Condition(f.Default, def)
}
