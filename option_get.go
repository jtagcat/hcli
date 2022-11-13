package harg

import (
	"fmt"
	"reflect"
	"time"
)

// Was there any option parsed matching the definition of slug
func (def *Definition) Touched() bool {
	return def.parsed.found
}

func (def *Definition) ok(needed Type) error {
	meta, ok := typeMetaM[def.Type]
	if !ok {
		return fmt.Errorf("Definition.ok(): Type %s not found in typeMetaM: %w", typeMetaM[def.Type].errName, ErrInternalBug)
	}

	if def.Type != needed { // not checking for needed Type to be in the map
		return fmt.Errorf("method needs enum Type %s, is incompatible with %s: %w", typeMetaM[needed].errName, typeMetaM[def.Type].errName, ErrIncompatibleMethod)
	}

	got := reflect.TypeOf(def.parsed.opt).Elem().Name()
	want := reflect.TypeOf(meta.emptyT).Elem().Name()
	if got != want {
		return fmt.Errorf("method for Type %s expected iface %s, is incompatible with parsed iface %s, %w", typeMetaM[def.Type].errName, want, got, ErrIncompatibleMethod)
	}

	return nil
}

// bool

func (def *Definition) SlBool() ([]bool, error) {
	if err := def.ok(Bool); err != nil {
		return nil, fmt.Errorf("SlBool: %w", err)
	}

	return def.parsed.opt.contents().(optBoolVal).value, nil
}

func (def *Definition) Bool() (bool, error) {
	sl, err := def.SlBool()
	if err != nil || len(sl) < 1 {
		return false, err
	}
	return sl[len(sl)-1], nil // last defined
}

// count is equal to count of consequtive true bools counting from right
//
// true false true true: 2,
// true false: 0,
// true: 1,
// true true true: 3,
func (def *Definition) Count() (int, error) {
	if err := def.ok(Bool); err != nil {
		return 0, fmt.Errorf("Count: %w", err)
	}
	return def.parsed.opt.contents().(optBoolVal).count, nil
}

// AlsoBool
func (def *Definition) IsBool() bool {
	return def.Type == Bool // type is changed to e_bool on parsing if
}

//// generatable ////

// string

func (def *Definition) SlString() ([]string, error) {
	if err := def.ok(String); err != nil {
		return nil, fmt.Errorf("SlString: %w", err)
	}
	return def.parsed.opt.contents().([]string), nil
}

func (def *Definition) String() (string, error) {
	sl, err := def.SlString()
	if err != nil || len(sl) < 1 {
		return "", err
	}
	return sl[len(sl)-1], nil // last defined
}

// int

func (def *Definition) SlInt() ([]int, error) {
	if err := def.ok(Int); err != nil {
		return nil, fmt.Errorf("SlInt: %w", err)
	}
	return def.parsed.opt.contents().([]int), nil
}

func (def *Definition) Int() (int, error) {
	sl, err := def.SlInt()
	if err != nil || len(sl) < 1 {
		return 0, err
	}
	return sl[len(sl)-1], nil // last defined
}

// int64

func (def *Definition) SlInt64() ([]int64, error) {
	if err := def.ok(Int); err != nil {
		return nil, fmt.Errorf("SlInt64: %w", err)
	}
	return def.parsed.opt.contents().([]int64), nil
}

func (def *Definition) Int64() (int64, error) {
	sl, err := def.SlInt64()
	if err != nil || len(sl) < 1 {
		return 0, err
	}
	return sl[len(sl)-1], nil // last defined
}

// uint

func (def *Definition) SlUint() ([]uint, error) {
	if err := def.ok(Int); err != nil {
		return nil, fmt.Errorf("SlUint: %w", err)
	}
	return def.parsed.opt.contents().([]uint), nil
}

func (def *Definition) Uint() (uint, error) {
	sl, err := def.SlUint()
	if err != nil || len(sl) < 1 {
		return 0, err
	}
	return sl[len(sl)-1], nil // last defined
}

// uint64

func (def *Definition) SlUint64() ([]uint64, error) {
	if err := def.ok(Int); err != nil {
		return nil, fmt.Errorf("SlUint64: %w", err)
	}
	return def.parsed.opt.contents().([]uint64), nil
}

func (def *Definition) Uint64() (uint64, error) {
	sl, err := def.SlUint64()
	if err != nil || len(sl) < 1 {
		return 0, err
	}
	return sl[len(sl)-1], nil // last defined
}

// float64

func (def *Definition) SlFloat64() ([]float64, error) {
	if err := def.ok(Int); err != nil {
		return nil, fmt.Errorf("SlFloat64: %w", err)
	}
	return def.parsed.opt.contents().([]float64), nil
}

func (def *Definition) Float64() (float64, error) {
	sl, err := def.SlFloat64()
	if err != nil || len(sl) < 1 {
		return 0, err
	}
	return sl[len(sl)-1], nil // last defined
}

// duration

func (def *Definition) SlDuration() ([]time.Duration, error) {
	if err := def.ok(Int); err != nil {
		return nil, fmt.Errorf("SlDuration: %w", err)
	}
	return def.parsed.opt.contents().([]time.Duration), nil
}

func (def *Definition) Duration() (time.Duration, error) {
	sl, err := def.SlDuration()
	if err != nil || len(sl) < 1 {
		return 0, err
	}
	return sl[len(sl)-1], nil // last defined
}
