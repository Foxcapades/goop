// Package option_test generated @ 2020-05-17 20:21:45.156432072 -0400 EDT m=+0.000834264
package option_test

import (
"testing"

. "github.com/Foxcapades/goop/v0/pkg/option"
. "github.com/smartystreets/goconvey/convey"
)

func TestNewComplex128(t *testing.T) {
	Convey("NewComplex128", t, func() {
		test := NewComplex128(58.98)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyComplex128(t *testing.T) {
	Convey("NewEmptyComplex128", t, func() {
		test := NewEmptyComplex128()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeComplex128(t *testing.T) {
	Convey("NewMaybeComplex128", t, func() {
		var val1 complex128

		test1 := NewMaybeComplex128(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeComplex128(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestComplex128_Get(t *testing.T) {
	Convey("Complex128.Get", t, func() {
		test1 := NewComplex128(58.98)
		So(test1.Get(), ShouldEqual, 58.98)

		test2 := NewEmptyComplex128()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestComplex128_Or(t *testing.T) {
	Convey("Complex128.Or", t, func() {
		test1 := NewComplex128(58.98)
		So(test1.Or(666.666), ShouldEqual, 58.98)

		test2 := NewEmptyComplex128()
		So(test2.Or(58.98), ShouldEqual, 58.98)
	})
}

func TestComplex128_OrGet(t *testing.T) {
	Convey("Complex128.OrGet", t, func() {
		test1 := NewComplex128(58.98)
		So(test1.OrGet(func() complex128 { return 666.666 }), ShouldEqual, 58.98)

		test2 := NewEmptyComplex128()
		So(test2.OrGet(func() complex128 { return 58.98 }), ShouldEqual, 58.98)
	})
}

func TestComplex128_OrPanicWith(t *testing.T) {
	Convey("Complex128.OrPanicWith", t, func() {
		test1 := NewComplex128(58.98)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 58.98)

		test2 := NewEmptyComplex128()
		So(func () {
			test2.OrPanicWith(func() interface{} { return "test value" })}, ShouldPanicWith, "test value")
	})
}

func TestComplex128_MapToNillable(t *testing.T) {
	Convey("Complex128.MapToNillable", t, func() {
		test1 := NewComplex128(58.98)
		So(test1.MapToNillable(func(b complex128) *complex128 { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewComplex128(58.98)
		So(test2.MapToNillable(func(b complex128) *complex128 { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyComplex128()
		So(func() {test3.MapToNillable(func(b complex128) *complex128 { panic("foo") }).IsNil()}, ShouldNotPanic)
	})
}

