// Package option_test generated @ 2020-05-17 20:31:35.070899067 -0400 EDT m=+0.000888110
package option_test

import (
	"testing"

	. "github.com/Foxcapades/goop/v0/pkg/option"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewByte(t *testing.T) {
	Convey("NewByte", t, func() {
		test := NewByte(255)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyByte(t *testing.T) {
	Convey("NewEmptyByte", t, func() {
		test := NewEmptyByte()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeByte(t *testing.T) {
	Convey("NewMaybeByte", t, func() {
		var val1 byte

		test1 := NewMaybeByte(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeByte(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestByte_Get(t *testing.T) {
	Convey("Byte.Get", t, func() {
		test1 := NewByte(255)
		So(test1.Get(), ShouldEqual, 255)

		test2 := NewEmptyByte()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestByte_Or(t *testing.T) {
	Convey("Byte.Or", t, func() {
		test1 := NewByte(255)
		So(test1.Or(56), ShouldEqual, 255)

		test2 := NewEmptyByte()
		So(test2.Or(255), ShouldEqual, 255)
	})
}

func TestByte_OrGet(t *testing.T) {
	Convey("Byte.OrGet", t, func() {
		test1 := NewByte(255)
		So(test1.OrGet(func() byte { return 56 }), ShouldEqual, 255)

		test2 := NewEmptyByte()
		So(test2.OrGet(func() byte { return 255 }), ShouldEqual, 255)
	})
}

func TestByte_OrPanicWith(t *testing.T) {
	Convey("Byte.OrPanicWith", t, func() {
		test1 := NewByte(255)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 255)

		test2 := NewEmptyByte()
		So(func() {
			test2.OrPanicWith(func() interface{} { return "test value" })
		}, ShouldPanicWith, "test value")
	})
}

func TestByte_MapToNillable(t *testing.T) {
	Convey("Byte.MapToNillable", t, func() {
		test1 := NewByte(255)
		So(test1.MapToNillable(func(b byte) *byte { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewByte(255)
		So(test2.MapToNillable(func(b byte) *byte { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyByte()
		So(func() {
			test3.MapToNillable(func(b byte) *byte { panic("foo") }).IsNil()
		}, ShouldNotPanic)
	})
}
