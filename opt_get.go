package harg

// stuff will panic on undefined slugs, expecting linter?

import "fmt"

// Was there any option parsed matching the definition of slug
func (optM *OptionsMap) Touched(slug string) bool {
	if v, ok := (*optM)[slug]; ok {
		return v.found(false)
	}

	panic(fmt.Sprintf("harg.OptionsTypedMap.Touched(): slug %s does not exist (undefined option)", slug))
}

// bool

func (optM *OptionsMap) SlBool(slug string) []bool {
	return (*optM)[slug].contents().(optBoolVal).value
}

func (optM *OptionsMap) Bool(slug string) bool {
	sl := optM.SlBool(slug)
	return sl[len(sl)-1] // last defined
}

// count is equal to count of consequtive true bools counting from right
//
// true false true true: 2,
// true false: 0,
// true: 1,
// true true true: 3,
func (optM *OptionsMap) Count(slug string) int {
	return (*optM)[slug].contents().(optBoolVal).count
}

// AlsoBool
func (optM *OptionsMap) IsBool(slug string) bool {
	return false // TODO:
	// return (*optM)[slug].contents().(optBoolVal).count
}

//// generatable ////

// string

func (optM *OptionsMap) SlString(slug string) []string {
	return (*optM)[slug].contents().([]string)
}

func (optM *OptionsMap) String(slug string) string {
	sl := optM.SlString(slug)
	return sl[len(sl)-1] // last defined
}

// int

func (optM *OptionsMap) SlInt(slug string) []int {
	return (*optM)[slug].contents().([]int)
}

func (optM *OptionsMap) Int(slug string) int {
	sl := optM.SlInt(slug)
	return sl[len(sl)-1] // last defined
}

// int64

func (optM *OptionsMap) SlInt64(slug string) []int64 {
	return (*optM)[slug].contents().([]int64)
}

func (optM *OptionsMap) Int64(slug string) int64 {
	sl := optM.SlInt64(slug)
	return sl[len(sl)-1] // last defined
}

// uint

func (optM *OptionsMap) SlUint(slug string) []uint {
	return (*optM)[slug].contents().([]uint)
}

func (optM *OptionsMap) Uint(slug string) uint {
	sl := optM.SlUint(slug)
	return sl[len(sl)-1] // last defined
}

// uint64

func (optM *OptionsMap) SlUint64(slug string) []uint64 {
	return (*optM)[slug].contents().([]uint64)
}

func (optM *OptionsMap) Uint64(slug string) uint64 {
	sl := optM.SlUint64(slug)
	return sl[len(sl)-1] // last defined
}

// float64

func (optM *OptionsMap) SlFloat64(slug string) []float64 {
	return (*optM)[slug].contents().([]float64)
}

func (optM *OptionsMap) Float64(slug string) float64 {
	sl := optM.SlFloat64(slug)
	return sl[len(sl)-1] // last defined
}

// duration

func (optM *OptionsMap) SlDuration(slug string) []float64 {
	return (*optM)[slug].contents().([]float64)
}

func (optM *OptionsMap) Duration(slug string) float64 {
	sl := optM.SlFloat64(slug)
	return sl[len(sl)-1] // last defined
}
