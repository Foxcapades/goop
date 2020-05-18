// Package option_test generated @ 2020-05-17 22:19:55.164878293 -0400 EDT m=+0.001439508
package option_test

import (
	"errors"
	"testing"

	. "github.com/Foxcapades/goop/v0/pkg/option"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewInt32(t *testing.T) {
	Convey("NewInt32", t, func() {
		test := NewInt32(1245)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyInt32(t *testing.T) {
	Convey("NewEmptyInt32", t, func() {
		test := NewEmptyInt32()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeInt32(t *testing.T) {
	Convey("NewMaybeInt32", t, func() {
		var val1 int32

		test1 := NewMaybeInt32(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeInt32(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestInt32_Get(t *testing.T) {
	Convey("Int32.Get", t, func() {
		test1 := NewInt32(1245)
		So(test1.Get(), ShouldEqual, 1245)

		test2 := NewEmptyInt32()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestInt32_Or(t *testing.T) {
	Convey("Int32.Or", t, func() {
		test1 := NewInt32(1245)
		So(test1.Or(1234), ShouldEqual, 1245)

		test2 := NewEmptyInt32()
		So(test2.Or(1245), ShouldEqual, 1245)
	})
}

func TestInt32_OrGet(t *testing.T) {
	Convey("Int32.OrGet", t, func() {
		test1 := NewInt32(1245)
		So(test1.OrGet(func() int32 { return 1234 }), ShouldEqual, 1245)

		test2 := NewEmptyInt32()
		So(test2.OrGet(func() int32 { return 1245 }), ShouldEqual, 1245)
	})
}

func TestInt32_OrPanicWith(t *testing.T) {
	Convey("Int32.OrPanicWith", t, func() {
		test1 := NewInt32(1245)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 1245)

		test2 := NewEmptyInt32()
		So(func() {
			test2.OrPanicWith(func() interface{} { return "test value" })
		}, ShouldPanicWith, "test value")
	})
}

func TestInt32_MapToNillable(t *testing.T) {
	Convey("Int32.MapToNillable", t, func() {
		test1 := NewInt32(1245)
		So(test1.MapToNillable(func(b int32) *int32 { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewInt32(1245)
		So(test2.MapToNillable(func(b int32) *int32 { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyInt32()
		So(func() {
			test3.MapToNillable(func(b int32) *int32 { panic("foo") }).IsNil()
		}, ShouldNotPanic)
	})
}

func TestInt32_ForceMapToBool(t *testing.T) {
	Convey("Int32.ForceMapToBool", t, func() {
		So(NewInt32(1245).ForceMapToBool(func(int32) bool { return true }).Get(),
			ShouldEqual, true)

		So(NewEmptyInt32().ForceMapToBool(func(int32) bool { return true }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToNillableBool(t *testing.T) {
	Convey("Int32.MapToNillableBool", t, func() {
		So(NewInt32(1245).MapToNillableBool(func(int32) *bool {
			var tmp bool = true
			return &tmp
		}).Get(), ShouldEqual, true)

		So(NewEmptyInt32().MapToNillableBool(func(int32) *bool {
			var tmp bool = true
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt32(1245).MapToNillableBool(func(int32) *bool { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToBool(t *testing.T) {
	Convey("Int32.MapToBool", t, func() {
		test1, err1 := NewInt32(1245).MapToBool(func(int32) (bool, error) {
			return true, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, true)

		test2, err2 := NewEmptyInt32().MapToBool(func(int32) (bool, error) {
			return true, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt32(1245).MapToBool(func(int32) (bool, error) {
			return true, errors.New("bool error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "bool error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt32_ForceMapToByte(t *testing.T) {
	Convey("Int32.ForceMapToByte", t, func() {
		So(NewInt32(1245).ForceMapToByte(func(int32) byte { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyInt32().ForceMapToByte(func(int32) byte { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToNillableByte(t *testing.T) {
	Convey("Int32.MapToNillableByte", t, func() {
		So(NewInt32(1245).MapToNillableByte(func(int32) *byte {
			var tmp byte = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyInt32().MapToNillableByte(func(int32) *byte {
			var tmp byte = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt32(1245).MapToNillableByte(func(int32) *byte { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToByte(t *testing.T) {
	Convey("Int32.MapToByte", t, func() {
		test1, err1 := NewInt32(1245).MapToByte(func(int32) (byte, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyInt32().MapToByte(func(int32) (byte, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt32(1245).MapToByte(func(int32) (byte, error) {
			return 222, errors.New("byte error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "byte error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt32_ForceMapToComplex128(t *testing.T) {
	Convey("Int32.ForceMapToComplex128", t, func() {
		So(NewInt32(1245).ForceMapToComplex128(func(int32) complex128 { return 987654 }).Get(),
			ShouldEqual, 987654)

		So(NewEmptyInt32().ForceMapToComplex128(func(int32) complex128 { return 987654 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToNillableComplex128(t *testing.T) {
	Convey("Int32.MapToNillableComplex128", t, func() {
		So(NewInt32(1245).MapToNillableComplex128(func(int32) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).Get(), ShouldEqual, 987654)

		So(NewEmptyInt32().MapToNillableComplex128(func(int32) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt32(1245).MapToNillableComplex128(func(int32) *complex128 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToComplex128(t *testing.T) {
	Convey("Int32.MapToComplex128", t, func() {
		test1, err1 := NewInt32(1245).MapToComplex128(func(int32) (complex128, error) {
			return 987654, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 987654)

		test2, err2 := NewEmptyInt32().MapToComplex128(func(int32) (complex128, error) {
			return 987654, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt32(1245).MapToComplex128(func(int32) (complex128, error) {
			return 987654, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt32_ForceMapToComplex64(t *testing.T) {
	Convey("Int32.ForceMapToComplex64", t, func() {
		So(NewInt32(1245).ForceMapToComplex64(func(int32) complex64 { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyInt32().ForceMapToComplex64(func(int32) complex64 { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToNillableComplex64(t *testing.T) {
	Convey("Int32.MapToNillableComplex64", t, func() {
		So(NewInt32(1245).MapToNillableComplex64(func(int32) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyInt32().MapToNillableComplex64(func(int32) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt32(1245).MapToNillableComplex64(func(int32) *complex64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToComplex64(t *testing.T) {
	Convey("Int32.MapToComplex64", t, func() {
		test1, err1 := NewInt32(1245).MapToComplex64(func(int32) (complex64, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyInt32().MapToComplex64(func(int32) (complex64, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt32(1245).MapToComplex64(func(int32) (complex64, error) {
			return 123456, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt32_ForceMapToFloat32(t *testing.T) {
	Convey("Int32.ForceMapToFloat32", t, func() {
		So(NewInt32(1245).ForceMapToFloat32(func(int32) float32 { return 32.3 }).Get(),
			ShouldEqual, 32.3)

		So(NewEmptyInt32().ForceMapToFloat32(func(int32) float32 { return 32.3 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToNillableFloat32(t *testing.T) {
	Convey("Int32.MapToNillableFloat32", t, func() {
		So(NewInt32(1245).MapToNillableFloat32(func(int32) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).Get(), ShouldEqual, 32.3)

		So(NewEmptyInt32().MapToNillableFloat32(func(int32) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt32(1245).MapToNillableFloat32(func(int32) *float32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToFloat32(t *testing.T) {
	Convey("Int32.MapToFloat32", t, func() {
		test1, err1 := NewInt32(1245).MapToFloat32(func(int32) (float32, error) {
			return 32.3, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 32.3)

		test2, err2 := NewEmptyInt32().MapToFloat32(func(int32) (float32, error) {
			return 32.3, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt32(1245).MapToFloat32(func(int32) (float32, error) {
			return 32.3, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt32_ForceMapToFloat64(t *testing.T) {
	Convey("Int32.ForceMapToFloat64", t, func() {
		So(NewInt32(1245).ForceMapToFloat64(func(int32) float64 { return 3.2455 }).Get(),
			ShouldEqual, 3.2455)

		So(NewEmptyInt32().ForceMapToFloat64(func(int32) float64 { return 3.2455 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToNillableFloat64(t *testing.T) {
	Convey("Int32.MapToNillableFloat64", t, func() {
		So(NewInt32(1245).MapToNillableFloat64(func(int32) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).Get(), ShouldEqual, 3.2455)

		So(NewEmptyInt32().MapToNillableFloat64(func(int32) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt32(1245).MapToNillableFloat64(func(int32) *float64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToFloat64(t *testing.T) {
	Convey("Int32.MapToFloat64", t, func() {
		test1, err1 := NewInt32(1245).MapToFloat64(func(int32) (float64, error) {
			return 3.2455, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 3.2455)

		test2, err2 := NewEmptyInt32().MapToFloat64(func(int32) (float64, error) {
			return 3.2455, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt32(1245).MapToFloat64(func(int32) (float64, error) {
			return 3.2455, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt32_ForceMapToInt(t *testing.T) {
	Convey("Int32.ForceMapToInt", t, func() {
		So(NewInt32(1245).ForceMapToInt(func(int32) int { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyInt32().ForceMapToInt(func(int32) int { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToNillableInt(t *testing.T) {
	Convey("Int32.MapToNillableInt", t, func() {
		So(NewInt32(1245).MapToNillableInt(func(int32) *int {
			var tmp int = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyInt32().MapToNillableInt(func(int32) *int {
			var tmp int = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt32(1245).MapToNillableInt(func(int32) *int { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToInt(t *testing.T) {
	Convey("Int32.MapToInt", t, func() {
		test1, err1 := NewInt32(1245).MapToInt(func(int32) (int, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyInt32().MapToInt(func(int32) (int, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt32(1245).MapToInt(func(int32) (int, error) {
			return 123456, errors.New("int error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt32_ForceMapToInt16(t *testing.T) {
	Convey("Int32.ForceMapToInt16", t, func() {
		So(NewInt32(1245).ForceMapToInt16(func(int32) int16 { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyInt32().ForceMapToInt16(func(int32) int16 { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToNillableInt16(t *testing.T) {
	Convey("Int32.MapToNillableInt16", t, func() {
		So(NewInt32(1245).MapToNillableInt16(func(int32) *int16 {
			var tmp int16 = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyInt32().MapToNillableInt16(func(int32) *int16 {
			var tmp int16 = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt32(1245).MapToNillableInt16(func(int32) *int16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToInt16(t *testing.T) {
	Convey("Int32.MapToInt16", t, func() {
		test1, err1 := NewInt32(1245).MapToInt16(func(int32) (int16, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyInt32().MapToInt16(func(int32) (int16, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt32(1245).MapToInt16(func(int32) (int16, error) {
			return 222, errors.New("int16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt32_ForceMapToInt64(t *testing.T) {
	Convey("Int32.ForceMapToInt64", t, func() {
		So(NewInt32(1245).ForceMapToInt64(func(int32) int64 { return 444 }).Get(),
			ShouldEqual, 444)

		So(NewEmptyInt32().ForceMapToInt64(func(int32) int64 { return 444 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToNillableInt64(t *testing.T) {
	Convey("Int32.MapToNillableInt64", t, func() {
		So(NewInt32(1245).MapToNillableInt64(func(int32) *int64 {
			var tmp int64 = 444
			return &tmp
		}).Get(), ShouldEqual, 444)

		So(NewEmptyInt32().MapToNillableInt64(func(int32) *int64 {
			var tmp int64 = 444
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt32(1245).MapToNillableInt64(func(int32) *int64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToInt64(t *testing.T) {
	Convey("Int32.MapToInt64", t, func() {
		test1, err1 := NewInt32(1245).MapToInt64(func(int32) (int64, error) {
			return 444, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 444)

		test2, err2 := NewEmptyInt32().MapToInt64(func(int32) (int64, error) {
			return 444, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt32(1245).MapToInt64(func(int32) (int64, error) {
			return 444, errors.New("int64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt32_ForceMapToInt8(t *testing.T) {
	Convey("Int32.ForceMapToInt8", t, func() {
		So(NewInt32(1245).ForceMapToInt8(func(int32) int8 { return 111 }).Get(),
			ShouldEqual, 111)

		So(NewEmptyInt32().ForceMapToInt8(func(int32) int8 { return 111 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToNillableInt8(t *testing.T) {
	Convey("Int32.MapToNillableInt8", t, func() {
		So(NewInt32(1245).MapToNillableInt8(func(int32) *int8 {
			var tmp int8 = 111
			return &tmp
		}).Get(), ShouldEqual, 111)

		So(NewEmptyInt32().MapToNillableInt8(func(int32) *int8 {
			var tmp int8 = 111
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt32(1245).MapToNillableInt8(func(int32) *int8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToInt8(t *testing.T) {
	Convey("Int32.MapToInt8", t, func() {
		test1, err1 := NewInt32(1245).MapToInt8(func(int32) (int8, error) {
			return 111, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 111)

		test2, err2 := NewEmptyInt32().MapToInt8(func(int32) (int8, error) {
			return 111, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt32(1245).MapToInt8(func(int32) (int8, error) {
			return 111, errors.New("int8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt32_ForceMapToUntyped(t *testing.T) {
	Convey("Int32.ForceMapToUntyped", t, func() {
		So(NewInt32(1245).ForceMapToUntyped(func(int32) interface{} { return "111111" }).Get(),
			ShouldEqual, "111111")

		So(NewEmptyInt32().ForceMapToUntyped(func(int32) interface{} { return "111111" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToNillableUntyped(t *testing.T) {
	Convey("Int32.MapToNillableUntyped", t, func() {
		So(NewInt32(1245).MapToNillableUntyped(func(int32) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).Get(), ShouldEqual, "111111")

		So(NewEmptyInt32().MapToNillableUntyped(func(int32) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt32(1245).MapToNillableUntyped(func(int32) *interface{} { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToUntyped(t *testing.T) {
	Convey("Int32.MapToUntyped", t, func() {
		test1, err1 := NewInt32(1245).MapToUntyped(func(int32) (interface{}, error) {
			return "111111", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "111111")

		test2, err2 := NewEmptyInt32().MapToUntyped(func(int32) (interface{}, error) {
			return "111111", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt32(1245).MapToUntyped(func(int32) (interface{}, error) {
			return "111111", errors.New("interface error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "interface error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt32_ForceMapToRune(t *testing.T) {
	Convey("Int32.ForceMapToRune", t, func() {
		So(NewInt32(1245).ForceMapToRune(func(int32) rune { return 555 }).Get(),
			ShouldEqual, 555)

		So(NewEmptyInt32().ForceMapToRune(func(int32) rune { return 555 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToNillableRune(t *testing.T) {
	Convey("Int32.MapToNillableRune", t, func() {
		So(NewInt32(1245).MapToNillableRune(func(int32) *rune {
			var tmp rune = 555
			return &tmp
		}).Get(), ShouldEqual, 555)

		So(NewEmptyInt32().MapToNillableRune(func(int32) *rune {
			var tmp rune = 555
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt32(1245).MapToNillableRune(func(int32) *rune { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToRune(t *testing.T) {
	Convey("Int32.MapToRune", t, func() {
		test1, err1 := NewInt32(1245).MapToRune(func(int32) (rune, error) {
			return 555, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 555)

		test2, err2 := NewEmptyInt32().MapToRune(func(int32) (rune, error) {
			return 555, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt32(1245).MapToRune(func(int32) (rune, error) {
			return 555, errors.New("rune error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "rune error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt32_ForceMapToString(t *testing.T) {
	Convey("Int32.ForceMapToString", t, func() {
		So(NewInt32(1245).ForceMapToString(func(int32) string { return "bonanza" }).Get(),
			ShouldEqual, "bonanza")

		So(NewEmptyInt32().ForceMapToString(func(int32) string { return "bonanza" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToNillableString(t *testing.T) {
	Convey("Int32.MapToNillableString", t, func() {
		So(NewInt32(1245).MapToNillableString(func(int32) *string {
			var tmp string = "bonanza"
			return &tmp
		}).Get(), ShouldEqual, "bonanza")

		So(NewEmptyInt32().MapToNillableString(func(int32) *string {
			var tmp string = "bonanza"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt32(1245).MapToNillableString(func(int32) *string { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToString(t *testing.T) {
	Convey("Int32.MapToString", t, func() {
		test1, err1 := NewInt32(1245).MapToString(func(int32) (string, error) {
			return "bonanza", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "bonanza")

		test2, err2 := NewEmptyInt32().MapToString(func(int32) (string, error) {
			return "bonanza", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt32(1245).MapToString(func(int32) (string, error) {
			return "bonanza", errors.New("string error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "string error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt32_ForceMapToUint(t *testing.T) {
	Convey("Int32.ForceMapToUint", t, func() {
		So(NewInt32(1245).ForceMapToUint(func(int32) uint { return 666 }).Get(),
			ShouldEqual, 666)

		So(NewEmptyInt32().ForceMapToUint(func(int32) uint { return 666 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToNillableUint(t *testing.T) {
	Convey("Int32.MapToNillableUint", t, func() {
		So(NewInt32(1245).MapToNillableUint(func(int32) *uint {
			var tmp uint = 666
			return &tmp
		}).Get(), ShouldEqual, 666)

		So(NewEmptyInt32().MapToNillableUint(func(int32) *uint {
			var tmp uint = 666
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt32(1245).MapToNillableUint(func(int32) *uint { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToUint(t *testing.T) {
	Convey("Int32.MapToUint", t, func() {
		test1, err1 := NewInt32(1245).MapToUint(func(int32) (uint, error) {
			return 666, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 666)

		test2, err2 := NewEmptyInt32().MapToUint(func(int32) (uint, error) {
			return 666, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt32(1245).MapToUint(func(int32) (uint, error) {
			return 666, errors.New("uint error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt32_ForceMapToUint16(t *testing.T) {
	Convey("Int32.ForceMapToUint16", t, func() {
		So(NewInt32(1245).ForceMapToUint16(func(int32) uint16 { return 888 }).Get(),
			ShouldEqual, 888)

		So(NewEmptyInt32().ForceMapToUint16(func(int32) uint16 { return 888 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToNillableUint16(t *testing.T) {
	Convey("Int32.MapToNillableUint16", t, func() {
		So(NewInt32(1245).MapToNillableUint16(func(int32) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).Get(), ShouldEqual, 888)

		So(NewEmptyInt32().MapToNillableUint16(func(int32) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt32(1245).MapToNillableUint16(func(int32) *uint16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToUint16(t *testing.T) {
	Convey("Int32.MapToUint16", t, func() {
		test1, err1 := NewInt32(1245).MapToUint16(func(int32) (uint16, error) {
			return 888, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 888)

		test2, err2 := NewEmptyInt32().MapToUint16(func(int32) (uint16, error) {
			return 888, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt32(1245).MapToUint16(func(int32) (uint16, error) {
			return 888, errors.New("uint16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt32_ForceMapToUint32(t *testing.T) {
	Convey("Int32.ForceMapToUint32", t, func() {
		So(NewInt32(1245).ForceMapToUint32(func(int32) uint32 { return 999 }).Get(),
			ShouldEqual, 999)

		So(NewEmptyInt32().ForceMapToUint32(func(int32) uint32 { return 999 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToNillableUint32(t *testing.T) {
	Convey("Int32.MapToNillableUint32", t, func() {
		So(NewInt32(1245).MapToNillableUint32(func(int32) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).Get(), ShouldEqual, 999)

		So(NewEmptyInt32().MapToNillableUint32(func(int32) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt32(1245).MapToNillableUint32(func(int32) *uint32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToUint32(t *testing.T) {
	Convey("Int32.MapToUint32", t, func() {
		test1, err1 := NewInt32(1245).MapToUint32(func(int32) (uint32, error) {
			return 999, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 999)

		test2, err2 := NewEmptyInt32().MapToUint32(func(int32) (uint32, error) {
			return 999, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt32(1245).MapToUint32(func(int32) (uint32, error) {
			return 999, errors.New("uint32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt32_ForceMapToUint64(t *testing.T) {
	Convey("Int32.ForceMapToUint64", t, func() {
		So(NewInt32(1245).ForceMapToUint64(func(int32) uint64 { return 101010 }).Get(),
			ShouldEqual, 101010)

		So(NewEmptyInt32().ForceMapToUint64(func(int32) uint64 { return 101010 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToNillableUint64(t *testing.T) {
	Convey("Int32.MapToNillableUint64", t, func() {
		So(NewInt32(1245).MapToNillableUint64(func(int32) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).Get(), ShouldEqual, 101010)

		So(NewEmptyInt32().MapToNillableUint64(func(int32) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt32(1245).MapToNillableUint64(func(int32) *uint64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToUint64(t *testing.T) {
	Convey("Int32.MapToUint64", t, func() {
		test1, err1 := NewInt32(1245).MapToUint64(func(int32) (uint64, error) {
			return 101010, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 101010)

		test2, err2 := NewEmptyInt32().MapToUint64(func(int32) (uint64, error) {
			return 101010, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt32(1245).MapToUint64(func(int32) (uint64, error) {
			return 101010, errors.New("uint64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt32_ForceMapToUint8(t *testing.T) {
	Convey("Int32.ForceMapToUint8", t, func() {
		So(NewInt32(1245).ForceMapToUint8(func(int32) uint8 { return 77 }).Get(),
			ShouldEqual, 77)

		So(NewEmptyInt32().ForceMapToUint8(func(int32) uint8 { return 77 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToNillableUint8(t *testing.T) {
	Convey("Int32.MapToNillableUint8", t, func() {
		So(NewInt32(1245).MapToNillableUint8(func(int32) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).Get(), ShouldEqual, 77)

		So(NewEmptyInt32().MapToNillableUint8(func(int32) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt32(1245).MapToNillableUint8(func(int32) *uint8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt32_MapToUint8(t *testing.T) {
	Convey("Int32.MapToUint8", t, func() {
		test1, err1 := NewInt32(1245).MapToUint8(func(int32) (uint8, error) {
			return 77, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 77)

		test2, err2 := NewEmptyInt32().MapToUint8(func(int32) (uint8, error) {
			return 77, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt32(1245).MapToUint8(func(int32) (uint8, error) {
			return 77, errors.New("uint8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}
