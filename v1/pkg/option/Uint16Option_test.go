// Package option_test generated @ 2020-05-17 22:19:55.164878293 -0400 EDT m=+0.001439508
package option_test

import (
	"errors"
	"testing"

	. "github.com/Foxcapades/goop/v1/pkg/option"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewUint16(t *testing.T) {
	Convey("NewUint16", t, func() {
		test := NewUint16(4562)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyUint16(t *testing.T) {
	Convey("NewEmptyUint16", t, func() {
		test := NewEmptyUint16()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeUint16(t *testing.T) {
	Convey("NewMaybeUint16", t, func() {
		var val1 uint16

		test1 := NewMaybeUint16(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeUint16(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestUint16_Get(t *testing.T) {
	Convey("Uint16.Get", t, func() {
		test1 := NewUint16(4562)
		So(test1.Get(), ShouldEqual, 4562)

		test2 := NewEmptyUint16()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestUint16_Or(t *testing.T) {
	Convey("Uint16.Or", t, func() {
		test1 := NewUint16(4562)
		So(test1.Or(7894), ShouldEqual, 4562)

		test2 := NewEmptyUint16()
		So(test2.Or(4562), ShouldEqual, 4562)
	})
}

func TestUint16_OrGet(t *testing.T) {
	Convey("Uint16.OrGet", t, func() {
		test1 := NewUint16(4562)
		So(test1.OrGet(func() uint16 { return 7894 }), ShouldEqual, 4562)

		test2 := NewEmptyUint16()
		So(test2.OrGet(func() uint16 { return 4562 }), ShouldEqual, 4562)
	})
}

func TestUint16_OrPanicWith(t *testing.T) {
	Convey("Uint16.OrPanicWith", t, func() {
		test1 := NewUint16(4562)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 4562)

		test2 := NewEmptyUint16()
		So(func() {
			test2.OrPanicWith(func() interface{} { return "test value" })
		}, ShouldPanicWith, "test value")
	})
}

func TestUint16_MapToNillable(t *testing.T) {
	Convey("Uint16.MapToNillable", t, func() {
		test1 := NewUint16(4562)
		So(test1.MapToNillable(func(b uint16) *uint16 { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewUint16(4562)
		So(test2.MapToNillable(func(b uint16) *uint16 { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyUint16()
		So(func() {
			test3.MapToNillable(func(b uint16) *uint16 { panic("foo") }).IsNil()
		}, ShouldNotPanic)
	})
}

func TestUint16_ForceMapToBool(t *testing.T) {
	Convey("Uint16.ForceMapToBool", t, func() {
		So(NewUint16(4562).ForceMapToBool(func(uint16) bool { return true }).Get(),
			ShouldEqual, true)

		So(NewEmptyUint16().ForceMapToBool(func(uint16) bool { return true }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToNillableBool(t *testing.T) {
	Convey("Uint16.MapToNillableBool", t, func() {
		So(NewUint16(4562).MapToNillableBool(func(uint16) *bool {
			var tmp bool = true
			return &tmp
		}).Get(), ShouldEqual, true)

		So(NewEmptyUint16().MapToNillableBool(func(uint16) *bool {
			var tmp bool = true
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint16(4562).MapToNillableBool(func(uint16) *bool { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToBool(t *testing.T) {
	Convey("Uint16.MapToBool", t, func() {
		test1, err1 := NewUint16(4562).MapToBool(func(uint16) (bool, error) {
			return true, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, true)

		test2, err2 := NewEmptyUint16().MapToBool(func(uint16) (bool, error) {
			return true, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint16(4562).MapToBool(func(uint16) (bool, error) {
			return true, errors.New("bool error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "bool error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint16_ForceMapToByte(t *testing.T) {
	Convey("Uint16.ForceMapToByte", t, func() {
		So(NewUint16(4562).ForceMapToByte(func(uint16) byte { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyUint16().ForceMapToByte(func(uint16) byte { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToNillableByte(t *testing.T) {
	Convey("Uint16.MapToNillableByte", t, func() {
		So(NewUint16(4562).MapToNillableByte(func(uint16) *byte {
			var tmp byte = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyUint16().MapToNillableByte(func(uint16) *byte {
			var tmp byte = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint16(4562).MapToNillableByte(func(uint16) *byte { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToByte(t *testing.T) {
	Convey("Uint16.MapToByte", t, func() {
		test1, err1 := NewUint16(4562).MapToByte(func(uint16) (byte, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyUint16().MapToByte(func(uint16) (byte, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint16(4562).MapToByte(func(uint16) (byte, error) {
			return 222, errors.New("byte error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "byte error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint16_ForceMapToComplex128(t *testing.T) {
	Convey("Uint16.ForceMapToComplex128", t, func() {
		So(NewUint16(4562).ForceMapToComplex128(func(uint16) complex128 { return 987654 }).Get(),
			ShouldEqual, 987654)

		So(NewEmptyUint16().ForceMapToComplex128(func(uint16) complex128 { return 987654 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToNillableComplex128(t *testing.T) {
	Convey("Uint16.MapToNillableComplex128", t, func() {
		So(NewUint16(4562).MapToNillableComplex128(func(uint16) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).Get(), ShouldEqual, 987654)

		So(NewEmptyUint16().MapToNillableComplex128(func(uint16) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint16(4562).MapToNillableComplex128(func(uint16) *complex128 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToComplex128(t *testing.T) {
	Convey("Uint16.MapToComplex128", t, func() {
		test1, err1 := NewUint16(4562).MapToComplex128(func(uint16) (complex128, error) {
			return 987654, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 987654)

		test2, err2 := NewEmptyUint16().MapToComplex128(func(uint16) (complex128, error) {
			return 987654, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint16(4562).MapToComplex128(func(uint16) (complex128, error) {
			return 987654, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint16_ForceMapToComplex64(t *testing.T) {
	Convey("Uint16.ForceMapToComplex64", t, func() {
		So(NewUint16(4562).ForceMapToComplex64(func(uint16) complex64 { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyUint16().ForceMapToComplex64(func(uint16) complex64 { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToNillableComplex64(t *testing.T) {
	Convey("Uint16.MapToNillableComplex64", t, func() {
		So(NewUint16(4562).MapToNillableComplex64(func(uint16) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyUint16().MapToNillableComplex64(func(uint16) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint16(4562).MapToNillableComplex64(func(uint16) *complex64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToComplex64(t *testing.T) {
	Convey("Uint16.MapToComplex64", t, func() {
		test1, err1 := NewUint16(4562).MapToComplex64(func(uint16) (complex64, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyUint16().MapToComplex64(func(uint16) (complex64, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint16(4562).MapToComplex64(func(uint16) (complex64, error) {
			return 123456, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint16_ForceMapToFloat32(t *testing.T) {
	Convey("Uint16.ForceMapToFloat32", t, func() {
		So(NewUint16(4562).ForceMapToFloat32(func(uint16) float32 { return 32.3 }).Get(),
			ShouldEqual, 32.3)

		So(NewEmptyUint16().ForceMapToFloat32(func(uint16) float32 { return 32.3 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToNillableFloat32(t *testing.T) {
	Convey("Uint16.MapToNillableFloat32", t, func() {
		So(NewUint16(4562).MapToNillableFloat32(func(uint16) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).Get(), ShouldEqual, 32.3)

		So(NewEmptyUint16().MapToNillableFloat32(func(uint16) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint16(4562).MapToNillableFloat32(func(uint16) *float32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToFloat32(t *testing.T) {
	Convey("Uint16.MapToFloat32", t, func() {
		test1, err1 := NewUint16(4562).MapToFloat32(func(uint16) (float32, error) {
			return 32.3, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 32.3)

		test2, err2 := NewEmptyUint16().MapToFloat32(func(uint16) (float32, error) {
			return 32.3, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint16(4562).MapToFloat32(func(uint16) (float32, error) {
			return 32.3, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint16_ForceMapToFloat64(t *testing.T) {
	Convey("Uint16.ForceMapToFloat64", t, func() {
		So(NewUint16(4562).ForceMapToFloat64(func(uint16) float64 { return 3.2455 }).Get(),
			ShouldEqual, 3.2455)

		So(NewEmptyUint16().ForceMapToFloat64(func(uint16) float64 { return 3.2455 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToNillableFloat64(t *testing.T) {
	Convey("Uint16.MapToNillableFloat64", t, func() {
		So(NewUint16(4562).MapToNillableFloat64(func(uint16) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).Get(), ShouldEqual, 3.2455)

		So(NewEmptyUint16().MapToNillableFloat64(func(uint16) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint16(4562).MapToNillableFloat64(func(uint16) *float64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToFloat64(t *testing.T) {
	Convey("Uint16.MapToFloat64", t, func() {
		test1, err1 := NewUint16(4562).MapToFloat64(func(uint16) (float64, error) {
			return 3.2455, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 3.2455)

		test2, err2 := NewEmptyUint16().MapToFloat64(func(uint16) (float64, error) {
			return 3.2455, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint16(4562).MapToFloat64(func(uint16) (float64, error) {
			return 3.2455, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint16_ForceMapToInt(t *testing.T) {
	Convey("Uint16.ForceMapToInt", t, func() {
		So(NewUint16(4562).ForceMapToInt(func(uint16) int { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyUint16().ForceMapToInt(func(uint16) int { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToNillableInt(t *testing.T) {
	Convey("Uint16.MapToNillableInt", t, func() {
		So(NewUint16(4562).MapToNillableInt(func(uint16) *int {
			var tmp int = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyUint16().MapToNillableInt(func(uint16) *int {
			var tmp int = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint16(4562).MapToNillableInt(func(uint16) *int { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToInt(t *testing.T) {
	Convey("Uint16.MapToInt", t, func() {
		test1, err1 := NewUint16(4562).MapToInt(func(uint16) (int, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyUint16().MapToInt(func(uint16) (int, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint16(4562).MapToInt(func(uint16) (int, error) {
			return 123456, errors.New("int error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint16_ForceMapToInt16(t *testing.T) {
	Convey("Uint16.ForceMapToInt16", t, func() {
		So(NewUint16(4562).ForceMapToInt16(func(uint16) int16 { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyUint16().ForceMapToInt16(func(uint16) int16 { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToNillableInt16(t *testing.T) {
	Convey("Uint16.MapToNillableInt16", t, func() {
		So(NewUint16(4562).MapToNillableInt16(func(uint16) *int16 {
			var tmp int16 = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyUint16().MapToNillableInt16(func(uint16) *int16 {
			var tmp int16 = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint16(4562).MapToNillableInt16(func(uint16) *int16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToInt16(t *testing.T) {
	Convey("Uint16.MapToInt16", t, func() {
		test1, err1 := NewUint16(4562).MapToInt16(func(uint16) (int16, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyUint16().MapToInt16(func(uint16) (int16, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint16(4562).MapToInt16(func(uint16) (int16, error) {
			return 222, errors.New("int16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint16_ForceMapToInt32(t *testing.T) {
	Convey("Uint16.ForceMapToInt32", t, func() {
		So(NewUint16(4562).ForceMapToInt32(func(uint16) int32 { return 333 }).Get(),
			ShouldEqual, 333)

		So(NewEmptyUint16().ForceMapToInt32(func(uint16) int32 { return 333 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToNillableInt32(t *testing.T) {
	Convey("Uint16.MapToNillableInt32", t, func() {
		So(NewUint16(4562).MapToNillableInt32(func(uint16) *int32 {
			var tmp int32 = 333
			return &tmp
		}).Get(), ShouldEqual, 333)

		So(NewEmptyUint16().MapToNillableInt32(func(uint16) *int32 {
			var tmp int32 = 333
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint16(4562).MapToNillableInt32(func(uint16) *int32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToInt32(t *testing.T) {
	Convey("Uint16.MapToInt32", t, func() {
		test1, err1 := NewUint16(4562).MapToInt32(func(uint16) (int32, error) {
			return 333, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 333)

		test2, err2 := NewEmptyUint16().MapToInt32(func(uint16) (int32, error) {
			return 333, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint16(4562).MapToInt32(func(uint16) (int32, error) {
			return 333, errors.New("int32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint16_ForceMapToInt64(t *testing.T) {
	Convey("Uint16.ForceMapToInt64", t, func() {
		So(NewUint16(4562).ForceMapToInt64(func(uint16) int64 { return 444 }).Get(),
			ShouldEqual, 444)

		So(NewEmptyUint16().ForceMapToInt64(func(uint16) int64 { return 444 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToNillableInt64(t *testing.T) {
	Convey("Uint16.MapToNillableInt64", t, func() {
		So(NewUint16(4562).MapToNillableInt64(func(uint16) *int64 {
			var tmp int64 = 444
			return &tmp
		}).Get(), ShouldEqual, 444)

		So(NewEmptyUint16().MapToNillableInt64(func(uint16) *int64 {
			var tmp int64 = 444
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint16(4562).MapToNillableInt64(func(uint16) *int64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToInt64(t *testing.T) {
	Convey("Uint16.MapToInt64", t, func() {
		test1, err1 := NewUint16(4562).MapToInt64(func(uint16) (int64, error) {
			return 444, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 444)

		test2, err2 := NewEmptyUint16().MapToInt64(func(uint16) (int64, error) {
			return 444, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint16(4562).MapToInt64(func(uint16) (int64, error) {
			return 444, errors.New("int64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint16_ForceMapToInt8(t *testing.T) {
	Convey("Uint16.ForceMapToInt8", t, func() {
		So(NewUint16(4562).ForceMapToInt8(func(uint16) int8 { return 111 }).Get(),
			ShouldEqual, 111)

		So(NewEmptyUint16().ForceMapToInt8(func(uint16) int8 { return 111 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToNillableInt8(t *testing.T) {
	Convey("Uint16.MapToNillableInt8", t, func() {
		So(NewUint16(4562).MapToNillableInt8(func(uint16) *int8 {
			var tmp int8 = 111
			return &tmp
		}).Get(), ShouldEqual, 111)

		So(NewEmptyUint16().MapToNillableInt8(func(uint16) *int8 {
			var tmp int8 = 111
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint16(4562).MapToNillableInt8(func(uint16) *int8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToInt8(t *testing.T) {
	Convey("Uint16.MapToInt8", t, func() {
		test1, err1 := NewUint16(4562).MapToInt8(func(uint16) (int8, error) {
			return 111, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 111)

		test2, err2 := NewEmptyUint16().MapToInt8(func(uint16) (int8, error) {
			return 111, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint16(4562).MapToInt8(func(uint16) (int8, error) {
			return 111, errors.New("int8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint16_ForceMapToUntyped(t *testing.T) {
	Convey("Uint16.ForceMapToUntyped", t, func() {
		So(NewUint16(4562).ForceMapToUntyped(func(uint16) interface{} { return "111111" }).Get(),
			ShouldEqual, "111111")

		So(NewEmptyUint16().ForceMapToUntyped(func(uint16) interface{} { return "111111" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToNillableUntyped(t *testing.T) {
	Convey("Uint16.MapToNillableUntyped", t, func() {
		So(NewUint16(4562).MapToNillableUntyped(func(uint16) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).Get(), ShouldEqual, "111111")

		So(NewEmptyUint16().MapToNillableUntyped(func(uint16) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint16(4562).MapToNillableUntyped(func(uint16) *interface{} { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToUntyped(t *testing.T) {
	Convey("Uint16.MapToUntyped", t, func() {
		test1, err1 := NewUint16(4562).MapToUntyped(func(uint16) (interface{}, error) {
			return "111111", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "111111")

		test2, err2 := NewEmptyUint16().MapToUntyped(func(uint16) (interface{}, error) {
			return "111111", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint16(4562).MapToUntyped(func(uint16) (interface{}, error) {
			return "111111", errors.New("interface error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "interface error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint16_ForceMapToRune(t *testing.T) {
	Convey("Uint16.ForceMapToRune", t, func() {
		So(NewUint16(4562).ForceMapToRune(func(uint16) rune { return 555 }).Get(),
			ShouldEqual, 555)

		So(NewEmptyUint16().ForceMapToRune(func(uint16) rune { return 555 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToNillableRune(t *testing.T) {
	Convey("Uint16.MapToNillableRune", t, func() {
		So(NewUint16(4562).MapToNillableRune(func(uint16) *rune {
			var tmp rune = 555
			return &tmp
		}).Get(), ShouldEqual, 555)

		So(NewEmptyUint16().MapToNillableRune(func(uint16) *rune {
			var tmp rune = 555
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint16(4562).MapToNillableRune(func(uint16) *rune { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToRune(t *testing.T) {
	Convey("Uint16.MapToRune", t, func() {
		test1, err1 := NewUint16(4562).MapToRune(func(uint16) (rune, error) {
			return 555, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 555)

		test2, err2 := NewEmptyUint16().MapToRune(func(uint16) (rune, error) {
			return 555, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint16(4562).MapToRune(func(uint16) (rune, error) {
			return 555, errors.New("rune error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "rune error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint16_ForceMapToString(t *testing.T) {
	Convey("Uint16.ForceMapToString", t, func() {
		So(NewUint16(4562).ForceMapToString(func(uint16) string { return "bonanza" }).Get(),
			ShouldEqual, "bonanza")

		So(NewEmptyUint16().ForceMapToString(func(uint16) string { return "bonanza" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToNillableString(t *testing.T) {
	Convey("Uint16.MapToNillableString", t, func() {
		So(NewUint16(4562).MapToNillableString(func(uint16) *string {
			var tmp string = "bonanza"
			return &tmp
		}).Get(), ShouldEqual, "bonanza")

		So(NewEmptyUint16().MapToNillableString(func(uint16) *string {
			var tmp string = "bonanza"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint16(4562).MapToNillableString(func(uint16) *string { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToString(t *testing.T) {
	Convey("Uint16.MapToString", t, func() {
		test1, err1 := NewUint16(4562).MapToString(func(uint16) (string, error) {
			return "bonanza", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "bonanza")

		test2, err2 := NewEmptyUint16().MapToString(func(uint16) (string, error) {
			return "bonanza", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint16(4562).MapToString(func(uint16) (string, error) {
			return "bonanza", errors.New("string error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "string error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint16_ForceMapToUint(t *testing.T) {
	Convey("Uint16.ForceMapToUint", t, func() {
		So(NewUint16(4562).ForceMapToUint(func(uint16) uint { return 666 }).Get(),
			ShouldEqual, 666)

		So(NewEmptyUint16().ForceMapToUint(func(uint16) uint { return 666 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToNillableUint(t *testing.T) {
	Convey("Uint16.MapToNillableUint", t, func() {
		So(NewUint16(4562).MapToNillableUint(func(uint16) *uint {
			var tmp uint = 666
			return &tmp
		}).Get(), ShouldEqual, 666)

		So(NewEmptyUint16().MapToNillableUint(func(uint16) *uint {
			var tmp uint = 666
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint16(4562).MapToNillableUint(func(uint16) *uint { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToUint(t *testing.T) {
	Convey("Uint16.MapToUint", t, func() {
		test1, err1 := NewUint16(4562).MapToUint(func(uint16) (uint, error) {
			return 666, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 666)

		test2, err2 := NewEmptyUint16().MapToUint(func(uint16) (uint, error) {
			return 666, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint16(4562).MapToUint(func(uint16) (uint, error) {
			return 666, errors.New("uint error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint16_ForceMapToUint32(t *testing.T) {
	Convey("Uint16.ForceMapToUint32", t, func() {
		So(NewUint16(4562).ForceMapToUint32(func(uint16) uint32 { return 999 }).Get(),
			ShouldEqual, 999)

		So(NewEmptyUint16().ForceMapToUint32(func(uint16) uint32 { return 999 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToNillableUint32(t *testing.T) {
	Convey("Uint16.MapToNillableUint32", t, func() {
		So(NewUint16(4562).MapToNillableUint32(func(uint16) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).Get(), ShouldEqual, 999)

		So(NewEmptyUint16().MapToNillableUint32(func(uint16) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint16(4562).MapToNillableUint32(func(uint16) *uint32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToUint32(t *testing.T) {
	Convey("Uint16.MapToUint32", t, func() {
		test1, err1 := NewUint16(4562).MapToUint32(func(uint16) (uint32, error) {
			return 999, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 999)

		test2, err2 := NewEmptyUint16().MapToUint32(func(uint16) (uint32, error) {
			return 999, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint16(4562).MapToUint32(func(uint16) (uint32, error) {
			return 999, errors.New("uint32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint16_ForceMapToUint64(t *testing.T) {
	Convey("Uint16.ForceMapToUint64", t, func() {
		So(NewUint16(4562).ForceMapToUint64(func(uint16) uint64 { return 101010 }).Get(),
			ShouldEqual, 101010)

		So(NewEmptyUint16().ForceMapToUint64(func(uint16) uint64 { return 101010 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToNillableUint64(t *testing.T) {
	Convey("Uint16.MapToNillableUint64", t, func() {
		So(NewUint16(4562).MapToNillableUint64(func(uint16) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).Get(), ShouldEqual, 101010)

		So(NewEmptyUint16().MapToNillableUint64(func(uint16) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint16(4562).MapToNillableUint64(func(uint16) *uint64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToUint64(t *testing.T) {
	Convey("Uint16.MapToUint64", t, func() {
		test1, err1 := NewUint16(4562).MapToUint64(func(uint16) (uint64, error) {
			return 101010, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 101010)

		test2, err2 := NewEmptyUint16().MapToUint64(func(uint16) (uint64, error) {
			return 101010, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint16(4562).MapToUint64(func(uint16) (uint64, error) {
			return 101010, errors.New("uint64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint16_ForceMapToUint8(t *testing.T) {
	Convey("Uint16.ForceMapToUint8", t, func() {
		So(NewUint16(4562).ForceMapToUint8(func(uint16) uint8 { return 77 }).Get(),
			ShouldEqual, 77)

		So(NewEmptyUint16().ForceMapToUint8(func(uint16) uint8 { return 77 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToNillableUint8(t *testing.T) {
	Convey("Uint16.MapToNillableUint8", t, func() {
		So(NewUint16(4562).MapToNillableUint8(func(uint16) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).Get(), ShouldEqual, 77)

		So(NewEmptyUint16().MapToNillableUint8(func(uint16) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint16(4562).MapToNillableUint8(func(uint16) *uint8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint16_MapToUint8(t *testing.T) {
	Convey("Uint16.MapToUint8", t, func() {
		test1, err1 := NewUint16(4562).MapToUint8(func(uint16) (uint8, error) {
			return 77, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 77)

		test2, err2 := NewEmptyUint16().MapToUint8(func(uint16) (uint8, error) {
			return 77, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint16(4562).MapToUint8(func(uint16) (uint8, error) {
			return 77, errors.New("uint8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}
