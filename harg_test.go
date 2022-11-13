package harg_test

import (
	"testing"

	"github.com/jtagcat/harg"
	"github.com/stretchr/testify/assert"
)

func TestNildefs(t *testing.T) {
	defs := harg.Definitions{}

	args, chokeReturn, err := defs.Parse(
		[]string{
			"hello", "world",
			"choke", "return",
		},
		[]string{"choke"},
	)

	assert.Nil(t, err)
	assert.Equal(t, []string{"hello", "world"}, args)
	assert.Equal(t, []string{"choke", "return"}, chokeReturn)
}
