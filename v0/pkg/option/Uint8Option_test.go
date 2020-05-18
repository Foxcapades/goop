// Package option_test generated @ 2020-05-17 20:21:45.156432072 -0400 EDT m=+0.000834264
package option_test

import (
"testing"

. "github.com/Foxcapades/goop/v0/pkg/option"
. "github.com/smartystreets/goconvey/convey"
)

func TestNewUint8(t *testing.T) {
	Convey("NewUint8", t, func() {
		test := NewUint8(123)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyUint8(t *testing.T) {
	Convey("NewEmptyUint8", t, func() {
		test := NewEmptyUint8()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeUint8(t *testing.T) {
	Convey("NewMaybeUint8", t, func() {
		var val1 uint8

		test1 := NewMaybeUint8(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeUint8(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestUint8_Get(t *testing.T) {
	Convey("Uint8.Get", t, func() {
		test1 := NewUint8(123)
		So(test1.Get(), ShouldEqual, 123)

		test2 := NewEmptyUint8()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestUint8_Or(t *testing.T) {
	Convey("Uint8.Or", t, func() {
		test1 := NewUint8(123)
		So(test1.Or(245), ShouldEqual, 123)

		test2 := NewEmptyUint8()
		So(test2.Or(123), ShouldEqual, 123)
	})
}

func TestUint8_OrGet(t *testing.T) {
	Convey("Uint8.OrGet", t, func() {
		test1 := NewUint8(123)
		So(test1.OrGet(func() uint8 { return 245 }), ShouldEqual, 123)

		test2 := NewEmptyUint8()
		So(test2.OrGet(func() uint8 { return 123 }), ShouldEqual, 123)
	})
}

func TestUint8_OrPanicWith(t *testing.T) {
	Convey("Uint8.OrPanicWith", t, func() {
		test1 := NewUint8(123)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 123)

		test2 := NewEmptyUint8()
		So(func () {
			test2.OrPanicWith(func() interface{} { return "test value" })}, ShouldPanicWith, "test value")
	})
}

func TestUint8_MapToNillable(t *testing.T) {
	Convey("Uint8.MapToNillable", t, func() {
		test1 := NewUint8(123)
		So(test1.MapToNillable(func(b uint8) *uint8 { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewUint8(123)
		So(test2.MapToNillable(func(b uint8) *uint8 { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyUint8()
		So(func() {test3.MapToNillable(func(b uint8) *uint8 { panic("foo") }).IsNil()}, ShouldNotPanic)
	})
}

