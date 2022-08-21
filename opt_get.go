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
