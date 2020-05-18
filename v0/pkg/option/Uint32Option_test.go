// Package option_test generated @ 2020-05-17 20:31:35.070899067 -0400 EDT m=+0.000888110
package option_test

import (
	"testing"

	. "github.com/Foxcapades/goop/v0/pkg/option"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewUint32(t *testing.T) {
	Convey("NewUint32", t, func() {
		test := NewUint32(123456)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyUint32(t *testing.T) {
	Convey("NewEmptyUint32", t, func() {
		test := NewEmptyUint32()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeUint32(t *testing.T) {
	Convey("NewMaybeUint32", t, func() {
		var val1 uint32

		test1 := NewMaybeUint32(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeUint32(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestUint32_Get(t *testing.T) {
	Convey("Uint32.Get", t, func() {
		test1 := NewUint32(123456)
		So(test1.Get(), ShouldEqual, 123456)

		test2 := NewEmptyUint32()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestUint32_Or(t *testing.T) {
	Convey("Uint32.Or", t, func() {
		test1 := NewUint32(123456)
		So(test1.Or(654321), ShouldEqual, 123456)

		test2 := NewEmptyUint32()
		So(test2.Or(123456), ShouldEqual, 123456)
	})
}

func TestUint32_OrGet(t *testing.T) {
	Convey("Uint32.OrGet", t, func() {
		test1 := NewUint32(123456)
		So(test1.OrGet(func() uint32 { return 654321 }), ShouldEqual, 123456)

		test2 := NewEmptyUint32()
		So(test2.OrGet(func() uint32 { return 123456 }), ShouldEqual, 123456)
	})
}

func TestUint32_OrPanicWith(t *testing.T) {
	Convey("Uint32.OrPanicWith", t, func() {
		test1 := NewUint32(123456)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 123456)

		test2 := NewEmptyUint32()
		So(func() {
			test2.OrPanicWith(func() interface{} { return "test value" })
		}, ShouldPanicWith, "test value")
	})
}

func TestUint32_MapToNillable(t *testing.T) {
	Convey("Uint32.MapToNillable", t, func() {
		test1 := NewUint32(123456)
		So(test1.MapToNillable(func(b uint32) *uint32 { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewUint32(123456)
		So(test2.MapToNillable(func(b uint32) *uint32 { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyUint32()
		So(func() {
			test3.MapToNillable(func(b uint32) *uint32 { panic("foo") }).IsNil()
		}, ShouldNotPanic)
	})
}
