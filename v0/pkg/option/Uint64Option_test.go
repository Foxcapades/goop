// Package option_test generated @ 2020-05-17 20:21:45.156432072 -0400 EDT m=+0.000834264
package option_test

import (
"testing"

. "github.com/Foxcapades/goop/v0/pkg/option"
. "github.com/smartystreets/goconvey/convey"
)

func TestNewUint64(t *testing.T) {
	Convey("NewUint64", t, func() {
		test := NewUint64(456789)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyUint64(t *testing.T) {
	Convey("NewEmptyUint64", t, func() {
		test := NewEmptyUint64()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeUint64(t *testing.T) {
	Convey("NewMaybeUint64", t, func() {
		var val1 uint64

		test1 := NewMaybeUint64(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeUint64(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestUint64_Get(t *testing.T) {
	Convey("Uint64.Get", t, func() {
		test1 := NewUint64(456789)
		So(test1.Get(), ShouldEqual, 456789)

		test2 := NewEmptyUint64()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestUint64_Or(t *testing.T) {
	Convey("Uint64.Or", t, func() {
		test1 := NewUint64(456789)
		So(test1.Or(9874654), ShouldEqual, 456789)

		test2 := NewEmptyUint64()
		So(test2.Or(456789), ShouldEqual, 456789)
	})
}

func TestUint64_OrGet(t *testing.T) {
	Convey("Uint64.OrGet", t, func() {
		test1 := NewUint64(456789)
		So(test1.OrGet(func() uint64 { return 9874654 }), ShouldEqual, 456789)

		test2 := NewEmptyUint64()
		So(test2.OrGet(func() uint64 { return 456789 }), ShouldEqual, 456789)
	})
}

func TestUint64_OrPanicWith(t *testing.T) {
	Convey("Uint64.OrPanicWith", t, func() {
		test1 := NewUint64(456789)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 456789)

		test2 := NewEmptyUint64()
		So(func () {
			test2.OrPanicWith(func() interface{} { return "test value" })}, ShouldPanicWith, "test value")
	})
}

func TestUint64_MapToNillable(t *testing.T) {
	Convey("Uint64.MapToNillable", t, func() {
		test1 := NewUint64(456789)
		So(test1.MapToNillable(func(b uint64) *uint64 { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewUint64(456789)
		So(test2.MapToNillable(func(b uint64) *uint64 { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyUint64()
		So(func() {test3.MapToNillable(func(b uint64) *uint64 { panic("foo") }).IsNil()}, ShouldNotPanic)
	})
}

