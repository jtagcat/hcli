package harg_test

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/jtagcat/hcli/harg"
	"github.com/stretchr/testify/assert"
)

// see FORMAT.md for what test is
// responsible for what part of the spec

func ExampleDefinitions_Parse() {
	kOne, kTwo, kThree := "o", "t", "three"
	defs := harg.Definitions{
		kOne:   {Type: harg.String},
		kTwo:   {Type: harg.Bool},
		kThree: {Type: harg.Duration},
	}

	osArgs := strings.Split("programName hello -to foo -o bar --three 5s --t -t -t world", " ")

	args, _, err := defs.Parse(osArgs[1:], nil)
	if err != nil {
		panic(fmt.Sprintf("parsing command-line arguments: %e", err))
	}

	fmt.Println(args) // [hello world]

	sl, ok := defs[kOne].SlString() // ok: is valid and set
	if ok {
		fmt.Println(sl) // [foo bar]
	}
	s, ok := defs[kOne].String()
	if ok {
		fmt.Println(s) // bar
	}

	two, _ := defs[kTwo].SlBool()
	fmt.Println(two) // [true false true true]

	count, _ := defs[kTwo].Count() // how many true in a row
	fmt.Println(count)             // 2

	dur, _ := defs[kThree].Duration()
	fmt.Println(dur) // 5s

	// Output:
	// [hello world]
	// [foo bar]
	// bar
	// [true false true true]
	// 2
	// 5s
}

func ExampleDefinitions_ParseEnv() {
	kOne, kTwo := "ONE", "two" // will be uppercased and joined with underscore
	_, _ = os.Setenv(kOne, "5s"), os.Setenv(kTwo, "hello,world")

	defs := harg.Definitions{
		kOne: {Type: harg.Duration},
		kTwo: {Type: harg.String, EnvCSV: true},
	}

	if err := defs.ParseEnv(); err != nil {
		panic(fmt.Sprintf("parsing environment: %e", err))
	}

	dur, _ := defs[kOne].Duration()
	fmt.Println(dur) // 5s

	str, _ := defs[kTwo].SlString()
	fmt.Println(str) // [hello world]

	// Output:
	// 5s
	// [hello world]
}

func TestParseNilDefs(t *testing.T) {
	t.Parallel()

	defs := harg.Definitions{}

	args, chokeReturn, err := defs.Parse([]string{
		"hello", "-", "world",
		"cHOKe", "return",
	},
		[]string{"choke"},
	)

	assert.Nil(t, err)
	assert.Equal(t, []string{"hello", "-", "world"}, args)
	assert.Equal(t, []string{"cHOKe", "return"}, chokeReturn)

	args, chokeReturn, err = defs.Parse(nil, nil)
	assert.Nil(t, err)
	assert.Nil(t, chokeReturn)
	assert.Nil(t, args)
}

