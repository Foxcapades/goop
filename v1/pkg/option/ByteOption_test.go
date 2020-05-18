// Package option_test generated @ 2020-05-17 22:19:55.164878293 -0400 EDT m=+0.001439508
package option_test

import (
	"errors"
	"testing"

	. "github.com/Foxcapades/goop/v1/pkg/option"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewByte(t *testing.T) {
	Convey("NewByte", t, func() {
		test := NewByte(255)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyByte(t *testing.T) {
	Convey("NewEmptyByte", t, func() {
		test := NewEmptyByte()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeByte(t *testing.T) {
	Convey("NewMaybeByte", t, func() {
		var val1 byte

		test1 := NewMaybeByte(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeByte(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestByte_Get(t *testing.T) {
	Convey("Byte.Get", t, func() {
		test1 := NewByte(255)
		So(test1.Get(), ShouldEqual, 255)

		test2 := NewEmptyByte()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestByte_Or(t *testing.T) {
	Convey("Byte.Or", t, func() {
		test1 := NewByte(255)
		So(test1.Or(56), ShouldEqual, 255)

		test2 := NewEmptyByte()
		So(test2.Or(255), ShouldEqual, 255)
	})
}

func TestByte_OrGet(t *testing.T) {
	Convey("Byte.OrGet", t, func() {
		test1 := NewByte(255)
		So(test1.OrGet(func() byte { return 56 }), ShouldEqual, 255)

		test2 := NewEmptyByte()
		So(test2.OrGet(func() byte { return 255 }), ShouldEqual, 255)
	})
}

func TestByte_OrPanicWith(t *testing.T) {
	Convey("Byte.OrPanicWith", t, func() {
		test1 := NewByte(255)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 255)

		test2 := NewEmptyByte()
		So(func() {
			test2.OrPanicWith(func() interface{} { return "test value" })
		}, ShouldPanicWith, "test value")
	})
}

func TestByte_MapToNillable(t *testing.T) {
	Convey("Byte.MapToNillable", t, func() {
		test1 := NewByte(255)
		So(test1.MapToNillable(func(b byte) *byte { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewByte(255)
		So(test2.MapToNillable(func(b byte) *byte { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyByte()
		So(func() {
			test3.MapToNillable(func(b byte) *byte { panic("foo") }).IsNil()
		}, ShouldNotPanic)
	})
}

func TestByte_ForceMapToBool(t *testing.T) {
	Convey("Byte.ForceMapToBool", t, func() {
		So(NewByte(255).ForceMapToBool(func(byte) bool { return true }).Get(),
			ShouldEqual, true)

		So(NewEmptyByte().ForceMapToBool(func(byte) bool { return true }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToNillableBool(t *testing.T) {
	Convey("Byte.MapToNillableBool", t, func() {
		So(NewByte(255).MapToNillableBool(func(byte) *bool {
			var tmp bool = true
			return &tmp
		}).Get(), ShouldEqual, true)

		So(NewEmptyByte().MapToNillableBool(func(byte) *bool {
			var tmp bool = true
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewByte(255).MapToNillableBool(func(byte) *bool { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToBool(t *testing.T) {
	Convey("Byte.MapToBool", t, func() {
		test1, err1 := NewByte(255).MapToBool(func(byte) (bool, error) {
			return true, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, true)

		test2, err2 := NewEmptyByte().MapToBool(func(byte) (bool, error) {
			return true, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewByte(255).MapToBool(func(byte) (bool, error) {
			return true, errors.New("bool error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "bool error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestByte_ForceMapToComplex128(t *testing.T) {
	Convey("Byte.ForceMapToComplex128", t, func() {
		So(NewByte(255).ForceMapToComplex128(func(byte) complex128 { return 987654 }).Get(),
			ShouldEqual, 987654)

		So(NewEmptyByte().ForceMapToComplex128(func(byte) complex128 { return 987654 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToNillableComplex128(t *testing.T) {
	Convey("Byte.MapToNillableComplex128", t, func() {
		So(NewByte(255).MapToNillableComplex128(func(byte) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).Get(), ShouldEqual, 987654)

		So(NewEmptyByte().MapToNillableComplex128(func(byte) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewByte(255).MapToNillableComplex128(func(byte) *complex128 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToComplex128(t *testing.T) {
	Convey("Byte.MapToComplex128", t, func() {
		test1, err1 := NewByte(255).MapToComplex128(func(byte) (complex128, error) {
			return 987654, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 987654)

		test2, err2 := NewEmptyByte().MapToComplex128(func(byte) (complex128, error) {
			return 987654, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewByte(255).MapToComplex128(func(byte) (complex128, error) {
			return 987654, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestByte_ForceMapToComplex64(t *testing.T) {
	Convey("Byte.ForceMapToComplex64", t, func() {
		So(NewByte(255).ForceMapToComplex64(func(byte) complex64 { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyByte().ForceMapToComplex64(func(byte) complex64 { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToNillableComplex64(t *testing.T) {
	Convey("Byte.MapToNillableComplex64", t, func() {
		So(NewByte(255).MapToNillableComplex64(func(byte) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyByte().MapToNillableComplex64(func(byte) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewByte(255).MapToNillableComplex64(func(byte) *complex64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToComplex64(t *testing.T) {
	Convey("Byte.MapToComplex64", t, func() {
		test1, err1 := NewByte(255).MapToComplex64(func(byte) (complex64, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyByte().MapToComplex64(func(byte) (complex64, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewByte(255).MapToComplex64(func(byte) (complex64, error) {
			return 123456, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestByte_ForceMapToFloat32(t *testing.T) {
	Convey("Byte.ForceMapToFloat32", t, func() {
		So(NewByte(255).ForceMapToFloat32(func(byte) float32 { return 32.3 }).Get(),
			ShouldEqual, 32.3)

		So(NewEmptyByte().ForceMapToFloat32(func(byte) float32 { return 32.3 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToNillableFloat32(t *testing.T) {
	Convey("Byte.MapToNillableFloat32", t, func() {
		So(NewByte(255).MapToNillableFloat32(func(byte) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).Get(), ShouldEqual, 32.3)

		So(NewEmptyByte().MapToNillableFloat32(func(byte) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewByte(255).MapToNillableFloat32(func(byte) *float32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToFloat32(t *testing.T) {
	Convey("Byte.MapToFloat32", t, func() {
		test1, err1 := NewByte(255).MapToFloat32(func(byte) (float32, error) {
			return 32.3, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 32.3)

		test2, err2 := NewEmptyByte().MapToFloat32(func(byte) (float32, error) {
			return 32.3, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewByte(255).MapToFloat32(func(byte) (float32, error) {
			return 32.3, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestByte_ForceMapToFloat64(t *testing.T) {
	Convey("Byte.ForceMapToFloat64", t, func() {
		So(NewByte(255).ForceMapToFloat64(func(byte) float64 { return 3.2455 }).Get(),
			ShouldEqual, 3.2455)

		So(NewEmptyByte().ForceMapToFloat64(func(byte) float64 { return 3.2455 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToNillableFloat64(t *testing.T) {
	Convey("Byte.MapToNillableFloat64", t, func() {
		So(NewByte(255).MapToNillableFloat64(func(byte) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).Get(), ShouldEqual, 3.2455)

		So(NewEmptyByte().MapToNillableFloat64(func(byte) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewByte(255).MapToNillableFloat64(func(byte) *float64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToFloat64(t *testing.T) {
	Convey("Byte.MapToFloat64", t, func() {
		test1, err1 := NewByte(255).MapToFloat64(func(byte) (float64, error) {
			return 3.2455, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 3.2455)

		test2, err2 := NewEmptyByte().MapToFloat64(func(byte) (float64, error) {
			return 3.2455, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewByte(255).MapToFloat64(func(byte) (float64, error) {
			return 3.2455, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestByte_ForceMapToInt(t *testing.T) {
	Convey("Byte.ForceMapToInt", t, func() {
		So(NewByte(255).ForceMapToInt(func(byte) int { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyByte().ForceMapToInt(func(byte) int { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToNillableInt(t *testing.T) {
	Convey("Byte.MapToNillableInt", t, func() {
		So(NewByte(255).MapToNillableInt(func(byte) *int {
			var tmp int = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyByte().MapToNillableInt(func(byte) *int {
			var tmp int = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewByte(255).MapToNillableInt(func(byte) *int { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToInt(t *testing.T) {
	Convey("Byte.MapToInt", t, func() {
		test1, err1 := NewByte(255).MapToInt(func(byte) (int, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyByte().MapToInt(func(byte) (int, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewByte(255).MapToInt(func(byte) (int, error) {
			return 123456, errors.New("int error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestByte_ForceMapToInt16(t *testing.T) {
	Convey("Byte.ForceMapToInt16", t, func() {
		So(NewByte(255).ForceMapToInt16(func(byte) int16 { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyByte().ForceMapToInt16(func(byte) int16 { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToNillableInt16(t *testing.T) {
	Convey("Byte.MapToNillableInt16", t, func() {
		So(NewByte(255).MapToNillableInt16(func(byte) *int16 {
			var tmp int16 = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyByte().MapToNillableInt16(func(byte) *int16 {
			var tmp int16 = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewByte(255).MapToNillableInt16(func(byte) *int16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToInt16(t *testing.T) {
	Convey("Byte.MapToInt16", t, func() {
		test1, err1 := NewByte(255).MapToInt16(func(byte) (int16, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyByte().MapToInt16(func(byte) (int16, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewByte(255).MapToInt16(func(byte) (int16, error) {
			return 222, errors.New("int16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestByte_ForceMapToInt32(t *testing.T) {
	Convey("Byte.ForceMapToInt32", t, func() {
		So(NewByte(255).ForceMapToInt32(func(byte) int32 { return 333 }).Get(),
			ShouldEqual, 333)

		So(NewEmptyByte().ForceMapToInt32(func(byte) int32 { return 333 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToNillableInt32(t *testing.T) {
	Convey("Byte.MapToNillableInt32", t, func() {
		So(NewByte(255).MapToNillableInt32(func(byte) *int32 {
			var tmp int32 = 333
			return &tmp
		}).Get(), ShouldEqual, 333)

		So(NewEmptyByte().MapToNillableInt32(func(byte) *int32 {
			var tmp int32 = 333
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewByte(255).MapToNillableInt32(func(byte) *int32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToInt32(t *testing.T) {
	Convey("Byte.MapToInt32", t, func() {
		test1, err1 := NewByte(255).MapToInt32(func(byte) (int32, error) {
			return 333, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 333)

		test2, err2 := NewEmptyByte().MapToInt32(func(byte) (int32, error) {
			return 333, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewByte(255).MapToInt32(func(byte) (int32, error) {
			return 333, errors.New("int32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestByte_ForceMapToInt64(t *testing.T) {
	Convey("Byte.ForceMapToInt64", t, func() {
		So(NewByte(255).ForceMapToInt64(func(byte) int64 { return 444 }).Get(),
			ShouldEqual, 444)

		So(NewEmptyByte().ForceMapToInt64(func(byte) int64 { return 444 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToNillableInt64(t *testing.T) {
	Convey("Byte.MapToNillableInt64", t, func() {
		So(NewByte(255).MapToNillableInt64(func(byte) *int64 {
			var tmp int64 = 444
			return &tmp
		}).Get(), ShouldEqual, 444)

		So(NewEmptyByte().MapToNillableInt64(func(byte) *int64 {
			var tmp int64 = 444
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewByte(255).MapToNillableInt64(func(byte) *int64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToInt64(t *testing.T) {
	Convey("Byte.MapToInt64", t, func() {
		test1, err1 := NewByte(255).MapToInt64(func(byte) (int64, error) {
			return 444, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 444)

		test2, err2 := NewEmptyByte().MapToInt64(func(byte) (int64, error) {
			return 444, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewByte(255).MapToInt64(func(byte) (int64, error) {
			return 444, errors.New("int64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestByte_ForceMapToInt8(t *testing.T) {
	Convey("Byte.ForceMapToInt8", t, func() {
		So(NewByte(255).ForceMapToInt8(func(byte) int8 { return 111 }).Get(),
			ShouldEqual, 111)

		So(NewEmptyByte().ForceMapToInt8(func(byte) int8 { return 111 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToNillableInt8(t *testing.T) {
	Convey("Byte.MapToNillableInt8", t, func() {
		So(NewByte(255).MapToNillableInt8(func(byte) *int8 {
			var tmp int8 = 111
			return &tmp
		}).Get(), ShouldEqual, 111)

		So(NewEmptyByte().MapToNillableInt8(func(byte) *int8 {
			var tmp int8 = 111
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewByte(255).MapToNillableInt8(func(byte) *int8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToInt8(t *testing.T) {
	Convey("Byte.MapToInt8", t, func() {
		test1, err1 := NewByte(255).MapToInt8(func(byte) (int8, error) {
			return 111, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 111)

		test2, err2 := NewEmptyByte().MapToInt8(func(byte) (int8, error) {
			return 111, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewByte(255).MapToInt8(func(byte) (int8, error) {
			return 111, errors.New("int8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestByte_ForceMapToUntyped(t *testing.T) {
	Convey("Byte.ForceMapToUntyped", t, func() {
		So(NewByte(255).ForceMapToUntyped(func(byte) interface{} { return "111111" }).Get(),
			ShouldEqual, "111111")

		So(NewEmptyByte().ForceMapToUntyped(func(byte) interface{} { return "111111" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToNillableUntyped(t *testing.T) {
	Convey("Byte.MapToNillableUntyped", t, func() {
		So(NewByte(255).MapToNillableUntyped(func(byte) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).Get(), ShouldEqual, "111111")

		So(NewEmptyByte().MapToNillableUntyped(func(byte) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewByte(255).MapToNillableUntyped(func(byte) *interface{} { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToUntyped(t *testing.T) {
	Convey("Byte.MapToUntyped", t, func() {
		test1, err1 := NewByte(255).MapToUntyped(func(byte) (interface{}, error) {
			return "111111", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "111111")

		test2, err2 := NewEmptyByte().MapToUntyped(func(byte) (interface{}, error) {
			return "111111", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewByte(255).MapToUntyped(func(byte) (interface{}, error) {
			return "111111", errors.New("interface error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "interface error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestByte_ForceMapToRune(t *testing.T) {
	Convey("Byte.ForceMapToRune", t, func() {
		So(NewByte(255).ForceMapToRune(func(byte) rune { return 555 }).Get(),
			ShouldEqual, 555)

		So(NewEmptyByte().ForceMapToRune(func(byte) rune { return 555 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToNillableRune(t *testing.T) {
	Convey("Byte.MapToNillableRune", t, func() {
		So(NewByte(255).MapToNillableRune(func(byte) *rune {
			var tmp rune = 555
			return &tmp
		}).Get(), ShouldEqual, 555)

		So(NewEmptyByte().MapToNillableRune(func(byte) *rune {
			var tmp rune = 555
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewByte(255).MapToNillableRune(func(byte) *rune { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToRune(t *testing.T) {
	Convey("Byte.MapToRune", t, func() {
		test1, err1 := NewByte(255).MapToRune(func(byte) (rune, error) {
			return 555, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 555)

		test2, err2 := NewEmptyByte().MapToRune(func(byte) (rune, error) {
			return 555, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewByte(255).MapToRune(func(byte) (rune, error) {
			return 555, errors.New("rune error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "rune error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestByte_ForceMapToString(t *testing.T) {
	Convey("Byte.ForceMapToString", t, func() {
		So(NewByte(255).ForceMapToString(func(byte) string { return "bonanza" }).Get(),
			ShouldEqual, "bonanza")

		So(NewEmptyByte().ForceMapToString(func(byte) string { return "bonanza" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToNillableString(t *testing.T) {
	Convey("Byte.MapToNillableString", t, func() {
		So(NewByte(255).MapToNillableString(func(byte) *string {
			var tmp string = "bonanza"
			return &tmp
		}).Get(), ShouldEqual, "bonanza")

		So(NewEmptyByte().MapToNillableString(func(byte) *string {
			var tmp string = "bonanza"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewByte(255).MapToNillableString(func(byte) *string { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToString(t *testing.T) {
	Convey("Byte.MapToString", t, func() {
		test1, err1 := NewByte(255).MapToString(func(byte) (string, error) {
			return "bonanza", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "bonanza")

		test2, err2 := NewEmptyByte().MapToString(func(byte) (string, error) {
			return "bonanza", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewByte(255).MapToString(func(byte) (string, error) {
			return "bonanza", errors.New("string error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "string error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestByte_ForceMapToUint(t *testing.T) {
	Convey("Byte.ForceMapToUint", t, func() {
		So(NewByte(255).ForceMapToUint(func(byte) uint { return 666 }).Get(),
			ShouldEqual, 666)

		So(NewEmptyByte().ForceMapToUint(func(byte) uint { return 666 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToNillableUint(t *testing.T) {
	Convey("Byte.MapToNillableUint", t, func() {
		So(NewByte(255).MapToNillableUint(func(byte) *uint {
			var tmp uint = 666
			return &tmp
		}).Get(), ShouldEqual, 666)

		So(NewEmptyByte().MapToNillableUint(func(byte) *uint {
			var tmp uint = 666
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewByte(255).MapToNillableUint(func(byte) *uint { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToUint(t *testing.T) {
	Convey("Byte.MapToUint", t, func() {
		test1, err1 := NewByte(255).MapToUint(func(byte) (uint, error) {
			return 666, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 666)

		test2, err2 := NewEmptyByte().MapToUint(func(byte) (uint, error) {
			return 666, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewByte(255).MapToUint(func(byte) (uint, error) {
			return 666, errors.New("uint error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestByte_ForceMapToUint16(t *testing.T) {
	Convey("Byte.ForceMapToUint16", t, func() {
		So(NewByte(255).ForceMapToUint16(func(byte) uint16 { return 888 }).Get(),
			ShouldEqual, 888)

		So(NewEmptyByte().ForceMapToUint16(func(byte) uint16 { return 888 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToNillableUint16(t *testing.T) {
	Convey("Byte.MapToNillableUint16", t, func() {
		So(NewByte(255).MapToNillableUint16(func(byte) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).Get(), ShouldEqual, 888)

		So(NewEmptyByte().MapToNillableUint16(func(byte) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewByte(255).MapToNillableUint16(func(byte) *uint16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToUint16(t *testing.T) {
	Convey("Byte.MapToUint16", t, func() {
		test1, err1 := NewByte(255).MapToUint16(func(byte) (uint16, error) {
			return 888, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 888)

		test2, err2 := NewEmptyByte().MapToUint16(func(byte) (uint16, error) {
			return 888, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewByte(255).MapToUint16(func(byte) (uint16, error) {
			return 888, errors.New("uint16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestByte_ForceMapToUint32(t *testing.T) {
	Convey("Byte.ForceMapToUint32", t, func() {
		So(NewByte(255).ForceMapToUint32(func(byte) uint32 { return 999 }).Get(),
			ShouldEqual, 999)

		So(NewEmptyByte().ForceMapToUint32(func(byte) uint32 { return 999 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToNillableUint32(t *testing.T) {
	Convey("Byte.MapToNillableUint32", t, func() {
		So(NewByte(255).MapToNillableUint32(func(byte) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).Get(), ShouldEqual, 999)

		So(NewEmptyByte().MapToNillableUint32(func(byte) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewByte(255).MapToNillableUint32(func(byte) *uint32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToUint32(t *testing.T) {
	Convey("Byte.MapToUint32", t, func() {
		test1, err1 := NewByte(255).MapToUint32(func(byte) (uint32, error) {
			return 999, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 999)

		test2, err2 := NewEmptyByte().MapToUint32(func(byte) (uint32, error) {
			return 999, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewByte(255).MapToUint32(func(byte) (uint32, error) {
			return 999, errors.New("uint32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestByte_ForceMapToUint64(t *testing.T) {
	Convey("Byte.ForceMapToUint64", t, func() {
		So(NewByte(255).ForceMapToUint64(func(byte) uint64 { return 101010 }).Get(),
			ShouldEqual, 101010)

		So(NewEmptyByte().ForceMapToUint64(func(byte) uint64 { return 101010 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToNillableUint64(t *testing.T) {
	Convey("Byte.MapToNillableUint64", t, func() {
		So(NewByte(255).MapToNillableUint64(func(byte) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).Get(), ShouldEqual, 101010)

		So(NewEmptyByte().MapToNillableUint64(func(byte) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewByte(255).MapToNillableUint64(func(byte) *uint64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToUint64(t *testing.T) {
	Convey("Byte.MapToUint64", t, func() {
		test1, err1 := NewByte(255).MapToUint64(func(byte) (uint64, error) {
			return 101010, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 101010)

		test2, err2 := NewEmptyByte().MapToUint64(func(byte) (uint64, error) {
			return 101010, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewByte(255).MapToUint64(func(byte) (uint64, error) {
			return 101010, errors.New("uint64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestByte_ForceMapToUint8(t *testing.T) {
	Convey("Byte.ForceMapToUint8", t, func() {
		So(NewByte(255).ForceMapToUint8(func(byte) uint8 { return 77 }).Get(),
			ShouldEqual, 77)

		So(NewEmptyByte().ForceMapToUint8(func(byte) uint8 { return 77 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToNillableUint8(t *testing.T) {
	Convey("Byte.MapToNillableUint8", t, func() {
		So(NewByte(255).MapToNillableUint8(func(byte) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).Get(), ShouldEqual, 77)

		So(NewEmptyByte().MapToNillableUint8(func(byte) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewByte(255).MapToNillableUint8(func(byte) *uint8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestByte_MapToUint8(t *testing.T) {
	Convey("Byte.MapToUint8", t, func() {
		test1, err1 := NewByte(255).MapToUint8(func(byte) (uint8, error) {
			return 77, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 77)

		test2, err2 := NewEmptyByte().MapToUint8(func(byte) (uint8, error) {
			return 77, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewByte(255).MapToUint8(func(byte) (uint8, error) {
			return 77, errors.New("uint8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}
