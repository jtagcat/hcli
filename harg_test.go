package harg_test

import (
	"testing"

	"github.com/jtagcat/harg"
	"github.com/stretchr/testify/assert"
)

func TestNildefs(t *testing.T) {
	t.Parallel()

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

func TestAliasParse(t *testing.T) {
	t.Parallel()

	oneKey := "one"

	defs := harg.Definitions{
		oneKey: {Type: harg.String},
	}
	defs.Alias("two", oneKey)

	args, chokeReturn, err := defs.Parse(
		[]string{
			"hello",
			"--one=one",
			"--two", "two",
			"world",
		}, nil,
	)

	assert.Nil(t, err)
	assert.Nil(t, chokeReturn)
	assert.Equal(t, []string{"hello", "world"}, args)

	sl, ok := defs[oneKey].SlString()
	assert.Equal(t, true, ok)
	assert.Equal(t, []string{"one", "two"}, sl)

	s, ok := defs[oneKey].String()
	assert.Equal(t, true, ok)
	assert.Equal(t, "two", s)
}
