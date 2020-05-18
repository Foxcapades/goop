// Package option_test generated @ 2020-05-17 20:21:45.156432072 -0400 EDT m=+0.000834264
package option_test

import (
"testing"

. "github.com/Foxcapades/goop/v0/pkg/option"
. "github.com/smartystreets/goconvey/convey"
)

func TestNewInt16(t *testing.T) {
	Convey("NewInt16", t, func() {
		test := NewInt16(256)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyInt16(t *testing.T) {
	Convey("NewEmptyInt16", t, func() {
		test := NewEmptyInt16()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeInt16(t *testing.T) {
	Convey("NewMaybeInt16", t, func() {
		var val1 int16

		test1 := NewMaybeInt16(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeInt16(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestInt16_Get(t *testing.T) {
	Convey("Int16.Get", t, func() {
		test1 := NewInt16(256)
		So(test1.Get(), ShouldEqual, 256)

		test2 := NewEmptyInt16()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestInt16_Or(t *testing.T) {
	Convey("Int16.Or", t, func() {
		test1 := NewInt16(256)
		So(test1.Or(585), ShouldEqual, 256)

		test2 := NewEmptyInt16()
		So(test2.Or(256), ShouldEqual, 256)
	})
}

func TestInt16_OrGet(t *testing.T) {
	Convey("Int16.OrGet", t, func() {
		test1 := NewInt16(256)
		So(test1.OrGet(func() int16 { return 585 }), ShouldEqual, 256)

		test2 := NewEmptyInt16()
		So(test2.OrGet(func() int16 { return 256 }), ShouldEqual, 256)
	})
}

func TestInt16_OrPanicWith(t *testing.T) {
	Convey("Int16.OrPanicWith", t, func() {
		test1 := NewInt16(256)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 256)

		test2 := NewEmptyInt16()
		So(func () {
			test2.OrPanicWith(func() interface{} { return "test value" })}, ShouldPanicWith, "test value")
	})
}

func TestInt16_MapToNillable(t *testing.T) {
	Convey("Int16.MapToNillable", t, func() {
		test1 := NewInt16(256)
		So(test1.MapToNillable(func(b int16) *int16 { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewInt16(256)
		So(test2.MapToNillable(func(b int16) *int16 { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyInt16()
		So(func() {test3.MapToNillable(func(b int16) *int16 { panic("foo") }).IsNil()}, ShouldNotPanic)
	})
}

