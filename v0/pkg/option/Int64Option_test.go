// Package option_test generated @ 2020-05-17 20:31:35.070899067 -0400 EDT m=+0.000888110
package option_test

import (
	"testing"

	. "github.com/Foxcapades/goop/v0/pkg/option"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewInt64(t *testing.T) {
	Convey("NewInt64", t, func() {
		test := NewInt64(55555)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyInt64(t *testing.T) {
	Convey("NewEmptyInt64", t, func() {
		test := NewEmptyInt64()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeInt64(t *testing.T) {
	Convey("NewMaybeInt64", t, func() {
		var val1 int64

		test1 := NewMaybeInt64(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeInt64(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestInt64_Get(t *testing.T) {
	Convey("Int64.Get", t, func() {
		test1 := NewInt64(55555)
		So(test1.Get(), ShouldEqual, 55555)

		test2 := NewEmptyInt64()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestInt64_Or(t *testing.T) {
	Convey("Int64.Or", t, func() {
		test1 := NewInt64(55555)
		So(test1.Or(99999), ShouldEqual, 55555)

		test2 := NewEmptyInt64()
		So(test2.Or(55555), ShouldEqual, 55555)
	})
}

func TestInt64_OrGet(t *testing.T) {
	Convey("Int64.OrGet", t, func() {
		test1 := NewInt64(55555)
		So(test1.OrGet(func() int64 { return 99999 }), ShouldEqual, 55555)

		test2 := NewEmptyInt64()
		So(test2.OrGet(func() int64 { return 55555 }), ShouldEqual, 55555)
	})
}

func TestInt64_OrPanicWith(t *testing.T) {
	Convey("Int64.OrPanicWith", t, func() {
		test1 := NewInt64(55555)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 55555)

		test2 := NewEmptyInt64()
		So(func() {
			test2.OrPanicWith(func() interface{} { return "test value" })
		}, ShouldPanicWith, "test value")
	})
}

func TestInt64_MapToNillable(t *testing.T) {
	Convey("Int64.MapToNillable", t, func() {
		test1 := NewInt64(55555)
		So(test1.MapToNillable(func(b int64) *int64 { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewInt64(55555)
		So(test2.MapToNillable(func(b int64) *int64 { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyInt64()
		So(func() {
			test3.MapToNillable(func(b int64) *int64 { panic("foo") }).IsNil()
		}, ShouldNotPanic)
	})
}
