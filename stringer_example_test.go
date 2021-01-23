package stringer_test

import (
	"fmt"

	"github.com/hemantjadon/stringer"
)

func ExampleString() {
	var v fmt.Stringer

	v = stringer.String("hello")

	fmt.Println(v)
	// Output: hello
}

func ExampleInt() {
	var v fmt.Stringer

	v = stringer.Int(10)

	fmt.Println(v)
	// Output: 10
}

func ExampleUint() {
	var v fmt.Stringer

	v = stringer.Uint(10)

	fmt.Println(v)
	// Output: 10
}

func ExampleInt8() {
	var v fmt.Stringer

	v = stringer.Int8(10)

	fmt.Println(v)
	// Output: 10
}

func ExampleUint8() {
	var v fmt.Stringer

	v = stringer.Uint8(10)

	fmt.Println(v)
	// Output: 10
}

func ExampleInt16() {
	var v fmt.Stringer

	v = stringer.Int16(10)

	fmt.Println(v)
	// Output: 10
}

func ExampleUint16() {
	var v fmt.Stringer

	v = stringer.Uint16(10)

	fmt.Println(v)
	// Output: 10
}

func ExampleInt32() {
	var v fmt.Stringer

	v = stringer.Int32(10)

	fmt.Println(v)
	// Output: 10
}

func ExampleUint32() {
	var v fmt.Stringer

	v = stringer.Uint32(10)

	fmt.Println(v)
	// Output: 10
}

func ExampleInt64() {
	var v fmt.Stringer

	v = stringer.Int64(10)

	fmt.Println(v)
	// Output: 10
}

func ExampleUint64() {
	var v fmt.Stringer

	v = stringer.Uint64(10)

	fmt.Println(v)
	// Output: 10
}

func ExampleFloat32() {
	var v, w fmt.Stringer

	v = stringer.Float32(25.3456)
	w = stringer.Float32(1000010000.987)

	fmt.Println(v)
	fmt.Println(w)
	// Output:
	// 25.3456
	// 1.00001e+09
}

func ExampleFloat64() {
	var v, w fmt.Stringer

	v = stringer.Float64(25.3456)
	w = stringer.Float64(1000010000.987)

	fmt.Println(v)
	fmt.Println(w)
	// Output:
	// 25.3456
	// 1.000010000987e+09
}

func ExampleBool() {
	var v, w fmt.Stringer

	v = stringer.Bool(true)
	w = stringer.Bool(false)

	fmt.Println(v)
	fmt.Println(w)
	// Output:
	// true
	// false
}
