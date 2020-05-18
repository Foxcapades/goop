// Package option_test generated @ 2020-05-17 20:31:35.070899067 -0400 EDT m=+0.000888110
package option_test

import (
	"testing"

	. "github.com/Foxcapades/goop/v0/pkg/option"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewInt32(t *testing.T) {
	Convey("NewInt32", t, func() {
		test := NewInt32(1245)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyInt32(t *testing.T) {
	Convey("NewEmptyInt32", t, func() {
		test := NewEmptyInt32()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeInt32(t *testing.T) {
	Convey("NewMaybeInt32", t, func() {
		var val1 int32

		test1 := NewMaybeInt32(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeInt32(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestInt32_Get(t *testing.T) {
	Convey("Int32.Get", t, func() {
		test1 := NewInt32(1245)
		So(test1.Get(), ShouldEqual, 1245)

		test2 := NewEmptyInt32()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestInt32_Or(t *testing.T) {
	Convey("Int32.Or", t, func() {
		test1 := NewInt32(1245)
		So(test1.Or(1234), ShouldEqual, 1245)

		test2 := NewEmptyInt32()
		So(test2.Or(1245), ShouldEqual, 1245)
	})
}

func TestInt32_OrGet(t *testing.T) {
	Convey("Int32.OrGet", t, func() {
		test1 := NewInt32(1245)
		So(test1.OrGet(func() int32 { return 1234 }), ShouldEqual, 1245)

		test2 := NewEmptyInt32()
		So(test2.OrGet(func() int32 { return 1245 }), ShouldEqual, 1245)
	})
}

func TestInt32_OrPanicWith(t *testing.T) {
	Convey("Int32.OrPanicWith", t, func() {
		test1 := NewInt32(1245)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 1245)

		test2 := NewEmptyInt32()
		So(func() {
			test2.OrPanicWith(func() interface{} { return "test value" })
		}, ShouldPanicWith, "test value")
	})
}

func TestInt32_MapToNillable(t *testing.T) {
	Convey("Int32.MapToNillable", t, func() {
		test1 := NewInt32(1245)
		So(test1.MapToNillable(func(b int32) *int32 { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewInt32(1245)
		So(test2.MapToNillable(func(b int32) *int32 { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyInt32()
		So(func() {
			test3.MapToNillable(func(b int32) *int32 { panic("foo") }).IsNil()
		}, ShouldNotPanic)
	})
}
