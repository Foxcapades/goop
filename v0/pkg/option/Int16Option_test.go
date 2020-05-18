// Package option_test generated @ 2020-05-17 22:19:55.164878293 -0400 EDT m=+0.001439508
package option_test

import (
	"errors"
	"testing"

	. "github.com/Foxcapades/goop/v0/pkg/option"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewInt16(t *testing.T) {
	Convey("NewInt16", t, func() {
		test := NewInt16(256)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyInt16(t *testing.T) {
	Convey("NewEmptyInt16", t, func() {
		test := NewEmptyInt16()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeInt16(t *testing.T) {
	Convey("NewMaybeInt16", t, func() {
		var val1 int16

		test1 := NewMaybeInt16(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeInt16(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestInt16_Get(t *testing.T) {
	Convey("Int16.Get", t, func() {
		test1 := NewInt16(256)
		So(test1.Get(), ShouldEqual, 256)

		test2 := NewEmptyInt16()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestInt16_Or(t *testing.T) {
	Convey("Int16.Or", t, func() {
		test1 := NewInt16(256)
		So(test1.Or(585), ShouldEqual, 256)

		test2 := NewEmptyInt16()
		So(test2.Or(256), ShouldEqual, 256)
	})
}

func TestInt16_OrGet(t *testing.T) {
	Convey("Int16.OrGet", t, func() {
		test1 := NewInt16(256)
		So(test1.OrGet(func() int16 { return 585 }), ShouldEqual, 256)

		test2 := NewEmptyInt16()
		So(test2.OrGet(func() int16 { return 256 }), ShouldEqual, 256)
	})
}

func TestInt16_OrPanicWith(t *testing.T) {
	Convey("Int16.OrPanicWith", t, func() {
		test1 := NewInt16(256)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 256)

		test2 := NewEmptyInt16()
		So(func() {
			test2.OrPanicWith(func() interface{} { return "test value" })
		}, ShouldPanicWith, "test value")
	})
}

func TestInt16_MapToNillable(t *testing.T) {
	Convey("Int16.MapToNillable", t, func() {
		test1 := NewInt16(256)
		So(test1.MapToNillable(func(b int16) *int16 { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewInt16(256)
		So(test2.MapToNillable(func(b int16) *int16 { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyInt16()
		So(func() {
			test3.MapToNillable(func(b int16) *int16 { panic("foo") }).IsNil()
		}, ShouldNotPanic)
	})
}

func TestInt16_ForceMapToBool(t *testing.T) {
	Convey("Int16.ForceMapToBool", t, func() {
		So(NewInt16(256).ForceMapToBool(func(int16) bool { return true }).Get(),
			ShouldEqual, true)

		So(NewEmptyInt16().ForceMapToBool(func(int16) bool { return true }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToNillableBool(t *testing.T) {
	Convey("Int16.MapToNillableBool", t, func() {
		So(NewInt16(256).MapToNillableBool(func(int16) *bool {
			var tmp bool = true
			return &tmp
		}).Get(), ShouldEqual, true)

		So(NewEmptyInt16().MapToNillableBool(func(int16) *bool {
			var tmp bool = true
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt16(256).MapToNillableBool(func(int16) *bool { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToBool(t *testing.T) {
	Convey("Int16.MapToBool", t, func() {
		test1, err1 := NewInt16(256).MapToBool(func(int16) (bool, error) {
			return true, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, true)

		test2, err2 := NewEmptyInt16().MapToBool(func(int16) (bool, error) {
			return true, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt16(256).MapToBool(func(int16) (bool, error) {
			return true, errors.New("bool error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "bool error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt16_ForceMapToByte(t *testing.T) {
	Convey("Int16.ForceMapToByte", t, func() {
		So(NewInt16(256).ForceMapToByte(func(int16) byte { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyInt16().ForceMapToByte(func(int16) byte { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToNillableByte(t *testing.T) {
	Convey("Int16.MapToNillableByte", t, func() {
		So(NewInt16(256).MapToNillableByte(func(int16) *byte {
			var tmp byte = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyInt16().MapToNillableByte(func(int16) *byte {
			var tmp byte = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt16(256).MapToNillableByte(func(int16) *byte { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToByte(t *testing.T) {
	Convey("Int16.MapToByte", t, func() {
		test1, err1 := NewInt16(256).MapToByte(func(int16) (byte, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyInt16().MapToByte(func(int16) (byte, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt16(256).MapToByte(func(int16) (byte, error) {
			return 222, errors.New("byte error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "byte error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt16_ForceMapToComplex128(t *testing.T) {
	Convey("Int16.ForceMapToComplex128", t, func() {
		So(NewInt16(256).ForceMapToComplex128(func(int16) complex128 { return 987654 }).Get(),
			ShouldEqual, 987654)

		So(NewEmptyInt16().ForceMapToComplex128(func(int16) complex128 { return 987654 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToNillableComplex128(t *testing.T) {
	Convey("Int16.MapToNillableComplex128", t, func() {
		So(NewInt16(256).MapToNillableComplex128(func(int16) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).Get(), ShouldEqual, 987654)

		So(NewEmptyInt16().MapToNillableComplex128(func(int16) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt16(256).MapToNillableComplex128(func(int16) *complex128 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToComplex128(t *testing.T) {
	Convey("Int16.MapToComplex128", t, func() {
		test1, err1 := NewInt16(256).MapToComplex128(func(int16) (complex128, error) {
			return 987654, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 987654)

		test2, err2 := NewEmptyInt16().MapToComplex128(func(int16) (complex128, error) {
			return 987654, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt16(256).MapToComplex128(func(int16) (complex128, error) {
			return 987654, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt16_ForceMapToComplex64(t *testing.T) {
	Convey("Int16.ForceMapToComplex64", t, func() {
		So(NewInt16(256).ForceMapToComplex64(func(int16) complex64 { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyInt16().ForceMapToComplex64(func(int16) complex64 { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToNillableComplex64(t *testing.T) {
	Convey("Int16.MapToNillableComplex64", t, func() {
		So(NewInt16(256).MapToNillableComplex64(func(int16) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyInt16().MapToNillableComplex64(func(int16) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt16(256).MapToNillableComplex64(func(int16) *complex64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToComplex64(t *testing.T) {
	Convey("Int16.MapToComplex64", t, func() {
		test1, err1 := NewInt16(256).MapToComplex64(func(int16) (complex64, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyInt16().MapToComplex64(func(int16) (complex64, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt16(256).MapToComplex64(func(int16) (complex64, error) {
			return 123456, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt16_ForceMapToFloat32(t *testing.T) {
	Convey("Int16.ForceMapToFloat32", t, func() {
		So(NewInt16(256).ForceMapToFloat32(func(int16) float32 { return 32.3 }).Get(),
			ShouldEqual, 32.3)

		So(NewEmptyInt16().ForceMapToFloat32(func(int16) float32 { return 32.3 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToNillableFloat32(t *testing.T) {
	Convey("Int16.MapToNillableFloat32", t, func() {
		So(NewInt16(256).MapToNillableFloat32(func(int16) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).Get(), ShouldEqual, 32.3)

		So(NewEmptyInt16().MapToNillableFloat32(func(int16) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt16(256).MapToNillableFloat32(func(int16) *float32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToFloat32(t *testing.T) {
	Convey("Int16.MapToFloat32", t, func() {
		test1, err1 := NewInt16(256).MapToFloat32(func(int16) (float32, error) {
			return 32.3, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 32.3)

		test2, err2 := NewEmptyInt16().MapToFloat32(func(int16) (float32, error) {
			return 32.3, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt16(256).MapToFloat32(func(int16) (float32, error) {
			return 32.3, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt16_ForceMapToFloat64(t *testing.T) {
	Convey("Int16.ForceMapToFloat64", t, func() {
		So(NewInt16(256).ForceMapToFloat64(func(int16) float64 { return 3.2455 }).Get(),
			ShouldEqual, 3.2455)

		So(NewEmptyInt16().ForceMapToFloat64(func(int16) float64 { return 3.2455 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToNillableFloat64(t *testing.T) {
	Convey("Int16.MapToNillableFloat64", t, func() {
		So(NewInt16(256).MapToNillableFloat64(func(int16) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).Get(), ShouldEqual, 3.2455)

		So(NewEmptyInt16().MapToNillableFloat64(func(int16) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt16(256).MapToNillableFloat64(func(int16) *float64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToFloat64(t *testing.T) {
	Convey("Int16.MapToFloat64", t, func() {
		test1, err1 := NewInt16(256).MapToFloat64(func(int16) (float64, error) {
			return 3.2455, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 3.2455)

		test2, err2 := NewEmptyInt16().MapToFloat64(func(int16) (float64, error) {
			return 3.2455, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt16(256).MapToFloat64(func(int16) (float64, error) {
			return 3.2455, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt16_ForceMapToInt(t *testing.T) {
	Convey("Int16.ForceMapToInt", t, func() {
		So(NewInt16(256).ForceMapToInt(func(int16) int { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyInt16().ForceMapToInt(func(int16) int { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToNillableInt(t *testing.T) {
	Convey("Int16.MapToNillableInt", t, func() {
		So(NewInt16(256).MapToNillableInt(func(int16) *int {
			var tmp int = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyInt16().MapToNillableInt(func(int16) *int {
			var tmp int = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt16(256).MapToNillableInt(func(int16) *int { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToInt(t *testing.T) {
	Convey("Int16.MapToInt", t, func() {
		test1, err1 := NewInt16(256).MapToInt(func(int16) (int, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyInt16().MapToInt(func(int16) (int, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt16(256).MapToInt(func(int16) (int, error) {
			return 123456, errors.New("int error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt16_ForceMapToInt32(t *testing.T) {
	Convey("Int16.ForceMapToInt32", t, func() {
		So(NewInt16(256).ForceMapToInt32(func(int16) int32 { return 333 }).Get(),
			ShouldEqual, 333)

		So(NewEmptyInt16().ForceMapToInt32(func(int16) int32 { return 333 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToNillableInt32(t *testing.T) {
	Convey("Int16.MapToNillableInt32", t, func() {
		So(NewInt16(256).MapToNillableInt32(func(int16) *int32 {
			var tmp int32 = 333
			return &tmp
		}).Get(), ShouldEqual, 333)

		So(NewEmptyInt16().MapToNillableInt32(func(int16) *int32 {
			var tmp int32 = 333
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt16(256).MapToNillableInt32(func(int16) *int32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToInt32(t *testing.T) {
	Convey("Int16.MapToInt32", t, func() {
		test1, err1 := NewInt16(256).MapToInt32(func(int16) (int32, error) {
			return 333, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 333)

		test2, err2 := NewEmptyInt16().MapToInt32(func(int16) (int32, error) {
			return 333, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt16(256).MapToInt32(func(int16) (int32, error) {
			return 333, errors.New("int32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt16_ForceMapToInt64(t *testing.T) {
	Convey("Int16.ForceMapToInt64", t, func() {
		So(NewInt16(256).ForceMapToInt64(func(int16) int64 { return 444 }).Get(),
			ShouldEqual, 444)

		So(NewEmptyInt16().ForceMapToInt64(func(int16) int64 { return 444 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToNillableInt64(t *testing.T) {
	Convey("Int16.MapToNillableInt64", t, func() {
		So(NewInt16(256).MapToNillableInt64(func(int16) *int64 {
			var tmp int64 = 444
			return &tmp
		}).Get(), ShouldEqual, 444)

		So(NewEmptyInt16().MapToNillableInt64(func(int16) *int64 {
			var tmp int64 = 444
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt16(256).MapToNillableInt64(func(int16) *int64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToInt64(t *testing.T) {
	Convey("Int16.MapToInt64", t, func() {
		test1, err1 := NewInt16(256).MapToInt64(func(int16) (int64, error) {
			return 444, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 444)

		test2, err2 := NewEmptyInt16().MapToInt64(func(int16) (int64, error) {
			return 444, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt16(256).MapToInt64(func(int16) (int64, error) {
			return 444, errors.New("int64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt16_ForceMapToInt8(t *testing.T) {
	Convey("Int16.ForceMapToInt8", t, func() {
		So(NewInt16(256).ForceMapToInt8(func(int16) int8 { return 111 }).Get(),
			ShouldEqual, 111)

		So(NewEmptyInt16().ForceMapToInt8(func(int16) int8 { return 111 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToNillableInt8(t *testing.T) {
	Convey("Int16.MapToNillableInt8", t, func() {
		So(NewInt16(256).MapToNillableInt8(func(int16) *int8 {
			var tmp int8 = 111
			return &tmp
		}).Get(), ShouldEqual, 111)

		So(NewEmptyInt16().MapToNillableInt8(func(int16) *int8 {
			var tmp int8 = 111
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt16(256).MapToNillableInt8(func(int16) *int8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToInt8(t *testing.T) {
	Convey("Int16.MapToInt8", t, func() {
		test1, err1 := NewInt16(256).MapToInt8(func(int16) (int8, error) {
			return 111, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 111)

		test2, err2 := NewEmptyInt16().MapToInt8(func(int16) (int8, error) {
			return 111, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt16(256).MapToInt8(func(int16) (int8, error) {
			return 111, errors.New("int8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt16_ForceMapToUntyped(t *testing.T) {
	Convey("Int16.ForceMapToUntyped", t, func() {
		So(NewInt16(256).ForceMapToUntyped(func(int16) interface{} { return "111111" }).Get(),
			ShouldEqual, "111111")

		So(NewEmptyInt16().ForceMapToUntyped(func(int16) interface{} { return "111111" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToNillableUntyped(t *testing.T) {
	Convey("Int16.MapToNillableUntyped", t, func() {
		So(NewInt16(256).MapToNillableUntyped(func(int16) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).Get(), ShouldEqual, "111111")

		So(NewEmptyInt16().MapToNillableUntyped(func(int16) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt16(256).MapToNillableUntyped(func(int16) *interface{} { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToUntyped(t *testing.T) {
	Convey("Int16.MapToUntyped", t, func() {
		test1, err1 := NewInt16(256).MapToUntyped(func(int16) (interface{}, error) {
			return "111111", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "111111")

		test2, err2 := NewEmptyInt16().MapToUntyped(func(int16) (interface{}, error) {
			return "111111", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt16(256).MapToUntyped(func(int16) (interface{}, error) {
			return "111111", errors.New("interface error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "interface error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt16_ForceMapToRune(t *testing.T) {
	Convey("Int16.ForceMapToRune", t, func() {
		So(NewInt16(256).ForceMapToRune(func(int16) rune { return 555 }).Get(),
			ShouldEqual, 555)

		So(NewEmptyInt16().ForceMapToRune(func(int16) rune { return 555 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToNillableRune(t *testing.T) {
	Convey("Int16.MapToNillableRune", t, func() {
		So(NewInt16(256).MapToNillableRune(func(int16) *rune {
			var tmp rune = 555
			return &tmp
		}).Get(), ShouldEqual, 555)

		So(NewEmptyInt16().MapToNillableRune(func(int16) *rune {
			var tmp rune = 555
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt16(256).MapToNillableRune(func(int16) *rune { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToRune(t *testing.T) {
	Convey("Int16.MapToRune", t, func() {
		test1, err1 := NewInt16(256).MapToRune(func(int16) (rune, error) {
			return 555, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 555)

		test2, err2 := NewEmptyInt16().MapToRune(func(int16) (rune, error) {
			return 555, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt16(256).MapToRune(func(int16) (rune, error) {
			return 555, errors.New("rune error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "rune error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt16_ForceMapToString(t *testing.T) {
	Convey("Int16.ForceMapToString", t, func() {
		So(NewInt16(256).ForceMapToString(func(int16) string { return "bonanza" }).Get(),
			ShouldEqual, "bonanza")

		So(NewEmptyInt16().ForceMapToString(func(int16) string { return "bonanza" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToNillableString(t *testing.T) {
	Convey("Int16.MapToNillableString", t, func() {
		So(NewInt16(256).MapToNillableString(func(int16) *string {
			var tmp string = "bonanza"
			return &tmp
		}).Get(), ShouldEqual, "bonanza")

		So(NewEmptyInt16().MapToNillableString(func(int16) *string {
			var tmp string = "bonanza"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt16(256).MapToNillableString(func(int16) *string { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToString(t *testing.T) {
	Convey("Int16.MapToString", t, func() {
		test1, err1 := NewInt16(256).MapToString(func(int16) (string, error) {
			return "bonanza", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "bonanza")

		test2, err2 := NewEmptyInt16().MapToString(func(int16) (string, error) {
			return "bonanza", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt16(256).MapToString(func(int16) (string, error) {
			return "bonanza", errors.New("string error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "string error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt16_ForceMapToUint(t *testing.T) {
	Convey("Int16.ForceMapToUint", t, func() {
		So(NewInt16(256).ForceMapToUint(func(int16) uint { return 666 }).Get(),
			ShouldEqual, 666)

		So(NewEmptyInt16().ForceMapToUint(func(int16) uint { return 666 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToNillableUint(t *testing.T) {
	Convey("Int16.MapToNillableUint", t, func() {
		So(NewInt16(256).MapToNillableUint(func(int16) *uint {
			var tmp uint = 666
			return &tmp
		}).Get(), ShouldEqual, 666)

		So(NewEmptyInt16().MapToNillableUint(func(int16) *uint {
			var tmp uint = 666
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt16(256).MapToNillableUint(func(int16) *uint { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToUint(t *testing.T) {
	Convey("Int16.MapToUint", t, func() {
		test1, err1 := NewInt16(256).MapToUint(func(int16) (uint, error) {
			return 666, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 666)

		test2, err2 := NewEmptyInt16().MapToUint(func(int16) (uint, error) {
			return 666, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt16(256).MapToUint(func(int16) (uint, error) {
			return 666, errors.New("uint error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt16_ForceMapToUint16(t *testing.T) {
	Convey("Int16.ForceMapToUint16", t, func() {
		So(NewInt16(256).ForceMapToUint16(func(int16) uint16 { return 888 }).Get(),
			ShouldEqual, 888)

		So(NewEmptyInt16().ForceMapToUint16(func(int16) uint16 { return 888 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToNillableUint16(t *testing.T) {
	Convey("Int16.MapToNillableUint16", t, func() {
		So(NewInt16(256).MapToNillableUint16(func(int16) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).Get(), ShouldEqual, 888)

		So(NewEmptyInt16().MapToNillableUint16(func(int16) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt16(256).MapToNillableUint16(func(int16) *uint16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToUint16(t *testing.T) {
	Convey("Int16.MapToUint16", t, func() {
		test1, err1 := NewInt16(256).MapToUint16(func(int16) (uint16, error) {
			return 888, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 888)

		test2, err2 := NewEmptyInt16().MapToUint16(func(int16) (uint16, error) {
			return 888, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt16(256).MapToUint16(func(int16) (uint16, error) {
			return 888, errors.New("uint16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt16_ForceMapToUint32(t *testing.T) {
	Convey("Int16.ForceMapToUint32", t, func() {
		So(NewInt16(256).ForceMapToUint32(func(int16) uint32 { return 999 }).Get(),
			ShouldEqual, 999)

		So(NewEmptyInt16().ForceMapToUint32(func(int16) uint32 { return 999 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToNillableUint32(t *testing.T) {
	Convey("Int16.MapToNillableUint32", t, func() {
		So(NewInt16(256).MapToNillableUint32(func(int16) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).Get(), ShouldEqual, 999)

		So(NewEmptyInt16().MapToNillableUint32(func(int16) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt16(256).MapToNillableUint32(func(int16) *uint32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToUint32(t *testing.T) {
	Convey("Int16.MapToUint32", t, func() {
		test1, err1 := NewInt16(256).MapToUint32(func(int16) (uint32, error) {
			return 999, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 999)

		test2, err2 := NewEmptyInt16().MapToUint32(func(int16) (uint32, error) {
			return 999, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt16(256).MapToUint32(func(int16) (uint32, error) {
			return 999, errors.New("uint32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt16_ForceMapToUint64(t *testing.T) {
	Convey("Int16.ForceMapToUint64", t, func() {
		So(NewInt16(256).ForceMapToUint64(func(int16) uint64 { return 101010 }).Get(),
			ShouldEqual, 101010)

		So(NewEmptyInt16().ForceMapToUint64(func(int16) uint64 { return 101010 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToNillableUint64(t *testing.T) {
	Convey("Int16.MapToNillableUint64", t, func() {
		So(NewInt16(256).MapToNillableUint64(func(int16) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).Get(), ShouldEqual, 101010)

		So(NewEmptyInt16().MapToNillableUint64(func(int16) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt16(256).MapToNillableUint64(func(int16) *uint64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToUint64(t *testing.T) {
	Convey("Int16.MapToUint64", t, func() {
		test1, err1 := NewInt16(256).MapToUint64(func(int16) (uint64, error) {
			return 101010, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 101010)

		test2, err2 := NewEmptyInt16().MapToUint64(func(int16) (uint64, error) {
			return 101010, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt16(256).MapToUint64(func(int16) (uint64, error) {
			return 101010, errors.New("uint64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt16_ForceMapToUint8(t *testing.T) {
	Convey("Int16.ForceMapToUint8", t, func() {
		So(NewInt16(256).ForceMapToUint8(func(int16) uint8 { return 77 }).Get(),
			ShouldEqual, 77)

		So(NewEmptyInt16().ForceMapToUint8(func(int16) uint8 { return 77 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToNillableUint8(t *testing.T) {
	Convey("Int16.MapToNillableUint8", t, func() {
		So(NewInt16(256).MapToNillableUint8(func(int16) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).Get(), ShouldEqual, 77)

		So(NewEmptyInt16().MapToNillableUint8(func(int16) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt16(256).MapToNillableUint8(func(int16) *uint8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt16_MapToUint8(t *testing.T) {
	Convey("Int16.MapToUint8", t, func() {
		test1, err1 := NewInt16(256).MapToUint8(func(int16) (uint8, error) {
			return 77, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 77)

		test2, err2 := NewEmptyInt16().MapToUint8(func(int16) (uint8, error) {
			return 77, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt16(256).MapToUint8(func(int16) (uint8, error) {
			return 77, errors.New("uint8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}
