package harg

import "strings"

// bool carries no meaning
func SliceToLowercaseMap(s []string) (m map[string]bool) {
	for _, str := range s {
		m[strings.ToLower(str)] = false
	}
	return m
}
