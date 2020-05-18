// Package option_test generated @ 2020-05-17 22:19:55.164878293 -0400 EDT m=+0.001439508
package option_test

import (
	"errors"
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

func TestFloat64_ForceMapToBool(t *testing.T) {
	Convey("Float64.ForceMapToBool", t, func() {
		So(NewFloat64(280.4).ForceMapToBool(func(float64) bool { return true }).Get(),
			ShouldEqual, true)

		So(NewEmptyFloat64().ForceMapToBool(func(float64) bool { return true }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToNillableBool(t *testing.T) {
	Convey("Float64.MapToNillableBool", t, func() {
		So(NewFloat64(280.4).MapToNillableBool(func(float64) *bool {
			var tmp bool = true
			return &tmp
		}).Get(), ShouldEqual, true)

		So(NewEmptyFloat64().MapToNillableBool(func(float64) *bool {
			var tmp bool = true
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat64(280.4).MapToNillableBool(func(float64) *bool { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToBool(t *testing.T) {
	Convey("Float64.MapToBool", t, func() {
		test1, err1 := NewFloat64(280.4).MapToBool(func(float64) (bool, error) {
			return true, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, true)

		test2, err2 := NewEmptyFloat64().MapToBool(func(float64) (bool, error) {
			return true, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat64(280.4).MapToBool(func(float64) (bool, error) {
			return true, errors.New("bool error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "bool error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat64_ForceMapToByte(t *testing.T) {
	Convey("Float64.ForceMapToByte", t, func() {
		So(NewFloat64(280.4).ForceMapToByte(func(float64) byte { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyFloat64().ForceMapToByte(func(float64) byte { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToNillableByte(t *testing.T) {
	Convey("Float64.MapToNillableByte", t, func() {
		So(NewFloat64(280.4).MapToNillableByte(func(float64) *byte {
			var tmp byte = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyFloat64().MapToNillableByte(func(float64) *byte {
			var tmp byte = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat64(280.4).MapToNillableByte(func(float64) *byte { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToByte(t *testing.T) {
	Convey("Float64.MapToByte", t, func() {
		test1, err1 := NewFloat64(280.4).MapToByte(func(float64) (byte, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyFloat64().MapToByte(func(float64) (byte, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat64(280.4).MapToByte(func(float64) (byte, error) {
			return 222, errors.New("byte error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "byte error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat64_ForceMapToComplex128(t *testing.T) {
	Convey("Float64.ForceMapToComplex128", t, func() {
		So(NewFloat64(280.4).ForceMapToComplex128(func(float64) complex128 { return 987654 }).Get(),
			ShouldEqual, 987654)

		So(NewEmptyFloat64().ForceMapToComplex128(func(float64) complex128 { return 987654 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToNillableComplex128(t *testing.T) {
	Convey("Float64.MapToNillableComplex128", t, func() {
		So(NewFloat64(280.4).MapToNillableComplex128(func(float64) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).Get(), ShouldEqual, 987654)

		So(NewEmptyFloat64().MapToNillableComplex128(func(float64) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat64(280.4).MapToNillableComplex128(func(float64) *complex128 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToComplex128(t *testing.T) {
	Convey("Float64.MapToComplex128", t, func() {
		test1, err1 := NewFloat64(280.4).MapToComplex128(func(float64) (complex128, error) {
			return 987654, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 987654)

		test2, err2 := NewEmptyFloat64().MapToComplex128(func(float64) (complex128, error) {
			return 987654, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat64(280.4).MapToComplex128(func(float64) (complex128, error) {
			return 987654, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat64_ForceMapToComplex64(t *testing.T) {
	Convey("Float64.ForceMapToComplex64", t, func() {
		So(NewFloat64(280.4).ForceMapToComplex64(func(float64) complex64 { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyFloat64().ForceMapToComplex64(func(float64) complex64 { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToNillableComplex64(t *testing.T) {
	Convey("Float64.MapToNillableComplex64", t, func() {
		So(NewFloat64(280.4).MapToNillableComplex64(func(float64) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyFloat64().MapToNillableComplex64(func(float64) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat64(280.4).MapToNillableComplex64(func(float64) *complex64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToComplex64(t *testing.T) {
	Convey("Float64.MapToComplex64", t, func() {
		test1, err1 := NewFloat64(280.4).MapToComplex64(func(float64) (complex64, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyFloat64().MapToComplex64(func(float64) (complex64, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat64(280.4).MapToComplex64(func(float64) (complex64, error) {
			return 123456, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat64_ForceMapToFloat32(t *testing.T) {
	Convey("Float64.ForceMapToFloat32", t, func() {
		So(NewFloat64(280.4).ForceMapToFloat32(func(float64) float32 { return 32.3 }).Get(),
			ShouldEqual, 32.3)

		So(NewEmptyFloat64().ForceMapToFloat32(func(float64) float32 { return 32.3 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToNillableFloat32(t *testing.T) {
	Convey("Float64.MapToNillableFloat32", t, func() {
		So(NewFloat64(280.4).MapToNillableFloat32(func(float64) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).Get(), ShouldEqual, 32.3)

		So(NewEmptyFloat64().MapToNillableFloat32(func(float64) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat64(280.4).MapToNillableFloat32(func(float64) *float32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToFloat32(t *testing.T) {
	Convey("Float64.MapToFloat32", t, func() {
		test1, err1 := NewFloat64(280.4).MapToFloat32(func(float64) (float32, error) {
			return 32.3, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 32.3)

		test2, err2 := NewEmptyFloat64().MapToFloat32(func(float64) (float32, error) {
			return 32.3, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat64(280.4).MapToFloat32(func(float64) (float32, error) {
			return 32.3, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat64_ForceMapToInt(t *testing.T) {
	Convey("Float64.ForceMapToInt", t, func() {
		So(NewFloat64(280.4).ForceMapToInt(func(float64) int { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyFloat64().ForceMapToInt(func(float64) int { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToNillableInt(t *testing.T) {
	Convey("Float64.MapToNillableInt", t, func() {
		So(NewFloat64(280.4).MapToNillableInt(func(float64) *int {
			var tmp int = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyFloat64().MapToNillableInt(func(float64) *int {
			var tmp int = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat64(280.4).MapToNillableInt(func(float64) *int { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToInt(t *testing.T) {
	Convey("Float64.MapToInt", t, func() {
		test1, err1 := NewFloat64(280.4).MapToInt(func(float64) (int, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyFloat64().MapToInt(func(float64) (int, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat64(280.4).MapToInt(func(float64) (int, error) {
			return 123456, errors.New("int error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat64_ForceMapToInt16(t *testing.T) {
	Convey("Float64.ForceMapToInt16", t, func() {
		So(NewFloat64(280.4).ForceMapToInt16(func(float64) int16 { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyFloat64().ForceMapToInt16(func(float64) int16 { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToNillableInt16(t *testing.T) {
	Convey("Float64.MapToNillableInt16", t, func() {
		So(NewFloat64(280.4).MapToNillableInt16(func(float64) *int16 {
			var tmp int16 = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyFloat64().MapToNillableInt16(func(float64) *int16 {
			var tmp int16 = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat64(280.4).MapToNillableInt16(func(float64) *int16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToInt16(t *testing.T) {
	Convey("Float64.MapToInt16", t, func() {
		test1, err1 := NewFloat64(280.4).MapToInt16(func(float64) (int16, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyFloat64().MapToInt16(func(float64) (int16, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat64(280.4).MapToInt16(func(float64) (int16, error) {
			return 222, errors.New("int16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat64_ForceMapToInt32(t *testing.T) {
	Convey("Float64.ForceMapToInt32", t, func() {
		So(NewFloat64(280.4).ForceMapToInt32(func(float64) int32 { return 333 }).Get(),
			ShouldEqual, 333)

		So(NewEmptyFloat64().ForceMapToInt32(func(float64) int32 { return 333 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToNillableInt32(t *testing.T) {
	Convey("Float64.MapToNillableInt32", t, func() {
		So(NewFloat64(280.4).MapToNillableInt32(func(float64) *int32 {
			var tmp int32 = 333
			return &tmp
		}).Get(), ShouldEqual, 333)

		So(NewEmptyFloat64().MapToNillableInt32(func(float64) *int32 {
			var tmp int32 = 333
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat64(280.4).MapToNillableInt32(func(float64) *int32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToInt32(t *testing.T) {
	Convey("Float64.MapToInt32", t, func() {
		test1, err1 := NewFloat64(280.4).MapToInt32(func(float64) (int32, error) {
			return 333, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 333)

		test2, err2 := NewEmptyFloat64().MapToInt32(func(float64) (int32, error) {
			return 333, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat64(280.4).MapToInt32(func(float64) (int32, error) {
			return 333, errors.New("int32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat64_ForceMapToInt64(t *testing.T) {
	Convey("Float64.ForceMapToInt64", t, func() {
		So(NewFloat64(280.4).ForceMapToInt64(func(float64) int64 { return 444 }).Get(),
			ShouldEqual, 444)

		So(NewEmptyFloat64().ForceMapToInt64(func(float64) int64 { return 444 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToNillableInt64(t *testing.T) {
	Convey("Float64.MapToNillableInt64", t, func() {
		So(NewFloat64(280.4).MapToNillableInt64(func(float64) *int64 {
			var tmp int64 = 444
			return &tmp
		}).Get(), ShouldEqual, 444)

		So(NewEmptyFloat64().MapToNillableInt64(func(float64) *int64 {
			var tmp int64 = 444
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat64(280.4).MapToNillableInt64(func(float64) *int64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToInt64(t *testing.T) {
	Convey("Float64.MapToInt64", t, func() {
		test1, err1 := NewFloat64(280.4).MapToInt64(func(float64) (int64, error) {
			return 444, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 444)

		test2, err2 := NewEmptyFloat64().MapToInt64(func(float64) (int64, error) {
			return 444, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat64(280.4).MapToInt64(func(float64) (int64, error) {
			return 444, errors.New("int64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat64_ForceMapToInt8(t *testing.T) {
	Convey("Float64.ForceMapToInt8", t, func() {
		So(NewFloat64(280.4).ForceMapToInt8(func(float64) int8 { return 111 }).Get(),
			ShouldEqual, 111)

		So(NewEmptyFloat64().ForceMapToInt8(func(float64) int8 { return 111 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToNillableInt8(t *testing.T) {
	Convey("Float64.MapToNillableInt8", t, func() {
		So(NewFloat64(280.4).MapToNillableInt8(func(float64) *int8 {
			var tmp int8 = 111
			return &tmp
		}).Get(), ShouldEqual, 111)

		So(NewEmptyFloat64().MapToNillableInt8(func(float64) *int8 {
			var tmp int8 = 111
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat64(280.4).MapToNillableInt8(func(float64) *int8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToInt8(t *testing.T) {
	Convey("Float64.MapToInt8", t, func() {
		test1, err1 := NewFloat64(280.4).MapToInt8(func(float64) (int8, error) {
			return 111, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 111)

		test2, err2 := NewEmptyFloat64().MapToInt8(func(float64) (int8, error) {
			return 111, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat64(280.4).MapToInt8(func(float64) (int8, error) {
			return 111, errors.New("int8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat64_ForceMapToUntyped(t *testing.T) {
	Convey("Float64.ForceMapToUntyped", t, func() {
		So(NewFloat64(280.4).ForceMapToUntyped(func(float64) interface{} { return "111111" }).Get(),
			ShouldEqual, "111111")

		So(NewEmptyFloat64().ForceMapToUntyped(func(float64) interface{} { return "111111" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToNillableUntyped(t *testing.T) {
	Convey("Float64.MapToNillableUntyped", t, func() {
		So(NewFloat64(280.4).MapToNillableUntyped(func(float64) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).Get(), ShouldEqual, "111111")

		So(NewEmptyFloat64().MapToNillableUntyped(func(float64) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat64(280.4).MapToNillableUntyped(func(float64) *interface{} { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToUntyped(t *testing.T) {
	Convey("Float64.MapToUntyped", t, func() {
		test1, err1 := NewFloat64(280.4).MapToUntyped(func(float64) (interface{}, error) {
			return "111111", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "111111")

		test2, err2 := NewEmptyFloat64().MapToUntyped(func(float64) (interface{}, error) {
			return "111111", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat64(280.4).MapToUntyped(func(float64) (interface{}, error) {
			return "111111", errors.New("interface error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "interface error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat64_ForceMapToRune(t *testing.T) {
	Convey("Float64.ForceMapToRune", t, func() {
		So(NewFloat64(280.4).ForceMapToRune(func(float64) rune { return 555 }).Get(),
			ShouldEqual, 555)

		So(NewEmptyFloat64().ForceMapToRune(func(float64) rune { return 555 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToNillableRune(t *testing.T) {
	Convey("Float64.MapToNillableRune", t, func() {
		So(NewFloat64(280.4).MapToNillableRune(func(float64) *rune {
			var tmp rune = 555
			return &tmp
		}).Get(), ShouldEqual, 555)

		So(NewEmptyFloat64().MapToNillableRune(func(float64) *rune {
			var tmp rune = 555
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat64(280.4).MapToNillableRune(func(float64) *rune { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToRune(t *testing.T) {
	Convey("Float64.MapToRune", t, func() {
		test1, err1 := NewFloat64(280.4).MapToRune(func(float64) (rune, error) {
			return 555, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 555)

		test2, err2 := NewEmptyFloat64().MapToRune(func(float64) (rune, error) {
			return 555, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat64(280.4).MapToRune(func(float64) (rune, error) {
			return 555, errors.New("rune error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "rune error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat64_ForceMapToString(t *testing.T) {
	Convey("Float64.ForceMapToString", t, func() {
		So(NewFloat64(280.4).ForceMapToString(func(float64) string { return "bonanza" }).Get(),
			ShouldEqual, "bonanza")

		So(NewEmptyFloat64().ForceMapToString(func(float64) string { return "bonanza" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToNillableString(t *testing.T) {
	Convey("Float64.MapToNillableString", t, func() {
		So(NewFloat64(280.4).MapToNillableString(func(float64) *string {
			var tmp string = "bonanza"
			return &tmp
		}).Get(), ShouldEqual, "bonanza")

		So(NewEmptyFloat64().MapToNillableString(func(float64) *string {
			var tmp string = "bonanza"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat64(280.4).MapToNillableString(func(float64) *string { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToString(t *testing.T) {
	Convey("Float64.MapToString", t, func() {
		test1, err1 := NewFloat64(280.4).MapToString(func(float64) (string, error) {
			return "bonanza", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "bonanza")

		test2, err2 := NewEmptyFloat64().MapToString(func(float64) (string, error) {
			return "bonanza", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat64(280.4).MapToString(func(float64) (string, error) {
			return "bonanza", errors.New("string error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "string error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat64_ForceMapToUint(t *testing.T) {
	Convey("Float64.ForceMapToUint", t, func() {
		So(NewFloat64(280.4).ForceMapToUint(func(float64) uint { return 666 }).Get(),
			ShouldEqual, 666)

		So(NewEmptyFloat64().ForceMapToUint(func(float64) uint { return 666 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToNillableUint(t *testing.T) {
	Convey("Float64.MapToNillableUint", t, func() {
		So(NewFloat64(280.4).MapToNillableUint(func(float64) *uint {
			var tmp uint = 666
			return &tmp
		}).Get(), ShouldEqual, 666)

		So(NewEmptyFloat64().MapToNillableUint(func(float64) *uint {
			var tmp uint = 666
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat64(280.4).MapToNillableUint(func(float64) *uint { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToUint(t *testing.T) {
	Convey("Float64.MapToUint", t, func() {
		test1, err1 := NewFloat64(280.4).MapToUint(func(float64) (uint, error) {
			return 666, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 666)

		test2, err2 := NewEmptyFloat64().MapToUint(func(float64) (uint, error) {
			return 666, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat64(280.4).MapToUint(func(float64) (uint, error) {
			return 666, errors.New("uint error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat64_ForceMapToUint16(t *testing.T) {
	Convey("Float64.ForceMapToUint16", t, func() {
		So(NewFloat64(280.4).ForceMapToUint16(func(float64) uint16 { return 888 }).Get(),
			ShouldEqual, 888)

		So(NewEmptyFloat64().ForceMapToUint16(func(float64) uint16 { return 888 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToNillableUint16(t *testing.T) {
	Convey("Float64.MapToNillableUint16", t, func() {
		So(NewFloat64(280.4).MapToNillableUint16(func(float64) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).Get(), ShouldEqual, 888)

		So(NewEmptyFloat64().MapToNillableUint16(func(float64) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat64(280.4).MapToNillableUint16(func(float64) *uint16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToUint16(t *testing.T) {
	Convey("Float64.MapToUint16", t, func() {
		test1, err1 := NewFloat64(280.4).MapToUint16(func(float64) (uint16, error) {
			return 888, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 888)

		test2, err2 := NewEmptyFloat64().MapToUint16(func(float64) (uint16, error) {
			return 888, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat64(280.4).MapToUint16(func(float64) (uint16, error) {
			return 888, errors.New("uint16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat64_ForceMapToUint32(t *testing.T) {
	Convey("Float64.ForceMapToUint32", t, func() {
		So(NewFloat64(280.4).ForceMapToUint32(func(float64) uint32 { return 999 }).Get(),
			ShouldEqual, 999)

		So(NewEmptyFloat64().ForceMapToUint32(func(float64) uint32 { return 999 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToNillableUint32(t *testing.T) {
	Convey("Float64.MapToNillableUint32", t, func() {
		So(NewFloat64(280.4).MapToNillableUint32(func(float64) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).Get(), ShouldEqual, 999)

		So(NewEmptyFloat64().MapToNillableUint32(func(float64) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat64(280.4).MapToNillableUint32(func(float64) *uint32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToUint32(t *testing.T) {
	Convey("Float64.MapToUint32", t, func() {
		test1, err1 := NewFloat64(280.4).MapToUint32(func(float64) (uint32, error) {
			return 999, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 999)

		test2, err2 := NewEmptyFloat64().MapToUint32(func(float64) (uint32, error) {
			return 999, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat64(280.4).MapToUint32(func(float64) (uint32, error) {
			return 999, errors.New("uint32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat64_ForceMapToUint64(t *testing.T) {
	Convey("Float64.ForceMapToUint64", t, func() {
		So(NewFloat64(280.4).ForceMapToUint64(func(float64) uint64 { return 101010 }).Get(),
			ShouldEqual, 101010)

		So(NewEmptyFloat64().ForceMapToUint64(func(float64) uint64 { return 101010 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToNillableUint64(t *testing.T) {
	Convey("Float64.MapToNillableUint64", t, func() {
		So(NewFloat64(280.4).MapToNillableUint64(func(float64) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).Get(), ShouldEqual, 101010)

		So(NewEmptyFloat64().MapToNillableUint64(func(float64) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat64(280.4).MapToNillableUint64(func(float64) *uint64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToUint64(t *testing.T) {
	Convey("Float64.MapToUint64", t, func() {
		test1, err1 := NewFloat64(280.4).MapToUint64(func(float64) (uint64, error) {
			return 101010, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 101010)

		test2, err2 := NewEmptyFloat64().MapToUint64(func(float64) (uint64, error) {
			return 101010, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat64(280.4).MapToUint64(func(float64) (uint64, error) {
			return 101010, errors.New("uint64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestFloat64_ForceMapToUint8(t *testing.T) {
	Convey("Float64.ForceMapToUint8", t, func() {
		So(NewFloat64(280.4).ForceMapToUint8(func(float64) uint8 { return 77 }).Get(),
			ShouldEqual, 77)

		So(NewEmptyFloat64().ForceMapToUint8(func(float64) uint8 { return 77 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToNillableUint8(t *testing.T) {
	Convey("Float64.MapToNillableUint8", t, func() {
		So(NewFloat64(280.4).MapToNillableUint8(func(float64) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).Get(), ShouldEqual, 77)

		So(NewEmptyFloat64().MapToNillableUint8(func(float64) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewFloat64(280.4).MapToNillableUint8(func(float64) *uint8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestFloat64_MapToUint8(t *testing.T) {
	Convey("Float64.MapToUint8", t, func() {
		test1, err1 := NewFloat64(280.4).MapToUint8(func(float64) (uint8, error) {
			return 77, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 77)

		test2, err2 := NewEmptyFloat64().MapToUint8(func(float64) (uint8, error) {
			return 77, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewFloat64(280.4).MapToUint8(func(float64) (uint8, error) {
			return 77, errors.New("uint8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}
