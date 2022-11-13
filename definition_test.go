package harg_test

import (
	"testing"

	"github.com/jtagcat/harg"
	"github.com/stretchr/testify/assert"
)

func TestSetAlias(t *testing.T) {
	defsOriginal := harg.Definitions{
		"foo": {
			Type: harg.String,
		},
	}
	defs := defsOriginal

	assert.ErrorIs(t,
		defs.SetAlias("alias", "invalid"),
		harg.ErrOptionHasNoDefinition,
	)
	assert.Equal(t, defsOriginal, defs)

	//

	defs = defsOriginal

	assert.Nil(t,
		defs.SetAlias("alias", "foo"))
	assert.Equal(t, harg.String, defs["alias"].Type)
}
