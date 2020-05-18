// Package option_test generated @ 2020-05-17 22:19:55.164878293 -0400 EDT m=+0.001439508
package option_test

import (
	"errors"
	"testing"

	. "github.com/Foxcapades/goop/v0/pkg/option"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewInt8(t *testing.T) {
	Convey("NewInt8", t, func() {
		test := NewInt8(127)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyInt8(t *testing.T) {
	Convey("NewEmptyInt8", t, func() {
		test := NewEmptyInt8()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeInt8(t *testing.T) {
	Convey("NewMaybeInt8", t, func() {
		var val1 int8

		test1 := NewMaybeInt8(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeInt8(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestInt8_Get(t *testing.T) {
	Convey("Int8.Get", t, func() {
		test1 := NewInt8(127)
		So(test1.Get(), ShouldEqual, 127)

		test2 := NewEmptyInt8()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestInt8_Or(t *testing.T) {
	Convey("Int8.Or", t, func() {
		test1 := NewInt8(127)
		So(test1.Or(-128), ShouldEqual, 127)

		test2 := NewEmptyInt8()
		So(test2.Or(127), ShouldEqual, 127)
	})
}

func TestInt8_OrGet(t *testing.T) {
	Convey("Int8.OrGet", t, func() {
		test1 := NewInt8(127)
		So(test1.OrGet(func() int8 { return -128 }), ShouldEqual, 127)

		test2 := NewEmptyInt8()
		So(test2.OrGet(func() int8 { return 127 }), ShouldEqual, 127)
	})
}

func TestInt8_OrPanicWith(t *testing.T) {
	Convey("Int8.OrPanicWith", t, func() {
		test1 := NewInt8(127)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 127)

		test2 := NewEmptyInt8()
		So(func() {
			test2.OrPanicWith(func() interface{} { return "test value" })
		}, ShouldPanicWith, "test value")
	})
}

func TestInt8_MapToNillable(t *testing.T) {
	Convey("Int8.MapToNillable", t, func() {
		test1 := NewInt8(127)
		So(test1.MapToNillable(func(b int8) *int8 { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewInt8(127)
		So(test2.MapToNillable(func(b int8) *int8 { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyInt8()
		So(func() {
			test3.MapToNillable(func(b int8) *int8 { panic("foo") }).IsNil()
		}, ShouldNotPanic)
	})
}

func TestInt8_ForceMapToBool(t *testing.T) {
	Convey("Int8.ForceMapToBool", t, func() {
		So(NewInt8(127).ForceMapToBool(func(int8) bool { return true }).Get(),
			ShouldEqual, true)

		So(NewEmptyInt8().ForceMapToBool(func(int8) bool { return true }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToNillableBool(t *testing.T) {
	Convey("Int8.MapToNillableBool", t, func() {
		So(NewInt8(127).MapToNillableBool(func(int8) *bool {
			var tmp bool = true
			return &tmp
		}).Get(), ShouldEqual, true)

		So(NewEmptyInt8().MapToNillableBool(func(int8) *bool {
			var tmp bool = true
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt8(127).MapToNillableBool(func(int8) *bool { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToBool(t *testing.T) {
	Convey("Int8.MapToBool", t, func() {
		test1, err1 := NewInt8(127).MapToBool(func(int8) (bool, error) {
			return true, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, true)

		test2, err2 := NewEmptyInt8().MapToBool(func(int8) (bool, error) {
			return true, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt8(127).MapToBool(func(int8) (bool, error) {
			return true, errors.New("bool error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "bool error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt8_ForceMapToByte(t *testing.T) {
	Convey("Int8.ForceMapToByte", t, func() {
		So(NewInt8(127).ForceMapToByte(func(int8) byte { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyInt8().ForceMapToByte(func(int8) byte { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToNillableByte(t *testing.T) {
	Convey("Int8.MapToNillableByte", t, func() {
		So(NewInt8(127).MapToNillableByte(func(int8) *byte {
			var tmp byte = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyInt8().MapToNillableByte(func(int8) *byte {
			var tmp byte = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt8(127).MapToNillableByte(func(int8) *byte { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToByte(t *testing.T) {
	Convey("Int8.MapToByte", t, func() {
		test1, err1 := NewInt8(127).MapToByte(func(int8) (byte, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyInt8().MapToByte(func(int8) (byte, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt8(127).MapToByte(func(int8) (byte, error) {
			return 222, errors.New("byte error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "byte error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt8_ForceMapToComplex128(t *testing.T) {
	Convey("Int8.ForceMapToComplex128", t, func() {
		So(NewInt8(127).ForceMapToComplex128(func(int8) complex128 { return 987654 }).Get(),
			ShouldEqual, 987654)

		So(NewEmptyInt8().ForceMapToComplex128(func(int8) complex128 { return 987654 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToNillableComplex128(t *testing.T) {
	Convey("Int8.MapToNillableComplex128", t, func() {
		So(NewInt8(127).MapToNillableComplex128(func(int8) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).Get(), ShouldEqual, 987654)

		So(NewEmptyInt8().MapToNillableComplex128(func(int8) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt8(127).MapToNillableComplex128(func(int8) *complex128 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToComplex128(t *testing.T) {
	Convey("Int8.MapToComplex128", t, func() {
		test1, err1 := NewInt8(127).MapToComplex128(func(int8) (complex128, error) {
			return 987654, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 987654)

		test2, err2 := NewEmptyInt8().MapToComplex128(func(int8) (complex128, error) {
			return 987654, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt8(127).MapToComplex128(func(int8) (complex128, error) {
			return 987654, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt8_ForceMapToComplex64(t *testing.T) {
	Convey("Int8.ForceMapToComplex64", t, func() {
		So(NewInt8(127).ForceMapToComplex64(func(int8) complex64 { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyInt8().ForceMapToComplex64(func(int8) complex64 { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToNillableComplex64(t *testing.T) {
	Convey("Int8.MapToNillableComplex64", t, func() {
		So(NewInt8(127).MapToNillableComplex64(func(int8) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyInt8().MapToNillableComplex64(func(int8) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt8(127).MapToNillableComplex64(func(int8) *complex64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToComplex64(t *testing.T) {
	Convey("Int8.MapToComplex64", t, func() {
		test1, err1 := NewInt8(127).MapToComplex64(func(int8) (complex64, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyInt8().MapToComplex64(func(int8) (complex64, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt8(127).MapToComplex64(func(int8) (complex64, error) {
			return 123456, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt8_ForceMapToFloat32(t *testing.T) {
	Convey("Int8.ForceMapToFloat32", t, func() {
		So(NewInt8(127).ForceMapToFloat32(func(int8) float32 { return 32.3 }).Get(),
			ShouldEqual, 32.3)

		So(NewEmptyInt8().ForceMapToFloat32(func(int8) float32 { return 32.3 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToNillableFloat32(t *testing.T) {
	Convey("Int8.MapToNillableFloat32", t, func() {
		So(NewInt8(127).MapToNillableFloat32(func(int8) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).Get(), ShouldEqual, 32.3)

		So(NewEmptyInt8().MapToNillableFloat32(func(int8) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt8(127).MapToNillableFloat32(func(int8) *float32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToFloat32(t *testing.T) {
	Convey("Int8.MapToFloat32", t, func() {
		test1, err1 := NewInt8(127).MapToFloat32(func(int8) (float32, error) {
			return 32.3, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 32.3)

		test2, err2 := NewEmptyInt8().MapToFloat32(func(int8) (float32, error) {
			return 32.3, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt8(127).MapToFloat32(func(int8) (float32, error) {
			return 32.3, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt8_ForceMapToFloat64(t *testing.T) {
	Convey("Int8.ForceMapToFloat64", t, func() {
		So(NewInt8(127).ForceMapToFloat64(func(int8) float64 { return 3.2455 }).Get(),
			ShouldEqual, 3.2455)

		So(NewEmptyInt8().ForceMapToFloat64(func(int8) float64 { return 3.2455 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToNillableFloat64(t *testing.T) {
	Convey("Int8.MapToNillableFloat64", t, func() {
		So(NewInt8(127).MapToNillableFloat64(func(int8) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).Get(), ShouldEqual, 3.2455)

		So(NewEmptyInt8().MapToNillableFloat64(func(int8) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt8(127).MapToNillableFloat64(func(int8) *float64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToFloat64(t *testing.T) {
	Convey("Int8.MapToFloat64", t, func() {
		test1, err1 := NewInt8(127).MapToFloat64(func(int8) (float64, error) {
			return 3.2455, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 3.2455)

		test2, err2 := NewEmptyInt8().MapToFloat64(func(int8) (float64, error) {
			return 3.2455, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt8(127).MapToFloat64(func(int8) (float64, error) {
			return 3.2455, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt8_ForceMapToInt(t *testing.T) {
	Convey("Int8.ForceMapToInt", t, func() {
		So(NewInt8(127).ForceMapToInt(func(int8) int { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyInt8().ForceMapToInt(func(int8) int { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToNillableInt(t *testing.T) {
	Convey("Int8.MapToNillableInt", t, func() {
		So(NewInt8(127).MapToNillableInt(func(int8) *int {
			var tmp int = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyInt8().MapToNillableInt(func(int8) *int {
			var tmp int = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt8(127).MapToNillableInt(func(int8) *int { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToInt(t *testing.T) {
	Convey("Int8.MapToInt", t, func() {
		test1, err1 := NewInt8(127).MapToInt(func(int8) (int, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyInt8().MapToInt(func(int8) (int, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt8(127).MapToInt(func(int8) (int, error) {
			return 123456, errors.New("int error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt8_ForceMapToInt16(t *testing.T) {
	Convey("Int8.ForceMapToInt16", t, func() {
		So(NewInt8(127).ForceMapToInt16(func(int8) int16 { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyInt8().ForceMapToInt16(func(int8) int16 { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToNillableInt16(t *testing.T) {
	Convey("Int8.MapToNillableInt16", t, func() {
		So(NewInt8(127).MapToNillableInt16(func(int8) *int16 {
			var tmp int16 = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyInt8().MapToNillableInt16(func(int8) *int16 {
			var tmp int16 = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt8(127).MapToNillableInt16(func(int8) *int16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToInt16(t *testing.T) {
	Convey("Int8.MapToInt16", t, func() {
		test1, err1 := NewInt8(127).MapToInt16(func(int8) (int16, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyInt8().MapToInt16(func(int8) (int16, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt8(127).MapToInt16(func(int8) (int16, error) {
			return 222, errors.New("int16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt8_ForceMapToInt32(t *testing.T) {
	Convey("Int8.ForceMapToInt32", t, func() {
		So(NewInt8(127).ForceMapToInt32(func(int8) int32 { return 333 }).Get(),
			ShouldEqual, 333)

		So(NewEmptyInt8().ForceMapToInt32(func(int8) int32 { return 333 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToNillableInt32(t *testing.T) {
	Convey("Int8.MapToNillableInt32", t, func() {
		So(NewInt8(127).MapToNillableInt32(func(int8) *int32 {
			var tmp int32 = 333
			return &tmp
		}).Get(), ShouldEqual, 333)

		So(NewEmptyInt8().MapToNillableInt32(func(int8) *int32 {
			var tmp int32 = 333
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt8(127).MapToNillableInt32(func(int8) *int32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToInt32(t *testing.T) {
	Convey("Int8.MapToInt32", t, func() {
		test1, err1 := NewInt8(127).MapToInt32(func(int8) (int32, error) {
			return 333, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 333)

		test2, err2 := NewEmptyInt8().MapToInt32(func(int8) (int32, error) {
			return 333, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt8(127).MapToInt32(func(int8) (int32, error) {
			return 333, errors.New("int32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt8_ForceMapToInt64(t *testing.T) {
	Convey("Int8.ForceMapToInt64", t, func() {
		So(NewInt8(127).ForceMapToInt64(func(int8) int64 { return 444 }).Get(),
			ShouldEqual, 444)

		So(NewEmptyInt8().ForceMapToInt64(func(int8) int64 { return 444 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToNillableInt64(t *testing.T) {
	Convey("Int8.MapToNillableInt64", t, func() {
		So(NewInt8(127).MapToNillableInt64(func(int8) *int64 {
			var tmp int64 = 444
			return &tmp
		}).Get(), ShouldEqual, 444)

		So(NewEmptyInt8().MapToNillableInt64(func(int8) *int64 {
			var tmp int64 = 444
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt8(127).MapToNillableInt64(func(int8) *int64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToInt64(t *testing.T) {
	Convey("Int8.MapToInt64", t, func() {
		test1, err1 := NewInt8(127).MapToInt64(func(int8) (int64, error) {
			return 444, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 444)

		test2, err2 := NewEmptyInt8().MapToInt64(func(int8) (int64, error) {
			return 444, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt8(127).MapToInt64(func(int8) (int64, error) {
			return 444, errors.New("int64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt8_ForceMapToUntyped(t *testing.T) {
	Convey("Int8.ForceMapToUntyped", t, func() {
		So(NewInt8(127).ForceMapToUntyped(func(int8) interface{} { return "111111" }).Get(),
			ShouldEqual, "111111")

		So(NewEmptyInt8().ForceMapToUntyped(func(int8) interface{} { return "111111" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToNillableUntyped(t *testing.T) {
	Convey("Int8.MapToNillableUntyped", t, func() {
		So(NewInt8(127).MapToNillableUntyped(func(int8) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).Get(), ShouldEqual, "111111")

		So(NewEmptyInt8().MapToNillableUntyped(func(int8) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt8(127).MapToNillableUntyped(func(int8) *interface{} { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToUntyped(t *testing.T) {
	Convey("Int8.MapToUntyped", t, func() {
		test1, err1 := NewInt8(127).MapToUntyped(func(int8) (interface{}, error) {
			return "111111", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "111111")

		test2, err2 := NewEmptyInt8().MapToUntyped(func(int8) (interface{}, error) {
			return "111111", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt8(127).MapToUntyped(func(int8) (interface{}, error) {
			return "111111", errors.New("interface error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "interface error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt8_ForceMapToRune(t *testing.T) {
	Convey("Int8.ForceMapToRune", t, func() {
		So(NewInt8(127).ForceMapToRune(func(int8) rune { return 555 }).Get(),
			ShouldEqual, 555)

		So(NewEmptyInt8().ForceMapToRune(func(int8) rune { return 555 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToNillableRune(t *testing.T) {
	Convey("Int8.MapToNillableRune", t, func() {
		So(NewInt8(127).MapToNillableRune(func(int8) *rune {
			var tmp rune = 555
			return &tmp
		}).Get(), ShouldEqual, 555)

		So(NewEmptyInt8().MapToNillableRune(func(int8) *rune {
			var tmp rune = 555
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt8(127).MapToNillableRune(func(int8) *rune { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToRune(t *testing.T) {
	Convey("Int8.MapToRune", t, func() {
		test1, err1 := NewInt8(127).MapToRune(func(int8) (rune, error) {
			return 555, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 555)

		test2, err2 := NewEmptyInt8().MapToRune(func(int8) (rune, error) {
			return 555, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt8(127).MapToRune(func(int8) (rune, error) {
			return 555, errors.New("rune error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "rune error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt8_ForceMapToString(t *testing.T) {
	Convey("Int8.ForceMapToString", t, func() {
		So(NewInt8(127).ForceMapToString(func(int8) string { return "bonanza" }).Get(),
			ShouldEqual, "bonanza")

		So(NewEmptyInt8().ForceMapToString(func(int8) string { return "bonanza" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToNillableString(t *testing.T) {
	Convey("Int8.MapToNillableString", t, func() {
		So(NewInt8(127).MapToNillableString(func(int8) *string {
			var tmp string = "bonanza"
			return &tmp
		}).Get(), ShouldEqual, "bonanza")

		So(NewEmptyInt8().MapToNillableString(func(int8) *string {
			var tmp string = "bonanza"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt8(127).MapToNillableString(func(int8) *string { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToString(t *testing.T) {
	Convey("Int8.MapToString", t, func() {
		test1, err1 := NewInt8(127).MapToString(func(int8) (string, error) {
			return "bonanza", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "bonanza")

		test2, err2 := NewEmptyInt8().MapToString(func(int8) (string, error) {
			return "bonanza", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt8(127).MapToString(func(int8) (string, error) {
			return "bonanza", errors.New("string error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "string error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt8_ForceMapToUint(t *testing.T) {
	Convey("Int8.ForceMapToUint", t, func() {
		So(NewInt8(127).ForceMapToUint(func(int8) uint { return 666 }).Get(),
			ShouldEqual, 666)

		So(NewEmptyInt8().ForceMapToUint(func(int8) uint { return 666 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToNillableUint(t *testing.T) {
	Convey("Int8.MapToNillableUint", t, func() {
		So(NewInt8(127).MapToNillableUint(func(int8) *uint {
			var tmp uint = 666
			return &tmp
		}).Get(), ShouldEqual, 666)

		So(NewEmptyInt8().MapToNillableUint(func(int8) *uint {
			var tmp uint = 666
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt8(127).MapToNillableUint(func(int8) *uint { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToUint(t *testing.T) {
	Convey("Int8.MapToUint", t, func() {
		test1, err1 := NewInt8(127).MapToUint(func(int8) (uint, error) {
			return 666, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 666)

		test2, err2 := NewEmptyInt8().MapToUint(func(int8) (uint, error) {
			return 666, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt8(127).MapToUint(func(int8) (uint, error) {
			return 666, errors.New("uint error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt8_ForceMapToUint16(t *testing.T) {
	Convey("Int8.ForceMapToUint16", t, func() {
		So(NewInt8(127).ForceMapToUint16(func(int8) uint16 { return 888 }).Get(),
			ShouldEqual, 888)

		So(NewEmptyInt8().ForceMapToUint16(func(int8) uint16 { return 888 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToNillableUint16(t *testing.T) {
	Convey("Int8.MapToNillableUint16", t, func() {
		So(NewInt8(127).MapToNillableUint16(func(int8) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).Get(), ShouldEqual, 888)

		So(NewEmptyInt8().MapToNillableUint16(func(int8) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt8(127).MapToNillableUint16(func(int8) *uint16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToUint16(t *testing.T) {
	Convey("Int8.MapToUint16", t, func() {
		test1, err1 := NewInt8(127).MapToUint16(func(int8) (uint16, error) {
			return 888, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 888)

		test2, err2 := NewEmptyInt8().MapToUint16(func(int8) (uint16, error) {
			return 888, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt8(127).MapToUint16(func(int8) (uint16, error) {
			return 888, errors.New("uint16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt8_ForceMapToUint32(t *testing.T) {
	Convey("Int8.ForceMapToUint32", t, func() {
		So(NewInt8(127).ForceMapToUint32(func(int8) uint32 { return 999 }).Get(),
			ShouldEqual, 999)

		So(NewEmptyInt8().ForceMapToUint32(func(int8) uint32 { return 999 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToNillableUint32(t *testing.T) {
	Convey("Int8.MapToNillableUint32", t, func() {
		So(NewInt8(127).MapToNillableUint32(func(int8) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).Get(), ShouldEqual, 999)

		So(NewEmptyInt8().MapToNillableUint32(func(int8) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt8(127).MapToNillableUint32(func(int8) *uint32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToUint32(t *testing.T) {
	Convey("Int8.MapToUint32", t, func() {
		test1, err1 := NewInt8(127).MapToUint32(func(int8) (uint32, error) {
			return 999, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 999)

		test2, err2 := NewEmptyInt8().MapToUint32(func(int8) (uint32, error) {
			return 999, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt8(127).MapToUint32(func(int8) (uint32, error) {
			return 999, errors.New("uint32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt8_ForceMapToUint64(t *testing.T) {
	Convey("Int8.ForceMapToUint64", t, func() {
		So(NewInt8(127).ForceMapToUint64(func(int8) uint64 { return 101010 }).Get(),
			ShouldEqual, 101010)

		So(NewEmptyInt8().ForceMapToUint64(func(int8) uint64 { return 101010 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToNillableUint64(t *testing.T) {
	Convey("Int8.MapToNillableUint64", t, func() {
		So(NewInt8(127).MapToNillableUint64(func(int8) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).Get(), ShouldEqual, 101010)

		So(NewEmptyInt8().MapToNillableUint64(func(int8) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt8(127).MapToNillableUint64(func(int8) *uint64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToUint64(t *testing.T) {
	Convey("Int8.MapToUint64", t, func() {
		test1, err1 := NewInt8(127).MapToUint64(func(int8) (uint64, error) {
			return 101010, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 101010)

		test2, err2 := NewEmptyInt8().MapToUint64(func(int8) (uint64, error) {
			return 101010, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt8(127).MapToUint64(func(int8) (uint64, error) {
			return 101010, errors.New("uint64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt8_ForceMapToUint8(t *testing.T) {
	Convey("Int8.ForceMapToUint8", t, func() {
		So(NewInt8(127).ForceMapToUint8(func(int8) uint8 { return 77 }).Get(),
			ShouldEqual, 77)

		So(NewEmptyInt8().ForceMapToUint8(func(int8) uint8 { return 77 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToNillableUint8(t *testing.T) {
	Convey("Int8.MapToNillableUint8", t, func() {
		So(NewInt8(127).MapToNillableUint8(func(int8) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).Get(), ShouldEqual, 77)

		So(NewEmptyInt8().MapToNillableUint8(func(int8) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt8(127).MapToNillableUint8(func(int8) *uint8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt8_MapToUint8(t *testing.T) {
	Convey("Int8.MapToUint8", t, func() {
		test1, err1 := NewInt8(127).MapToUint8(func(int8) (uint8, error) {
			return 77, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 77)

		test2, err2 := NewEmptyInt8().MapToUint8(func(int8) (uint8, error) {
			return 77, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt8(127).MapToUint8(func(int8) (uint8, error) {
			return 77, errors.New("uint8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}
