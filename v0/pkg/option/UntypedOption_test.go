// Package option_test generated @ 2020-05-17 20:21:45.156432072 -0400 EDT m=+0.000834264
package option_test

import (
"testing"

. "github.com/Foxcapades/goop/v0/pkg/option"
. "github.com/smartystreets/goconvey/convey"
)

func TestNewUntyped(t *testing.T) {
	Convey("NewUntyped", t, func() {
		test := NewUntyped("clam")
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyUntyped(t *testing.T) {
	Convey("NewEmptyUntyped", t, func() {
		test := NewEmptyUntyped()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeUntyped(t *testing.T) {
	Convey("NewMaybeUntyped", t, func() {
		var val1 interface{}

		test1 := NewMaybeUntyped(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeUntyped(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestUntyped_Get(t *testing.T) {
	Convey("Untyped.Get", t, func() {
		test1 := NewUntyped("clam")
		So(test1.Get(), ShouldEqual, "clam")

		test2 := NewEmptyUntyped()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestUntyped_Or(t *testing.T) {
	Convey("Untyped.Or", t, func() {
		test1 := NewUntyped("clam")
		So(test1.Or(69), ShouldEqual, "clam")

		test2 := NewEmptyUntyped()
		So(test2.Or("clam"), ShouldEqual, "clam")
	})
}

func TestUntyped_OrGet(t *testing.T) {
	Convey("Untyped.OrGet", t, func() {
		test1 := NewUntyped("clam")
		So(test1.OrGet(func() interface{} { return 69 }), ShouldEqual, "clam")

		test2 := NewEmptyUntyped()
		So(test2.OrGet(func() interface{} { return "clam" }), ShouldEqual, "clam")
	})
}

func TestUntyped_OrPanicWith(t *testing.T) {
	Convey("Untyped.OrPanicWith", t, func() {
		test1 := NewUntyped("clam")
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, "clam")

		test2 := NewEmptyUntyped()
		So(func () {
			test2.OrPanicWith(func() interface{} { return "test value" })}, ShouldPanicWith, "test value")
	})
}

func TestUntyped_MapToNillable(t *testing.T) {
	Convey("Untyped.MapToNillable", t, func() {
		test1 := NewUntyped("clam")
		So(test1.MapToNillable(func(b interface{}) *interface{} { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewUntyped("clam")
		So(test2.MapToNillable(func(b interface{}) *interface{} { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyUntyped()
		So(func() {test3.MapToNillable(func(b interface{}) *interface{} { panic("foo") }).IsNil()}, ShouldNotPanic)
	})
}

