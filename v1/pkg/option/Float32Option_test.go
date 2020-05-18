// Package option_test generated @ 2020-05-17 22:19:55.164878293 -0400 EDT m=+0.001439508
package option_test

import (
	"errors"
	"testing"

	. "github.com/Foxcapades/goop/v0/pkg/option"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewFloat32(t *testing.T) {
	Convey("NewFloat32", t, func() {
		test := NewFloat32(28.3)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyFloat32(t *testing.T) {
	Convey("NewEmptyFloat32", t, func() {
		test := NewEmptyFloat32()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeFloat32(t *testing.T) {
	Convey("NewMaybeFloat32", t, func() {
		var val1 float32

		test1 := NewMaybeFloat32(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeFloat32(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestFloat32_Get(t *testing.T) {
	Convey("Float32.Get", t, func() {
		test1 := NewFloat32(28.3)
		So(test1.Get(), ShouldEqual, 28.3)

		test2 := NewEmptyFloat32()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestFloat32_Or(t *testing.T) {
	Convey("Float32.Or", t, func() {
		test1 := NewFloat32(28.3)
		So(test1.Or(69.9), ShouldEqual, 28.3)

		test2 := NewEmptyFloat32()
		So(test2.Or(28.3), ShouldEqual, 28.3)
	})
}

func TestFloat32_OrGet(t *testing.T) {
	Convey("Float32.OrGet", t, func() {
		test1 := NewFloat32(28.3)
		So(test1.OrGet(func() float32 { return 69.9 }), ShouldEqual, 28.3)

		test2 := NewEmptyFloat32()
		So(test2.OrGet(func() float32 { return 28.3 }), ShouldEqual, 28.3)
	})
}

func TestFloat32_OrPanicWith(t *testing.T) {
	Convey("Float32.OrPanicWith", t, func() {
		test1 := NewFloat32(28.3)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 28.3)

		test2 := NewEmptyFloat32()
		So(func() {
			test2.OrPanicWith(func() interface{} { return "test value" })
		}, ShouldPanicWith, "test value")
	})
}

func TestFloat32_MapToNillable(t *testing.T) {
	Convey("Float32.MapToNillable", t, func() {
		test1 := NewFloat32(28.3)
		So(test1.MapToNillable(func(b float32) *float32 { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewFloat32(28.3)
		So(test2.MapToNillable(func(b float32) *float32 { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyFloat32()
		So(func() {
			test3.MapToNillable(func(b float32) *float32 { panic("foo") }).IsNil()
		}, ShouldNotPanic)
	})
}

func TestFloat32_ForceMapToBool(t *testing.T) {
	Convey("Float32.ForceMapToBool", t, func() {
		So(NewFloat32(28.3).ForceMapToBool(func(float32) bool { return true }).Get(),
			ShouldEqual, true)

		So(NewEmptyFloat32().ForceMapToBool(func(float32) bool { return true }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToNillableBool(t *testing.T) {
	Convey("Float32.MapToNillableBool", t, func() {
		So(NewFloat32(28.3).MapToNillableBool(func(float32) *bool {
			var tmp bool = true
			return &tmp
		}).Get(), ShouldEqual, true)

		So(NewEmptyFloat32().MapToNillableBool(func(float32) *bool {
			var tmp bool = true
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat32(28.3).MapToNillableBool(func(float32) *bool { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToBool(t *testing.T) {
	Convey("Float32.MapToBool", t, func() {
		test1, err1 := NewFloat32(28.3).MapToBool(func(float32) (bool, error) {
			return true, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, true)

		test2, err2 := NewEmptyFloat32().MapToBool(func(float32) (bool, error) {
			return true, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat32(28.3).MapToBool(func(float32) (bool, error) {
			return true, errors.New("bool error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "bool error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat32_ForceMapToByte(t *testing.T) {
	Convey("Float32.ForceMapToByte", t, func() {
		So(NewFloat32(28.3).ForceMapToByte(func(float32) byte { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyFloat32().ForceMapToByte(func(float32) byte { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToNillableByte(t *testing.T) {
	Convey("Float32.MapToNillableByte", t, func() {
		So(NewFloat32(28.3).MapToNillableByte(func(float32) *byte {
			var tmp byte = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyFloat32().MapToNillableByte(func(float32) *byte {
			var tmp byte = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat32(28.3).MapToNillableByte(func(float32) *byte { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToByte(t *testing.T) {
	Convey("Float32.MapToByte", t, func() {
		test1, err1 := NewFloat32(28.3).MapToByte(func(float32) (byte, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyFloat32().MapToByte(func(float32) (byte, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat32(28.3).MapToByte(func(float32) (byte, error) {
			return 222, errors.New("byte error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "byte error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat32_ForceMapToComplex128(t *testing.T) {
	Convey("Float32.ForceMapToComplex128", t, func() {
		So(NewFloat32(28.3).ForceMapToComplex128(func(float32) complex128 { return 987654 }).Get(),
			ShouldEqual, 987654)

		So(NewEmptyFloat32().ForceMapToComplex128(func(float32) complex128 { return 987654 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToNillableComplex128(t *testing.T) {
	Convey("Float32.MapToNillableComplex128", t, func() {
		So(NewFloat32(28.3).MapToNillableComplex128(func(float32) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).Get(), ShouldEqual, 987654)

		So(NewEmptyFloat32().MapToNillableComplex128(func(float32) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat32(28.3).MapToNillableComplex128(func(float32) *complex128 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToComplex128(t *testing.T) {
	Convey("Float32.MapToComplex128", t, func() {
		test1, err1 := NewFloat32(28.3).MapToComplex128(func(float32) (complex128, error) {
			return 987654, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 987654)

		test2, err2 := NewEmptyFloat32().MapToComplex128(func(float32) (complex128, error) {
			return 987654, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat32(28.3).MapToComplex128(func(float32) (complex128, error) {
			return 987654, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat32_ForceMapToComplex64(t *testing.T) {
	Convey("Float32.ForceMapToComplex64", t, func() {
		So(NewFloat32(28.3).ForceMapToComplex64(func(float32) complex64 { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyFloat32().ForceMapToComplex64(func(float32) complex64 { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToNillableComplex64(t *testing.T) {
	Convey("Float32.MapToNillableComplex64", t, func() {
		So(NewFloat32(28.3).MapToNillableComplex64(func(float32) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyFloat32().MapToNillableComplex64(func(float32) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat32(28.3).MapToNillableComplex64(func(float32) *complex64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToComplex64(t *testing.T) {
	Convey("Float32.MapToComplex64", t, func() {
		test1, err1 := NewFloat32(28.3).MapToComplex64(func(float32) (complex64, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyFloat32().MapToComplex64(func(float32) (complex64, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat32(28.3).MapToComplex64(func(float32) (complex64, error) {
			return 123456, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat32_ForceMapToFloat64(t *testing.T) {
	Convey("Float32.ForceMapToFloat64", t, func() {
		So(NewFloat32(28.3).ForceMapToFloat64(func(float32) float64 { return 3.2455 }).Get(),
			ShouldEqual, 3.2455)

		So(NewEmptyFloat32().ForceMapToFloat64(func(float32) float64 { return 3.2455 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToNillableFloat64(t *testing.T) {
	Convey("Float32.MapToNillableFloat64", t, func() {
		So(NewFloat32(28.3).MapToNillableFloat64(func(float32) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).Get(), ShouldEqual, 3.2455)

		So(NewEmptyFloat32().MapToNillableFloat64(func(float32) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat32(28.3).MapToNillableFloat64(func(float32) *float64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToFloat64(t *testing.T) {
	Convey("Float32.MapToFloat64", t, func() {
		test1, err1 := NewFloat32(28.3).MapToFloat64(func(float32) (float64, error) {
			return 3.2455, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 3.2455)

		test2, err2 := NewEmptyFloat32().MapToFloat64(func(float32) (float64, error) {
			return 3.2455, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat32(28.3).MapToFloat64(func(float32) (float64, error) {
			return 3.2455, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat32_ForceMapToInt(t *testing.T) {
	Convey("Float32.ForceMapToInt", t, func() {
		So(NewFloat32(28.3).ForceMapToInt(func(float32) int { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyFloat32().ForceMapToInt(func(float32) int { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToNillableInt(t *testing.T) {
	Convey("Float32.MapToNillableInt", t, func() {
		So(NewFloat32(28.3).MapToNillableInt(func(float32) *int {
			var tmp int = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyFloat32().MapToNillableInt(func(float32) *int {
			var tmp int = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat32(28.3).MapToNillableInt(func(float32) *int { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToInt(t *testing.T) {
	Convey("Float32.MapToInt", t, func() {
		test1, err1 := NewFloat32(28.3).MapToInt(func(float32) (int, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyFloat32().MapToInt(func(float32) (int, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat32(28.3).MapToInt(func(float32) (int, error) {
			return 123456, errors.New("int error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat32_ForceMapToInt16(t *testing.T) {
	Convey("Float32.ForceMapToInt16", t, func() {
		So(NewFloat32(28.3).ForceMapToInt16(func(float32) int16 { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyFloat32().ForceMapToInt16(func(float32) int16 { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToNillableInt16(t *testing.T) {
	Convey("Float32.MapToNillableInt16", t, func() {
		So(NewFloat32(28.3).MapToNillableInt16(func(float32) *int16 {
			var tmp int16 = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyFloat32().MapToNillableInt16(func(float32) *int16 {
			var tmp int16 = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat32(28.3).MapToNillableInt16(func(float32) *int16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToInt16(t *testing.T) {
	Convey("Float32.MapToInt16", t, func() {
		test1, err1 := NewFloat32(28.3).MapToInt16(func(float32) (int16, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyFloat32().MapToInt16(func(float32) (int16, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat32(28.3).MapToInt16(func(float32) (int16, error) {
			return 222, errors.New("int16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat32_ForceMapToInt32(t *testing.T) {
	Convey("Float32.ForceMapToInt32", t, func() {
		So(NewFloat32(28.3).ForceMapToInt32(func(float32) int32 { return 333 }).Get(),
			ShouldEqual, 333)

		So(NewEmptyFloat32().ForceMapToInt32(func(float32) int32 { return 333 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToNillableInt32(t *testing.T) {
	Convey("Float32.MapToNillableInt32", t, func() {
		So(NewFloat32(28.3).MapToNillableInt32(func(float32) *int32 {
			var tmp int32 = 333
			return &tmp
		}).Get(), ShouldEqual, 333)

		So(NewEmptyFloat32().MapToNillableInt32(func(float32) *int32 {
			var tmp int32 = 333
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat32(28.3).MapToNillableInt32(func(float32) *int32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToInt32(t *testing.T) {
	Convey("Float32.MapToInt32", t, func() {
		test1, err1 := NewFloat32(28.3).MapToInt32(func(float32) (int32, error) {
			return 333, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 333)

		test2, err2 := NewEmptyFloat32().MapToInt32(func(float32) (int32, error) {
			return 333, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat32(28.3).MapToInt32(func(float32) (int32, error) {
			return 333, errors.New("int32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat32_ForceMapToInt64(t *testing.T) {
	Convey("Float32.ForceMapToInt64", t, func() {
		So(NewFloat32(28.3).ForceMapToInt64(func(float32) int64 { return 444 }).Get(),
			ShouldEqual, 444)

		So(NewEmptyFloat32().ForceMapToInt64(func(float32) int64 { return 444 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToNillableInt64(t *testing.T) {
	Convey("Float32.MapToNillableInt64", t, func() {
		So(NewFloat32(28.3).MapToNillableInt64(func(float32) *int64 {
			var tmp int64 = 444
			return &tmp
		}).Get(), ShouldEqual, 444)

		So(NewEmptyFloat32().MapToNillableInt64(func(float32) *int64 {
			var tmp int64 = 444
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat32(28.3).MapToNillableInt64(func(float32) *int64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToInt64(t *testing.T) {
	Convey("Float32.MapToInt64", t, func() {
		test1, err1 := NewFloat32(28.3).MapToInt64(func(float32) (int64, error) {
			return 444, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 444)

		test2, err2 := NewEmptyFloat32().MapToInt64(func(float32) (int64, error) {
			return 444, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat32(28.3).MapToInt64(func(float32) (int64, error) {
			return 444, errors.New("int64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat32_ForceMapToInt8(t *testing.T) {
	Convey("Float32.ForceMapToInt8", t, func() {
		So(NewFloat32(28.3).ForceMapToInt8(func(float32) int8 { return 111 }).Get(),
			ShouldEqual, 111)

		So(NewEmptyFloat32().ForceMapToInt8(func(float32) int8 { return 111 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToNillableInt8(t *testing.T) {
	Convey("Float32.MapToNillableInt8", t, func() {
		So(NewFloat32(28.3).MapToNillableInt8(func(float32) *int8 {
			var tmp int8 = 111
			return &tmp
		}).Get(), ShouldEqual, 111)

		So(NewEmptyFloat32().MapToNillableInt8(func(float32) *int8 {
			var tmp int8 = 111
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat32(28.3).MapToNillableInt8(func(float32) *int8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToInt8(t *testing.T) {
	Convey("Float32.MapToInt8", t, func() {
		test1, err1 := NewFloat32(28.3).MapToInt8(func(float32) (int8, error) {
			return 111, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 111)

		test2, err2 := NewEmptyFloat32().MapToInt8(func(float32) (int8, error) {
			return 111, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat32(28.3).MapToInt8(func(float32) (int8, error) {
			return 111, errors.New("int8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat32_ForceMapToUntyped(t *testing.T) {
	Convey("Float32.ForceMapToUntyped", t, func() {
		So(NewFloat32(28.3).ForceMapToUntyped(func(float32) interface{} { return "111111" }).Get(),
			ShouldEqual, "111111")

		So(NewEmptyFloat32().ForceMapToUntyped(func(float32) interface{} { return "111111" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToNillableUntyped(t *testing.T) {
	Convey("Float32.MapToNillableUntyped", t, func() {
		So(NewFloat32(28.3).MapToNillableUntyped(func(float32) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).Get(), ShouldEqual, "111111")

		So(NewEmptyFloat32().MapToNillableUntyped(func(float32) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat32(28.3).MapToNillableUntyped(func(float32) *interface{} { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToUntyped(t *testing.T) {
	Convey("Float32.MapToUntyped", t, func() {
		test1, err1 := NewFloat32(28.3).MapToUntyped(func(float32) (interface{}, error) {
			return "111111", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "111111")

		test2, err2 := NewEmptyFloat32().MapToUntyped(func(float32) (interface{}, error) {
			return "111111", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat32(28.3).MapToUntyped(func(float32) (interface{}, error) {
			return "111111", errors.New("interface error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "interface error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat32_ForceMapToRune(t *testing.T) {
	Convey("Float32.ForceMapToRune", t, func() {
		So(NewFloat32(28.3).ForceMapToRune(func(float32) rune { return 555 }).Get(),
			ShouldEqual, 555)

		So(NewEmptyFloat32().ForceMapToRune(func(float32) rune { return 555 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToNillableRune(t *testing.T) {
	Convey("Float32.MapToNillableRune", t, func() {
		So(NewFloat32(28.3).MapToNillableRune(func(float32) *rune {
			var tmp rune = 555
			return &tmp
		}).Get(), ShouldEqual, 555)

		So(NewEmptyFloat32().MapToNillableRune(func(float32) *rune {
			var tmp rune = 555
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat32(28.3).MapToNillableRune(func(float32) *rune { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToRune(t *testing.T) {
	Convey("Float32.MapToRune", t, func() {
		test1, err1 := NewFloat32(28.3).MapToRune(func(float32) (rune, error) {
			return 555, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 555)

		test2, err2 := NewEmptyFloat32().MapToRune(func(float32) (rune, error) {
			return 555, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat32(28.3).MapToRune(func(float32) (rune, error) {
			return 555, errors.New("rune error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "rune error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat32_ForceMapToString(t *testing.T) {
	Convey("Float32.ForceMapToString", t, func() {
		So(NewFloat32(28.3).ForceMapToString(func(float32) string { return "bonanza" }).Get(),
			ShouldEqual, "bonanza")

		So(NewEmptyFloat32().ForceMapToString(func(float32) string { return "bonanza" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToNillableString(t *testing.T) {
	Convey("Float32.MapToNillableString", t, func() {
		So(NewFloat32(28.3).MapToNillableString(func(float32) *string {
			var tmp string = "bonanza"
			return &tmp
		}).Get(), ShouldEqual, "bonanza")

		So(NewEmptyFloat32().MapToNillableString(func(float32) *string {
			var tmp string = "bonanza"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat32(28.3).MapToNillableString(func(float32) *string { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToString(t *testing.T) {
	Convey("Float32.MapToString", t, func() {
		test1, err1 := NewFloat32(28.3).MapToString(func(float32) (string, error) {
			return "bonanza", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "bonanza")

		test2, err2 := NewEmptyFloat32().MapToString(func(float32) (string, error) {
			return "bonanza", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat32(28.3).MapToString(func(float32) (string, error) {
			return "bonanza", errors.New("string error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "string error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat32_ForceMapToUint(t *testing.T) {
	Convey("Float32.ForceMapToUint", t, func() {
		So(NewFloat32(28.3).ForceMapToUint(func(float32) uint { return 666 }).Get(),
			ShouldEqual, 666)

		So(NewEmptyFloat32().ForceMapToUint(func(float32) uint { return 666 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToNillableUint(t *testing.T) {
	Convey("Float32.MapToNillableUint", t, func() {
		So(NewFloat32(28.3).MapToNillableUint(func(float32) *uint {
			var tmp uint = 666
			return &tmp
		}).Get(), ShouldEqual, 666)

		So(NewEmptyFloat32().MapToNillableUint(func(float32) *uint {
			var tmp uint = 666
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat32(28.3).MapToNillableUint(func(float32) *uint { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToUint(t *testing.T) {
	Convey("Float32.MapToUint", t, func() {
		test1, err1 := NewFloat32(28.3).MapToUint(func(float32) (uint, error) {
			return 666, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 666)

		test2, err2 := NewEmptyFloat32().MapToUint(func(float32) (uint, error) {
			return 666, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat32(28.3).MapToUint(func(float32) (uint, error) {
			return 666, errors.New("uint error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat32_ForceMapToUint16(t *testing.T) {
	Convey("Float32.ForceMapToUint16", t, func() {
		So(NewFloat32(28.3).ForceMapToUint16(func(float32) uint16 { return 888 }).Get(),
			ShouldEqual, 888)

		So(NewEmptyFloat32().ForceMapToUint16(func(float32) uint16 { return 888 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToNillableUint16(t *testing.T) {
	Convey("Float32.MapToNillableUint16", t, func() {
		So(NewFloat32(28.3).MapToNillableUint16(func(float32) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).Get(), ShouldEqual, 888)

		So(NewEmptyFloat32().MapToNillableUint16(func(float32) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat32(28.3).MapToNillableUint16(func(float32) *uint16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToUint16(t *testing.T) {
	Convey("Float32.MapToUint16", t, func() {
		test1, err1 := NewFloat32(28.3).MapToUint16(func(float32) (uint16, error) {
			return 888, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 888)

		test2, err2 := NewEmptyFloat32().MapToUint16(func(float32) (uint16, error) {
			return 888, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat32(28.3).MapToUint16(func(float32) (uint16, error) {
			return 888, errors.New("uint16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat32_ForceMapToUint32(t *testing.T) {
	Convey("Float32.ForceMapToUint32", t, func() {
		So(NewFloat32(28.3).ForceMapToUint32(func(float32) uint32 { return 999 }).Get(),
			ShouldEqual, 999)

		So(NewEmptyFloat32().ForceMapToUint32(func(float32) uint32 { return 999 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToNillableUint32(t *testing.T) {
	Convey("Float32.MapToNillableUint32", t, func() {
		So(NewFloat32(28.3).MapToNillableUint32(func(float32) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).Get(), ShouldEqual, 999)

		So(NewEmptyFloat32().MapToNillableUint32(func(float32) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat32(28.3).MapToNillableUint32(func(float32) *uint32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToUint32(t *testing.T) {
	Convey("Float32.MapToUint32", t, func() {
		test1, err1 := NewFloat32(28.3).MapToUint32(func(float32) (uint32, error) {
			return 999, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 999)

		test2, err2 := NewEmptyFloat32().MapToUint32(func(float32) (uint32, error) {
			return 999, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat32(28.3).MapToUint32(func(float32) (uint32, error) {
			return 999, errors.New("uint32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat32_ForceMapToUint64(t *testing.T) {
	Convey("Float32.ForceMapToUint64", t, func() {
		So(NewFloat32(28.3).ForceMapToUint64(func(float32) uint64 { return 101010 }).Get(),
			ShouldEqual, 101010)

		So(NewEmptyFloat32().ForceMapToUint64(func(float32) uint64 { return 101010 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToNillableUint64(t *testing.T) {
	Convey("Float32.MapToNillableUint64", t, func() {
		So(NewFloat32(28.3).MapToNillableUint64(func(float32) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).Get(), ShouldEqual, 101010)

		So(NewEmptyFloat32().MapToNillableUint64(func(float32) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat32(28.3).MapToNillableUint64(func(float32) *uint64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToUint64(t *testing.T) {
	Convey("Float32.MapToUint64", t, func() {
		test1, err1 := NewFloat32(28.3).MapToUint64(func(float32) (uint64, error) {
			return 101010, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 101010)

		test2, err2 := NewEmptyFloat32().MapToUint64(func(float32) (uint64, error) {
			return 101010, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat32(28.3).MapToUint64(func(float32) (uint64, error) {
			return 101010, errors.New("uint64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat32_ForceMapToUint8(t *testing.T) {
	Convey("Float32.ForceMapToUint8", t, func() {
		So(NewFloat32(28.3).ForceMapToUint8(func(float32) uint8 { return 77 }).Get(),
			ShouldEqual, 77)

		So(NewEmptyFloat32().ForceMapToUint8(func(float32) uint8 { return 77 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToNillableUint8(t *testing.T) {
	Convey("Float32.MapToNillableUint8", t, func() {
		So(NewFloat32(28.3).MapToNillableUint8(func(float32) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).Get(), ShouldEqual, 77)

		So(NewEmptyFloat32().MapToNillableUint8(func(float32) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat32(28.3).MapToNillableUint8(func(float32) *uint8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat32_MapToUint8(t *testing.T) {
	Convey("Float32.MapToUint8", t, func() {
		test1, err1 := NewFloat32(28.3).MapToUint8(func(float32) (uint8, error) {
			return 77, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 77)

		test2, err2 := NewEmptyFloat32().MapToUint8(func(float32) (uint8, error) {
			return 77, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat32(28.3).MapToUint8(func(float32) (uint8, error) {
			return 77, errors.New("uint8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}
