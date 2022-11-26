package harg

import (
	"errors"
	"fmt"

	internal "github.com/jtagcat/harg/internal"
)

func (def *Definition) parseOptionValue(key, value string) error { // errContext provided
	// restore
	if def.AlsoBool && def.originalType != Bool {
		def.parsed, def.Type = nil, def.originalType
	}

	// initialize option interface
	if def.parsed == nil {
		def.parsed = typeMetaM[def.Type].new()
	}

	if err := def.parsed.add(value); err != nil {
		return fmt.Errorf("parsing %s as %s: %w", internal.KeyErrorName(key), typeMetaM[def.Type].errName, internal.GenericErr{
			Err:     ErrIncompatibleValue,
			Wrapped: err,
		})
	}

	return nil
}

func (def *Definition) parseBoolValue(key string, val bool) error {
	// defs.normalize(): actual Type == Bool can never be AlsoBool

	if def.parsed == nil {
		def.parsed = typeMetaM[Bool].new()

		if def.AlsoBool {
			def.originalType = def.Type
			def.Type = Bool
		}
	}

	if def.Type != Bool {
		return fmt.Errorf("parsing %s as %s: %w", internal.KeyErrorName(key), typeMetaM[def.Type].errName, internal.GenericErr{
			Err:     ErrIncompatibleValue,
			Wrapped: errors.New("AlsoBool must not have a Bool value after non-Bool value"),
		})
	}

	def.parsed.(*optBool).addT(val)
	return nil
}

func (o *optBool) addT(v bool) {
	o.value = append(o.value, v)
}
