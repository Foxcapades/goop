// Package option_test generated @ 2020-05-17 20:21:45.156432072 -0400 EDT m=+0.000834264
package option_test

import (
"testing"

. "github.com/Foxcapades/goop/v0/pkg/option"
. "github.com/smartystreets/goconvey/convey"
)

func TestNewFloat32(t *testing.T) {
	Convey("NewFloat32", t, func() {
		test := NewFloat32(28.3)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyFloat32(t *testing.T) {
	Convey("NewEmptyFloat32", t, func() {
		test := NewEmptyFloat32()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeFloat32(t *testing.T) {
	Convey("NewMaybeFloat32", t, func() {
		var val1 float32

		test1 := NewMaybeFloat32(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeFloat32(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestFloat32_Get(t *testing.T) {
	Convey("Float32.Get", t, func() {
		test1 := NewFloat32(28.3)
		So(test1.Get(), ShouldEqual, 28.3)

		test2 := NewEmptyFloat32()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestFloat32_Or(t *testing.T) {
	Convey("Float32.Or", t, func() {
		test1 := NewFloat32(28.3)
		So(test1.Or(69.9), ShouldEqual, 28.3)

		test2 := NewEmptyFloat32()
		So(test2.Or(28.3), ShouldEqual, 28.3)
	})
}

func TestFloat32_OrGet(t *testing.T) {
	Convey("Float32.OrGet", t, func() {
		test1 := NewFloat32(28.3)
		So(test1.OrGet(func() float32 { return 69.9 }), ShouldEqual, 28.3)

		test2 := NewEmptyFloat32()
		So(test2.OrGet(func() float32 { return 28.3 }), ShouldEqual, 28.3)
	})
}

func TestFloat32_OrPanicWith(t *testing.T) {
	Convey("Float32.OrPanicWith", t, func() {
		test1 := NewFloat32(28.3)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 28.3)

		test2 := NewEmptyFloat32()
		So(func () {
			test2.OrPanicWith(func() interface{} { return "test value" })}, ShouldPanicWith, "test value")
	})
}

func TestFloat32_MapToNillable(t *testing.T) {
	Convey("Float32.MapToNillable", t, func() {
		test1 := NewFloat32(28.3)
		So(test1.MapToNillable(func(b float32) *float32 { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewFloat32(28.3)
		So(test2.MapToNillable(func(b float32) *float32 { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyFloat32()
		So(func() {test3.MapToNillable(func(b float32) *float32 { panic("foo") }).IsNil()}, ShouldNotPanic)
	})
}

