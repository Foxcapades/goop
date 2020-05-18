// Package option_test generated @ 2020-05-17 20:31:35.070899067 -0400 EDT m=+0.000888110
package option_test

import (
	"testing"

	. "github.com/Foxcapades/goop/v0/pkg/option"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewUint(t *testing.T) {
	Convey("NewUint", t, func() {
		test := NewUint(5898)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyUint(t *testing.T) {
	Convey("NewEmptyUint", t, func() {
		test := NewEmptyUint()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeUint(t *testing.T) {
	Convey("NewMaybeUint", t, func() {
		var val1 uint

		test1 := NewMaybeUint(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeUint(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestUint_Get(t *testing.T) {
	Convey("Uint.Get", t, func() {
		test1 := NewUint(5898)
		So(test1.Get(), ShouldEqual, 5898)

		test2 := NewEmptyUint()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestUint_Or(t *testing.T) {
	Convey("Uint.Or", t, func() {
		test1 := NewUint(5898)
		So(test1.Or(36985), ShouldEqual, 5898)

		test2 := NewEmptyUint()
		So(test2.Or(5898), ShouldEqual, 5898)
	})
}

func TestUint_OrGet(t *testing.T) {
	Convey("Uint.OrGet", t, func() {
		test1 := NewUint(5898)
		So(test1.OrGet(func() uint { return 36985 }), ShouldEqual, 5898)

		test2 := NewEmptyUint()
		So(test2.OrGet(func() uint { return 5898 }), ShouldEqual, 5898)
	})
}

func TestUint_OrPanicWith(t *testing.T) {
	Convey("Uint.OrPanicWith", t, func() {
		test1 := NewUint(5898)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 5898)

		test2 := NewEmptyUint()
		So(func() {
			test2.OrPanicWith(func() interface{} { return "test value" })
		}, ShouldPanicWith, "test value")
	})
}

func TestUint_MapToNillable(t *testing.T) {
	Convey("Uint.MapToNillable", t, func() {
		test1 := NewUint(5898)
		So(test1.MapToNillable(func(b uint) *uint { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewUint(5898)
		So(test2.MapToNillable(func(b uint) *uint { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyUint()
		So(func() {
			test3.MapToNillable(func(b uint) *uint { panic("foo") }).IsNil()
		}, ShouldNotPanic)
	})
}
