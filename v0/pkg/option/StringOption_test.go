// Package option_test generated @ 2020-05-17 20:21:45.156432072 -0400 EDT m=+0.000834264
package option_test

import (
"testing"

. "github.com/Foxcapades/goop/v0/pkg/option"
. "github.com/smartystreets/goconvey/convey"
)

func TestNewString(t *testing.T) {
	Convey("NewString", t, func() {
		test := NewString("hi")
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyString(t *testing.T) {
	Convey("NewEmptyString", t, func() {
		test := NewEmptyString()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeString(t *testing.T) {
	Convey("NewMaybeString", t, func() {
		var val1 string

		test1 := NewMaybeString(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeString(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestString_Get(t *testing.T) {
	Convey("String.Get", t, func() {
		test1 := NewString("hi")
		So(test1.Get(), ShouldEqual, "hi")

		test2 := NewEmptyString()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestString_Or(t *testing.T) {
	Convey("String.Or", t, func() {
		test1 := NewString("hi")
		So(test1.Or("bye"), ShouldEqual, "hi")

		test2 := NewEmptyString()
		So(test2.Or("hi"), ShouldEqual, "hi")
	})
}

func TestString_OrGet(t *testing.T) {
	Convey("String.OrGet", t, func() {
		test1 := NewString("hi")
		So(test1.OrGet(func() string { return "bye" }), ShouldEqual, "hi")

		test2 := NewEmptyString()
		So(test2.OrGet(func() string { return "hi" }), ShouldEqual, "hi")
	})
}

func TestString_OrPanicWith(t *testing.T) {
	Convey("String.OrPanicWith", t, func() {
		test1 := NewString("hi")
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, "hi")

		test2 := NewEmptyString()
		So(func () {
			test2.OrPanicWith(func() interface{} { return "test value" })}, ShouldPanicWith, "test value")
	})
}

func TestString_MapToNillable(t *testing.T) {
	Convey("String.MapToNillable", t, func() {
		test1 := NewString("hi")
		So(test1.MapToNillable(func(b string) *string { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewString("hi")
		So(test2.MapToNillable(func(b string) *string { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyString()
		So(func() {test3.MapToNillable(func(b string) *string { panic("foo") }).IsNil()}, ShouldNotPanic)
	})
}

