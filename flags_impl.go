package hcli

import (
	"fmt"

	"github.com/jtagcat/hcli/harg"
	"golang.org/x/exp/slices"
)

// implements Condition
//
// Requires the user to set Flag. Setting to default is valid.
func Defined[T comparable](_ []T, def *harg.Definition) error {
	if def.Default() {
		return fmt.Errorf("must be set")
	}

	return nil
}

func getDefault[T any]() (defaultValue T) {
	return
}

// implements Condition
//
// Requires the user to set Flag to something other than the default.
// NOTE: if default is []string{"hello"}, then non-default []string{"hello", "hello"} will return the "default" with .String() (not .SlString)
func NotDefaultSl[T comparable](defaultValue []T, def *harg.Definition) error {
	gotAny, _ := def.Any()
	got, _ := gotAny.([]T) // caller ensures defaultValue and definition match

	if slices.Equal(got, defaultValue) {
		return fmt.Errorf("must be non-default, default: %v", defaultValue)
	}

	return nil
}

// implements Condition
//
// Requires the user to set the flag's first value to something other than the default.
func NotDefault[T comparable](defaultValue []T, def *harg.Definition) error {
	gotAny, _ := def.Any()
	got, _ := gotAny.([]T) // caller ensures defaultValue and definition match

	if got[0] == defaultValue[0] {
		return fmt.Errorf("must be non-default, default: %v", defaultValue[0])
	}

	return nil
}

// implements Flag for bool
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

// TODO: Generatable:

// implements Flag for string
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
