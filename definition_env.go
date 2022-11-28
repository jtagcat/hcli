package harg

import (
	"strings"
)

// adds a prefix to all Definitions
func (defs Definitions) AddPrefix(prefix string) Definitions {
	new := make(Definitions)

	for name, def := range defs {
		new[prefix+name] = def
	}

	return new
}

func (defs Definitions) normalizeEnv() error {
	return defs.genericNormalize(func(key string, def *Definition) string {
		key = strings.ReplaceAll(key, " ", "_")

		// capitalize all keys
		return strings.ToUpper(key)
	})
}
