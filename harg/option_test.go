package harg_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/jtagcat/hcli/harg"
	"github.com/stretchr/testify/require"
)

func TestOptBool(t *testing.T) {
	key := "k"
	defs := harg.Definitions{
		key: {Type: harg.Bool},
	}

	want := []bool{true, true}

	parsed, chokeReturn, err := defs.Parse([]string{
		"-k",
		"-k",
	}, nil)
	require.Nil(t, err)
	require.Zero(t, len(parsed))
	require.Zero(t, len(chokeReturn))

	require.False(t, false, defs[key].Default())
	require.True(t, defs[key].IsBool())

	v, ok := defs[key].SlBool()
	require.True(t, ok)
	require.Equal(t, len(want), len(v))
	require.Equal(t, want[0], v[0])
	require.Equal(t, want[1], v[1])
	a, ok := defs[key].SlAny()
	require.True(t, ok)
	require.Equal(t, v, a)

	sv, ok := defs[key].Bool()
	require.True(t, ok)
	require.Equal(t, want[1], sv)
	av, ok := defs[key].Any()
	require.True(t, ok)
	require.Equal(t, sv, av)

	c, ok := defs[key].Count()
	require.True(t, ok)
	require.Equal(t, len(want), c)
}

func TestOptString(t *testing.T) {
	key := "k"
	defs := harg.Definitions{
		key: {Type: harg.String},
	}

	want := []string{"hello", "world"}

	parsed, chokeReturn, err := defs.Parse([]string{
		"-k", want[0],
		"-k", want[1],
	}, nil)
	require.Nil(t, err)
	require.Zero(t, len(parsed))
	require.Zero(t, len(chokeReturn))

	require.False(t, false, defs[key].Default())
	require.False(t, defs[key].IsBool())

	v, ok := defs[key].SlString()
	require.True(t, ok)
	require.Equal(t, len(want), len(v))
	require.Equal(t, want[0], v[0])
	require.Equal(t, want[1], v[1])
	a, ok := defs[key].SlAny()
	require.True(t, ok)
	require.Equal(t, v, a)

	sv, ok := defs[key].String()
	require.True(t, ok)
	require.Equal(t, want[1], sv)
	av, ok := defs[key].Any()
	require.True(t, ok)
	require.Equal(t, sv, av)
}

func TestOptInt(t *testing.T) {
	key := "k"
	defs := harg.Definitions{
		key: {Type: harg.Int},
	}

	want := []int{-1, 1}

	parsed, chokeReturn, err := defs.Parse([]string{
		"-k", strconv.Itoa(want[0]),
		"-k", strconv.Itoa(want[1]),
	}, nil)
	require.Nil(t, err)
	require.Zero(t, len(parsed))
	require.Zero(t, len(chokeReturn))

	require.False(t, false, defs[key].Default())
	require.False(t, defs[key].IsBool())

	v, ok := defs[key].SlInt()
	require.True(t, ok)
	require.Equal(t, len(want), len(v))
	require.Equal(t, want[0], v[0])
	require.Equal(t, want[1], v[1])
	a, ok := defs[key].SlAny()
	require.True(t, ok)
	require.Equal(t, v, a)

	sv, ok := defs[key].Int()
	require.True(t, ok)
	require.Equal(t, want[1], sv)
	av, ok := defs[key].Any()
	require.True(t, ok)
	require.Equal(t, sv, av)
}

func TestOptInt64(t *testing.T) {
	key := "k"
	defs := harg.Definitions{
		key: {Type: harg.Int64},
	}

	want := []int64{-1, 1}

	parsed, chokeReturn, err := defs.Parse([]string{
		"-k", strconv.Itoa(int(want[0])),
		"-k", strconv.Itoa(int(want[1])),
	}, nil)
	require.Nil(t, err)
	require.Zero(t, len(parsed))
	require.Zero(t, len(chokeReturn))

	require.False(t, false, defs[key].Default())
	require.False(t, defs[key].IsBool())

	v, ok := defs[key].SlInt64()
	require.True(t, ok)
	require.Equal(t, len(want), len(v))
	require.Equal(t, want[0], v[0])
	require.Equal(t, want[1], v[1])
	a, ok := defs[key].SlAny()
	require.True(t, ok)
	require.Equal(t, v, a)

	sv, ok := defs[key].Int64()
	require.True(t, ok)
	require.Equal(t, want[1], sv)
	av, ok := defs[key].Any()
	require.True(t, ok)
	require.Equal(t, sv, av)
}

