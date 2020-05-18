// Package option_test generated @ 2020-05-17 22:19:55.164878293 -0400 EDT m=+0.001439508
package option_test

import (
	"errors"
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

func TestInt64_ForceMapToBool(t *testing.T) {
	Convey("Int64.ForceMapToBool", t, func() {
		So(NewInt64(55555).ForceMapToBool(func(int64) bool { return true }).Get(),
			ShouldEqual, true)

		So(NewEmptyInt64().ForceMapToBool(func(int64) bool { return true }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToNillableBool(t *testing.T) {
	Convey("Int64.MapToNillableBool", t, func() {
		So(NewInt64(55555).MapToNillableBool(func(int64) *bool {
			var tmp bool = true
			return &tmp
		}).Get(), ShouldEqual, true)

		So(NewEmptyInt64().MapToNillableBool(func(int64) *bool {
			var tmp bool = true
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt64(55555).MapToNillableBool(func(int64) *bool { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToBool(t *testing.T) {
	Convey("Int64.MapToBool", t, func() {
		test1, err1 := NewInt64(55555).MapToBool(func(int64) (bool, error) {
			return true, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, true)

		test2, err2 := NewEmptyInt64().MapToBool(func(int64) (bool, error) {
			return true, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt64(55555).MapToBool(func(int64) (bool, error) {
			return true, errors.New("bool error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "bool error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt64_ForceMapToByte(t *testing.T) {
	Convey("Int64.ForceMapToByte", t, func() {
		So(NewInt64(55555).ForceMapToByte(func(int64) byte { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyInt64().ForceMapToByte(func(int64) byte { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToNillableByte(t *testing.T) {
	Convey("Int64.MapToNillableByte", t, func() {
		So(NewInt64(55555).MapToNillableByte(func(int64) *byte {
			var tmp byte = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyInt64().MapToNillableByte(func(int64) *byte {
			var tmp byte = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt64(55555).MapToNillableByte(func(int64) *byte { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToByte(t *testing.T) {
	Convey("Int64.MapToByte", t, func() {
		test1, err1 := NewInt64(55555).MapToByte(func(int64) (byte, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyInt64().MapToByte(func(int64) (byte, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt64(55555).MapToByte(func(int64) (byte, error) {
			return 222, errors.New("byte error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "byte error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt64_ForceMapToComplex128(t *testing.T) {
	Convey("Int64.ForceMapToComplex128", t, func() {
		So(NewInt64(55555).ForceMapToComplex128(func(int64) complex128 { return 987654 }).Get(),
			ShouldEqual, 987654)

		So(NewEmptyInt64().ForceMapToComplex128(func(int64) complex128 { return 987654 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToNillableComplex128(t *testing.T) {
	Convey("Int64.MapToNillableComplex128", t, func() {
		So(NewInt64(55555).MapToNillableComplex128(func(int64) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).Get(), ShouldEqual, 987654)

		So(NewEmptyInt64().MapToNillableComplex128(func(int64) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt64(55555).MapToNillableComplex128(func(int64) *complex128 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToComplex128(t *testing.T) {
	Convey("Int64.MapToComplex128", t, func() {
		test1, err1 := NewInt64(55555).MapToComplex128(func(int64) (complex128, error) {
			return 987654, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 987654)

		test2, err2 := NewEmptyInt64().MapToComplex128(func(int64) (complex128, error) {
			return 987654, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt64(55555).MapToComplex128(func(int64) (complex128, error) {
			return 987654, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt64_ForceMapToComplex64(t *testing.T) {
	Convey("Int64.ForceMapToComplex64", t, func() {
		So(NewInt64(55555).ForceMapToComplex64(func(int64) complex64 { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyInt64().ForceMapToComplex64(func(int64) complex64 { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToNillableComplex64(t *testing.T) {
	Convey("Int64.MapToNillableComplex64", t, func() {
		So(NewInt64(55555).MapToNillableComplex64(func(int64) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyInt64().MapToNillableComplex64(func(int64) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt64(55555).MapToNillableComplex64(func(int64) *complex64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToComplex64(t *testing.T) {
	Convey("Int64.MapToComplex64", t, func() {
		test1, err1 := NewInt64(55555).MapToComplex64(func(int64) (complex64, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyInt64().MapToComplex64(func(int64) (complex64, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt64(55555).MapToComplex64(func(int64) (complex64, error) {
			return 123456, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt64_ForceMapToFloat32(t *testing.T) {
	Convey("Int64.ForceMapToFloat32", t, func() {
		So(NewInt64(55555).ForceMapToFloat32(func(int64) float32 { return 32.3 }).Get(),
			ShouldEqual, 32.3)

		So(NewEmptyInt64().ForceMapToFloat32(func(int64) float32 { return 32.3 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToNillableFloat32(t *testing.T) {
	Convey("Int64.MapToNillableFloat32", t, func() {
		So(NewInt64(55555).MapToNillableFloat32(func(int64) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).Get(), ShouldEqual, 32.3)

		So(NewEmptyInt64().MapToNillableFloat32(func(int64) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt64(55555).MapToNillableFloat32(func(int64) *float32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToFloat32(t *testing.T) {
	Convey("Int64.MapToFloat32", t, func() {
		test1, err1 := NewInt64(55555).MapToFloat32(func(int64) (float32, error) {
			return 32.3, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 32.3)

		test2, err2 := NewEmptyInt64().MapToFloat32(func(int64) (float32, error) {
			return 32.3, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt64(55555).MapToFloat32(func(int64) (float32, error) {
			return 32.3, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt64_ForceMapToFloat64(t *testing.T) {
	Convey("Int64.ForceMapToFloat64", t, func() {
		So(NewInt64(55555).ForceMapToFloat64(func(int64) float64 { return 3.2455 }).Get(),
			ShouldEqual, 3.2455)

		So(NewEmptyInt64().ForceMapToFloat64(func(int64) float64 { return 3.2455 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToNillableFloat64(t *testing.T) {
	Convey("Int64.MapToNillableFloat64", t, func() {
		So(NewInt64(55555).MapToNillableFloat64(func(int64) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).Get(), ShouldEqual, 3.2455)

		So(NewEmptyInt64().MapToNillableFloat64(func(int64) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt64(55555).MapToNillableFloat64(func(int64) *float64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToFloat64(t *testing.T) {
	Convey("Int64.MapToFloat64", t, func() {
		test1, err1 := NewInt64(55555).MapToFloat64(func(int64) (float64, error) {
			return 3.2455, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 3.2455)

		test2, err2 := NewEmptyInt64().MapToFloat64(func(int64) (float64, error) {
			return 3.2455, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt64(55555).MapToFloat64(func(int64) (float64, error) {
			return 3.2455, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt64_ForceMapToInt(t *testing.T) {
	Convey("Int64.ForceMapToInt", t, func() {
		So(NewInt64(55555).ForceMapToInt(func(int64) int { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyInt64().ForceMapToInt(func(int64) int { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToNillableInt(t *testing.T) {
	Convey("Int64.MapToNillableInt", t, func() {
		So(NewInt64(55555).MapToNillableInt(func(int64) *int {
			var tmp int = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyInt64().MapToNillableInt(func(int64) *int {
			var tmp int = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt64(55555).MapToNillableInt(func(int64) *int { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToInt(t *testing.T) {
	Convey("Int64.MapToInt", t, func() {
		test1, err1 := NewInt64(55555).MapToInt(func(int64) (int, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyInt64().MapToInt(func(int64) (int, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt64(55555).MapToInt(func(int64) (int, error) {
			return 123456, errors.New("int error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt64_ForceMapToInt16(t *testing.T) {
	Convey("Int64.ForceMapToInt16", t, func() {
		So(NewInt64(55555).ForceMapToInt16(func(int64) int16 { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyInt64().ForceMapToInt16(func(int64) int16 { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToNillableInt16(t *testing.T) {
	Convey("Int64.MapToNillableInt16", t, func() {
		So(NewInt64(55555).MapToNillableInt16(func(int64) *int16 {
			var tmp int16 = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyInt64().MapToNillableInt16(func(int64) *int16 {
			var tmp int16 = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt64(55555).MapToNillableInt16(func(int64) *int16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToInt16(t *testing.T) {
	Convey("Int64.MapToInt16", t, func() {
		test1, err1 := NewInt64(55555).MapToInt16(func(int64) (int16, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyInt64().MapToInt16(func(int64) (int16, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt64(55555).MapToInt16(func(int64) (int16, error) {
			return 222, errors.New("int16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt64_ForceMapToInt32(t *testing.T) {
	Convey("Int64.ForceMapToInt32", t, func() {
		So(NewInt64(55555).ForceMapToInt32(func(int64) int32 { return 333 }).Get(),
			ShouldEqual, 333)

		So(NewEmptyInt64().ForceMapToInt32(func(int64) int32 { return 333 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToNillableInt32(t *testing.T) {
	Convey("Int64.MapToNillableInt32", t, func() {
		So(NewInt64(55555).MapToNillableInt32(func(int64) *int32 {
			var tmp int32 = 333
			return &tmp
		}).Get(), ShouldEqual, 333)

		So(NewEmptyInt64().MapToNillableInt32(func(int64) *int32 {
			var tmp int32 = 333
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt64(55555).MapToNillableInt32(func(int64) *int32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToInt32(t *testing.T) {
	Convey("Int64.MapToInt32", t, func() {
		test1, err1 := NewInt64(55555).MapToInt32(func(int64) (int32, error) {
			return 333, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 333)

		test2, err2 := NewEmptyInt64().MapToInt32(func(int64) (int32, error) {
			return 333, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt64(55555).MapToInt32(func(int64) (int32, error) {
			return 333, errors.New("int32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt64_ForceMapToInt8(t *testing.T) {
	Convey("Int64.ForceMapToInt8", t, func() {
		So(NewInt64(55555).ForceMapToInt8(func(int64) int8 { return 111 }).Get(),
			ShouldEqual, 111)

		So(NewEmptyInt64().ForceMapToInt8(func(int64) int8 { return 111 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToNillableInt8(t *testing.T) {
	Convey("Int64.MapToNillableInt8", t, func() {
		So(NewInt64(55555).MapToNillableInt8(func(int64) *int8 {
			var tmp int8 = 111
			return &tmp
		}).Get(), ShouldEqual, 111)

		So(NewEmptyInt64().MapToNillableInt8(func(int64) *int8 {
			var tmp int8 = 111
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt64(55555).MapToNillableInt8(func(int64) *int8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToInt8(t *testing.T) {
	Convey("Int64.MapToInt8", t, func() {
		test1, err1 := NewInt64(55555).MapToInt8(func(int64) (int8, error) {
			return 111, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 111)

		test2, err2 := NewEmptyInt64().MapToInt8(func(int64) (int8, error) {
			return 111, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt64(55555).MapToInt8(func(int64) (int8, error) {
			return 111, errors.New("int8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt64_ForceMapToUntyped(t *testing.T) {
	Convey("Int64.ForceMapToUntyped", t, func() {
		So(NewInt64(55555).ForceMapToUntyped(func(int64) interface{} { return "111111" }).Get(),
			ShouldEqual, "111111")

		So(NewEmptyInt64().ForceMapToUntyped(func(int64) interface{} { return "111111" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToNillableUntyped(t *testing.T) {
	Convey("Int64.MapToNillableUntyped", t, func() {
		So(NewInt64(55555).MapToNillableUntyped(func(int64) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).Get(), ShouldEqual, "111111")

		So(NewEmptyInt64().MapToNillableUntyped(func(int64) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt64(55555).MapToNillableUntyped(func(int64) *interface{} { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToUntyped(t *testing.T) {
	Convey("Int64.MapToUntyped", t, func() {
		test1, err1 := NewInt64(55555).MapToUntyped(func(int64) (interface{}, error) {
			return "111111", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "111111")

		test2, err2 := NewEmptyInt64().MapToUntyped(func(int64) (interface{}, error) {
			return "111111", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt64(55555).MapToUntyped(func(int64) (interface{}, error) {
			return "111111", errors.New("interface error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "interface error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt64_ForceMapToRune(t *testing.T) {
	Convey("Int64.ForceMapToRune", t, func() {
		So(NewInt64(55555).ForceMapToRune(func(int64) rune { return 555 }).Get(),
			ShouldEqual, 555)

		So(NewEmptyInt64().ForceMapToRune(func(int64) rune { return 555 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToNillableRune(t *testing.T) {
	Convey("Int64.MapToNillableRune", t, func() {
		So(NewInt64(55555).MapToNillableRune(func(int64) *rune {
			var tmp rune = 555
			return &tmp
		}).Get(), ShouldEqual, 555)

		So(NewEmptyInt64().MapToNillableRune(func(int64) *rune {
			var tmp rune = 555
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt64(55555).MapToNillableRune(func(int64) *rune { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToRune(t *testing.T) {
	Convey("Int64.MapToRune", t, func() {
		test1, err1 := NewInt64(55555).MapToRune(func(int64) (rune, error) {
			return 555, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 555)

		test2, err2 := NewEmptyInt64().MapToRune(func(int64) (rune, error) {
			return 555, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt64(55555).MapToRune(func(int64) (rune, error) {
			return 555, errors.New("rune error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "rune error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt64_ForceMapToString(t *testing.T) {
	Convey("Int64.ForceMapToString", t, func() {
		So(NewInt64(55555).ForceMapToString(func(int64) string { return "bonanza" }).Get(),
			ShouldEqual, "bonanza")

		So(NewEmptyInt64().ForceMapToString(func(int64) string { return "bonanza" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToNillableString(t *testing.T) {
	Convey("Int64.MapToNillableString", t, func() {
		So(NewInt64(55555).MapToNillableString(func(int64) *string {
			var tmp string = "bonanza"
			return &tmp
		}).Get(), ShouldEqual, "bonanza")

		So(NewEmptyInt64().MapToNillableString(func(int64) *string {
			var tmp string = "bonanza"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt64(55555).MapToNillableString(func(int64) *string { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToString(t *testing.T) {
	Convey("Int64.MapToString", t, func() {
		test1, err1 := NewInt64(55555).MapToString(func(int64) (string, error) {
			return "bonanza", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "bonanza")

		test2, err2 := NewEmptyInt64().MapToString(func(int64) (string, error) {
			return "bonanza", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt64(55555).MapToString(func(int64) (string, error) {
			return "bonanza", errors.New("string error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "string error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt64_ForceMapToUint(t *testing.T) {
	Convey("Int64.ForceMapToUint", t, func() {
		So(NewInt64(55555).ForceMapToUint(func(int64) uint { return 666 }).Get(),
			ShouldEqual, 666)

		So(NewEmptyInt64().ForceMapToUint(func(int64) uint { return 666 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToNillableUint(t *testing.T) {
	Convey("Int64.MapToNillableUint", t, func() {
		So(NewInt64(55555).MapToNillableUint(func(int64) *uint {
			var tmp uint = 666
			return &tmp
		}).Get(), ShouldEqual, 666)

		So(NewEmptyInt64().MapToNillableUint(func(int64) *uint {
			var tmp uint = 666
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt64(55555).MapToNillableUint(func(int64) *uint { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToUint(t *testing.T) {
	Convey("Int64.MapToUint", t, func() {
		test1, err1 := NewInt64(55555).MapToUint(func(int64) (uint, error) {
			return 666, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 666)

		test2, err2 := NewEmptyInt64().MapToUint(func(int64) (uint, error) {
			return 666, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt64(55555).MapToUint(func(int64) (uint, error) {
			return 666, errors.New("uint error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt64_ForceMapToUint16(t *testing.T) {
	Convey("Int64.ForceMapToUint16", t, func() {
		So(NewInt64(55555).ForceMapToUint16(func(int64) uint16 { return 888 }).Get(),
			ShouldEqual, 888)

		So(NewEmptyInt64().ForceMapToUint16(func(int64) uint16 { return 888 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToNillableUint16(t *testing.T) {
	Convey("Int64.MapToNillableUint16", t, func() {
		So(NewInt64(55555).MapToNillableUint16(func(int64) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).Get(), ShouldEqual, 888)

		So(NewEmptyInt64().MapToNillableUint16(func(int64) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt64(55555).MapToNillableUint16(func(int64) *uint16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToUint16(t *testing.T) {
	Convey("Int64.MapToUint16", t, func() {
		test1, err1 := NewInt64(55555).MapToUint16(func(int64) (uint16, error) {
			return 888, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 888)

		test2, err2 := NewEmptyInt64().MapToUint16(func(int64) (uint16, error) {
			return 888, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt64(55555).MapToUint16(func(int64) (uint16, error) {
			return 888, errors.New("uint16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt64_ForceMapToUint32(t *testing.T) {
	Convey("Int64.ForceMapToUint32", t, func() {
		So(NewInt64(55555).ForceMapToUint32(func(int64) uint32 { return 999 }).Get(),
			ShouldEqual, 999)

		So(NewEmptyInt64().ForceMapToUint32(func(int64) uint32 { return 999 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToNillableUint32(t *testing.T) {
	Convey("Int64.MapToNillableUint32", t, func() {
		So(NewInt64(55555).MapToNillableUint32(func(int64) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).Get(), ShouldEqual, 999)

		So(NewEmptyInt64().MapToNillableUint32(func(int64) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt64(55555).MapToNillableUint32(func(int64) *uint32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToUint32(t *testing.T) {
	Convey("Int64.MapToUint32", t, func() {
		test1, err1 := NewInt64(55555).MapToUint32(func(int64) (uint32, error) {
			return 999, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 999)

		test2, err2 := NewEmptyInt64().MapToUint32(func(int64) (uint32, error) {
			return 999, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt64(55555).MapToUint32(func(int64) (uint32, error) {
			return 999, errors.New("uint32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt64_ForceMapToUint64(t *testing.T) {
	Convey("Int64.ForceMapToUint64", t, func() {
		So(NewInt64(55555).ForceMapToUint64(func(int64) uint64 { return 101010 }).Get(),
			ShouldEqual, 101010)

		So(NewEmptyInt64().ForceMapToUint64(func(int64) uint64 { return 101010 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToNillableUint64(t *testing.T) {
	Convey("Int64.MapToNillableUint64", t, func() {
		So(NewInt64(55555).MapToNillableUint64(func(int64) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).Get(), ShouldEqual, 101010)

		So(NewEmptyInt64().MapToNillableUint64(func(int64) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt64(55555).MapToNillableUint64(func(int64) *uint64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToUint64(t *testing.T) {
	Convey("Int64.MapToUint64", t, func() {
		test1, err1 := NewInt64(55555).MapToUint64(func(int64) (uint64, error) {
			return 101010, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 101010)

		test2, err2 := NewEmptyInt64().MapToUint64(func(int64) (uint64, error) {
			return 101010, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt64(55555).MapToUint64(func(int64) (uint64, error) {
			return 101010, errors.New("uint64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt64_ForceMapToUint8(t *testing.T) {
	Convey("Int64.ForceMapToUint8", t, func() {
		So(NewInt64(55555).ForceMapToUint8(func(int64) uint8 { return 77 }).Get(),
			ShouldEqual, 77)

		So(NewEmptyInt64().ForceMapToUint8(func(int64) uint8 { return 77 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToNillableUint8(t *testing.T) {
	Convey("Int64.MapToNillableUint8", t, func() {
		So(NewInt64(55555).MapToNillableUint8(func(int64) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).Get(), ShouldEqual, 77)

		So(NewEmptyInt64().MapToNillableUint8(func(int64) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt64(55555).MapToNillableUint8(func(int64) *uint8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt64_MapToUint8(t *testing.T) {
	Convey("Int64.MapToUint8", t, func() {
		test1, err1 := NewInt64(55555).MapToUint8(func(int64) (uint8, error) {
			return 77, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 77)

		test2, err2 := NewEmptyInt64().MapToUint8(func(int64) (uint8, error) {
			return 77, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt64(55555).MapToUint8(func(int64) (uint8, error) {
			return 77, errors.New("uint8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}
