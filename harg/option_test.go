package harg_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/jtagcat/hcli/harg"
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
	if err != nil {
		t.Fatal(err)
	}
	if len(parsed) != 0 || len(chokeReturn) != 0 {
		t.Fatal("parsed or chokeReturn is not empty")
	}

	if defs[key].Default() {
		t.Fatal("value is default")
	}

	v, ok := defs[key].SlBool()
	if !ok {
		t.Fatal("sl call not ok")
	}
	if len(v) != len(want) || v[0] != want[0] || v[1] != want[1] {
		t.Fatal("did not get wanted")
	}

	sv, ok := defs[key].Bool()
	if !ok {
		t.Fatal("single value not ok")
	}
	if sv != want[1] {
		t.Fatal("single value did not match wanted")
	}

	if !defs[key].IsBool() {
		t.Fatal("value should be bool")
	}

	c, ok := defs[key].Count()
	if !ok {
		t.Fatal("single value not ok")
	}
	if c != len(want) {
		t.Fatal("count did not match wanted")
	}
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
	if err != nil {
		t.Fatal(err)
	}
	if len(parsed) != 0 || len(chokeReturn) != 0 {
		t.Fatal("parsed or chokeReturn is not empty")
	}

	if defs[key].Default() {
		t.Fatal("value is default")
	}

	v, ok := defs[key].SlString()
	if !ok {
		t.Fatal("sl call not ok")
	}
	if len(v) != len(want) || v[0] != want[0] || v[1] != want[1] {
		t.Fatal("did not get wanted")
	}

	sv, ok := defs[key].String()
	if !ok {
		t.Fatal("single value not ok")
	}
	if sv != want[1] {
		t.Fatal("single value did not match wanted")
	}

	if defs[key].IsBool() {
		t.Fatal("value should not be bool")
	}
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
	if err != nil {
		t.Fatal(err)
	}
	if len(parsed) != 0 || len(chokeReturn) != 0 {
		t.Fatal("parsed or chokeReturn is not empty")
	}

	if defs[key].Default() {
		t.Fatal("value is default")
	}

	v, ok := defs[key].SlInt()
	if !ok {
		t.Fatal("sl call not ok")
	}
	if len(v) != len(want) || v[0] != want[0] || v[1] != want[1] {
		t.Fatal("did not get wanted")
	}

	sv, ok := defs[key].Int()
	if !ok {
		t.Fatal("single value not ok")
	}
	if sv != want[1] {
		t.Fatal("single value did not match wanted")
	}

	if defs[key].IsBool() {
		t.Fatal("value should not be bool")
	}
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
	if err != nil {
		t.Fatal(err)
	}
	if len(parsed) != 0 || len(chokeReturn) != 0 {
		t.Fatal("parsed or chokeReturn is not empty")
	}

	if defs[key].Default() {
		t.Fatal("value is default")
	}

	v, ok := defs[key].SlInt64()
	if !ok {
		t.Fatal("sl call not ok")
	}
	if len(v) != len(want) || v[0] != want[0] || v[1] != want[1] {
		t.Fatal("did not get wanted")
	}

	sv, ok := defs[key].Int64()
	if !ok {
		t.Fatal("single value not ok")
	}
	if sv != want[1] {
		t.Fatal("single value did not match wanted")
	}

	if defs[key].IsBool() {
		t.Fatal("value should not be bool")
	}
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
	if err != nil {
		t.Fatal(err)
	}
	if len(parsed) != 0 || len(chokeReturn) != 0 {
		t.Fatal("parsed or chokeReturn is not empty")
	}

	if defs[key].Default() {
		t.Fatal("value is default")
	}

	v, ok := defs[key].SlUint()
	if !ok {
		t.Fatal("sl call not ok")
	}
	if len(v) != len(want) || v[0] != want[0] || v[1] != want[1] {
		t.Fatal("did not get wanted")
	}

	sv, ok := defs[key].Uint()
	if !ok {
		t.Fatal("single value not ok")
	}
	if sv != want[1] {
		t.Fatal("single value did not match wanted")
	}

	if defs[key].IsBool() {
		t.Fatal("value should not be bool")
	}
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
	if err != nil {
		t.Fatal(err)
	}
	if len(parsed) != 0 || len(chokeReturn) != 0 {
		t.Fatal("parsed or chokeReturn is not empty")
	}

	if defs[key].Default() {
		t.Fatal("value is default")
	}

	v, ok := defs[key].SlUint64()
	if !ok {
		t.Fatal("sl call not ok")
	}
	if len(v) != len(want) || v[0] != want[0] || v[1] != want[1] {
		t.Fatal("did not get wanted")
	}

	sv, ok := defs[key].Uint64()
	if !ok {
		t.Fatal("single value not ok")
	}
	if sv != want[1] {
		t.Fatal("single value did not match wanted")
	}

	if defs[key].IsBool() {
		t.Fatal("value should not be bool")
	}
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
	if err != nil {
		t.Fatal(err)
	}
	if len(parsed) != 0 || len(chokeReturn) != 0 {
		t.Fatal("parsed or chokeReturn is not empty")
	}

	if defs[key].Default() {
		t.Fatal("value is default")
	}

	v, ok := defs[key].SlFloat64()
	if !ok {
		t.Fatal("sl call not ok")
	}
	if len(v) != len(want) || v[0] != want[0] || v[1] != want[1] {
		t.Fatal("did not get wanted")
	}

	sv, ok := defs[key].Float64()
	if !ok {
		t.Fatal("single value not ok")
	}
	if sv != want[1] {
		t.Fatal("single value did not match wanted")
	}

	if defs[key].IsBool() {
		t.Fatal("value should not be bool")
	}
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
	if err != nil {
		t.Fatal(err)
	}
	if len(parsed) != 0 || len(chokeReturn) != 0 {
		t.Fatal("parsed or chokeReturn is not empty")
	}

	if defs[key].Default() {
		t.Fatal("value is default")
	}

	v, ok := defs[key].SlDuration()
	if !ok {
		t.Fatal("sl call not ok")
	}
	if len(v) != len(want) || v[0] != want[0] || v[1] != want[1] {
		t.Fatal("did not get wanted")
	}

	sv, ok := defs[key].Duration()
	if !ok {
		t.Fatal("single value not ok")
	}
	if sv != want[1] {
		t.Fatal("single value did not match wanted")
	}

	if defs[key].IsBool() {
		t.Fatal("value should not be bool")
	}
}
