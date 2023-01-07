package harg

import (
	"time"
)

// Whether definition's value is a default value (was it set via Parse())
func (def *Definition) Default() bool {
	if def == nil {
		return false
	}
	return def.parsed == nil
}

func (def *Definition) SlAny() (v any, ok bool) {
	switch def.Type {
	case Bool:
		return def.SlBool()
	case String:
		return def.SlString()
	case Int:
		return def.SlInt()
	case Int64:
		return def.SlInt64()
	case Uint:
		return def.SlUint()
	case Uint64:
		return def.SlUint64()
	case Float64:
		return def.SlFloat64()
	case Duration:
		return def.SlDuration()
	default:
		panic("unreachable") // tested against
	}
}

func (def *Definition) Any() (v any, ok bool) {
	switch def.Type {
	case Bool:
		return def.Bool()
	case String:
		return def.String()
	case Int:
		return def.Int()
	case Int64:
		return def.Int64()
	case Uint:
		return def.Uint()
	case Uint64:
		return def.Uint64()
	case Float64:
		return def.Float64()
	case Duration:
		return def.Duration()
	default:
		panic("unreachable") // tested against
	}
}

// For checking if AlsoBool's type was changed to Bool on parsing.
func (def *Definition) IsBool() bool {
	if def == nil {
		return false
	}
	return def.Type == Bool
}

// Count of consecutive true values read from right/last
//
// Examples:
//
//	true false true true: 2
//	true false: 0
//	true: 1
//	true true true: 3
func (def *Definition) Count() (v int, ok bool) {
	sl, ok := def.SlBool()
	if !ok || len(sl) == 0 {
		return
	}

	for i := len(sl) - 1; i >= 0; i-- {
		if !sl[i] {
			break
		}
		v++
	}
	return v, true
}

//// generatable ////

// bool

func (o optBool) contents() any {
	return o.value
}

func (def *Definition) SlBool() ([]bool, bool) {
	// not seen/parsed or mismatched type
	if def.Default() || def.Type != Bool {
		return nil, false
	}

	return def.parsed.contents().([]bool), true
}

func (def *Definition) Bool() (v bool, ok bool) {
	sl, ok := def.SlBool()
	if !ok || len(sl) == 0 {
		return
	}
	return sl[len(sl)-1], true // last
}

// string

func (o optString) contents() any {
	return o.value
}

func (def *Definition) SlString() ([]string, bool) {
	// not seen/parsed or mismatched type
	if def.Default() || def.Type != String {
		return nil, false
	}

	return def.parsed.contents().([]string), true
}

func (def *Definition) String() (v string, ok bool) {
	sl, ok := def.SlString()
	if !ok || len(sl) == 0 {
		return
	}
	return sl[len(sl)-1], true // last
}

// int

func (o optInt) contents() any {
	return o.value
}

func (def *Definition) SlInt() ([]int, bool) {
	// not seen/parsed or mismatched type
	if def.Default() || def.Type != Int {
		return nil, false
	}

	return def.parsed.contents().([]int), true
}

func (def *Definition) Int() (v int, ok bool) {
	sl, ok := def.SlInt()
	if !ok || len(sl) == 0 {
		return
	}
	return sl[len(sl)-1], true // last
}

// int64

func (o optInt64) contents() any {
	return o.value
}

func (def *Definition) SlInt64() ([]int64, bool) {
	// not seen/parsed or mismatched type
	if def.Default() || def.Type != Int64 {
		return nil, false
	}

	return def.parsed.contents().([]int64), true
}

func (def *Definition) Int64() (v int64, ok bool) {
	sl, ok := def.SlInt64()
	if !ok || len(sl) == 0 {
		return
	}
	return sl[len(sl)-1], true // last
}

// uint

func (o optUint) contents() any {
	return o.value
}

func (def *Definition) SlUint() ([]uint, bool) {
	// not seen/parsed or mismatched type
	if def.Default() || def.Type != Uint {
		return nil, false
	}

	return def.parsed.contents().([]uint), true
}

func (def *Definition) Uint() (v uint, ok bool) {
	sl, ok := def.SlUint()
	if !ok || len(sl) == 0 {
		return
	}
	return sl[len(sl)-1], true // last
}

// uint64

func (o optUint64) contents() any {
	return o.value
}

func (def *Definition) SlUint64() ([]uint64, bool) {
	// not seen/parsed or mismatched type
	if def.Default() || def.Type != Uint64 {
		return nil, false
	}

	return def.parsed.contents().([]uint64), true
}

func (def *Definition) Uint64() (v uint64, ok bool) {
	sl, ok := def.SlUint64()
	if !ok || len(sl) == 0 {
		return
	}
	return sl[len(sl)-1], true // last
}

// float64

func (o optFloat64) contents() any {
	return o.value
}

func (def *Definition) SlFloat64() ([]float64, bool) {
	// not seen/parsed or mismatched type
	if def.Default() || def.Type != Float64 {
		return nil, false
	}

	return def.parsed.contents().([]float64), true
}

func (def *Definition) Float64() (v float64, ok bool) {
	sl, ok := def.SlFloat64()
	if !ok || len(sl) == 0 {
		return
	}
	return sl[len(sl)-1], true // last
}

// duration

func (o optDuration) contents() any {
	return o.value
}

func (def *Definition) SlDuration() ([]time.Duration, bool) {
	// not seen/parsed or mismatched type
	if def.Default() || def.Type != Duration {
		return nil, false
	}

	return def.parsed.contents().([]time.Duration), true
}

func (def *Definition) Duration() (v time.Duration, ok bool) {
	sl, ok := def.SlDuration()
	if !ok || len(sl) == 0 {
		return
	}
	return sl[len(sl)-1], true // last
}
