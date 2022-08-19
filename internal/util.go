package harg

import (
	"strings"
	"unicode/utf8"
)

// bool carries no meaning
func SliceToLowercaseMap(s []string) (m map[string]bool) {
	for _, str := range s {
		m[strings.ToLower(str)] = false
	}
	return m
}

// TODO: test
func LowercaseLongMapNames[T any](m map[string]T) map[string]T {
	for name, data := range m {
		// short args are case sensitive, skip
		if utf8.RuneCountInString(name) == 1 {
			continue
		}

		// case insensitivize long args
		lower := strings.ToLower(name)
		if name != lower {
			m[lower] = data
			delete(m, name)
		}
	}
	return m
}
