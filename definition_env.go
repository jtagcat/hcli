package harg

import (
	"fmt"
	"strings"
	"unicode"
)

func (defs Definitions) normalizeEnv() error {
	return defs.genericNormalize(func(key string, def *Definition) (string, error) {
		key = strings.ReplaceAll(key, " ", "_")

		for _, r := range key {
			if r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r) {
				continue
			}

			return key, fmt.Errorf("must contain only underscores, letters and/or digits")
		}

		// capitalize all keys
		return strings.ToUpper(key), nil
	})
}