func TestOptUint(t *testing.T) {
	key := "k"
	defs := harg.Definitions{
		key: {Type: harg.Uint},
	}

	want := []uint{0, 1}

	parsed, chokeReturn, err := defs.Parse([]string{
		"-k", strconv.Itoa(int(want[0])),
		"-k", strconv.Itoa(int(want[1])),
	}, nil)
	require.Nil(t, err)
	require.Zero(t, len(parsed))
	require.Zero(t, len(chokeReturn))

	require.False(t, false, defs[key].Default())
	require.False(t, defs[key].IsBool())

	v, ok := defs[key].SlUint()
	require.True(t, ok)
	require.Equal(t, len(want), len(v))
	require.Equal(t, want[0], v[0])
	require.Equal(t, want[1], v[1])
	a, ok := defs[key].SlAny()
	require.True(t, ok)
	require.Equal(t, v, a)

	sv, ok := defs[key].Uint()
	require.True(t, ok)
	require.Equal(t, want[1], sv)
	av, ok := defs[key].Any()
	require.True(t, ok)
	require.Equal(t, sv, av)
}

func TestOptUint64(t *testing.T) {
	key := "k"
	defs := harg.Definitions{
		key: {Type: harg.Uint64},
	}

	want := []uint64{0, 1}

	parsed, chokeReturn, err := defs.Parse([]string{
		"-k", strconv.Itoa(int(want[0])),
		"-k", strconv.Itoa(int(want[1])),
	}, nil)
	require.Nil(t, err)
	require.Zero(t, len(parsed))
	require.Zero(t, len(chokeReturn))

	require.False(t, false, defs[key].Default())
	require.False(t, defs[key].IsBool())

	v, ok := defs[key].SlUint64()
	require.True(t, ok)
	require.Equal(t, len(want), len(v))
	require.Equal(t, want[0], v[0])
	require.Equal(t, want[1], v[1])
	a, ok := defs[key].SlAny()
	require.True(t, ok)
	require.Equal(t, v, a)

	sv, ok := defs[key].Uint64()
	require.True(t, ok)
	require.Equal(t, want[1], sv)
	av, ok := defs[key].Any()
	require.True(t, ok)
	require.Equal(t, sv, av)
}

func TestOptFloat64(t *testing.T) {
	key := "k"
	defs := harg.Definitions{
		key: {Type: harg.Float64},
	}

	want := []float64{-0.5, 0.5}

	parsed, chokeReturn, err := defs.Parse([]string{
		"-k", strconv.FormatFloat(want[0], 'f', -1, 64),
		"-k", strconv.FormatFloat(want[1], 'f', -1, 64),
	}, nil)
	require.Nil(t, err)
	require.Zero(t, len(parsed))
	require.Zero(t, len(chokeReturn))

	require.False(t, false, defs[key].Default())
	require.False(t, defs[key].IsBool())

	v, ok := defs[key].SlFloat64()
	require.True(t, ok)
	require.Equal(t, len(want), len(v))
	require.Equal(t, want[0], v[0])
	require.Equal(t, want[1], v[1])
	a, ok := defs[key].SlAny()
	require.True(t, ok)
	require.Equal(t, v, a)

	sv, ok := defs[key].Float64()
	require.True(t, ok)
	require.Equal(t, want[1], sv)
	av, ok := defs[key].Any()
	require.True(t, ok)
	require.Equal(t, sv, av)
}

func TestOptDuration(t *testing.T) {
	key := "k"
	defs := harg.Definitions{
		key: {Type: harg.Duration},
	}

	want := []time.Duration{3600000000000, 15000000000}

	parsed, chokeReturn, err := defs.Parse([]string{
		"-k", want[0].String(),
		"-k", want[1].String(),
	}, nil)
	require.Nil(t, err)
	require.Zero(t, len(parsed))
	require.Zero(t, len(chokeReturn))

	require.False(t, false, defs[key].Default())
	require.False(t, defs[key].IsBool())

	v, ok := defs[key].SlDuration()
	require.True(t, ok)
	require.Equal(t, len(want), len(v))
	require.Equal(t, want[0], v[0])
	require.Equal(t, want[1], v[1])
	a, ok := defs[key].SlAny()
	require.True(t, ok)
	require.Equal(t, v, a)

	sv, ok := defs[key].Duration()
	require.True(t, ok)
	require.Equal(t, want[1], sv)
	av, ok := defs[key].Any()
	require.True(t, ok)
	require.Equal(t, sv, av)
}
