// Package option_test generated @ 2020-05-17 22:19:55.164878293 -0400 EDT m=+0.001439508
package option_test

import (
	"errors"
	"testing"

	. "github.com/Foxcapades/goop/v0/pkg/option"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewUint8(t *testing.T) {
	Convey("NewUint8", t, func() {
		test := NewUint8(123)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyUint8(t *testing.T) {
	Convey("NewEmptyUint8", t, func() {
		test := NewEmptyUint8()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeUint8(t *testing.T) {
	Convey("NewMaybeUint8", t, func() {
		var val1 uint8

		test1 := NewMaybeUint8(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeUint8(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestUint8_Get(t *testing.T) {
	Convey("Uint8.Get", t, func() {
		test1 := NewUint8(123)
		So(test1.Get(), ShouldEqual, 123)

		test2 := NewEmptyUint8()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestUint8_Or(t *testing.T) {
	Convey("Uint8.Or", t, func() {
		test1 := NewUint8(123)
		So(test1.Or(245), ShouldEqual, 123)

		test2 := NewEmptyUint8()
		So(test2.Or(123), ShouldEqual, 123)
	})
}

func TestUint8_OrGet(t *testing.T) {
	Convey("Uint8.OrGet", t, func() {
		test1 := NewUint8(123)
		So(test1.OrGet(func() uint8 { return 245 }), ShouldEqual, 123)

		test2 := NewEmptyUint8()
		So(test2.OrGet(func() uint8 { return 123 }), ShouldEqual, 123)
	})
}

func TestUint8_OrPanicWith(t *testing.T) {
	Convey("Uint8.OrPanicWith", t, func() {
		test1 := NewUint8(123)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 123)

		test2 := NewEmptyUint8()
		So(func() {
			test2.OrPanicWith(func() interface{} { return "test value" })
		}, ShouldPanicWith, "test value")
	})
}

func TestUint8_MapToNillable(t *testing.T) {
	Convey("Uint8.MapToNillable", t, func() {
		test1 := NewUint8(123)
		So(test1.MapToNillable(func(b uint8) *uint8 { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewUint8(123)
		So(test2.MapToNillable(func(b uint8) *uint8 { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyUint8()
		So(func() {
			test3.MapToNillable(func(b uint8) *uint8 { panic("foo") }).IsNil()
		}, ShouldNotPanic)
	})
}

func TestUint8_ForceMapToBool(t *testing.T) {
	Convey("Uint8.ForceMapToBool", t, func() {
		So(NewUint8(123).ForceMapToBool(func(uint8) bool { return true }).Get(),
			ShouldEqual, true)

		So(NewEmptyUint8().ForceMapToBool(func(uint8) bool { return true }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToNillableBool(t *testing.T) {
	Convey("Uint8.MapToNillableBool", t, func() {
		So(NewUint8(123).MapToNillableBool(func(uint8) *bool {
			var tmp bool = true
			return &tmp
		}).Get(), ShouldEqual, true)

		So(NewEmptyUint8().MapToNillableBool(func(uint8) *bool {
			var tmp bool = true
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint8(123).MapToNillableBool(func(uint8) *bool { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToBool(t *testing.T) {
	Convey("Uint8.MapToBool", t, func() {
		test1, err1 := NewUint8(123).MapToBool(func(uint8) (bool, error) {
			return true, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, true)

		test2, err2 := NewEmptyUint8().MapToBool(func(uint8) (bool, error) {
			return true, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint8(123).MapToBool(func(uint8) (bool, error) {
			return true, errors.New("bool error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "bool error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint8_ForceMapToByte(t *testing.T) {
	Convey("Uint8.ForceMapToByte", t, func() {
		So(NewUint8(123).ForceMapToByte(func(uint8) byte { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyUint8().ForceMapToByte(func(uint8) byte { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToNillableByte(t *testing.T) {
	Convey("Uint8.MapToNillableByte", t, func() {
		So(NewUint8(123).MapToNillableByte(func(uint8) *byte {
			var tmp byte = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyUint8().MapToNillableByte(func(uint8) *byte {
			var tmp byte = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint8(123).MapToNillableByte(func(uint8) *byte { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToByte(t *testing.T) {
	Convey("Uint8.MapToByte", t, func() {
		test1, err1 := NewUint8(123).MapToByte(func(uint8) (byte, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyUint8().MapToByte(func(uint8) (byte, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint8(123).MapToByte(func(uint8) (byte, error) {
			return 222, errors.New("byte error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "byte error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint8_ForceMapToComplex128(t *testing.T) {
	Convey("Uint8.ForceMapToComplex128", t, func() {
		So(NewUint8(123).ForceMapToComplex128(func(uint8) complex128 { return 987654 }).Get(),
			ShouldEqual, 987654)

		So(NewEmptyUint8().ForceMapToComplex128(func(uint8) complex128 { return 987654 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToNillableComplex128(t *testing.T) {
	Convey("Uint8.MapToNillableComplex128", t, func() {
		So(NewUint8(123).MapToNillableComplex128(func(uint8) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).Get(), ShouldEqual, 987654)

		So(NewEmptyUint8().MapToNillableComplex128(func(uint8) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint8(123).MapToNillableComplex128(func(uint8) *complex128 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToComplex128(t *testing.T) {
	Convey("Uint8.MapToComplex128", t, func() {
		test1, err1 := NewUint8(123).MapToComplex128(func(uint8) (complex128, error) {
			return 987654, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 987654)

		test2, err2 := NewEmptyUint8().MapToComplex128(func(uint8) (complex128, error) {
			return 987654, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint8(123).MapToComplex128(func(uint8) (complex128, error) {
			return 987654, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint8_ForceMapToComplex64(t *testing.T) {
	Convey("Uint8.ForceMapToComplex64", t, func() {
		So(NewUint8(123).ForceMapToComplex64(func(uint8) complex64 { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyUint8().ForceMapToComplex64(func(uint8) complex64 { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToNillableComplex64(t *testing.T) {
	Convey("Uint8.MapToNillableComplex64", t, func() {
		So(NewUint8(123).MapToNillableComplex64(func(uint8) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyUint8().MapToNillableComplex64(func(uint8) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint8(123).MapToNillableComplex64(func(uint8) *complex64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToComplex64(t *testing.T) {
	Convey("Uint8.MapToComplex64", t, func() {
		test1, err1 := NewUint8(123).MapToComplex64(func(uint8) (complex64, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyUint8().MapToComplex64(func(uint8) (complex64, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint8(123).MapToComplex64(func(uint8) (complex64, error) {
			return 123456, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint8_ForceMapToFloat32(t *testing.T) {
	Convey("Uint8.ForceMapToFloat32", t, func() {
		So(NewUint8(123).ForceMapToFloat32(func(uint8) float32 { return 32.3 }).Get(),
			ShouldEqual, 32.3)

		So(NewEmptyUint8().ForceMapToFloat32(func(uint8) float32 { return 32.3 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToNillableFloat32(t *testing.T) {
	Convey("Uint8.MapToNillableFloat32", t, func() {
		So(NewUint8(123).MapToNillableFloat32(func(uint8) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).Get(), ShouldEqual, 32.3)

		So(NewEmptyUint8().MapToNillableFloat32(func(uint8) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint8(123).MapToNillableFloat32(func(uint8) *float32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToFloat32(t *testing.T) {
	Convey("Uint8.MapToFloat32", t, func() {
		test1, err1 := NewUint8(123).MapToFloat32(func(uint8) (float32, error) {
			return 32.3, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 32.3)

		test2, err2 := NewEmptyUint8().MapToFloat32(func(uint8) (float32, error) {
			return 32.3, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint8(123).MapToFloat32(func(uint8) (float32, error) {
			return 32.3, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint8_ForceMapToFloat64(t *testing.T) {
	Convey("Uint8.ForceMapToFloat64", t, func() {
		So(NewUint8(123).ForceMapToFloat64(func(uint8) float64 { return 3.2455 }).Get(),
			ShouldEqual, 3.2455)

		So(NewEmptyUint8().ForceMapToFloat64(func(uint8) float64 { return 3.2455 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToNillableFloat64(t *testing.T) {
	Convey("Uint8.MapToNillableFloat64", t, func() {
		So(NewUint8(123).MapToNillableFloat64(func(uint8) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).Get(), ShouldEqual, 3.2455)

		So(NewEmptyUint8().MapToNillableFloat64(func(uint8) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint8(123).MapToNillableFloat64(func(uint8) *float64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToFloat64(t *testing.T) {
	Convey("Uint8.MapToFloat64", t, func() {
		test1, err1 := NewUint8(123).MapToFloat64(func(uint8) (float64, error) {
			return 3.2455, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 3.2455)

		test2, err2 := NewEmptyUint8().MapToFloat64(func(uint8) (float64, error) {
			return 3.2455, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint8(123).MapToFloat64(func(uint8) (float64, error) {
			return 3.2455, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint8_ForceMapToInt(t *testing.T) {
	Convey("Uint8.ForceMapToInt", t, func() {
		So(NewUint8(123).ForceMapToInt(func(uint8) int { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyUint8().ForceMapToInt(func(uint8) int { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToNillableInt(t *testing.T) {
	Convey("Uint8.MapToNillableInt", t, func() {
		So(NewUint8(123).MapToNillableInt(func(uint8) *int {
			var tmp int = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyUint8().MapToNillableInt(func(uint8) *int {
			var tmp int = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint8(123).MapToNillableInt(func(uint8) *int { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToInt(t *testing.T) {
	Convey("Uint8.MapToInt", t, func() {
		test1, err1 := NewUint8(123).MapToInt(func(uint8) (int, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyUint8().MapToInt(func(uint8) (int, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint8(123).MapToInt(func(uint8) (int, error) {
			return 123456, errors.New("int error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint8_ForceMapToInt16(t *testing.T) {
	Convey("Uint8.ForceMapToInt16", t, func() {
		So(NewUint8(123).ForceMapToInt16(func(uint8) int16 { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyUint8().ForceMapToInt16(func(uint8) int16 { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToNillableInt16(t *testing.T) {
	Convey("Uint8.MapToNillableInt16", t, func() {
		So(NewUint8(123).MapToNillableInt16(func(uint8) *int16 {
			var tmp int16 = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyUint8().MapToNillableInt16(func(uint8) *int16 {
			var tmp int16 = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint8(123).MapToNillableInt16(func(uint8) *int16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToInt16(t *testing.T) {
	Convey("Uint8.MapToInt16", t, func() {
		test1, err1 := NewUint8(123).MapToInt16(func(uint8) (int16, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyUint8().MapToInt16(func(uint8) (int16, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint8(123).MapToInt16(func(uint8) (int16, error) {
			return 222, errors.New("int16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint8_ForceMapToInt32(t *testing.T) {
	Convey("Uint8.ForceMapToInt32", t, func() {
		So(NewUint8(123).ForceMapToInt32(func(uint8) int32 { return 333 }).Get(),
			ShouldEqual, 333)

		So(NewEmptyUint8().ForceMapToInt32(func(uint8) int32 { return 333 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToNillableInt32(t *testing.T) {
	Convey("Uint8.MapToNillableInt32", t, func() {
		So(NewUint8(123).MapToNillableInt32(func(uint8) *int32 {
			var tmp int32 = 333
			return &tmp
		}).Get(), ShouldEqual, 333)

		So(NewEmptyUint8().MapToNillableInt32(func(uint8) *int32 {
			var tmp int32 = 333
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint8(123).MapToNillableInt32(func(uint8) *int32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToInt32(t *testing.T) {
	Convey("Uint8.MapToInt32", t, func() {
		test1, err1 := NewUint8(123).MapToInt32(func(uint8) (int32, error) {
			return 333, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 333)

		test2, err2 := NewEmptyUint8().MapToInt32(func(uint8) (int32, error) {
			return 333, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint8(123).MapToInt32(func(uint8) (int32, error) {
			return 333, errors.New("int32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint8_ForceMapToInt64(t *testing.T) {
	Convey("Uint8.ForceMapToInt64", t, func() {
		So(NewUint8(123).ForceMapToInt64(func(uint8) int64 { return 444 }).Get(),
			ShouldEqual, 444)

		So(NewEmptyUint8().ForceMapToInt64(func(uint8) int64 { return 444 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToNillableInt64(t *testing.T) {
	Convey("Uint8.MapToNillableInt64", t, func() {
		So(NewUint8(123).MapToNillableInt64(func(uint8) *int64 {
			var tmp int64 = 444
			return &tmp
		}).Get(), ShouldEqual, 444)

		So(NewEmptyUint8().MapToNillableInt64(func(uint8) *int64 {
			var tmp int64 = 444
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint8(123).MapToNillableInt64(func(uint8) *int64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToInt64(t *testing.T) {
	Convey("Uint8.MapToInt64", t, func() {
		test1, err1 := NewUint8(123).MapToInt64(func(uint8) (int64, error) {
			return 444, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 444)

		test2, err2 := NewEmptyUint8().MapToInt64(func(uint8) (int64, error) {
			return 444, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint8(123).MapToInt64(func(uint8) (int64, error) {
			return 444, errors.New("int64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint8_ForceMapToInt8(t *testing.T) {
	Convey("Uint8.ForceMapToInt8", t, func() {
		So(NewUint8(123).ForceMapToInt8(func(uint8) int8 { return 111 }).Get(),
			ShouldEqual, 111)

		So(NewEmptyUint8().ForceMapToInt8(func(uint8) int8 { return 111 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToNillableInt8(t *testing.T) {
	Convey("Uint8.MapToNillableInt8", t, func() {
		So(NewUint8(123).MapToNillableInt8(func(uint8) *int8 {
			var tmp int8 = 111
			return &tmp
		}).Get(), ShouldEqual, 111)

		So(NewEmptyUint8().MapToNillableInt8(func(uint8) *int8 {
			var tmp int8 = 111
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint8(123).MapToNillableInt8(func(uint8) *int8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToInt8(t *testing.T) {
	Convey("Uint8.MapToInt8", t, func() {
		test1, err1 := NewUint8(123).MapToInt8(func(uint8) (int8, error) {
			return 111, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 111)

		test2, err2 := NewEmptyUint8().MapToInt8(func(uint8) (int8, error) {
			return 111, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint8(123).MapToInt8(func(uint8) (int8, error) {
			return 111, errors.New("int8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint8_ForceMapToUntyped(t *testing.T) {
	Convey("Uint8.ForceMapToUntyped", t, func() {
		So(NewUint8(123).ForceMapToUntyped(func(uint8) interface{} { return "111111" }).Get(),
			ShouldEqual, "111111")

		So(NewEmptyUint8().ForceMapToUntyped(func(uint8) interface{} { return "111111" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToNillableUntyped(t *testing.T) {
	Convey("Uint8.MapToNillableUntyped", t, func() {
		So(NewUint8(123).MapToNillableUntyped(func(uint8) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).Get(), ShouldEqual, "111111")

		So(NewEmptyUint8().MapToNillableUntyped(func(uint8) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint8(123).MapToNillableUntyped(func(uint8) *interface{} { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToUntyped(t *testing.T) {
	Convey("Uint8.MapToUntyped", t, func() {
		test1, err1 := NewUint8(123).MapToUntyped(func(uint8) (interface{}, error) {
			return "111111", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "111111")

		test2, err2 := NewEmptyUint8().MapToUntyped(func(uint8) (interface{}, error) {
			return "111111", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint8(123).MapToUntyped(func(uint8) (interface{}, error) {
			return "111111", errors.New("interface error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "interface error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint8_ForceMapToRune(t *testing.T) {
	Convey("Uint8.ForceMapToRune", t, func() {
		So(NewUint8(123).ForceMapToRune(func(uint8) rune { return 555 }).Get(),
			ShouldEqual, 555)

		So(NewEmptyUint8().ForceMapToRune(func(uint8) rune { return 555 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToNillableRune(t *testing.T) {
	Convey("Uint8.MapToNillableRune", t, func() {
		So(NewUint8(123).MapToNillableRune(func(uint8) *rune {
			var tmp rune = 555
			return &tmp
		}).Get(), ShouldEqual, 555)

		So(NewEmptyUint8().MapToNillableRune(func(uint8) *rune {
			var tmp rune = 555
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint8(123).MapToNillableRune(func(uint8) *rune { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToRune(t *testing.T) {
	Convey("Uint8.MapToRune", t, func() {
		test1, err1 := NewUint8(123).MapToRune(func(uint8) (rune, error) {
			return 555, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 555)

		test2, err2 := NewEmptyUint8().MapToRune(func(uint8) (rune, error) {
			return 555, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint8(123).MapToRune(func(uint8) (rune, error) {
			return 555, errors.New("rune error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "rune error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint8_ForceMapToString(t *testing.T) {
	Convey("Uint8.ForceMapToString", t, func() {
		So(NewUint8(123).ForceMapToString(func(uint8) string { return "bonanza" }).Get(),
			ShouldEqual, "bonanza")

		So(NewEmptyUint8().ForceMapToString(func(uint8) string { return "bonanza" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToNillableString(t *testing.T) {
	Convey("Uint8.MapToNillableString", t, func() {
		So(NewUint8(123).MapToNillableString(func(uint8) *string {
			var tmp string = "bonanza"
			return &tmp
		}).Get(), ShouldEqual, "bonanza")

		So(NewEmptyUint8().MapToNillableString(func(uint8) *string {
			var tmp string = "bonanza"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint8(123).MapToNillableString(func(uint8) *string { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToString(t *testing.T) {
	Convey("Uint8.MapToString", t, func() {
		test1, err1 := NewUint8(123).MapToString(func(uint8) (string, error) {
			return "bonanza", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "bonanza")

		test2, err2 := NewEmptyUint8().MapToString(func(uint8) (string, error) {
			return "bonanza", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint8(123).MapToString(func(uint8) (string, error) {
			return "bonanza", errors.New("string error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "string error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint8_ForceMapToUint(t *testing.T) {
	Convey("Uint8.ForceMapToUint", t, func() {
		So(NewUint8(123).ForceMapToUint(func(uint8) uint { return 666 }).Get(),
			ShouldEqual, 666)

		So(NewEmptyUint8().ForceMapToUint(func(uint8) uint { return 666 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToNillableUint(t *testing.T) {
	Convey("Uint8.MapToNillableUint", t, func() {
		So(NewUint8(123).MapToNillableUint(func(uint8) *uint {
			var tmp uint = 666
			return &tmp
		}).Get(), ShouldEqual, 666)

		So(NewEmptyUint8().MapToNillableUint(func(uint8) *uint {
			var tmp uint = 666
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint8(123).MapToNillableUint(func(uint8) *uint { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToUint(t *testing.T) {
	Convey("Uint8.MapToUint", t, func() {
		test1, err1 := NewUint8(123).MapToUint(func(uint8) (uint, error) {
			return 666, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 666)

		test2, err2 := NewEmptyUint8().MapToUint(func(uint8) (uint, error) {
			return 666, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint8(123).MapToUint(func(uint8) (uint, error) {
			return 666, errors.New("uint error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint8_ForceMapToUint16(t *testing.T) {
	Convey("Uint8.ForceMapToUint16", t, func() {
		So(NewUint8(123).ForceMapToUint16(func(uint8) uint16 { return 888 }).Get(),
			ShouldEqual, 888)

		So(NewEmptyUint8().ForceMapToUint16(func(uint8) uint16 { return 888 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToNillableUint16(t *testing.T) {
	Convey("Uint8.MapToNillableUint16", t, func() {
		So(NewUint8(123).MapToNillableUint16(func(uint8) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).Get(), ShouldEqual, 888)

		So(NewEmptyUint8().MapToNillableUint16(func(uint8) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint8(123).MapToNillableUint16(func(uint8) *uint16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToUint16(t *testing.T) {
	Convey("Uint8.MapToUint16", t, func() {
		test1, err1 := NewUint8(123).MapToUint16(func(uint8) (uint16, error) {
			return 888, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 888)

		test2, err2 := NewEmptyUint8().MapToUint16(func(uint8) (uint16, error) {
			return 888, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint8(123).MapToUint16(func(uint8) (uint16, error) {
			return 888, errors.New("uint16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint8_ForceMapToUint32(t *testing.T) {
	Convey("Uint8.ForceMapToUint32", t, func() {
		So(NewUint8(123).ForceMapToUint32(func(uint8) uint32 { return 999 }).Get(),
			ShouldEqual, 999)

		So(NewEmptyUint8().ForceMapToUint32(func(uint8) uint32 { return 999 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToNillableUint32(t *testing.T) {
	Convey("Uint8.MapToNillableUint32", t, func() {
		So(NewUint8(123).MapToNillableUint32(func(uint8) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).Get(), ShouldEqual, 999)

		So(NewEmptyUint8().MapToNillableUint32(func(uint8) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint8(123).MapToNillableUint32(func(uint8) *uint32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToUint32(t *testing.T) {
	Convey("Uint8.MapToUint32", t, func() {
		test1, err1 := NewUint8(123).MapToUint32(func(uint8) (uint32, error) {
			return 999, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 999)

		test2, err2 := NewEmptyUint8().MapToUint32(func(uint8) (uint32, error) {
			return 999, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint8(123).MapToUint32(func(uint8) (uint32, error) {
			return 999, errors.New("uint32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint8_ForceMapToUint64(t *testing.T) {
	Convey("Uint8.ForceMapToUint64", t, func() {
		So(NewUint8(123).ForceMapToUint64(func(uint8) uint64 { return 101010 }).Get(),
			ShouldEqual, 101010)

		So(NewEmptyUint8().ForceMapToUint64(func(uint8) uint64 { return 101010 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToNillableUint64(t *testing.T) {
	Convey("Uint8.MapToNillableUint64", t, func() {
		So(NewUint8(123).MapToNillableUint64(func(uint8) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).Get(), ShouldEqual, 101010)

		So(NewEmptyUint8().MapToNillableUint64(func(uint8) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint8(123).MapToNillableUint64(func(uint8) *uint64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint8_MapToUint64(t *testing.T) {
	Convey("Uint8.MapToUint64", t, func() {
		test1, err1 := NewUint8(123).MapToUint64(func(uint8) (uint64, error) {
			return 101010, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 101010)

		test2, err2 := NewEmptyUint8().MapToUint64(func(uint8) (uint64, error) {
			return 101010, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint8(123).MapToUint64(func(uint8) (uint64, error) {
			return 101010, errors.New("uint64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}
