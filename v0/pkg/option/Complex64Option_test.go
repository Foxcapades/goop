// Package option_test generated @ 2020-05-17 22:19:55.164878293 -0400 EDT m=+0.001439508
package option_test

import (
	"errors"
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

func TestComplex64_ForceMapToBool(t *testing.T) {
	Convey("Complex64.ForceMapToBool", t, func() {
		So(NewComplex64(2548).ForceMapToBool(func(complex64) bool { return true }).Get(),
			ShouldEqual, true)

		So(NewEmptyComplex64().ForceMapToBool(func(complex64) bool { return true }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToNillableBool(t *testing.T) {
	Convey("Complex64.MapToNillableBool", t, func() {
		So(NewComplex64(2548).MapToNillableBool(func(complex64) *bool {
			var tmp bool = true
			return &tmp
		}).Get(), ShouldEqual, true)

		So(NewEmptyComplex64().MapToNillableBool(func(complex64) *bool {
			var tmp bool = true
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex64(2548).MapToNillableBool(func(complex64) *bool { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToBool(t *testing.T) {
	Convey("Complex64.MapToBool", t, func() {
		test1, err1 := NewComplex64(2548).MapToBool(func(complex64) (bool, error) {
			return true, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, true)

		test2, err2 := NewEmptyComplex64().MapToBool(func(complex64) (bool, error) {
			return true, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex64(2548).MapToBool(func(complex64) (bool, error) {
			return true, errors.New("bool error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "bool error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex64_ForceMapToByte(t *testing.T) {
	Convey("Complex64.ForceMapToByte", t, func() {
		So(NewComplex64(2548).ForceMapToByte(func(complex64) byte { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyComplex64().ForceMapToByte(func(complex64) byte { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToNillableByte(t *testing.T) {
	Convey("Complex64.MapToNillableByte", t, func() {
		So(NewComplex64(2548).MapToNillableByte(func(complex64) *byte {
			var tmp byte = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyComplex64().MapToNillableByte(func(complex64) *byte {
			var tmp byte = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex64(2548).MapToNillableByte(func(complex64) *byte { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToByte(t *testing.T) {
	Convey("Complex64.MapToByte", t, func() {
		test1, err1 := NewComplex64(2548).MapToByte(func(complex64) (byte, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyComplex64().MapToByte(func(complex64) (byte, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex64(2548).MapToByte(func(complex64) (byte, error) {
			return 222, errors.New("byte error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "byte error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex64_ForceMapToComplex128(t *testing.T) {
	Convey("Complex64.ForceMapToComplex128", t, func() {
		So(NewComplex64(2548).ForceMapToComplex128(func(complex64) complex128 { return 987654 }).Get(),
			ShouldEqual, 987654)

		So(NewEmptyComplex64().ForceMapToComplex128(func(complex64) complex128 { return 987654 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToNillableComplex128(t *testing.T) {
	Convey("Complex64.MapToNillableComplex128", t, func() {
		So(NewComplex64(2548).MapToNillableComplex128(func(complex64) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).Get(), ShouldEqual, 987654)

		So(NewEmptyComplex64().MapToNillableComplex128(func(complex64) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex64(2548).MapToNillableComplex128(func(complex64) *complex128 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToComplex128(t *testing.T) {
	Convey("Complex64.MapToComplex128", t, func() {
		test1, err1 := NewComplex64(2548).MapToComplex128(func(complex64) (complex128, error) {
			return 987654, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 987654)

		test2, err2 := NewEmptyComplex64().MapToComplex128(func(complex64) (complex128, error) {
			return 987654, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex64(2548).MapToComplex128(func(complex64) (complex128, error) {
			return 987654, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex64_ForceMapToFloat32(t *testing.T) {
	Convey("Complex64.ForceMapToFloat32", t, func() {
		So(NewComplex64(2548).ForceMapToFloat32(func(complex64) float32 { return 32.3 }).Get(),
			ShouldEqual, 32.3)

		So(NewEmptyComplex64().ForceMapToFloat32(func(complex64) float32 { return 32.3 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToNillableFloat32(t *testing.T) {
	Convey("Complex64.MapToNillableFloat32", t, func() {
		So(NewComplex64(2548).MapToNillableFloat32(func(complex64) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).Get(), ShouldEqual, 32.3)

		So(NewEmptyComplex64().MapToNillableFloat32(func(complex64) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex64(2548).MapToNillableFloat32(func(complex64) *float32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToFloat32(t *testing.T) {
	Convey("Complex64.MapToFloat32", t, func() {
		test1, err1 := NewComplex64(2548).MapToFloat32(func(complex64) (float32, error) {
			return 32.3, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 32.3)

		test2, err2 := NewEmptyComplex64().MapToFloat32(func(complex64) (float32, error) {
			return 32.3, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex64(2548).MapToFloat32(func(complex64) (float32, error) {
			return 32.3, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex64_ForceMapToFloat64(t *testing.T) {
	Convey("Complex64.ForceMapToFloat64", t, func() {
		So(NewComplex64(2548).ForceMapToFloat64(func(complex64) float64 { return 3.2455 }).Get(),
			ShouldEqual, 3.2455)

		So(NewEmptyComplex64().ForceMapToFloat64(func(complex64) float64 { return 3.2455 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToNillableFloat64(t *testing.T) {
	Convey("Complex64.MapToNillableFloat64", t, func() {
		So(NewComplex64(2548).MapToNillableFloat64(func(complex64) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).Get(), ShouldEqual, 3.2455)

		So(NewEmptyComplex64().MapToNillableFloat64(func(complex64) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex64(2548).MapToNillableFloat64(func(complex64) *float64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToFloat64(t *testing.T) {
	Convey("Complex64.MapToFloat64", t, func() {
		test1, err1 := NewComplex64(2548).MapToFloat64(func(complex64) (float64, error) {
			return 3.2455, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 3.2455)

		test2, err2 := NewEmptyComplex64().MapToFloat64(func(complex64) (float64, error) {
			return 3.2455, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex64(2548).MapToFloat64(func(complex64) (float64, error) {
			return 3.2455, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex64_ForceMapToInt(t *testing.T) {
	Convey("Complex64.ForceMapToInt", t, func() {
		So(NewComplex64(2548).ForceMapToInt(func(complex64) int { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyComplex64().ForceMapToInt(func(complex64) int { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToNillableInt(t *testing.T) {
	Convey("Complex64.MapToNillableInt", t, func() {
		So(NewComplex64(2548).MapToNillableInt(func(complex64) *int {
			var tmp int = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyComplex64().MapToNillableInt(func(complex64) *int {
			var tmp int = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex64(2548).MapToNillableInt(func(complex64) *int { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToInt(t *testing.T) {
	Convey("Complex64.MapToInt", t, func() {
		test1, err1 := NewComplex64(2548).MapToInt(func(complex64) (int, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyComplex64().MapToInt(func(complex64) (int, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex64(2548).MapToInt(func(complex64) (int, error) {
			return 123456, errors.New("int error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex64_ForceMapToInt16(t *testing.T) {
	Convey("Complex64.ForceMapToInt16", t, func() {
		So(NewComplex64(2548).ForceMapToInt16(func(complex64) int16 { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyComplex64().ForceMapToInt16(func(complex64) int16 { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToNillableInt16(t *testing.T) {
	Convey("Complex64.MapToNillableInt16", t, func() {
		So(NewComplex64(2548).MapToNillableInt16(func(complex64) *int16 {
			var tmp int16 = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyComplex64().MapToNillableInt16(func(complex64) *int16 {
			var tmp int16 = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex64(2548).MapToNillableInt16(func(complex64) *int16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToInt16(t *testing.T) {
	Convey("Complex64.MapToInt16", t, func() {
		test1, err1 := NewComplex64(2548).MapToInt16(func(complex64) (int16, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyComplex64().MapToInt16(func(complex64) (int16, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex64(2548).MapToInt16(func(complex64) (int16, error) {
			return 222, errors.New("int16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex64_ForceMapToInt32(t *testing.T) {
	Convey("Complex64.ForceMapToInt32", t, func() {
		So(NewComplex64(2548).ForceMapToInt32(func(complex64) int32 { return 333 }).Get(),
			ShouldEqual, 333)

		So(NewEmptyComplex64().ForceMapToInt32(func(complex64) int32 { return 333 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToNillableInt32(t *testing.T) {
	Convey("Complex64.MapToNillableInt32", t, func() {
		So(NewComplex64(2548).MapToNillableInt32(func(complex64) *int32 {
			var tmp int32 = 333
			return &tmp
		}).Get(), ShouldEqual, 333)

		So(NewEmptyComplex64().MapToNillableInt32(func(complex64) *int32 {
			var tmp int32 = 333
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex64(2548).MapToNillableInt32(func(complex64) *int32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToInt32(t *testing.T) {
	Convey("Complex64.MapToInt32", t, func() {
		test1, err1 := NewComplex64(2548).MapToInt32(func(complex64) (int32, error) {
			return 333, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 333)

		test2, err2 := NewEmptyComplex64().MapToInt32(func(complex64) (int32, error) {
			return 333, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex64(2548).MapToInt32(func(complex64) (int32, error) {
			return 333, errors.New("int32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex64_ForceMapToInt64(t *testing.T) {
	Convey("Complex64.ForceMapToInt64", t, func() {
		So(NewComplex64(2548).ForceMapToInt64(func(complex64) int64 { return 444 }).Get(),
			ShouldEqual, 444)

		So(NewEmptyComplex64().ForceMapToInt64(func(complex64) int64 { return 444 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToNillableInt64(t *testing.T) {
	Convey("Complex64.MapToNillableInt64", t, func() {
		So(NewComplex64(2548).MapToNillableInt64(func(complex64) *int64 {
			var tmp int64 = 444
			return &tmp
		}).Get(), ShouldEqual, 444)

		So(NewEmptyComplex64().MapToNillableInt64(func(complex64) *int64 {
			var tmp int64 = 444
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex64(2548).MapToNillableInt64(func(complex64) *int64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToInt64(t *testing.T) {
	Convey("Complex64.MapToInt64", t, func() {
		test1, err1 := NewComplex64(2548).MapToInt64(func(complex64) (int64, error) {
			return 444, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 444)

		test2, err2 := NewEmptyComplex64().MapToInt64(func(complex64) (int64, error) {
			return 444, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex64(2548).MapToInt64(func(complex64) (int64, error) {
			return 444, errors.New("int64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex64_ForceMapToInt8(t *testing.T) {
	Convey("Complex64.ForceMapToInt8", t, func() {
		So(NewComplex64(2548).ForceMapToInt8(func(complex64) int8 { return 111 }).Get(),
			ShouldEqual, 111)

		So(NewEmptyComplex64().ForceMapToInt8(func(complex64) int8 { return 111 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToNillableInt8(t *testing.T) {
	Convey("Complex64.MapToNillableInt8", t, func() {
		So(NewComplex64(2548).MapToNillableInt8(func(complex64) *int8 {
			var tmp int8 = 111
			return &tmp
		}).Get(), ShouldEqual, 111)

		So(NewEmptyComplex64().MapToNillableInt8(func(complex64) *int8 {
			var tmp int8 = 111
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex64(2548).MapToNillableInt8(func(complex64) *int8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToInt8(t *testing.T) {
	Convey("Complex64.MapToInt8", t, func() {
		test1, err1 := NewComplex64(2548).MapToInt8(func(complex64) (int8, error) {
			return 111, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 111)

		test2, err2 := NewEmptyComplex64().MapToInt8(func(complex64) (int8, error) {
			return 111, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex64(2548).MapToInt8(func(complex64) (int8, error) {
			return 111, errors.New("int8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex64_ForceMapToUntyped(t *testing.T) {
	Convey("Complex64.ForceMapToUntyped", t, func() {
		So(NewComplex64(2548).ForceMapToUntyped(func(complex64) interface{} { return "111111" }).Get(),
			ShouldEqual, "111111")

		So(NewEmptyComplex64().ForceMapToUntyped(func(complex64) interface{} { return "111111" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToNillableUntyped(t *testing.T) {
	Convey("Complex64.MapToNillableUntyped", t, func() {
		So(NewComplex64(2548).MapToNillableUntyped(func(complex64) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).Get(), ShouldEqual, "111111")

		So(NewEmptyComplex64().MapToNillableUntyped(func(complex64) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex64(2548).MapToNillableUntyped(func(complex64) *interface{} { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToUntyped(t *testing.T) {
	Convey("Complex64.MapToUntyped", t, func() {
		test1, err1 := NewComplex64(2548).MapToUntyped(func(complex64) (interface{}, error) {
			return "111111", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "111111")

		test2, err2 := NewEmptyComplex64().MapToUntyped(func(complex64) (interface{}, error) {
			return "111111", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex64(2548).MapToUntyped(func(complex64) (interface{}, error) {
			return "111111", errors.New("interface error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "interface error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex64_ForceMapToRune(t *testing.T) {
	Convey("Complex64.ForceMapToRune", t, func() {
		So(NewComplex64(2548).ForceMapToRune(func(complex64) rune { return 555 }).Get(),
			ShouldEqual, 555)

		So(NewEmptyComplex64().ForceMapToRune(func(complex64) rune { return 555 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToNillableRune(t *testing.T) {
	Convey("Complex64.MapToNillableRune", t, func() {
		So(NewComplex64(2548).MapToNillableRune(func(complex64) *rune {
			var tmp rune = 555
			return &tmp
		}).Get(), ShouldEqual, 555)

		So(NewEmptyComplex64().MapToNillableRune(func(complex64) *rune {
			var tmp rune = 555
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex64(2548).MapToNillableRune(func(complex64) *rune { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToRune(t *testing.T) {
	Convey("Complex64.MapToRune", t, func() {
		test1, err1 := NewComplex64(2548).MapToRune(func(complex64) (rune, error) {
			return 555, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 555)

		test2, err2 := NewEmptyComplex64().MapToRune(func(complex64) (rune, error) {
			return 555, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex64(2548).MapToRune(func(complex64) (rune, error) {
			return 555, errors.New("rune error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "rune error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex64_ForceMapToString(t *testing.T) {
	Convey("Complex64.ForceMapToString", t, func() {
		So(NewComplex64(2548).ForceMapToString(func(complex64) string { return "bonanza" }).Get(),
			ShouldEqual, "bonanza")

		So(NewEmptyComplex64().ForceMapToString(func(complex64) string { return "bonanza" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToNillableString(t *testing.T) {
	Convey("Complex64.MapToNillableString", t, func() {
		So(NewComplex64(2548).MapToNillableString(func(complex64) *string {
			var tmp string = "bonanza"
			return &tmp
		}).Get(), ShouldEqual, "bonanza")

		So(NewEmptyComplex64().MapToNillableString(func(complex64) *string {
			var tmp string = "bonanza"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex64(2548).MapToNillableString(func(complex64) *string { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToString(t *testing.T) {
	Convey("Complex64.MapToString", t, func() {
		test1, err1 := NewComplex64(2548).MapToString(func(complex64) (string, error) {
			return "bonanza", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "bonanza")

		test2, err2 := NewEmptyComplex64().MapToString(func(complex64) (string, error) {
			return "bonanza", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex64(2548).MapToString(func(complex64) (string, error) {
			return "bonanza", errors.New("string error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "string error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex64_ForceMapToUint(t *testing.T) {
	Convey("Complex64.ForceMapToUint", t, func() {
		So(NewComplex64(2548).ForceMapToUint(func(complex64) uint { return 666 }).Get(),
			ShouldEqual, 666)

		So(NewEmptyComplex64().ForceMapToUint(func(complex64) uint { return 666 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToNillableUint(t *testing.T) {
	Convey("Complex64.MapToNillableUint", t, func() {
		So(NewComplex64(2548).MapToNillableUint(func(complex64) *uint {
			var tmp uint = 666
			return &tmp
		}).Get(), ShouldEqual, 666)

		So(NewEmptyComplex64().MapToNillableUint(func(complex64) *uint {
			var tmp uint = 666
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex64(2548).MapToNillableUint(func(complex64) *uint { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToUint(t *testing.T) {
	Convey("Complex64.MapToUint", t, func() {
		test1, err1 := NewComplex64(2548).MapToUint(func(complex64) (uint, error) {
			return 666, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 666)

		test2, err2 := NewEmptyComplex64().MapToUint(func(complex64) (uint, error) {
			return 666, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex64(2548).MapToUint(func(complex64) (uint, error) {
			return 666, errors.New("uint error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex64_ForceMapToUint16(t *testing.T) {
	Convey("Complex64.ForceMapToUint16", t, func() {
		So(NewComplex64(2548).ForceMapToUint16(func(complex64) uint16 { return 888 }).Get(),
			ShouldEqual, 888)

		So(NewEmptyComplex64().ForceMapToUint16(func(complex64) uint16 { return 888 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToNillableUint16(t *testing.T) {
	Convey("Complex64.MapToNillableUint16", t, func() {
		So(NewComplex64(2548).MapToNillableUint16(func(complex64) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).Get(), ShouldEqual, 888)

		So(NewEmptyComplex64().MapToNillableUint16(func(complex64) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex64(2548).MapToNillableUint16(func(complex64) *uint16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToUint16(t *testing.T) {
	Convey("Complex64.MapToUint16", t, func() {
		test1, err1 := NewComplex64(2548).MapToUint16(func(complex64) (uint16, error) {
			return 888, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 888)

		test2, err2 := NewEmptyComplex64().MapToUint16(func(complex64) (uint16, error) {
			return 888, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex64(2548).MapToUint16(func(complex64) (uint16, error) {
			return 888, errors.New("uint16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex64_ForceMapToUint32(t *testing.T) {
	Convey("Complex64.ForceMapToUint32", t, func() {
		So(NewComplex64(2548).ForceMapToUint32(func(complex64) uint32 { return 999 }).Get(),
			ShouldEqual, 999)

		So(NewEmptyComplex64().ForceMapToUint32(func(complex64) uint32 { return 999 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToNillableUint32(t *testing.T) {
	Convey("Complex64.MapToNillableUint32", t, func() {
		So(NewComplex64(2548).MapToNillableUint32(func(complex64) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).Get(), ShouldEqual, 999)

		So(NewEmptyComplex64().MapToNillableUint32(func(complex64) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex64(2548).MapToNillableUint32(func(complex64) *uint32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToUint32(t *testing.T) {
	Convey("Complex64.MapToUint32", t, func() {
		test1, err1 := NewComplex64(2548).MapToUint32(func(complex64) (uint32, error) {
			return 999, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 999)

		test2, err2 := NewEmptyComplex64().MapToUint32(func(complex64) (uint32, error) {
			return 999, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex64(2548).MapToUint32(func(complex64) (uint32, error) {
			return 999, errors.New("uint32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex64_ForceMapToUint64(t *testing.T) {
	Convey("Complex64.ForceMapToUint64", t, func() {
		So(NewComplex64(2548).ForceMapToUint64(func(complex64) uint64 { return 101010 }).Get(),
			ShouldEqual, 101010)

		So(NewEmptyComplex64().ForceMapToUint64(func(complex64) uint64 { return 101010 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToNillableUint64(t *testing.T) {
	Convey("Complex64.MapToNillableUint64", t, func() {
		So(NewComplex64(2548).MapToNillableUint64(func(complex64) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).Get(), ShouldEqual, 101010)

		So(NewEmptyComplex64().MapToNillableUint64(func(complex64) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex64(2548).MapToNillableUint64(func(complex64) *uint64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToUint64(t *testing.T) {
	Convey("Complex64.MapToUint64", t, func() {
		test1, err1 := NewComplex64(2548).MapToUint64(func(complex64) (uint64, error) {
			return 101010, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 101010)

		test2, err2 := NewEmptyComplex64().MapToUint64(func(complex64) (uint64, error) {
			return 101010, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex64(2548).MapToUint64(func(complex64) (uint64, error) {
			return 101010, errors.New("uint64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestComplex64_ForceMapToUint8(t *testing.T) {
	Convey("Complex64.ForceMapToUint8", t, func() {
		So(NewComplex64(2548).ForceMapToUint8(func(complex64) uint8 { return 77 }).Get(),
			ShouldEqual, 77)

		So(NewEmptyComplex64().ForceMapToUint8(func(complex64) uint8 { return 77 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToNillableUint8(t *testing.T) {
	Convey("Complex64.MapToNillableUint8", t, func() {
		So(NewComplex64(2548).MapToNillableUint8(func(complex64) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).Get(), ShouldEqual, 77)

		So(NewEmptyComplex64().MapToNillableUint8(func(complex64) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewComplex64(2548).MapToNillableUint8(func(complex64) *uint8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestComplex64_MapToUint8(t *testing.T) {
	Convey("Complex64.MapToUint8", t, func() {
		test1, err1 := NewComplex64(2548).MapToUint8(func(complex64) (uint8, error) {
			return 77, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 77)

		test2, err2 := NewEmptyComplex64().MapToUint8(func(complex64) (uint8, error) {
			return 77, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewComplex64(2548).MapToUint8(func(complex64) (uint8, error) {
			return 77, errors.New("uint8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}
