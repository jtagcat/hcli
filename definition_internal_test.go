package harg

import (
	"errors"
	"testing"
)

func TestDefinitionNormalizeNilMap(t *testing.T) {
	t.Parallel()

	defs := Definitions{}

	if err := defs.normalize(); err != nil {
		t.Error(err)
	}

	if _, ok := defs["test"]; ok {
		t.Error("Definitions should be empty, but still usable")
	}
}

func TestDefinitionNormalize(t *testing.T) {
	t.Parallel()

	defs := Definitions{
		"nil": nil,
		"Uppercase": &Definition{
			Type:     String,
			AlsoBool: true,
		},
		"lowercase": &Definition{
			Type:     Bool,
			AlsoBool: true,
		},
		"S": &Definition{
			Type:     String,
			AlsoBool: true,
		},
		"s": &Definition{
			Type:     Bool,
			AlsoBool: true,
		},
	}

	if err := defs.normalize(); err != nil {
		t.Fatal(err)
	}

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

	if def := defs["uppercase"]; def.AlsoBool == false {
		t.Errorf("uppercase AlsoBool should be false")
	}
	for _, name := range []string{"lowercase", "S", "s"} {
		if def := defs[name]; def.AlsoBool == true {
			t.Errorf("%s AlsoBool should be false", name)
		}
	}
}

func TestDefinitionOverMax(t *testing.T) {
	t.Parallel()

	defs := Definitions{
		"Bad": &Definition{
			Type: Type(int(typeMax) + 1),
		},
	}

	err := defs.normalize()
	if !errors.Is(err, ErrInvalidDefinition) {
		t.Fatalf("error not %e, is %e", ErrInvalidDefinition, err)
	}
}

func TestTypeMetaMLen(t *testing.T) {
	t.Parallel()

	if len(typeMetaM) != int(typeMax)+1 {
		t.Fatalf("expected typeMetaM (%d) to be equal to Type(Max) (%d)", len(typeMetaM), int(typeMax)+1)
	}
}

func TestDefinitionDigits(t *testing.T) {
	t.Parallel()

	defs := Definitions{
		"1": &Definition{
			Type: Bool,
		},
	}

	err := defs.normalize()
	if !errors.Is(err, ErrInvalidDefinition) {
		t.Fatalf("error not %e, is %e", ErrInvalidDefinition, err)
	}
}
