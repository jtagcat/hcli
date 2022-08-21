package harg

import (
	"strconv"
	"time"
)

func (optM *OptionsMap) parseOptionContent(
	originalKey string, // may be alias name or == effectiveKey
	effectiveKey string, def *Definition,
	value string, // "" means literally empty, caller has already defaulted booleans to true
) error { // errContext provided

	// if def.Type == e_bool || def.AlsoBool {
	// 	// TODO: use loops, bool.add()
	// 	// boolVal, err := strconv.ParseBool(value) // TODO: drop "t", "f", add "yes", "no", maybe also "y", "n"
	// 	// if err == nil {
	// 	// 	typeM := (*optM)[e_bool]

	// 	// 	// TODO:

	// 	// 	if def.Type == e_bool {
	// 	// 	}

	// 	// 	return nil
	// 	// }

	// 	// // err != nil
	// 	// if def.Type == e_bool {
	// 	// 	return fmt.Errorf("parsing option %s with definition %s as type bool: %w", originalKey, effectiveKey, err)
	// 	// }

	// 	// // AlsoBool continues to switch
	// }

	// // valueful
	// typeM := (*optM)[def.Type]
	// opt, ok := typeM[effectiveKey]
	// if !ok {
	// 	opt = typeEmptyM[def.Type]
	// }

	// err := opt.add(value)
	// if err != nil {
	// 	return fmt.Errorf("parsing option %s with definition %s: %e: %w", originalKey, effectiveKey, ErrIncompatibleValue, err)
	// }

	// typeM[effectiveKey] = opt
	// return nil
}

type (
	OptionsMap map[string]option // parallel to Definitions.D
	option     interface {
		typeName() string
		found(write bool) bool
		contents() any           // resolved with option.Sl
		add(rawOpt string) error // string: type name (to use in error)
	}
)

type optCommon struct {
	foundV bool
}

func (o *optCommon) found(write bool) bool {
	if write {
		o.foundV = true
	}
	return o.foundV
}

type Type int // enum

const ( // enum
	e_bool Type = iota
	// doesn't seem the best way, but let's try
	e_string
	e_int
	e_int64
	e_uint
	e_uint64
	e_float64
	e_duration
)

var typeEmptyM = map[Type]option{
	// e_bool is handled in parseOptionContent
	e_string:   &optString{},
	e_int:      &optInt{},
	e_int64:    &optInt64{},
	e_uint:     &optUint{},
	e_uint64:   &optUint64{},
	e_float64:  &optFloat64{},
	e_duration: &optDuration{},
}

// bool / count

type (
	optBool struct {
		optCommon
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
	optCommon
	value []string
}

func (o *optString) contents() any {
	return o.value
}

func (o *optString) add(s string) error {
	o.value = append(o.value, s)
	return nil
}

// int

type optInt struct {
	optCommon
	value []int
}

func (o *optInt) contents() any {
	return o.value
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
	optCommon
	value []int64
}

func (o *optInt64) contents() any {
	return o.value
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
	optCommon
	value []uint
}

func (o *optUint) contents() any {
	return o.value
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
	optCommon
	value []uint64
}

func (o *optUint64) contents() any {
	return o.value
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
	optCommon
	value []float64
}

func (o *optFloat64) contents() any {
	return o.value
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
	optCommon
	value []time.Duration
}

func (o *optDuration) contents() any {
	return o.value
}

func (o *optDuration) add(s string) error {
	v, err := time.ParseDuration(s)
	if err != nil {
		return err
	}

	o.value = append(o.value, time.Duration(v))
	return err
}

// timestamp
// TODO:

// ip
// ipv4
// ipv6
// TODO:
