package hcli

import (
	"fmt"
	"strings"

	"github.com/jtagcat/hcli/harg"
)

type (
	Flag interface {
		Type() harg.Type
		options() []string
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
		for _, o := range f.options() {
			if strings.EqualFold(o, name) {
				return true
			}
		}
	}
	return false
}

// Flag structs contain:
//	Options []string

//	Env    []string
//	EnvCSV bool

//	AlsoBool bool // only for non-bools
//	Priority FlagSource
//	Local  bool // if local, isn't available for subcommands

//	Default *<flag type>
//	Condition FlagCondition

//	Usage string

// Generatable:

// bool

type BoolFlag struct {
	Options []string

	Env    []string
	EnvCSV bool

	Priority FlagSource
	Local    bool // if local, isn't available for subcommands

	Default   bool
	Condition FlagCondition

	Usage string
}

func (_ BoolFlag) Type() harg.Type {
	return harg.Bool
}

func (f BoolFlag) options() []string {
	return f.Options
}

// string

type StringFlag struct {
	Options []string

	Env    []string
	EnvCSV bool

	AlsoBool bool
	Priority FlagSource
	Local    bool // if local, isn't available for subcommands

	Default   string
	Condition FlagCondition

	Usage string
}

func (_ StringFlag) Type() harg.Type {
	return harg.String
}

func (f StringFlag) options() []string {
	return f.Options
}

// TODO: ...
