package harg

import (
	"fmt"
	"os"
	"strconv"
)

// Parses Definitions from the Environment.
//
// All definitions will be transformed to uppercase. Spaces are replaced with underscores.
func (defs *Definitions) ParseEnv() error {
	if err := defs.normalizeEnv(); err != nil {
		return err
	}

	for key, def := range *defs {
		errContext := func() string { return fmt.Sprintf("environment %s", key) }

		val, ok := os.LookupEnv(key)
		if !ok {
			continue
		}

		if def.AlsoBool {
			boolVal, err := strconv.ParseBool(val)
			if err == nil {
				def.parseBoolValue(boolVal, errContext)
			}
		}

		if err := def.parseValue(val, errContext); err != nil {
			return err
		}
	}

	return nil
}
