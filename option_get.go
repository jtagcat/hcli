package harg

import (
	"time"
)

// Whether there was any option parsed matching the definition.
func (def *Definition) Touched() bool {
	return def.parsed != nil
}

// Whether AlsoBool's type was changed to Bool on parsing.
func (def *Definition) IsBool() bool {
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
	if def.Type != Bool || def.parsed == nil {
		return
	}

	return def.parsed.contents().(optBoolVal).count, true
}

//// generatable ////

// bool

func (def *Definition) SlBool() ([]bool, bool) {
	// mismatched type or !def.Touched()
	if def.Type != Bool || def.parsed == nil {
		return nil, false
	}

	return def.parsed.contents().(optBoolVal).value, true
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
	// mismatched type or !def.Touched()
	if def.Type != String || def.parsed == nil {
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
	// mismatched type or !def.Touched()
	if def.Type != Int || def.parsed == nil {
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

func (o *optInt64) contents() any {
	return o.value
}

func (def *Definition) SlInt64() ([]int64, bool) {
	// mismatched type or !def.Touched()
	if def.Type != Int64 || def.parsed == nil {
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

func (o *optUint) contents() any {
	return o.value
}

func (def *Definition) SlUint() ([]uint, bool) {
	// mismatched type or !def.Touched()
	if def.Type != Uint || def.parsed == nil {
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

func (o *optUint64) contents() any {
	return o.value
}

func (def *Definition) SlUint64() ([]uint64, bool) {
	// mismatched type or !def.Touched()
	if def.Type != Uint64 || def.parsed == nil {
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

func (o *optFloat64) contents() any {
	return o.value
}

func (def *Definition) SlFloat64() ([]float64, bool) {
	// mismatched type or !def.Touched()
	if def.Type != Float64 || def.parsed == nil {
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

func (o *optDuration) contents() any {
	return o.value
}

func (def *Definition) SlDuration() ([]time.Duration, bool) {
	// mismatched type or !def.Touched()
	if def.Type != Duration || def.parsed == nil {
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
