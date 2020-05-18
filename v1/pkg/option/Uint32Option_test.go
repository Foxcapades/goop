// Package option_test generated @ 2020-05-17 22:19:55.164878293 -0400 EDT m=+0.001439508
package option_test

import (
	"errors"
	"testing"

	. "github.com/Foxcapades/goop/v1/pkg/option"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewUint32(t *testing.T) {
	Convey("NewUint32", t, func() {
		test := NewUint32(123456)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyUint32(t *testing.T) {
	Convey("NewEmptyUint32", t, func() {
		test := NewEmptyUint32()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeUint32(t *testing.T) {
	Convey("NewMaybeUint32", t, func() {
		var val1 uint32

		test1 := NewMaybeUint32(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeUint32(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestUint32_Get(t *testing.T) {
	Convey("Uint32.Get", t, func() {
		test1 := NewUint32(123456)
		So(test1.Get(), ShouldEqual, 123456)

		test2 := NewEmptyUint32()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestUint32_Or(t *testing.T) {
	Convey("Uint32.Or", t, func() {
		test1 := NewUint32(123456)
		So(test1.Or(654321), ShouldEqual, 123456)

		test2 := NewEmptyUint32()
		So(test2.Or(123456), ShouldEqual, 123456)
	})
}

func TestUint32_OrGet(t *testing.T) {
	Convey("Uint32.OrGet", t, func() {
		test1 := NewUint32(123456)
		So(test1.OrGet(func() uint32 { return 654321 }), ShouldEqual, 123456)

		test2 := NewEmptyUint32()
		So(test2.OrGet(func() uint32 { return 123456 }), ShouldEqual, 123456)
	})
}

func TestUint32_OrPanicWith(t *testing.T) {
	Convey("Uint32.OrPanicWith", t, func() {
		test1 := NewUint32(123456)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 123456)

		test2 := NewEmptyUint32()
		So(func() {
			test2.OrPanicWith(func() interface{} { return "test value" })
		}, ShouldPanicWith, "test value")
	})
}

func TestUint32_MapToNillable(t *testing.T) {
	Convey("Uint32.MapToNillable", t, func() {
		test1 := NewUint32(123456)
		So(test1.MapToNillable(func(b uint32) *uint32 { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewUint32(123456)
		So(test2.MapToNillable(func(b uint32) *uint32 { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyUint32()
		So(func() {
			test3.MapToNillable(func(b uint32) *uint32 { panic("foo") }).IsNil()
		}, ShouldNotPanic)
	})
}

func TestUint32_ForceMapToBool(t *testing.T) {
	Convey("Uint32.ForceMapToBool", t, func() {
		So(NewUint32(123456).ForceMapToBool(func(uint32) bool { return true }).Get(),
			ShouldEqual, true)

		So(NewEmptyUint32().ForceMapToBool(func(uint32) bool { return true }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToNillableBool(t *testing.T) {
	Convey("Uint32.MapToNillableBool", t, func() {
		So(NewUint32(123456).MapToNillableBool(func(uint32) *bool {
			var tmp bool = true
			return &tmp
		}).Get(), ShouldEqual, true)

		So(NewEmptyUint32().MapToNillableBool(func(uint32) *bool {
			var tmp bool = true
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint32(123456).MapToNillableBool(func(uint32) *bool { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToBool(t *testing.T) {
	Convey("Uint32.MapToBool", t, func() {
		test1, err1 := NewUint32(123456).MapToBool(func(uint32) (bool, error) {
			return true, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, true)

		test2, err2 := NewEmptyUint32().MapToBool(func(uint32) (bool, error) {
			return true, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint32(123456).MapToBool(func(uint32) (bool, error) {
			return true, errors.New("bool error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "bool error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint32_ForceMapToByte(t *testing.T) {
	Convey("Uint32.ForceMapToByte", t, func() {
		So(NewUint32(123456).ForceMapToByte(func(uint32) byte { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyUint32().ForceMapToByte(func(uint32) byte { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToNillableByte(t *testing.T) {
	Convey("Uint32.MapToNillableByte", t, func() {
		So(NewUint32(123456).MapToNillableByte(func(uint32) *byte {
			var tmp byte = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyUint32().MapToNillableByte(func(uint32) *byte {
			var tmp byte = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint32(123456).MapToNillableByte(func(uint32) *byte { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToByte(t *testing.T) {
	Convey("Uint32.MapToByte", t, func() {
		test1, err1 := NewUint32(123456).MapToByte(func(uint32) (byte, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyUint32().MapToByte(func(uint32) (byte, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint32(123456).MapToByte(func(uint32) (byte, error) {
			return 222, errors.New("byte error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "byte error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint32_ForceMapToComplex128(t *testing.T) {
	Convey("Uint32.ForceMapToComplex128", t, func() {
		So(NewUint32(123456).ForceMapToComplex128(func(uint32) complex128 { return 987654 }).Get(),
			ShouldEqual, 987654)

		So(NewEmptyUint32().ForceMapToComplex128(func(uint32) complex128 { return 987654 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToNillableComplex128(t *testing.T) {
	Convey("Uint32.MapToNillableComplex128", t, func() {
		So(NewUint32(123456).MapToNillableComplex128(func(uint32) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).Get(), ShouldEqual, 987654)

		So(NewEmptyUint32().MapToNillableComplex128(func(uint32) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint32(123456).MapToNillableComplex128(func(uint32) *complex128 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToComplex128(t *testing.T) {
	Convey("Uint32.MapToComplex128", t, func() {
		test1, err1 := NewUint32(123456).MapToComplex128(func(uint32) (complex128, error) {
			return 987654, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 987654)

		test2, err2 := NewEmptyUint32().MapToComplex128(func(uint32) (complex128, error) {
			return 987654, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint32(123456).MapToComplex128(func(uint32) (complex128, error) {
			return 987654, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint32_ForceMapToComplex64(t *testing.T) {
	Convey("Uint32.ForceMapToComplex64", t, func() {
		So(NewUint32(123456).ForceMapToComplex64(func(uint32) complex64 { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyUint32().ForceMapToComplex64(func(uint32) complex64 { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToNillableComplex64(t *testing.T) {
	Convey("Uint32.MapToNillableComplex64", t, func() {
		So(NewUint32(123456).MapToNillableComplex64(func(uint32) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyUint32().MapToNillableComplex64(func(uint32) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint32(123456).MapToNillableComplex64(func(uint32) *complex64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToComplex64(t *testing.T) {
	Convey("Uint32.MapToComplex64", t, func() {
		test1, err1 := NewUint32(123456).MapToComplex64(func(uint32) (complex64, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyUint32().MapToComplex64(func(uint32) (complex64, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint32(123456).MapToComplex64(func(uint32) (complex64, error) {
			return 123456, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint32_ForceMapToFloat32(t *testing.T) {
	Convey("Uint32.ForceMapToFloat32", t, func() {
		So(NewUint32(123456).ForceMapToFloat32(func(uint32) float32 { return 32.3 }).Get(),
			ShouldEqual, 32.3)

		So(NewEmptyUint32().ForceMapToFloat32(func(uint32) float32 { return 32.3 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToNillableFloat32(t *testing.T) {
	Convey("Uint32.MapToNillableFloat32", t, func() {
		So(NewUint32(123456).MapToNillableFloat32(func(uint32) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).Get(), ShouldEqual, 32.3)

		So(NewEmptyUint32().MapToNillableFloat32(func(uint32) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint32(123456).MapToNillableFloat32(func(uint32) *float32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToFloat32(t *testing.T) {
	Convey("Uint32.MapToFloat32", t, func() {
		test1, err1 := NewUint32(123456).MapToFloat32(func(uint32) (float32, error) {
			return 32.3, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 32.3)

		test2, err2 := NewEmptyUint32().MapToFloat32(func(uint32) (float32, error) {
			return 32.3, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint32(123456).MapToFloat32(func(uint32) (float32, error) {
			return 32.3, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint32_ForceMapToFloat64(t *testing.T) {
	Convey("Uint32.ForceMapToFloat64", t, func() {
		So(NewUint32(123456).ForceMapToFloat64(func(uint32) float64 { return 3.2455 }).Get(),
			ShouldEqual, 3.2455)

		So(NewEmptyUint32().ForceMapToFloat64(func(uint32) float64 { return 3.2455 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToNillableFloat64(t *testing.T) {
	Convey("Uint32.MapToNillableFloat64", t, func() {
		So(NewUint32(123456).MapToNillableFloat64(func(uint32) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).Get(), ShouldEqual, 3.2455)

		So(NewEmptyUint32().MapToNillableFloat64(func(uint32) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint32(123456).MapToNillableFloat64(func(uint32) *float64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToFloat64(t *testing.T) {
	Convey("Uint32.MapToFloat64", t, func() {
		test1, err1 := NewUint32(123456).MapToFloat64(func(uint32) (float64, error) {
			return 3.2455, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 3.2455)

		test2, err2 := NewEmptyUint32().MapToFloat64(func(uint32) (float64, error) {
			return 3.2455, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint32(123456).MapToFloat64(func(uint32) (float64, error) {
			return 3.2455, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint32_ForceMapToInt(t *testing.T) {
	Convey("Uint32.ForceMapToInt", t, func() {
		So(NewUint32(123456).ForceMapToInt(func(uint32) int { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyUint32().ForceMapToInt(func(uint32) int { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToNillableInt(t *testing.T) {
	Convey("Uint32.MapToNillableInt", t, func() {
		So(NewUint32(123456).MapToNillableInt(func(uint32) *int {
			var tmp int = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyUint32().MapToNillableInt(func(uint32) *int {
			var tmp int = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint32(123456).MapToNillableInt(func(uint32) *int { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToInt(t *testing.T) {
	Convey("Uint32.MapToInt", t, func() {
		test1, err1 := NewUint32(123456).MapToInt(func(uint32) (int, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyUint32().MapToInt(func(uint32) (int, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint32(123456).MapToInt(func(uint32) (int, error) {
			return 123456, errors.New("int error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint32_ForceMapToInt16(t *testing.T) {
	Convey("Uint32.ForceMapToInt16", t, func() {
		So(NewUint32(123456).ForceMapToInt16(func(uint32) int16 { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyUint32().ForceMapToInt16(func(uint32) int16 { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToNillableInt16(t *testing.T) {
	Convey("Uint32.MapToNillableInt16", t, func() {
		So(NewUint32(123456).MapToNillableInt16(func(uint32) *int16 {
			var tmp int16 = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyUint32().MapToNillableInt16(func(uint32) *int16 {
			var tmp int16 = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint32(123456).MapToNillableInt16(func(uint32) *int16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToInt16(t *testing.T) {
	Convey("Uint32.MapToInt16", t, func() {
		test1, err1 := NewUint32(123456).MapToInt16(func(uint32) (int16, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyUint32().MapToInt16(func(uint32) (int16, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint32(123456).MapToInt16(func(uint32) (int16, error) {
			return 222, errors.New("int16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint32_ForceMapToInt32(t *testing.T) {
	Convey("Uint32.ForceMapToInt32", t, func() {
		So(NewUint32(123456).ForceMapToInt32(func(uint32) int32 { return 333 }).Get(),
			ShouldEqual, 333)

		So(NewEmptyUint32().ForceMapToInt32(func(uint32) int32 { return 333 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToNillableInt32(t *testing.T) {
	Convey("Uint32.MapToNillableInt32", t, func() {
		So(NewUint32(123456).MapToNillableInt32(func(uint32) *int32 {
			var tmp int32 = 333
			return &tmp
		}).Get(), ShouldEqual, 333)

		So(NewEmptyUint32().MapToNillableInt32(func(uint32) *int32 {
			var tmp int32 = 333
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint32(123456).MapToNillableInt32(func(uint32) *int32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToInt32(t *testing.T) {
	Convey("Uint32.MapToInt32", t, func() {
		test1, err1 := NewUint32(123456).MapToInt32(func(uint32) (int32, error) {
			return 333, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 333)

		test2, err2 := NewEmptyUint32().MapToInt32(func(uint32) (int32, error) {
			return 333, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint32(123456).MapToInt32(func(uint32) (int32, error) {
			return 333, errors.New("int32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint32_ForceMapToInt64(t *testing.T) {
	Convey("Uint32.ForceMapToInt64", t, func() {
		So(NewUint32(123456).ForceMapToInt64(func(uint32) int64 { return 444 }).Get(),
			ShouldEqual, 444)

		So(NewEmptyUint32().ForceMapToInt64(func(uint32) int64 { return 444 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToNillableInt64(t *testing.T) {
	Convey("Uint32.MapToNillableInt64", t, func() {
		So(NewUint32(123456).MapToNillableInt64(func(uint32) *int64 {
			var tmp int64 = 444
			return &tmp
		}).Get(), ShouldEqual, 444)

		So(NewEmptyUint32().MapToNillableInt64(func(uint32) *int64 {
			var tmp int64 = 444
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint32(123456).MapToNillableInt64(func(uint32) *int64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToInt64(t *testing.T) {
	Convey("Uint32.MapToInt64", t, func() {
		test1, err1 := NewUint32(123456).MapToInt64(func(uint32) (int64, error) {
			return 444, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 444)

		test2, err2 := NewEmptyUint32().MapToInt64(func(uint32) (int64, error) {
			return 444, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint32(123456).MapToInt64(func(uint32) (int64, error) {
			return 444, errors.New("int64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint32_ForceMapToInt8(t *testing.T) {
	Convey("Uint32.ForceMapToInt8", t, func() {
		So(NewUint32(123456).ForceMapToInt8(func(uint32) int8 { return 111 }).Get(),
			ShouldEqual, 111)

		So(NewEmptyUint32().ForceMapToInt8(func(uint32) int8 { return 111 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToNillableInt8(t *testing.T) {
	Convey("Uint32.MapToNillableInt8", t, func() {
		So(NewUint32(123456).MapToNillableInt8(func(uint32) *int8 {
			var tmp int8 = 111
			return &tmp
		}).Get(), ShouldEqual, 111)

		So(NewEmptyUint32().MapToNillableInt8(func(uint32) *int8 {
			var tmp int8 = 111
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint32(123456).MapToNillableInt8(func(uint32) *int8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToInt8(t *testing.T) {
	Convey("Uint32.MapToInt8", t, func() {
		test1, err1 := NewUint32(123456).MapToInt8(func(uint32) (int8, error) {
			return 111, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 111)

		test2, err2 := NewEmptyUint32().MapToInt8(func(uint32) (int8, error) {
			return 111, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint32(123456).MapToInt8(func(uint32) (int8, error) {
			return 111, errors.New("int8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint32_ForceMapToUntyped(t *testing.T) {
	Convey("Uint32.ForceMapToUntyped", t, func() {
		So(NewUint32(123456).ForceMapToUntyped(func(uint32) interface{} { return "111111" }).Get(),
			ShouldEqual, "111111")

		So(NewEmptyUint32().ForceMapToUntyped(func(uint32) interface{} { return "111111" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToNillableUntyped(t *testing.T) {
	Convey("Uint32.MapToNillableUntyped", t, func() {
		So(NewUint32(123456).MapToNillableUntyped(func(uint32) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).Get(), ShouldEqual, "111111")

		So(NewEmptyUint32().MapToNillableUntyped(func(uint32) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint32(123456).MapToNillableUntyped(func(uint32) *interface{} { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToUntyped(t *testing.T) {
	Convey("Uint32.MapToUntyped", t, func() {
		test1, err1 := NewUint32(123456).MapToUntyped(func(uint32) (interface{}, error) {
			return "111111", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "111111")

		test2, err2 := NewEmptyUint32().MapToUntyped(func(uint32) (interface{}, error) {
			return "111111", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint32(123456).MapToUntyped(func(uint32) (interface{}, error) {
			return "111111", errors.New("interface error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "interface error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint32_ForceMapToRune(t *testing.T) {
	Convey("Uint32.ForceMapToRune", t, func() {
		So(NewUint32(123456).ForceMapToRune(func(uint32) rune { return 555 }).Get(),
			ShouldEqual, 555)

		So(NewEmptyUint32().ForceMapToRune(func(uint32) rune { return 555 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToNillableRune(t *testing.T) {
	Convey("Uint32.MapToNillableRune", t, func() {
		So(NewUint32(123456).MapToNillableRune(func(uint32) *rune {
			var tmp rune = 555
			return &tmp
		}).Get(), ShouldEqual, 555)

		So(NewEmptyUint32().MapToNillableRune(func(uint32) *rune {
			var tmp rune = 555
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint32(123456).MapToNillableRune(func(uint32) *rune { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToRune(t *testing.T) {
	Convey("Uint32.MapToRune", t, func() {
		test1, err1 := NewUint32(123456).MapToRune(func(uint32) (rune, error) {
			return 555, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 555)

		test2, err2 := NewEmptyUint32().MapToRune(func(uint32) (rune, error) {
			return 555, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint32(123456).MapToRune(func(uint32) (rune, error) {
			return 555, errors.New("rune error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "rune error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint32_ForceMapToString(t *testing.T) {
	Convey("Uint32.ForceMapToString", t, func() {
		So(NewUint32(123456).ForceMapToString(func(uint32) string { return "bonanza" }).Get(),
			ShouldEqual, "bonanza")

		So(NewEmptyUint32().ForceMapToString(func(uint32) string { return "bonanza" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToNillableString(t *testing.T) {
	Convey("Uint32.MapToNillableString", t, func() {
		So(NewUint32(123456).MapToNillableString(func(uint32) *string {
			var tmp string = "bonanza"
			return &tmp
		}).Get(), ShouldEqual, "bonanza")

		So(NewEmptyUint32().MapToNillableString(func(uint32) *string {
			var tmp string = "bonanza"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint32(123456).MapToNillableString(func(uint32) *string { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToString(t *testing.T) {
	Convey("Uint32.MapToString", t, func() {
		test1, err1 := NewUint32(123456).MapToString(func(uint32) (string, error) {
			return "bonanza", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "bonanza")

		test2, err2 := NewEmptyUint32().MapToString(func(uint32) (string, error) {
			return "bonanza", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint32(123456).MapToString(func(uint32) (string, error) {
			return "bonanza", errors.New("string error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "string error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint32_ForceMapToUint(t *testing.T) {
	Convey("Uint32.ForceMapToUint", t, func() {
		So(NewUint32(123456).ForceMapToUint(func(uint32) uint { return 666 }).Get(),
			ShouldEqual, 666)

		So(NewEmptyUint32().ForceMapToUint(func(uint32) uint { return 666 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToNillableUint(t *testing.T) {
	Convey("Uint32.MapToNillableUint", t, func() {
		So(NewUint32(123456).MapToNillableUint(func(uint32) *uint {
			var tmp uint = 666
			return &tmp
		}).Get(), ShouldEqual, 666)

		So(NewEmptyUint32().MapToNillableUint(func(uint32) *uint {
			var tmp uint = 666
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint32(123456).MapToNillableUint(func(uint32) *uint { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToUint(t *testing.T) {
	Convey("Uint32.MapToUint", t, func() {
		test1, err1 := NewUint32(123456).MapToUint(func(uint32) (uint, error) {
			return 666, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 666)

		test2, err2 := NewEmptyUint32().MapToUint(func(uint32) (uint, error) {
			return 666, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint32(123456).MapToUint(func(uint32) (uint, error) {
			return 666, errors.New("uint error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint32_ForceMapToUint16(t *testing.T) {
	Convey("Uint32.ForceMapToUint16", t, func() {
		So(NewUint32(123456).ForceMapToUint16(func(uint32) uint16 { return 888 }).Get(),
			ShouldEqual, 888)

		So(NewEmptyUint32().ForceMapToUint16(func(uint32) uint16 { return 888 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToNillableUint16(t *testing.T) {
	Convey("Uint32.MapToNillableUint16", t, func() {
		So(NewUint32(123456).MapToNillableUint16(func(uint32) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).Get(), ShouldEqual, 888)

		So(NewEmptyUint32().MapToNillableUint16(func(uint32) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint32(123456).MapToNillableUint16(func(uint32) *uint16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToUint16(t *testing.T) {
	Convey("Uint32.MapToUint16", t, func() {
		test1, err1 := NewUint32(123456).MapToUint16(func(uint32) (uint16, error) {
			return 888, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 888)

		test2, err2 := NewEmptyUint32().MapToUint16(func(uint32) (uint16, error) {
			return 888, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint32(123456).MapToUint16(func(uint32) (uint16, error) {
			return 888, errors.New("uint16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint32_ForceMapToUint64(t *testing.T) {
	Convey("Uint32.ForceMapToUint64", t, func() {
		So(NewUint32(123456).ForceMapToUint64(func(uint32) uint64 { return 101010 }).Get(),
			ShouldEqual, 101010)

		So(NewEmptyUint32().ForceMapToUint64(func(uint32) uint64 { return 101010 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToNillableUint64(t *testing.T) {
	Convey("Uint32.MapToNillableUint64", t, func() {
		So(NewUint32(123456).MapToNillableUint64(func(uint32) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).Get(), ShouldEqual, 101010)

		So(NewEmptyUint32().MapToNillableUint64(func(uint32) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint32(123456).MapToNillableUint64(func(uint32) *uint64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToUint64(t *testing.T) {
	Convey("Uint32.MapToUint64", t, func() {
		test1, err1 := NewUint32(123456).MapToUint64(func(uint32) (uint64, error) {
			return 101010, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 101010)

		test2, err2 := NewEmptyUint32().MapToUint64(func(uint32) (uint64, error) {
			return 101010, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint32(123456).MapToUint64(func(uint32) (uint64, error) {
			return 101010, errors.New("uint64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint32_ForceMapToUint8(t *testing.T) {
	Convey("Uint32.ForceMapToUint8", t, func() {
		So(NewUint32(123456).ForceMapToUint8(func(uint32) uint8 { return 77 }).Get(),
			ShouldEqual, 77)

		So(NewEmptyUint32().ForceMapToUint8(func(uint32) uint8 { return 77 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToNillableUint8(t *testing.T) {
	Convey("Uint32.MapToNillableUint8", t, func() {
		So(NewUint32(123456).MapToNillableUint8(func(uint32) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).Get(), ShouldEqual, 77)

		So(NewEmptyUint32().MapToNillableUint8(func(uint32) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint32(123456).MapToNillableUint8(func(uint32) *uint8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint32_MapToUint8(t *testing.T) {
	Convey("Uint32.MapToUint8", t, func() {
		test1, err1 := NewUint32(123456).MapToUint8(func(uint32) (uint8, error) {
			return 77, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 77)

		test2, err2 := NewEmptyUint32().MapToUint8(func(uint32) (uint8, error) {
			return 77, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint32(123456).MapToUint8(func(uint32) (uint8, error) {
			return 77, errors.New("uint8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}
