package harg

import (
	"strconv"
	"time"
)

type option interface {
	contents() any           // resolved with option.Sl
	add(rawOpt string) error // string: type name (to use in error)
}

type Type uint8 // enum:
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

type typeMeta struct {
	name string
	new  func() option
} //
var typeMetaM = map[Type]typeMeta{
	Bool:     {"bool", func() option { return &optBool{} }},
	String:   {"string", func() option { return &optString{} }},
	Int:      {"int", func() option { return &optInt{} }},
	Int64:    {"int64", func() option { return &optInt64{} }},
	Uint:     {"uint", func() option { return &optUint{} }},
	Uint64:   {"uint64", func() option { return &optUint64{} }},
	Float64:  {"float64", func() option { return &optFloat64{} }},
	Duration: {"duration", func() option { return &optDuration{} }},
}

// bool / count

type (
	optBool struct {
		value []bool
	}
)

func (o *optBool) add(s string) error {
	v, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}

	o.value = append(o.value, v)
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

	o.value = append(o.value, v)
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

	o.value = append(o.value, v)
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

	o.value = append(o.value, v)
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
