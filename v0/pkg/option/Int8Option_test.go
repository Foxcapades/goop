// Package option_test generated @ 2020-05-17 20:31:35.070899067 -0400 EDT m=+0.000888110
package option_test

import (
	"testing"

	. "github.com/Foxcapades/goop/v0/pkg/option"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewInt8(t *testing.T) {
	Convey("NewInt8", t, func() {
		test := NewInt8(127)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyInt8(t *testing.T) {
	Convey("NewEmptyInt8", t, func() {
		test := NewEmptyInt8()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeInt8(t *testing.T) {
	Convey("NewMaybeInt8", t, func() {
		var val1 int8

		test1 := NewMaybeInt8(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeInt8(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestInt8_Get(t *testing.T) {
	Convey("Int8.Get", t, func() {
		test1 := NewInt8(127)
		So(test1.Get(), ShouldEqual, 127)

		test2 := NewEmptyInt8()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestInt8_Or(t *testing.T) {
	Convey("Int8.Or", t, func() {
		test1 := NewInt8(127)
		So(test1.Or(-128), ShouldEqual, 127)

		test2 := NewEmptyInt8()
		So(test2.Or(127), ShouldEqual, 127)
	})
}

func TestInt8_OrGet(t *testing.T) {
	Convey("Int8.OrGet", t, func() {
		test1 := NewInt8(127)
		So(test1.OrGet(func() int8 { return -128 }), ShouldEqual, 127)

		test2 := NewEmptyInt8()
		So(test2.OrGet(func() int8 { return 127 }), ShouldEqual, 127)
	})
}

func TestInt8_OrPanicWith(t *testing.T) {
	Convey("Int8.OrPanicWith", t, func() {
		test1 := NewInt8(127)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 127)

		test2 := NewEmptyInt8()
		So(func() {
			test2.OrPanicWith(func() interface{} { return "test value" })
		}, ShouldPanicWith, "test value")
	})
}

func TestInt8_MapToNillable(t *testing.T) {
	Convey("Int8.MapToNillable", t, func() {
		test1 := NewInt8(127)
		So(test1.MapToNillable(func(b int8) *int8 { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewInt8(127)
		So(test2.MapToNillable(func(b int8) *int8 { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyInt8()
		So(func() {
			test3.MapToNillable(func(b int8) *int8 { panic("foo") }).IsNil()
		}, ShouldNotPanic)
	})
}
