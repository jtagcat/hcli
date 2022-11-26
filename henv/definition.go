//go:build dev

package henv

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/jtagcat/harg"
	"github.com/jtagcat/harg/internal"
)

type Definitions harg.Definitions

// adds a prefix to all Definitions
func (defs Definitions) AddPrefix(prefix string) Definitions {
	new := make(Definitions)

	for name, def := range defs {
		new[prefix+name] = def
	}

	return new
}

func (defs Definitions) normalize() error {
	for name, def := range defs {
		if def == nil || name == "" {

			delete(defs, name)
			continue
		}

		if def.Type > typeMax {
			return fmt.Errorf("%s: %w", internal.KeyErrorName(name), internal.GenericErr{
				Err: harg.ErrInvalidDefinition, Wrapped: errors.New("Type does not exist"),
			})
		}

		if unicode.IsDigit(rune(name[0])) {
			return fmt.Errorf("%s: %w", internal.KeyErrorName(name), internal.GenericErr{
				Err: harg.ErrInvalidDefinition, Wrapped: errors.New("Definition name can't start with a digit"),
			})
		}

		if def.Type == harg.Bool && def.AlsoBool {
			def.AlsoBool = false
		}

		// capitalize all keys
		upper := strings.ToUpper(name)
		if name != upper {
			defs[upper] = def
			delete(defs, name)
		}
	}

	return nil
}
