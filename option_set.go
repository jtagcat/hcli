package harg

import (
	"fmt"
	"strconv"
	"time"

	internal "github.com/jtagcat/harg/internal"
)

func (def *Definition) parseOptionContent(
	key string,
	value string, // "" means literally empty, caller has already defaulted valueless booleans to true
) error { // errContext provided
	// defs.normalize(): actual Type == Bool can never be AlsoBool

	var notAlsoBool bool
	if def.AlsoBool &&
		// try parsing as AlsoBool only when not already parsed as non-bool
		(def.parsed.opt == nil || def.Type == Bool) {

		if def.parsed.opt == nil {
			def.parsed.opt = typeMetaM[Bool].emptyT
		}

		if err := def.parsed.opt.add(value); err == nil {
			def.Type = Bool
			return nil
		}

		notAlsoBool = true
		// on err continue to parse normally:
	}

	// initialize option interface
	if def.parsed.opt == nil || notAlsoBool {
		def.parsed.opt = typeMetaM[def.Type].emptyT
	}

	if err := def.parsed.opt.add(value); err != nil {

		err = fmt.Errorf("%e: %w", ErrIncompatibleValue, err) // add ErrIncompatibleValue, as it is universally comparable
		return fmt.Errorf("parsing %s as %s: %w", internal.KeyErrorName(key), typeMetaM[def.Type].errName, err)
	}

	return nil
}

type option interface {
	contents() any           // resolved with option.Sl
	add(rawOpt string) error // string: type name (to use in error)
}

type Type uint32 // enum:
const (
	Bool Type = iota
	String
	Int
	Int64
	Uint
	Uint64
	Float64
	Duration
) //
var typeMax = Duration

var typeMetaM = map[Type]struct {
	errName string
	emptyT  option
}{
	Bool:     {"bool", &optBool{}},
	String:   {"string", &optString{}},
	Int:      {"int", &optInt{}},
	Int64:    {"int64", &optInt64{}},
	Uint:     {"uint", &optUint{}},
	Uint64:   {"uint64", &optUint64{}},
	Float64:  {"float64", &optFloat64{}},
	Duration: {"duration", &optDuration{}},
}

// bool / count

type (
	optBool struct {
		value optBoolVal
	}
	optBoolVal struct {
		count int
		value []bool
	}
)

func (o *optBool) contents() any {
	return o.value
}

func (o *optBool) add(s string) error {
	v, err := strconv.ParseBool(s) // TODO: drop "t", "f", add "yes", "no", maybe also "y", "n"
	if err != nil {
		return err
	}

	if v == true {
		o.value.count++
	} else {
		o.value.count = 0
	}

	o.value.value = append(o.value.value, v)
	return nil
}

// string

type optString struct {
	value []string
}

func (o *optString) add(s string) error {
	o.value = append(o.value, s)
	return nil
}

// int

type optInt struct {
	value []int
}

func (o *optInt) add(s string) error {
	v, err := strconv.ParseInt(s, 0, strconv.IntSize)
	if err != nil {
		return err
	}

	o.value = append(o.value, int(v))
	return err
}

// int64

type optInt64 struct {
	value []int64
}

func (o *optInt64) add(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return err
	}

	o.value = append(o.value, int64(v))
	return err
}

// uint

type optUint struct {
	value []uint
}

func (o *optUint) add(s string) error {
	v, err := strconv.ParseInt(s, 0, strconv.IntSize)
	if err != nil {
		return err
	}

	o.value = append(o.value, uint(v))
	return err
}

// uint64

type optUint64 struct {
	value []uint64
}

func (o *optUint64) add(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return err
	}

	o.value = append(o.value, uint64(v))
	return err
}

// float64

type optFloat64 struct {
	value []float64
}

func (o *optFloat64) add(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}

	o.value = append(o.value, float64(v))
	return err
}

// duration

type optDuration struct {
	value []time.Duration
}

func (o *optDuration) add(s string) error {
	v, err := time.ParseDuration(s)
	if err != nil {
		return err
	}

	o.value = append(o.value, time.Duration(v))
	return err
}

// TODO: more Types
// add to: Types enum; option_set (2); option_get (3)
//
// timestamp
// ip
// ipv4
// ipv6
// ...?