func TestParseDoubledash(t *testing.T) {
	t.Parallel()

	defs := harg.Definitions{}

	args, chokeReturn, err := defs.Parse([]string{
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

	kOne := "one"

	defs := harg.Definitions{
		kOne: {Type: harg.String},
	}
	assert.Nil(t, defs.Alias("twõか", kOne))

	args, chokeReturn, err := defs.Parse([]string{
		"hello",
		"--one=one",
		"--twõか", "two",
		"world",
	}, nil,
	)

	assert.Nil(t, err)
	assert.Nil(t, chokeReturn)
	assert.Equal(t, []string{"hello", "world"}, args)

	sl, ok := defs[kOne].SlString()
	assert.Equal(t, true, ok)
	assert.Equal(t, []string{"one", "two"}, sl)

	s, ok := defs[kOne].String()
	assert.Equal(t, true, ok)
	assert.Equal(t, "two", s)
}

func TestParseLongOptEat(t *testing.T) {
	t.Parallel()

	kOne, kTwo, kFoo := "oかe", "t", "f"

	defs := harg.Definitions{
		kOne: {Type: harg.String},
		kTwo: {},
		kFoo: {},
	}

	args, chokeReturn, err := defs.Parse([]string{
		"hello",
		"--OかE=-t",
		"--oかE", "-f",
		"--oかe",
		"world",
	}, []string{"world"},
	)

	assert.Nil(t, err)
	assert.Nil(t, chokeReturn)
	assert.Equal(t, []string{"hello"}, args)

	sl, ok := defs[kOne].SlString()
	assert.Equal(t, true, ok)
	assert.Equal(t, []string{"-t", "", "world"}, sl)

	assert.Equal(t, true, defs[kTwo].Default())
	assert.Equal(t, false, defs[kFoo].Default())
}

func TestParseShortOptEat(t *testing.T) {
	t.Parallel()

	kOne, kTwo, kFoo := "か", "t", "f"

	defs := harg.Definitions{
		kOne: {Type: harg.String},
		kTwo: {},
		kFoo: {},
	}

	args, chokeReturn, err := defs.Parse([]string{
		"hello",
		"-かt",
		"-か=-t",
		"-か", "=-t",
		"-か", "-f",
		"-か",
		"world",
	}, []string{"world"},
	)

	assert.Nil(t, err)
	assert.Nil(t, chokeReturn)
	assert.Equal(t, []string{"hello"}, args)

	sl, ok := defs[kOne].SlString()
	assert.Equal(t, true, ok)
	assert.Equal(t, []string{"t", "-t", "=-t", "", "world"}, sl)

	assert.Equal(t, true, defs[kTwo].Default())
	assert.Equal(t, false, defs[kFoo].Default())
}

func TestParseShortBoolOpt(t *testing.T) {
	t.Parallel()

	kZero, kOne, kTwo := "か", "õ", "x"
	kUnset := "u"

	defs := harg.Definitions{
		kZero:  {},
		kOne:   {},
		kTwo:   {},
		kUnset: {},
	}
	assert.Nil(t, defs.Alias("õx", kZero))

	for in, want := range map[string][]bool{
		"-か":      {true, false, false},
		"-か\n--か": {false, false, false},
		"-かõ-x":   {true, true, false},
		"-か-õx":   {true, false, true},
		"-か-õ-x":  {true, false, false},
		"--õx":    {true, false, false},
		"---õx":   {false, false, false},
	} {
		defs := defs

		args, chokeReturn, err := defs.Parse(
			strings.Split(in, "\n"), nil,
		)

		assert.Nil(t, err)
		assert.Nil(t, chokeReturn)
		assert.Nil(t, args)

		set := defs[kUnset].Default()
		assert.Equal(t, true, set)

		b, ok := defs[kZero].Bool()
		assert.Equal(t, true, ok)
		assert.Equal(t, want[0], b)

		b, _ = defs[kOne].Bool()
		assert.Equal(t, want[1], b)

		b, _ = defs[kTwo].Bool()
		assert.Equal(t, want[2], b)
	}
}

func TestParseCount(t *testing.T) {
	t.Parallel()

	// also responsible for testing if typeMap.new() actually copies or no

	kZero, kOne := "a", "b"
	defs := harg.Definitions{
		kZero: {},
		kOne:  {},
	}

	args, chokeReturn, err := defs.Parse([]string{
		"-a-aaaa-a",
		"--b", "-b-b-bbb",
	}, nil,
	)

	assert.Nil(t, err)
	assert.Nil(t, chokeReturn)
	assert.Nil(t, args)

	sl, ok := defs[kZero].SlBool()
	assert.Equal(t, true, ok)
	assert.Equal(t, []bool{true, false, true, true, true, false}, sl)
	c, ok := defs[kZero].Count()
	assert.Equal(t, true, ok)
	assert.Equal(t, 0, c)

	sl, ok = defs[kOne].SlBool()
	assert.Equal(t, true, ok)
	assert.Equal(t, []bool{false, true, false, false, true, true}, sl)
	c, ok = defs[kOne].Count()
	assert.Equal(t, true, ok)
	assert.Equal(t, 2, c)
}

func TestParseLongOptAlsoBool(t *testing.T) {
	t.Parallel()

	kOne, kTwo := "foo", "bar"

	defs := harg.Definitions{
		kOne: {Type: harg.String, AlsoBool: true},
		kTwo: {Type: harg.String, AlsoBool: true},
	}

	args, chokeReturn, err := defs.Parse([]string{
		"---foo", "bar", // false
		"--foo", "bar", // true
		"--bar=true", // "true", not true
	}, nil,
	)

	assert.Nil(t, err)
	assert.Nil(t, chokeReturn)
	assert.Equal(t, []string{"bar", "bar"}, args)

	sl, ok := defs[kOne].SlBool()
	assert.Equal(t, true, ok)
	assert.Equal(t, []bool{false, true}, sl)

	s, ok := defs[kTwo].String()
	assert.Equal(t, true, ok)
	assert.Equal(t, "true", s)
}

func TestParseError(t *testing.T) {
	t.Parallel()

	defs := harg.Definitions{
		"str":      {Type: harg.String},
		"bool":     {},
		"alsobool": {Type: harg.String, AlsoBool: true},
	}

	for _, test := range []errTest{
		// Negating long option
		{in: []string{"---str"}, errIs: harg.ErrIncompatibleValue},       // not bool
		{in: []string{"---bool=true"}, errIs: harg.ErrIncompatibleValue}, // bool with value

		// AlsoBool after Value
		{in: []string{"--alsobool=val", "--alsobool"}, errIs: harg.ErrIncompatibleValue},

		// No definition
		{in: []string{"--nodef"}, errIs: harg.ErrOptionHasNoDefinition},
		{in: []string{"-n"}, errIs: harg.ErrOptionHasNoDefinition},

		// Some errors are tested in definition tests.
	} {
		defs := defs

		args, chokeReturn, err := defs.Parse(
			test.in, nil,
		)

		assert.ErrorIs(t, err, test.errIs)
		assert.Nil(t, chokeReturn)
		assert.Nil(t, args)
	}
}

type errTest struct {
	in    []string
	errIs error
}

func TestGetNormalizedKey(t *testing.T) {
	kOne := "hElLO" // will be lowercased
	defs := harg.Definitions{
		kOne: {},
	}

	args, chokeReturn, err := defs.Parse([]string{
		"--hello", "--HELlO", // any case should work
	}, nil)
	assert.Nil(t, err)
	assert.Nil(t, chokeReturn)
	assert.Nil(t, args)

	c, ok := defs[kOne].Count()
	assert.Equal(t, true, ok)
	assert.Equal(t, 2, c)
}

func TestGetNormalizedEnvKey(t *testing.T) {
	kOne := "hElLO_world" // will be uppercased
	defs := harg.Definitions{
		kOne: {},
	}

	assert.Nil(t, os.Setenv("HELLO_wORlD", "true"))

	assert.Nil(t, defs.ParseEnv())

	b, ok := defs[kOne].Bool()
	assert.Equal(t, true, ok)
	assert.Equal(t, true, b)
}

func TestParseEnv(t *testing.T) {
	kOne, kTwo := "ONE", "two" // will be uppercased and joined with underscore
	defs := harg.Definitions{
		kOne: {Type: harg.Duration},
		kTwo: {EnvCSV: true},
	}

	assert.Nil(t, os.Setenv(kOne, "5s"))
	assert.Nil(t, os.Setenv(kTwo, "true,true"))

	assert.Nil(t, defs.ParseEnv())

	dur, ok := defs[kOne].Duration()
	assert.Equal(t, true, ok)
	assert.Equal(t, time.Duration(5000000000), dur)

	c, ok := defs[kTwo].Count()
	assert.Equal(t, true, ok)
	assert.Equal(t, 2, c)
}
