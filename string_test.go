package tostring

import (
	"errors"
	"fmt"
	"math"
	"testing"
)

const (
	mInt   int   = math.MaxInt
	mInt8  int8  = math.MaxInt8
	mInt16 int16 = math.MaxInt16
	mInt32 int32 = math.MaxInt32
	mInt64 int64 = math.MaxInt64

	mUint   uint   = math.MaxUint
	mUint8  uint8  = math.MaxUint8
	mUint16 uint16 = math.MaxUint16
	mUint32 uint32 = math.MaxUint32
	mUint64 uint64 = math.MaxUint64

	mFloat32 float32 = math.MaxFloat32
	mFloat64 float64 = math.MaxFloat64
)

type test struct {
	name  string
	value any
}

var integers = []test{
	{"Int", mInt},
	{"Int8", mInt8},
	{"Int16", mInt16},
	{"Int32", mInt32},
	{"Int64", mInt64},
	{"Uint", mUint},
	{"Uint8", mUint8},
	{"Uint16", mUint16},
	{"Uint32", mUint32},
	{"Uint64", mUint64},
	{"Float32", mFloat32},
	{"Float64", mFloat64},
	{"Float32 (small)", float32(1.520)},
	{"Float64 (small)", float64(1.520)},
}

var others = []test{
	{"String", str},
	{"Stringer", newStringer(str)},
	{"bool", true},
	{"boolf", false},
	{"nil", nil},
	{"error", errors.New("err message")},
}

func TestAny(t *testing.T) {
	var suite []test

	suite = append(suite, others...)
	suite = append(suite, integers...)

	for _, s := range suite {
		t.Run(s.name, func(t *testing.T) {
			expect := fmt.Sprint(s.value)
			got, gotSafe := Any(s.value), AnyUnsafe(s.value)

			equal(t, expect, got)
			equal(t, expect, gotSafe)
		})
	}
}

func TestAnyBytes(t *testing.T) {
	b := []byte(str)
	equal(t, str, Any(b))
	equal(t, str, AnyUnsafe(b))
}

var str = "value"

type stringer struct {
	s string
}

func newStringer(s string) *stringer {
	return &stringer{s}
}

func (ms *stringer) String() string {
	return ms.s
}

func equal(t *testing.T, expect, got any) {
	if expect != got {
		t.Fatalf("expected %s, got %s", expect, got)
	}
}
