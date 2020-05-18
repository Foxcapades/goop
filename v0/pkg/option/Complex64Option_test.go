// Package option_test generated @ 2020-05-17 20:31:35.070899067 -0400 EDT m=+0.000888110
package option_test

import (
	"testing"

	. "github.com/Foxcapades/goop/v0/pkg/option"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewComplex64(t *testing.T) {
	Convey("NewComplex64", t, func() {
		test := NewComplex64(2548)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyComplex64(t *testing.T) {
	Convey("NewEmptyComplex64", t, func() {
		test := NewEmptyComplex64()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeComplex64(t *testing.T) {
	Convey("NewMaybeComplex64", t, func() {
		var val1 complex64

		test1 := NewMaybeComplex64(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeComplex64(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestComplex64_Get(t *testing.T) {
	Convey("Complex64.Get", t, func() {
		test1 := NewComplex64(2548)
		So(test1.Get(), ShouldEqual, 2548)

		test2 := NewEmptyComplex64()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestComplex64_Or(t *testing.T) {
	Convey("Complex64.Or", t, func() {
		test1 := NewComplex64(2548)
		So(test1.Or(77.777), ShouldEqual, 2548)

		test2 := NewEmptyComplex64()
		So(test2.Or(2548), ShouldEqual, 2548)
	})
}

func TestComplex64_OrGet(t *testing.T) {
	Convey("Complex64.OrGet", t, func() {
		test1 := NewComplex64(2548)
		So(test1.OrGet(func() complex64 { return 77.777 }), ShouldEqual, 2548)

		test2 := NewEmptyComplex64()
		So(test2.OrGet(func() complex64 { return 2548 }), ShouldEqual, 2548)
	})
}

func TestComplex64_OrPanicWith(t *testing.T) {
	Convey("Complex64.OrPanicWith", t, func() {
		test1 := NewComplex64(2548)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 2548)

		test2 := NewEmptyComplex64()
		So(func() {
			test2.OrPanicWith(func() interface{} { return "test value" })
		}, ShouldPanicWith, "test value")
	})
}

func TestComplex64_MapToNillable(t *testing.T) {
	Convey("Complex64.MapToNillable", t, func() {
		test1 := NewComplex64(2548)
		So(test1.MapToNillable(func(b complex64) *complex64 { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewComplex64(2548)
		So(test2.MapToNillable(func(b complex64) *complex64 { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyComplex64()
		So(func() {
			test3.MapToNillable(func(b complex64) *complex64 { panic("foo") }).IsNil()
		}, ShouldNotPanic)
	})
}
