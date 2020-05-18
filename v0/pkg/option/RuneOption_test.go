// Package option_test generated @ 2020-05-17 20:21:45.156432072 -0400 EDT m=+0.000834264
package option_test

import (
"testing"

. "github.com/Foxcapades/goop/v0/pkg/option"
. "github.com/smartystreets/goconvey/convey"
)

func TestNewRune(t *testing.T) {
	Convey("NewRune", t, func() {
		test := NewRune(454545)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyRune(t *testing.T) {
	Convey("NewEmptyRune", t, func() {
		test := NewEmptyRune()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeRune(t *testing.T) {
	Convey("NewMaybeRune", t, func() {
		var val1 rune

		test1 := NewMaybeRune(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeRune(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestRune_Get(t *testing.T) {
	Convey("Rune.Get", t, func() {
		test1 := NewRune(454545)
		So(test1.Get(), ShouldEqual, 454545)

		test2 := NewEmptyRune()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestRune_Or(t *testing.T) {
	Convey("Rune.Or", t, func() {
		test1 := NewRune(454545)
		So(test1.Or(898989), ShouldEqual, 454545)

		test2 := NewEmptyRune()
		So(test2.Or(454545), ShouldEqual, 454545)
	})
}

func TestRune_OrGet(t *testing.T) {
	Convey("Rune.OrGet", t, func() {
		test1 := NewRune(454545)
		So(test1.OrGet(func() rune { return 898989 }), ShouldEqual, 454545)

		test2 := NewEmptyRune()
		So(test2.OrGet(func() rune { return 454545 }), ShouldEqual, 454545)
	})
}

func TestRune_OrPanicWith(t *testing.T) {
	Convey("Rune.OrPanicWith", t, func() {
		test1 := NewRune(454545)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 454545)

		test2 := NewEmptyRune()
		So(func () {
			test2.OrPanicWith(func() interface{} { return "test value" })}, ShouldPanicWith, "test value")
	})
}

func TestRune_MapToNillable(t *testing.T) {
	Convey("Rune.MapToNillable", t, func() {
		test1 := NewRune(454545)
		So(test1.MapToNillable(func(b rune) *rune { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewRune(454545)
		So(test2.MapToNillable(func(b rune) *rune { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyRune()
		So(func() {test3.MapToNillable(func(b rune) *rune { panic("foo") }).IsNil()}, ShouldNotPanic)
	})
}

