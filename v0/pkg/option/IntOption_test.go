// Package option_test generated @ 2020-05-17 20:31:35.070899067 -0400 EDT m=+0.000888110
package option_test

import (
	"testing"

	. "github.com/Foxcapades/goop/v0/pkg/option"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewInt(t *testing.T) {
	Convey("NewInt", t, func() {
		test := NewInt(123)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyInt(t *testing.T) {
	Convey("NewEmptyInt", t, func() {
		test := NewEmptyInt()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeInt(t *testing.T) {
	Convey("NewMaybeInt", t, func() {
		var val1 int

		test1 := NewMaybeInt(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeInt(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestInt_Get(t *testing.T) {
	Convey("Int.Get", t, func() {
		test1 := NewInt(123)
		So(test1.Get(), ShouldEqual, 123)

		test2 := NewEmptyInt()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestInt_Or(t *testing.T) {
	Convey("Int.Or", t, func() {
		test1 := NewInt(123)
		So(test1.Or(321), ShouldEqual, 123)

		test2 := NewEmptyInt()
		So(test2.Or(123), ShouldEqual, 123)
	})
}

func TestInt_OrGet(t *testing.T) {
	Convey("Int.OrGet", t, func() {
		test1 := NewInt(123)
		So(test1.OrGet(func() int { return 321 }), ShouldEqual, 123)

		test2 := NewEmptyInt()
		So(test2.OrGet(func() int { return 123 }), ShouldEqual, 123)
	})
}

func TestInt_OrPanicWith(t *testing.T) {
	Convey("Int.OrPanicWith", t, func() {
		test1 := NewInt(123)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 123)

		test2 := NewEmptyInt()
		So(func() {
			test2.OrPanicWith(func() interface{} { return "test value" })
		}, ShouldPanicWith, "test value")
	})
}

func TestInt_MapToNillable(t *testing.T) {
	Convey("Int.MapToNillable", t, func() {
		test1 := NewInt(123)
		So(test1.MapToNillable(func(b int) *int { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewInt(123)
		So(test2.MapToNillable(func(b int) *int { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyInt()
		So(func() {
			test3.MapToNillable(func(b int) *int { panic("foo") }).IsNil()
		}, ShouldNotPanic)
	})
}
