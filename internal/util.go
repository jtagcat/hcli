package internal

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

func SliceLowercaseIndex(s []string) map[string]struct{} {
	index := make(map[string]struct{})

	for _, str := range s {
		index[strings.ToLower(str)] = struct{}{}
	}

	return index
}

func LowercaseLongKey(key string) string {
	if utf8.RuneCountInString(key) == 1 {
		return key
	}

	return strings.ToLower(key)
}

func OptErrorName(key string) string {
	var keyType string
	if utf8.RuneCountInString(key) > 1 {
		keyType = "long"
	} else {
		keyType = "short"
	}

	return fmt.Sprintf("%s option %q", keyType, key)
}

type GenericErr struct {
	Err     error
	Wrapped error
}

func (a GenericErr) Is(target error) bool {
	return errors.Is(a.Err, target)
}

func (a GenericErr) Unwrap() error {
	return a.Wrapped
}

func (a GenericErr) Error() string {
	return a.Err.Error() + ": " + a.Wrapped.Error()
}
