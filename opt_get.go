package harg

// stuff will panic on undefined slugs, expecting linter?

import "fmt"

// Was there any option parsed matching the definition of slug
func (optM *OptionsTypedMap) Touched(slug string) bool {
	for t := range *optM {
		if v, ok := (*optM)[t][slug]; ok {
			return v.found(false)
		}
	}
	panic(fmt.Sprintf("harg.OptionsTypedMap.Touched(): slug %s does not exist (undefined option)", slug))
}

// bool

func (optM *OptionsTypedMap) SlBool(slug string) []bool {
	return (*optM)[e_string][slug].contents().(optBoolVal).value
}

func (optM *OptionsTypedMap) Bool(slug string) bool {
	sl := optM.SlBool(slug)
	return sl[len(sl)-1] // last defined
}

// count is equal to count of consequtive true bools counting from right
//
// true false true true: 2,
// true false: 0,
// true: 1,
// true true true: 3,
func (optM *OptionsTypedMap) Count(slug string) int {
	return (*optM)[e_string][slug].contents().(optBoolVal).count
}

// AlsoBool
func (optM *OptionsTypedMap) IsBool(slug string) bool {
	return false // TODO:
	// return (*optM)[e_string][slug].contents().(optBoolVal).count
}

//// generatable ////

// string

func (optM *OptionsTypedMap) SlString(slug string) []string {
	return (*optM)[e_string][slug].contents().([]string)
}

func (optM *OptionsTypedMap) String(slug string) string {
	sl := optM.SlString(slug)
	return sl[len(sl)-1] // last defined
}

// int

func (optM *OptionsTypedMap) SlInt(slug string) []int {
	return (*optM)[e_int][slug].contents().([]int)
}

func (optM *OptionsTypedMap) Int(slug string) int {
	sl := optM.SlInt(slug)
	return sl[len(sl)-1] // last defined
}

// int64

func (optM *OptionsTypedMap) SlInt64(slug string) []int64 {
	return (*optM)[e_int64][slug].contents().([]int64)
}

func (optM *OptionsTypedMap) Int64(slug string) int64 {
	sl := optM.SlInt64(slug)
	return sl[len(sl)-1] // last defined
}

// uint

func (optM *OptionsTypedMap) SlUint(slug string) []uint {
	return (*optM)[e_uint][slug].contents().([]uint)
}

func (optM *OptionsTypedMap) Uint(slug string) uint {
	sl := optM.SlUint(slug)
	return sl[len(sl)-1] // last defined
}

// uint64

func (optM *OptionsTypedMap) SlUint64(slug string) []uint64 {
	return (*optM)[e_uint64][slug].contents().([]uint64)
}

func (optM *OptionsTypedMap) Uint64(slug string) uint64 {
	sl := optM.SlUint64(slug)
	return sl[len(sl)-1] // last defined
}

// float64

func (optM *OptionsTypedMap) SlFloat64(slug string) []float64 {
	return (*optM)[e_float64][slug].contents().([]float64)
}

func (optM *OptionsTypedMap) Float64(slug string) float64 {
	sl := optM.SlFloat64(slug)
	return sl[len(sl)-1] // last defined
}

// duration

func (optM *OptionsTypedMap) SlDuration(slug string) []float64 {
	return (*optM)[e_duration][slug].contents().([]float64)
}

func (optM *OptionsTypedMap) Duration(slug string) float64 {
	sl := optM.SlFloat64(slug)
	return sl[len(sl)-1] // last defined
}
