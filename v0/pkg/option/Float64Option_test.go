// Package option_test generated @ 2020-05-17 20:31:35.070899067 -0400 EDT m=+0.000888110
package option_test

import (
	"testing"

	. "github.com/Foxcapades/goop/v0/pkg/option"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewFloat64(t *testing.T) {
	Convey("NewFloat64", t, func() {
		test := NewFloat64(280.4)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyFloat64(t *testing.T) {
	Convey("NewEmptyFloat64", t, func() {
		test := NewEmptyFloat64()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeFloat64(t *testing.T) {
	Convey("NewMaybeFloat64", t, func() {
		var val1 float64

		test1 := NewMaybeFloat64(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeFloat64(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestFloat64_Get(t *testing.T) {
	Convey("Float64.Get", t, func() {
		test1 := NewFloat64(280.4)
		So(test1.Get(), ShouldEqual, 280.4)

		test2 := NewEmptyFloat64()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestFloat64_Or(t *testing.T) {
	Convey("Float64.Or", t, func() {
		test1 := NewFloat64(280.4)
		So(test1.Or(325.1118), ShouldEqual, 280.4)

		test2 := NewEmptyFloat64()
		So(test2.Or(280.4), ShouldEqual, 280.4)
	})
}

func TestFloat64_OrGet(t *testing.T) {
	Convey("Float64.OrGet", t, func() {
		test1 := NewFloat64(280.4)
		So(test1.OrGet(func() float64 { return 325.1118 }), ShouldEqual, 280.4)

		test2 := NewEmptyFloat64()
		So(test2.OrGet(func() float64 { return 280.4 }), ShouldEqual, 280.4)
	})
}

func TestFloat64_OrPanicWith(t *testing.T) {
	Convey("Float64.OrPanicWith", t, func() {
		test1 := NewFloat64(280.4)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 280.4)

		test2 := NewEmptyFloat64()
		So(func() {
			test2.OrPanicWith(func() interface{} { return "test value" })
		}, ShouldPanicWith, "test value")
	})
}

func TestFloat64_MapToNillable(t *testing.T) {
	Convey("Float64.MapToNillable", t, func() {
		test1 := NewFloat64(280.4)
		So(test1.MapToNillable(func(b float64) *float64 { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewFloat64(280.4)
		So(test2.MapToNillable(func(b float64) *float64 { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyFloat64()
		So(func() {
			test3.MapToNillable(func(b float64) *float64 { panic("foo") }).IsNil()
		}, ShouldNotPanic)
	})
}
