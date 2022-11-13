package harg_test

import (
	"testing"

	"github.com/jtagcat/harg"
	"github.com/stretchr/testify/assert"
)

// see FORMAT.md for what test is
// responsible for what part of the spec

func TestParseNilDefs(t *testing.T) {
	t.Parallel()

	defs := harg.Definitions{}

	args, chokeReturn, err := defs.Parse(
		[]string{
			"hello", "-", "world",
			"cHOKe", "return",
		},
		[]string{"choke"},
	)

	assert.Nil(t, err)
	assert.Equal(t, []string{"hello", "-", "world"}, args)
	assert.Equal(t, []string{"cHOKe", "return"}, chokeReturn)
}

func TestParseDoubledash(t *testing.T) {
	t.Parallel()

	defs := harg.Definitions{}

	args, chokeReturn, err := defs.Parse(
		[]string{
			"hello", "world",
			"--",
			"choke",
			"--argument",
			"-a",
		},
		[]string{"choke"},
	)

	assert.Nil(t, err)
	assert.Nil(t, chokeReturn)
	assert.Equal(t, []string{"hello", "world", "choke", "--argument", "-a"}, args)
}

func TestAliasParse(t *testing.T) {
	t.Parallel()

	oneKey := "one"

	defs := harg.Definitions{
		oneKey: {Type: harg.String},
	}
	assert.Nil(t,
		defs.Alias("twõ", oneKey))

	args, chokeReturn, err := defs.Parse(
		[]string{
			"hello",
			"--one=one",
			"--twõ", "two",
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

func TestParseLongOptEat(t *testing.T) {
	t.Parallel()

	oneKey := "one"
	twoKey := "t"

	defs := harg.Definitions{
		oneKey: {Type: harg.String},
		twoKey: {},
	}

	args, chokeReturn, err := defs.Parse(
		[]string{
			"hello",
			"--oNE", "-t",
			"--one",
			"world",
		}, []string{"world"},
	)

	assert.Nil(t, err)
	assert.Nil(t, chokeReturn)
	assert.Equal(t, []string{"hello"}, args)

	sl, ok := defs[oneKey].SlString()
	assert.Equal(t, true, ok)
	assert.Equal(t, []string{"-t", "world"}, sl)

	assert.Equal(t, false, defs[twoKey].Touched())
}

func TestParseShortOptEat(t *testing.T) {
	t.Parallel()

	// - Chokes are not detected as part of options (no choking:`--foo choke` foo:`choke`; no choking:`-o choke -b` o:`choke` b:`true`)
	// - Short options are 1 utf8 character, case sensitive. [^TestParseShortOptEat]

	t.Fatal("not implemented")
}

func TestParseLongOptNotSingleChar(t *testing.T) {
	t.Parallel()
	// "--a" doesn't go to long
	t.Fatal("not implemented")
}

func TestParseShortOptClustering(t *testing.T) {
	t.Parallel()

	// - Short options can be clustered after the prefix. (`-abc` a:`true` b:`true` c:`true`) [^TestParseShortOptClustering]
	// - Preceeding `-` negates the following bool, otherwise ignored. (`--a` a:`false`; `-a-bc` a:`true` b:`false` c:`true`) [^TestParseShortOptClustering]
	//     - If `-` is used for the first short option, short options can't be clustered. (invalid:`--ab`; invalid:`--a-b` (seen as long options)) [^TestParseShortOptClustering]
	// - Non-bools take arguments until space or from the next argument. (`-aovalue`, `-ao value` a:`false` o:`value`) [^TestParseShortOptClustering]

	t.Fatal("not implemented")
}

func TestParseLongOptAlsoBool(t *testing.T) {
	t.Parallel()

	// - AlsoBool treats a valueless valueful option as a bool. (`--foo`; `--foo=value`) [^TestParseLongOptAlsoBool]
	//     - Space-seperated syntax is unavailable. (invalid:`--foo value`) [^TestParseLongOptAlsoBool]
	//     - Bools in values are parsed as booleans. (`--foo=true` is bool, not string "true") [^TestParseLongOptAlsoBool]
	//     - Given multiple mixed bool/value same-slug options, bools before values are ignored, and bools after value error. [^TestParseLongOptAlsoBool]

	t.Fatal("not implemented")
}

func TestParseCount(t *testing.T) {
	t.Parallel()

	// - `Count()`: Equal to the count of consecutive true values read from right/last [^TestParseCount]

	t.Fatal("not implemented")
}
