// Package option_test generated @ 2020-05-17 20:21:45.156432072 -0400 EDT m=+0.000834264
package option_test

import (
"testing"

. "github.com/Foxcapades/goop/v0/pkg/option"
. "github.com/smartystreets/goconvey/convey"
)

func TestNewBool(t *testing.T) {
	Convey("NewBool", t, func() {
		test := NewBool(true)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyBool(t *testing.T) {
	Convey("NewEmptyBool", t, func() {
		test := NewEmptyBool()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeBool(t *testing.T) {
	Convey("NewMaybeBool", t, func() {
		var val1 bool

		test1 := NewMaybeBool(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeBool(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestBool_Get(t *testing.T) {
	Convey("Bool.Get", t, func() {
		test1 := NewBool(true)
		So(test1.Get(), ShouldEqual, true)

		test2 := NewEmptyBool()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestBool_Or(t *testing.T) {
	Convey("Bool.Or", t, func() {
		test1 := NewBool(true)
		So(test1.Or(false), ShouldEqual, true)

		test2 := NewEmptyBool()
		So(test2.Or(true), ShouldEqual, true)
	})
}

func TestBool_OrGet(t *testing.T) {
	Convey("Bool.OrGet", t, func() {
		test1 := NewBool(true)
		So(test1.OrGet(func() bool { return false }), ShouldEqual, true)

		test2 := NewEmptyBool()
		So(test2.OrGet(func() bool { return true }), ShouldEqual, true)
	})
}

func TestBool_OrPanicWith(t *testing.T) {
	Convey("Bool.OrPanicWith", t, func() {
		test1 := NewBool(true)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, true)

		test2 := NewEmptyBool()
		So(func () {
			test2.OrPanicWith(func() interface{} { return "test value" })}, ShouldPanicWith, "test value")
	})
}

func TestBool_MapToNillable(t *testing.T) {
	Convey("Bool.MapToNillable", t, func() {
		test1 := NewBool(true)
		So(test1.MapToNillable(func(b bool) *bool { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewBool(true)
		So(test2.MapToNillable(func(b bool) *bool { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyBool()
		So(func() {test3.MapToNillable(func(b bool) *bool { panic("foo") }).IsNil()}, ShouldNotPanic)
	})
}

