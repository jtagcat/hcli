package harg

import (
	"testing"
)

func TestDefinitionNormalizeNilMap(t *testing.T) {
	t.Parallel()

	defs := Definitions{}

	defs.normalize()

	if _, ok := defs["test"]; ok {
		t.Error("Definitions should be empty, but still usable")
	}
}

func TestDefinitionNormalize(t *testing.T) {
	t.Parallel()

	defs := Definitions{
		"nil":       nil,
		"Uppercase": &Definition{},
		"lowercase": &Definition{},
		"S":         &Definition{},
		"s":         &Definition{},
	}
	defs.normalize()

	for _, name := range []string{"test", "nil", "Uppercase"} {
		if _, ok := defs[name]; ok {
			t.Errorf("%s should not be in Definitions", name)
		}
	}

	for _, name := range []string{"uppercase", "lowercase", "S", "s"} {
		if _, ok := defs[name]; !ok {
			t.Errorf("%s should be in Definitions", name)
		}
	}
}
