// Package option_test generated @ 2020-05-17 22:19:55.164878293 -0400 EDT m=+0.001439508
package option_test

import (
	"errors"
	"testing"

	. "github.com/Foxcapades/goop/v0/pkg/option"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewComplex128(t *testing.T) {
	Convey("NewComplex128", t, func() {
		test := NewComplex128(58.98)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyComplex128(t *testing.T) {
	Convey("NewEmptyComplex128", t, func() {
		test := NewEmptyComplex128()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeComplex128(t *testing.T) {
	Convey("NewMaybeComplex128", t, func() {
		var val1 complex128

		test1 := NewMaybeComplex128(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeComplex128(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestComplex128_Get(t *testing.T) {
	Convey("Complex128.Get", t, func() {
		test1 := NewComplex128(58.98)
		So(test1.Get(), ShouldEqual, 58.98)

		test2 := NewEmptyComplex128()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestComplex128_Or(t *testing.T) {
	Convey("Complex128.Or", t, func() {
		test1 := NewComplex128(58.98)
		So(test1.Or(666.666), ShouldEqual, 58.98)

		test2 := NewEmptyComplex128()
		So(test2.Or(58.98), ShouldEqual, 58.98)
	})
}

func TestComplex128_OrGet(t *testing.T) {
	Convey("Complex128.OrGet", t, func() {
		test1 := NewComplex128(58.98)
		So(test1.OrGet(func() complex128 { return 666.666 }), ShouldEqual, 58.98)

		test2 := NewEmptyComplex128()
		So(test2.OrGet(func() complex128 { return 58.98 }), ShouldEqual, 58.98)
	})
}

func TestComplex128_OrPanicWith(t *testing.T) {
	Convey("Complex128.OrPanicWith", t, func() {
		test1 := NewComplex128(58.98)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 58.98)

		test2 := NewEmptyComplex128()
		So(func() {
			test2.OrPanicWith(func() interface{} { return "test value" })
		}, ShouldPanicWith, "test value")
	})
}

func TestComplex128_MapToNillable(t *testing.T) {
	Convey("Complex128.MapToNillable", t, func() {
		test1 := NewComplex128(58.98)
		So(test1.MapToNillable(func(b complex128) *complex128 { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewComplex128(58.98)
		So(test2.MapToNillable(func(b complex128) *complex128 { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyComplex128()
		So(func() {
			test3.MapToNillable(func(b complex128) *complex128 { panic("foo") }).IsNil()
		}, ShouldNotPanic)
	})
}

func TestComplex128_ForceMapToBool(t *testing.T) {
	Convey("Complex128.ForceMapToBool", t, func() {
		So(NewComplex128(58.98).ForceMapToBool(func(complex128) bool { return true }).Get(),
			ShouldEqual, true)

		So(NewEmptyComplex128().ForceMapToBool(func(complex128) bool { return true }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToNillableBool(t *testing.T) {
	Convey("Complex128.MapToNillableBool", t, func() {
		So(NewComplex128(58.98).MapToNillableBool(func(complex128) *bool {
			var tmp bool = true
			return &tmp
		}).Get(), ShouldEqual, true)

		So(NewEmptyComplex128().MapToNillableBool(func(complex128) *bool {
			var tmp bool = true
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex128(58.98).MapToNillableBool(func(complex128) *bool { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToBool(t *testing.T) {
	Convey("Complex128.MapToBool", t, func() {
		test1, err1 := NewComplex128(58.98).MapToBool(func(complex128) (bool, error) {
			return true, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, true)

		test2, err2 := NewEmptyComplex128().MapToBool(func(complex128) (bool, error) {
			return true, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex128(58.98).MapToBool(func(complex128) (bool, error) {
			return true, errors.New("bool error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "bool error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex128_ForceMapToByte(t *testing.T) {
	Convey("Complex128.ForceMapToByte", t, func() {
		So(NewComplex128(58.98).ForceMapToByte(func(complex128) byte { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyComplex128().ForceMapToByte(func(complex128) byte { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToNillableByte(t *testing.T) {
	Convey("Complex128.MapToNillableByte", t, func() {
		So(NewComplex128(58.98).MapToNillableByte(func(complex128) *byte {
			var tmp byte = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyComplex128().MapToNillableByte(func(complex128) *byte {
			var tmp byte = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex128(58.98).MapToNillableByte(func(complex128) *byte { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToByte(t *testing.T) {
	Convey("Complex128.MapToByte", t, func() {
		test1, err1 := NewComplex128(58.98).MapToByte(func(complex128) (byte, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyComplex128().MapToByte(func(complex128) (byte, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex128(58.98).MapToByte(func(complex128) (byte, error) {
			return 222, errors.New("byte error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "byte error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex128_ForceMapToComplex64(t *testing.T) {
	Convey("Complex128.ForceMapToComplex64", t, func() {
		So(NewComplex128(58.98).ForceMapToComplex64(func(complex128) complex64 { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyComplex128().ForceMapToComplex64(func(complex128) complex64 { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToNillableComplex64(t *testing.T) {
	Convey("Complex128.MapToNillableComplex64", t, func() {
		So(NewComplex128(58.98).MapToNillableComplex64(func(complex128) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyComplex128().MapToNillableComplex64(func(complex128) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex128(58.98).MapToNillableComplex64(func(complex128) *complex64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToComplex64(t *testing.T) {
	Convey("Complex128.MapToComplex64", t, func() {
		test1, err1 := NewComplex128(58.98).MapToComplex64(func(complex128) (complex64, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyComplex128().MapToComplex64(func(complex128) (complex64, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex128(58.98).MapToComplex64(func(complex128) (complex64, error) {
			return 123456, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex128_ForceMapToFloat32(t *testing.T) {
	Convey("Complex128.ForceMapToFloat32", t, func() {
		So(NewComplex128(58.98).ForceMapToFloat32(func(complex128) float32 { return 32.3 }).Get(),
			ShouldEqual, 32.3)

		So(NewEmptyComplex128().ForceMapToFloat32(func(complex128) float32 { return 32.3 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToNillableFloat32(t *testing.T) {
	Convey("Complex128.MapToNillableFloat32", t, func() {
		So(NewComplex128(58.98).MapToNillableFloat32(func(complex128) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).Get(), ShouldEqual, 32.3)

		So(NewEmptyComplex128().MapToNillableFloat32(func(complex128) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex128(58.98).MapToNillableFloat32(func(complex128) *float32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToFloat32(t *testing.T) {
	Convey("Complex128.MapToFloat32", t, func() {
		test1, err1 := NewComplex128(58.98).MapToFloat32(func(complex128) (float32, error) {
			return 32.3, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 32.3)

		test2, err2 := NewEmptyComplex128().MapToFloat32(func(complex128) (float32, error) {
			return 32.3, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex128(58.98).MapToFloat32(func(complex128) (float32, error) {
			return 32.3, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex128_ForceMapToFloat64(t *testing.T) {
	Convey("Complex128.ForceMapToFloat64", t, func() {
		So(NewComplex128(58.98).ForceMapToFloat64(func(complex128) float64 { return 3.2455 }).Get(),
			ShouldEqual, 3.2455)

		So(NewEmptyComplex128().ForceMapToFloat64(func(complex128) float64 { return 3.2455 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToNillableFloat64(t *testing.T) {
	Convey("Complex128.MapToNillableFloat64", t, func() {
		So(NewComplex128(58.98).MapToNillableFloat64(func(complex128) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).Get(), ShouldEqual, 3.2455)

		So(NewEmptyComplex128().MapToNillableFloat64(func(complex128) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex128(58.98).MapToNillableFloat64(func(complex128) *float64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToFloat64(t *testing.T) {
	Convey("Complex128.MapToFloat64", t, func() {
		test1, err1 := NewComplex128(58.98).MapToFloat64(func(complex128) (float64, error) {
			return 3.2455, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 3.2455)

		test2, err2 := NewEmptyComplex128().MapToFloat64(func(complex128) (float64, error) {
			return 3.2455, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex128(58.98).MapToFloat64(func(complex128) (float64, error) {
			return 3.2455, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex128_ForceMapToInt(t *testing.T) {
	Convey("Complex128.ForceMapToInt", t, func() {
		So(NewComplex128(58.98).ForceMapToInt(func(complex128) int { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyComplex128().ForceMapToInt(func(complex128) int { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToNillableInt(t *testing.T) {
	Convey("Complex128.MapToNillableInt", t, func() {
		So(NewComplex128(58.98).MapToNillableInt(func(complex128) *int {
			var tmp int = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyComplex128().MapToNillableInt(func(complex128) *int {
			var tmp int = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex128(58.98).MapToNillableInt(func(complex128) *int { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToInt(t *testing.T) {
	Convey("Complex128.MapToInt", t, func() {
		test1, err1 := NewComplex128(58.98).MapToInt(func(complex128) (int, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyComplex128().MapToInt(func(complex128) (int, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex128(58.98).MapToInt(func(complex128) (int, error) {
			return 123456, errors.New("int error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex128_ForceMapToInt16(t *testing.T) {
	Convey("Complex128.ForceMapToInt16", t, func() {
		So(NewComplex128(58.98).ForceMapToInt16(func(complex128) int16 { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyComplex128().ForceMapToInt16(func(complex128) int16 { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToNillableInt16(t *testing.T) {
	Convey("Complex128.MapToNillableInt16", t, func() {
		So(NewComplex128(58.98).MapToNillableInt16(func(complex128) *int16 {
			var tmp int16 = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyComplex128().MapToNillableInt16(func(complex128) *int16 {
			var tmp int16 = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex128(58.98).MapToNillableInt16(func(complex128) *int16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToInt16(t *testing.T) {
	Convey("Complex128.MapToInt16", t, func() {
		test1, err1 := NewComplex128(58.98).MapToInt16(func(complex128) (int16, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyComplex128().MapToInt16(func(complex128) (int16, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex128(58.98).MapToInt16(func(complex128) (int16, error) {
			return 222, errors.New("int16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex128_ForceMapToInt32(t *testing.T) {
	Convey("Complex128.ForceMapToInt32", t, func() {
		So(NewComplex128(58.98).ForceMapToInt32(func(complex128) int32 { return 333 }).Get(),
			ShouldEqual, 333)

		So(NewEmptyComplex128().ForceMapToInt32(func(complex128) int32 { return 333 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToNillableInt32(t *testing.T) {
	Convey("Complex128.MapToNillableInt32", t, func() {
		So(NewComplex128(58.98).MapToNillableInt32(func(complex128) *int32 {
			var tmp int32 = 333
			return &tmp
		}).Get(), ShouldEqual, 333)

		So(NewEmptyComplex128().MapToNillableInt32(func(complex128) *int32 {
			var tmp int32 = 333
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex128(58.98).MapToNillableInt32(func(complex128) *int32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToInt32(t *testing.T) {
	Convey("Complex128.MapToInt32", t, func() {
		test1, err1 := NewComplex128(58.98).MapToInt32(func(complex128) (int32, error) {
			return 333, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 333)

		test2, err2 := NewEmptyComplex128().MapToInt32(func(complex128) (int32, error) {
			return 333, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex128(58.98).MapToInt32(func(complex128) (int32, error) {
			return 333, errors.New("int32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex128_ForceMapToInt64(t *testing.T) {
	Convey("Complex128.ForceMapToInt64", t, func() {
		So(NewComplex128(58.98).ForceMapToInt64(func(complex128) int64 { return 444 }).Get(),
			ShouldEqual, 444)

		So(NewEmptyComplex128().ForceMapToInt64(func(complex128) int64 { return 444 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToNillableInt64(t *testing.T) {
	Convey("Complex128.MapToNillableInt64", t, func() {
		So(NewComplex128(58.98).MapToNillableInt64(func(complex128) *int64 {
			var tmp int64 = 444
			return &tmp
		}).Get(), ShouldEqual, 444)

		So(NewEmptyComplex128().MapToNillableInt64(func(complex128) *int64 {
			var tmp int64 = 444
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex128(58.98).MapToNillableInt64(func(complex128) *int64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToInt64(t *testing.T) {
	Convey("Complex128.MapToInt64", t, func() {
		test1, err1 := NewComplex128(58.98).MapToInt64(func(complex128) (int64, error) {
			return 444, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 444)

		test2, err2 := NewEmptyComplex128().MapToInt64(func(complex128) (int64, error) {
			return 444, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex128(58.98).MapToInt64(func(complex128) (int64, error) {
			return 444, errors.New("int64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex128_ForceMapToInt8(t *testing.T) {
	Convey("Complex128.ForceMapToInt8", t, func() {
		So(NewComplex128(58.98).ForceMapToInt8(func(complex128) int8 { return 111 }).Get(),
			ShouldEqual, 111)

		So(NewEmptyComplex128().ForceMapToInt8(func(complex128) int8 { return 111 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToNillableInt8(t *testing.T) {
	Convey("Complex128.MapToNillableInt8", t, func() {
		So(NewComplex128(58.98).MapToNillableInt8(func(complex128) *int8 {
			var tmp int8 = 111
			return &tmp
		}).Get(), ShouldEqual, 111)

		So(NewEmptyComplex128().MapToNillableInt8(func(complex128) *int8 {
			var tmp int8 = 111
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex128(58.98).MapToNillableInt8(func(complex128) *int8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToInt8(t *testing.T) {
	Convey("Complex128.MapToInt8", t, func() {
		test1, err1 := NewComplex128(58.98).MapToInt8(func(complex128) (int8, error) {
			return 111, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 111)

		test2, err2 := NewEmptyComplex128().MapToInt8(func(complex128) (int8, error) {
			return 111, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex128(58.98).MapToInt8(func(complex128) (int8, error) {
			return 111, errors.New("int8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex128_ForceMapToUntyped(t *testing.T) {
	Convey("Complex128.ForceMapToUntyped", t, func() {
		So(NewComplex128(58.98).ForceMapToUntyped(func(complex128) interface{} { return "111111" }).Get(),
			ShouldEqual, "111111")

		So(NewEmptyComplex128().ForceMapToUntyped(func(complex128) interface{} { return "111111" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToNillableUntyped(t *testing.T) {
	Convey("Complex128.MapToNillableUntyped", t, func() {
		So(NewComplex128(58.98).MapToNillableUntyped(func(complex128) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).Get(), ShouldEqual, "111111")

		So(NewEmptyComplex128().MapToNillableUntyped(func(complex128) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex128(58.98).MapToNillableUntyped(func(complex128) *interface{} { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToUntyped(t *testing.T) {
	Convey("Complex128.MapToUntyped", t, func() {
		test1, err1 := NewComplex128(58.98).MapToUntyped(func(complex128) (interface{}, error) {
			return "111111", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "111111")

		test2, err2 := NewEmptyComplex128().MapToUntyped(func(complex128) (interface{}, error) {
			return "111111", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex128(58.98).MapToUntyped(func(complex128) (interface{}, error) {
			return "111111", errors.New("interface error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "interface error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex128_ForceMapToRune(t *testing.T) {
	Convey("Complex128.ForceMapToRune", t, func() {
		So(NewComplex128(58.98).ForceMapToRune(func(complex128) rune { return 555 }).Get(),
			ShouldEqual, 555)

		So(NewEmptyComplex128().ForceMapToRune(func(complex128) rune { return 555 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToNillableRune(t *testing.T) {
	Convey("Complex128.MapToNillableRune", t, func() {
		So(NewComplex128(58.98).MapToNillableRune(func(complex128) *rune {
			var tmp rune = 555
			return &tmp
		}).Get(), ShouldEqual, 555)

		So(NewEmptyComplex128().MapToNillableRune(func(complex128) *rune {
			var tmp rune = 555
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex128(58.98).MapToNillableRune(func(complex128) *rune { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToRune(t *testing.T) {
	Convey("Complex128.MapToRune", t, func() {
		test1, err1 := NewComplex128(58.98).MapToRune(func(complex128) (rune, error) {
			return 555, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 555)

		test2, err2 := NewEmptyComplex128().MapToRune(func(complex128) (rune, error) {
			return 555, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex128(58.98).MapToRune(func(complex128) (rune, error) {
			return 555, errors.New("rune error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "rune error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex128_ForceMapToString(t *testing.T) {
	Convey("Complex128.ForceMapToString", t, func() {
		So(NewComplex128(58.98).ForceMapToString(func(complex128) string { return "bonanza" }).Get(),
			ShouldEqual, "bonanza")

		So(NewEmptyComplex128().ForceMapToString(func(complex128) string { return "bonanza" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToNillableString(t *testing.T) {
	Convey("Complex128.MapToNillableString", t, func() {
		So(NewComplex128(58.98).MapToNillableString(func(complex128) *string {
			var tmp string = "bonanza"
			return &tmp
		}).Get(), ShouldEqual, "bonanza")

		So(NewEmptyComplex128().MapToNillableString(func(complex128) *string {
			var tmp string = "bonanza"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex128(58.98).MapToNillableString(func(complex128) *string { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToString(t *testing.T) {
	Convey("Complex128.MapToString", t, func() {
		test1, err1 := NewComplex128(58.98).MapToString(func(complex128) (string, error) {
			return "bonanza", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "bonanza")

		test2, err2 := NewEmptyComplex128().MapToString(func(complex128) (string, error) {
			return "bonanza", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex128(58.98).MapToString(func(complex128) (string, error) {
			return "bonanza", errors.New("string error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "string error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex128_ForceMapToUint(t *testing.T) {
	Convey("Complex128.ForceMapToUint", t, func() {
		So(NewComplex128(58.98).ForceMapToUint(func(complex128) uint { return 666 }).Get(),
			ShouldEqual, 666)

		So(NewEmptyComplex128().ForceMapToUint(func(complex128) uint { return 666 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToNillableUint(t *testing.T) {
	Convey("Complex128.MapToNillableUint", t, func() {
		So(NewComplex128(58.98).MapToNillableUint(func(complex128) *uint {
			var tmp uint = 666
			return &tmp
		}).Get(), ShouldEqual, 666)

		So(NewEmptyComplex128().MapToNillableUint(func(complex128) *uint {
			var tmp uint = 666
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex128(58.98).MapToNillableUint(func(complex128) *uint { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToUint(t *testing.T) {
	Convey("Complex128.MapToUint", t, func() {
		test1, err1 := NewComplex128(58.98).MapToUint(func(complex128) (uint, error) {
			return 666, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 666)

		test2, err2 := NewEmptyComplex128().MapToUint(func(complex128) (uint, error) {
			return 666, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex128(58.98).MapToUint(func(complex128) (uint, error) {
			return 666, errors.New("uint error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex128_ForceMapToUint16(t *testing.T) {
	Convey("Complex128.ForceMapToUint16", t, func() {
		So(NewComplex128(58.98).ForceMapToUint16(func(complex128) uint16 { return 888 }).Get(),
			ShouldEqual, 888)

		So(NewEmptyComplex128().ForceMapToUint16(func(complex128) uint16 { return 888 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToNillableUint16(t *testing.T) {
	Convey("Complex128.MapToNillableUint16", t, func() {
		So(NewComplex128(58.98).MapToNillableUint16(func(complex128) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).Get(), ShouldEqual, 888)

		So(NewEmptyComplex128().MapToNillableUint16(func(complex128) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex128(58.98).MapToNillableUint16(func(complex128) *uint16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToUint16(t *testing.T) {
	Convey("Complex128.MapToUint16", t, func() {
		test1, err1 := NewComplex128(58.98).MapToUint16(func(complex128) (uint16, error) {
			return 888, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 888)

		test2, err2 := NewEmptyComplex128().MapToUint16(func(complex128) (uint16, error) {
			return 888, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex128(58.98).MapToUint16(func(complex128) (uint16, error) {
			return 888, errors.New("uint16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex128_ForceMapToUint32(t *testing.T) {
	Convey("Complex128.ForceMapToUint32", t, func() {
		So(NewComplex128(58.98).ForceMapToUint32(func(complex128) uint32 { return 999 }).Get(),
			ShouldEqual, 999)

		So(NewEmptyComplex128().ForceMapToUint32(func(complex128) uint32 { return 999 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToNillableUint32(t *testing.T) {
	Convey("Complex128.MapToNillableUint32", t, func() {
		So(NewComplex128(58.98).MapToNillableUint32(func(complex128) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).Get(), ShouldEqual, 999)

		So(NewEmptyComplex128().MapToNillableUint32(func(complex128) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex128(58.98).MapToNillableUint32(func(complex128) *uint32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToUint32(t *testing.T) {
	Convey("Complex128.MapToUint32", t, func() {
		test1, err1 := NewComplex128(58.98).MapToUint32(func(complex128) (uint32, error) {
			return 999, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 999)

		test2, err2 := NewEmptyComplex128().MapToUint32(func(complex128) (uint32, error) {
			return 999, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex128(58.98).MapToUint32(func(complex128) (uint32, error) {
			return 999, errors.New("uint32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex128_ForceMapToUint64(t *testing.T) {
	Convey("Complex128.ForceMapToUint64", t, func() {
		So(NewComplex128(58.98).ForceMapToUint64(func(complex128) uint64 { return 101010 }).Get(),
			ShouldEqual, 101010)

		So(NewEmptyComplex128().ForceMapToUint64(func(complex128) uint64 { return 101010 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToNillableUint64(t *testing.T) {
	Convey("Complex128.MapToNillableUint64", t, func() {
		So(NewComplex128(58.98).MapToNillableUint64(func(complex128) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).Get(), ShouldEqual, 101010)

		So(NewEmptyComplex128().MapToNillableUint64(func(complex128) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex128(58.98).MapToNillableUint64(func(complex128) *uint64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToUint64(t *testing.T) {
	Convey("Complex128.MapToUint64", t, func() {
		test1, err1 := NewComplex128(58.98).MapToUint64(func(complex128) (uint64, error) {
			return 101010, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 101010)

		test2, err2 := NewEmptyComplex128().MapToUint64(func(complex128) (uint64, error) {
			return 101010, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex128(58.98).MapToUint64(func(complex128) (uint64, error) {
			return 101010, errors.New("uint64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex128_ForceMapToUint8(t *testing.T) {
	Convey("Complex128.ForceMapToUint8", t, func() {
		So(NewComplex128(58.98).ForceMapToUint8(func(complex128) uint8 { return 77 }).Get(),
			ShouldEqual, 77)

		So(NewEmptyComplex128().ForceMapToUint8(func(complex128) uint8 { return 77 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToNillableUint8(t *testing.T) {
	Convey("Complex128.MapToNillableUint8", t, func() {
		So(NewComplex128(58.98).MapToNillableUint8(func(complex128) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).Get(), ShouldEqual, 77)

		So(NewEmptyComplex128().MapToNillableUint8(func(complex128) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex128(58.98).MapToNillableUint8(func(complex128) *uint8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex128_MapToUint8(t *testing.T) {
	Convey("Complex128.MapToUint8", t, func() {
		test1, err1 := NewComplex128(58.98).MapToUint8(func(complex128) (uint8, error) {
			return 77, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 77)

		test2, err2 := NewEmptyComplex128().MapToUint8(func(complex128) (uint8, error) {
			return 77, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex128(58.98).MapToUint8(func(complex128) (uint8, error) {
			return 77, errors.New("uint8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}
