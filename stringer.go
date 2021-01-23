// Package stringer provides fmt.Stringer wrappers for builtin go types.
package stringer

import (
	"fmt"
	"strconv"
)

// String wraps the given string in a fmt.Stringer.
func String(s string) fmt.Stringer {
	return str(s)
}

type str string

func (s str) String() string {
	return string(s)
}

// Int wraps the given int in a fmt.Stringer.
func Int(n int) fmt.Stringer {
	return i(n)
}

type i int

func (n i) String() string {
	return strconv.FormatInt(int64(n), 10)
}

// Uint wraps the given uint in a fmt.Stringer.
func Uint(n uint) fmt.Stringer {
	return ui(n)
}

type ui uint

func (n ui) String() string {
	return strconv.FormatInt(int64(n), 10)
}

// Int8 wraps given int8 in a fmt.Stringer.
func Int8(n int8) fmt.Stringer {
	return i8(n)
}

type i8 int8

func (n i8) String() string {
	return strconv.FormatInt(int64(n), 10)
}

// Uint8 wraps given uint8 in a fmt.Stringer.
func Uint8(n uint8) fmt.Stringer {
	return ui8(n)
}

type ui8 uint8

func (n ui8) String() string {
	return strconv.FormatUint(uint64(n), 10)
}

// Int16 wraps given int16 in a fmt.Stringer.
func Int16(n int16) fmt.Stringer {
	return i16(n)
}

type i16 int16

func (n i16) String() string {
	return strconv.FormatInt(int64(n), 10)
}

// Uint16 wraps given uint16 in a fmt.Stringer.
func Uint16(n uint16) fmt.Stringer {
	return ui16(n)
}

type ui16 uint16

func (n ui16) String() string {
	return strconv.FormatUint(uint64(n), 10)
}

// Int32 wraps given int32 in a fmt.Stringer.
func Int32(n int32) fmt.Stringer {
	return i32(n)
}

type i32 int32

func (n i32) String() string {
	return strconv.FormatInt(int64(n), 10)
}

// Uint32 wraps given uint32 in a fmt.Stringer.
func Uint32(n uint32) fmt.Stringer {
	return ui32(n)
}

type ui32 uint32

func (n ui32) String() string {
	return strconv.FormatUint(uint64(n), 10)
}

// Int64 wraps given int64 in a fmt.Stringer.
func Int64(n int64) fmt.Stringer {
	return i64(n)
}

type i64 int64

func (n i64) String() string {
	return strconv.FormatInt(int64(n), 10)
}

// Uint64 wraps given uint64 in a fmt.Stringer.
func Uint64(n uint64) fmt.Stringer {
	return ui64(n)
}

type ui64 uint64

func (n ui64) String() string {
	return strconv.FormatUint(uint64(n), 10)
}

// Float32 wraps given float32 in a fmt.Stringer.
func Float32(n float32) fmt.Stringer {
	return f32(n)
}

type f32 float32

func (f f32) String() string {
	return strconv.FormatFloat(float64(f), 'g', -1, 32)
}

// Float64 wraps given float64 in a fmt.Stringer.
func Float64(n float64) fmt.Stringer {
	return f64(n)
}

type f64 float64

func (f f64) String() string {
	return strconv.FormatFloat(float64(f), 'g', -1, 64)
}

// Bool wraps given bool in a fmt.Stringer.
func Bool(n bool) fmt.Stringer {
	return boolean(n)
}

type boolean bool

func (b boolean) String() string {
	return strconv.FormatBool(bool(b))
}
