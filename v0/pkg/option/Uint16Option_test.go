// Package option_test generated @ 2020-05-17 20:31:35.070899067 -0400 EDT m=+0.000888110
package option_test

import (
	"testing"

	. "github.com/Foxcapades/goop/v0/pkg/option"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewUint16(t *testing.T) {
	Convey("NewUint16", t, func() {
		test := NewUint16(4562)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyUint16(t *testing.T) {
	Convey("NewEmptyUint16", t, func() {
		test := NewEmptyUint16()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeUint16(t *testing.T) {
	Convey("NewMaybeUint16", t, func() {
		var val1 uint16

		test1 := NewMaybeUint16(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeUint16(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestUint16_Get(t *testing.T) {
	Convey("Uint16.Get", t, func() {
		test1 := NewUint16(4562)
		So(test1.Get(), ShouldEqual, 4562)

		test2 := NewEmptyUint16()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestUint16_Or(t *testing.T) {
	Convey("Uint16.Or", t, func() {
		test1 := NewUint16(4562)
		So(test1.Or(7894), ShouldEqual, 4562)

		test2 := NewEmptyUint16()
		So(test2.Or(4562), ShouldEqual, 4562)
	})
}

func TestUint16_OrGet(t *testing.T) {
	Convey("Uint16.OrGet", t, func() {
		test1 := NewUint16(4562)
		So(test1.OrGet(func() uint16 { return 7894 }), ShouldEqual, 4562)

		test2 := NewEmptyUint16()
		So(test2.OrGet(func() uint16 { return 4562 }), ShouldEqual, 4562)
	})
}

func TestUint16_OrPanicWith(t *testing.T) {
	Convey("Uint16.OrPanicWith", t, func() {
		test1 := NewUint16(4562)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 4562)

		test2 := NewEmptyUint16()
		So(func() {
			test2.OrPanicWith(func() interface{} { return "test value" })
		}, ShouldPanicWith, "test value")
	})
}

func TestUint16_MapToNillable(t *testing.T) {
	Convey("Uint16.MapToNillable", t, func() {
		test1 := NewUint16(4562)
		So(test1.MapToNillable(func(b uint16) *uint16 { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewUint16(4562)
		So(test2.MapToNillable(func(b uint16) *uint16 { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyUint16()
		So(func() {
			test3.MapToNillable(func(b uint16) *uint16 { panic("foo") }).IsNil()
		}, ShouldNotPanic)
	})
}
