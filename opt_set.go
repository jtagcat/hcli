package harg

import (
	"fmt"
	"strconv"
	"time"
)

func (defM *DefinitionMap) toEmptyOptM() OptionsTypedMap {
	return OptionsTypedMap{} // TODO:
}

// TODO:
type (
	OptionsTypedMap map[Type]optionsMap
	optionsMap      map[string]option // parallel to Definitions.D
	option          interface {
		found(write bool) bool
		contents() any // resolved with option.Sl
		add(rawOpt string) error
	}
)

func (o *optCommon) found(write bool) bool {
	if write {
		o.foundV = true
	}
	return o.foundV
}

type optCommon struct {
	foundV bool
}

// bool

// type optBool struct {
// 	optCommon
// 	count     int  //
// 	lastValue bool // result
// }

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
	return nil
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
	return nil
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
	return nil
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
	return nil
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
	return nil
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
	return nil
}

// TODO:
func (optM *OptionsTypedMap) parseOptionContent(def *Definition, value *string, valueFound *bool) error {
	return fmt.Errorf("not implemented") // TODO:
}
