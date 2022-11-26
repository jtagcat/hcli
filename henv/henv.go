//go:build dev

package henv

import (
	"errors"
	"os"
	"strconv"
)

// TODO: communicate to def.parse* errors that it is not short/long option, but environment

// Parses Definitions against the Environment.
func (defs *Definitions) Parse() error {
	return errors.New("not implemented")

	if err := defs.normalize(); err != nil {
		return err
	}

	for key, def := range *defs {
		val, ok := os.LookupEnv(key)
		if !ok {
			continue
		}

		if def.AlsoBool {
			boolVal, err := strconv.ParseBool(val)
			if err == nil {
				def.parseBoolValue(key, boolVal)
			}
		}

		if err := def.parseOptionValue(key, val); err != nil {
			return err
		}
	}

	return nil
}
