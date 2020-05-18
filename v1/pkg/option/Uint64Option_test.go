// Package option_test generated @ 2020-05-17 22:19:55.164878293 -0400 EDT m=+0.001439508
package option_test

import (
	"errors"
	"testing"

	. "github.com/Foxcapades/goop/v1/pkg/option"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewUint64(t *testing.T) {
	Convey("NewUint64", t, func() {
		test := NewUint64(456789)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyUint64(t *testing.T) {
	Convey("NewEmptyUint64", t, func() {
		test := NewEmptyUint64()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeUint64(t *testing.T) {
	Convey("NewMaybeUint64", t, func() {
		var val1 uint64

		test1 := NewMaybeUint64(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeUint64(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestUint64_Get(t *testing.T) {
	Convey("Uint64.Get", t, func() {
		test1 := NewUint64(456789)
		So(test1.Get(), ShouldEqual, 456789)

		test2 := NewEmptyUint64()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestUint64_Or(t *testing.T) {
	Convey("Uint64.Or", t, func() {
		test1 := NewUint64(456789)
		So(test1.Or(9874654), ShouldEqual, 456789)

		test2 := NewEmptyUint64()
		So(test2.Or(456789), ShouldEqual, 456789)
	})
}

func TestUint64_OrGet(t *testing.T) {
	Convey("Uint64.OrGet", t, func() {
		test1 := NewUint64(456789)
		So(test1.OrGet(func() uint64 { return 9874654 }), ShouldEqual, 456789)

		test2 := NewEmptyUint64()
		So(test2.OrGet(func() uint64 { return 456789 }), ShouldEqual, 456789)
	})
}

func TestUint64_OrPanicWith(t *testing.T) {
	Convey("Uint64.OrPanicWith", t, func() {
		test1 := NewUint64(456789)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 456789)

		test2 := NewEmptyUint64()
		So(func() {
			test2.OrPanicWith(func() interface{} { return "test value" })
		}, ShouldPanicWith, "test value")
	})
}

func TestUint64_MapToNillable(t *testing.T) {
	Convey("Uint64.MapToNillable", t, func() {
		test1 := NewUint64(456789)
		So(test1.MapToNillable(func(b uint64) *uint64 { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewUint64(456789)
		So(test2.MapToNillable(func(b uint64) *uint64 { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyUint64()
		So(func() {
			test3.MapToNillable(func(b uint64) *uint64 { panic("foo") }).IsNil()
		}, ShouldNotPanic)
	})
}

func TestUint64_ForceMapToBool(t *testing.T) {
	Convey("Uint64.ForceMapToBool", t, func() {
		So(NewUint64(456789).ForceMapToBool(func(uint64) bool { return true }).Get(),
			ShouldEqual, true)

		So(NewEmptyUint64().ForceMapToBool(func(uint64) bool { return true }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToNillableBool(t *testing.T) {
	Convey("Uint64.MapToNillableBool", t, func() {
		So(NewUint64(456789).MapToNillableBool(func(uint64) *bool {
			var tmp bool = true
			return &tmp
		}).Get(), ShouldEqual, true)

		So(NewEmptyUint64().MapToNillableBool(func(uint64) *bool {
			var tmp bool = true
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint64(456789).MapToNillableBool(func(uint64) *bool { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToBool(t *testing.T) {
	Convey("Uint64.MapToBool", t, func() {
		test1, err1 := NewUint64(456789).MapToBool(func(uint64) (bool, error) {
			return true, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, true)

		test2, err2 := NewEmptyUint64().MapToBool(func(uint64) (bool, error) {
			return true, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint64(456789).MapToBool(func(uint64) (bool, error) {
			return true, errors.New("bool error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "bool error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint64_ForceMapToByte(t *testing.T) {
	Convey("Uint64.ForceMapToByte", t, func() {
		So(NewUint64(456789).ForceMapToByte(func(uint64) byte { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyUint64().ForceMapToByte(func(uint64) byte { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToNillableByte(t *testing.T) {
	Convey("Uint64.MapToNillableByte", t, func() {
		So(NewUint64(456789).MapToNillableByte(func(uint64) *byte {
			var tmp byte = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyUint64().MapToNillableByte(func(uint64) *byte {
			var tmp byte = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint64(456789).MapToNillableByte(func(uint64) *byte { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToByte(t *testing.T) {
	Convey("Uint64.MapToByte", t, func() {
		test1, err1 := NewUint64(456789).MapToByte(func(uint64) (byte, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyUint64().MapToByte(func(uint64) (byte, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint64(456789).MapToByte(func(uint64) (byte, error) {
			return 222, errors.New("byte error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "byte error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint64_ForceMapToComplex128(t *testing.T) {
	Convey("Uint64.ForceMapToComplex128", t, func() {
		So(NewUint64(456789).ForceMapToComplex128(func(uint64) complex128 { return 987654 }).Get(),
			ShouldEqual, 987654)

		So(NewEmptyUint64().ForceMapToComplex128(func(uint64) complex128 { return 987654 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToNillableComplex128(t *testing.T) {
	Convey("Uint64.MapToNillableComplex128", t, func() {
		So(NewUint64(456789).MapToNillableComplex128(func(uint64) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).Get(), ShouldEqual, 987654)

		So(NewEmptyUint64().MapToNillableComplex128(func(uint64) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint64(456789).MapToNillableComplex128(func(uint64) *complex128 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToComplex128(t *testing.T) {
	Convey("Uint64.MapToComplex128", t, func() {
		test1, err1 := NewUint64(456789).MapToComplex128(func(uint64) (complex128, error) {
			return 987654, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 987654)

		test2, err2 := NewEmptyUint64().MapToComplex128(func(uint64) (complex128, error) {
			return 987654, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint64(456789).MapToComplex128(func(uint64) (complex128, error) {
			return 987654, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint64_ForceMapToComplex64(t *testing.T) {
	Convey("Uint64.ForceMapToComplex64", t, func() {
		So(NewUint64(456789).ForceMapToComplex64(func(uint64) complex64 { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyUint64().ForceMapToComplex64(func(uint64) complex64 { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToNillableComplex64(t *testing.T) {
	Convey("Uint64.MapToNillableComplex64", t, func() {
		So(NewUint64(456789).MapToNillableComplex64(func(uint64) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyUint64().MapToNillableComplex64(func(uint64) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint64(456789).MapToNillableComplex64(func(uint64) *complex64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToComplex64(t *testing.T) {
	Convey("Uint64.MapToComplex64", t, func() {
		test1, err1 := NewUint64(456789).MapToComplex64(func(uint64) (complex64, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyUint64().MapToComplex64(func(uint64) (complex64, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint64(456789).MapToComplex64(func(uint64) (complex64, error) {
			return 123456, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint64_ForceMapToFloat32(t *testing.T) {
	Convey("Uint64.ForceMapToFloat32", t, func() {
		So(NewUint64(456789).ForceMapToFloat32(func(uint64) float32 { return 32.3 }).Get(),
			ShouldEqual, 32.3)

		So(NewEmptyUint64().ForceMapToFloat32(func(uint64) float32 { return 32.3 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToNillableFloat32(t *testing.T) {
	Convey("Uint64.MapToNillableFloat32", t, func() {
		So(NewUint64(456789).MapToNillableFloat32(func(uint64) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).Get(), ShouldEqual, 32.3)

		So(NewEmptyUint64().MapToNillableFloat32(func(uint64) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint64(456789).MapToNillableFloat32(func(uint64) *float32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToFloat32(t *testing.T) {
	Convey("Uint64.MapToFloat32", t, func() {
		test1, err1 := NewUint64(456789).MapToFloat32(func(uint64) (float32, error) {
			return 32.3, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 32.3)

		test2, err2 := NewEmptyUint64().MapToFloat32(func(uint64) (float32, error) {
			return 32.3, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint64(456789).MapToFloat32(func(uint64) (float32, error) {
			return 32.3, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint64_ForceMapToFloat64(t *testing.T) {
	Convey("Uint64.ForceMapToFloat64", t, func() {
		So(NewUint64(456789).ForceMapToFloat64(func(uint64) float64 { return 3.2455 }).Get(),
			ShouldEqual, 3.2455)

		So(NewEmptyUint64().ForceMapToFloat64(func(uint64) float64 { return 3.2455 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToNillableFloat64(t *testing.T) {
	Convey("Uint64.MapToNillableFloat64", t, func() {
		So(NewUint64(456789).MapToNillableFloat64(func(uint64) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).Get(), ShouldEqual, 3.2455)

		So(NewEmptyUint64().MapToNillableFloat64(func(uint64) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint64(456789).MapToNillableFloat64(func(uint64) *float64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToFloat64(t *testing.T) {
	Convey("Uint64.MapToFloat64", t, func() {
		test1, err1 := NewUint64(456789).MapToFloat64(func(uint64) (float64, error) {
			return 3.2455, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 3.2455)

		test2, err2 := NewEmptyUint64().MapToFloat64(func(uint64) (float64, error) {
			return 3.2455, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint64(456789).MapToFloat64(func(uint64) (float64, error) {
			return 3.2455, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint64_ForceMapToInt(t *testing.T) {
	Convey("Uint64.ForceMapToInt", t, func() {
		So(NewUint64(456789).ForceMapToInt(func(uint64) int { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyUint64().ForceMapToInt(func(uint64) int { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToNillableInt(t *testing.T) {
	Convey("Uint64.MapToNillableInt", t, func() {
		So(NewUint64(456789).MapToNillableInt(func(uint64) *int {
			var tmp int = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyUint64().MapToNillableInt(func(uint64) *int {
			var tmp int = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint64(456789).MapToNillableInt(func(uint64) *int { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToInt(t *testing.T) {
	Convey("Uint64.MapToInt", t, func() {
		test1, err1 := NewUint64(456789).MapToInt(func(uint64) (int, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyUint64().MapToInt(func(uint64) (int, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint64(456789).MapToInt(func(uint64) (int, error) {
			return 123456, errors.New("int error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint64_ForceMapToInt16(t *testing.T) {
	Convey("Uint64.ForceMapToInt16", t, func() {
		So(NewUint64(456789).ForceMapToInt16(func(uint64) int16 { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyUint64().ForceMapToInt16(func(uint64) int16 { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToNillableInt16(t *testing.T) {
	Convey("Uint64.MapToNillableInt16", t, func() {
		So(NewUint64(456789).MapToNillableInt16(func(uint64) *int16 {
			var tmp int16 = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyUint64().MapToNillableInt16(func(uint64) *int16 {
			var tmp int16 = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint64(456789).MapToNillableInt16(func(uint64) *int16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToInt16(t *testing.T) {
	Convey("Uint64.MapToInt16", t, func() {
		test1, err1 := NewUint64(456789).MapToInt16(func(uint64) (int16, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyUint64().MapToInt16(func(uint64) (int16, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint64(456789).MapToInt16(func(uint64) (int16, error) {
			return 222, errors.New("int16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint64_ForceMapToInt32(t *testing.T) {
	Convey("Uint64.ForceMapToInt32", t, func() {
		So(NewUint64(456789).ForceMapToInt32(func(uint64) int32 { return 333 }).Get(),
			ShouldEqual, 333)

		So(NewEmptyUint64().ForceMapToInt32(func(uint64) int32 { return 333 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToNillableInt32(t *testing.T) {
	Convey("Uint64.MapToNillableInt32", t, func() {
		So(NewUint64(456789).MapToNillableInt32(func(uint64) *int32 {
			var tmp int32 = 333
			return &tmp
		}).Get(), ShouldEqual, 333)

		So(NewEmptyUint64().MapToNillableInt32(func(uint64) *int32 {
			var tmp int32 = 333
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint64(456789).MapToNillableInt32(func(uint64) *int32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToInt32(t *testing.T) {
	Convey("Uint64.MapToInt32", t, func() {
		test1, err1 := NewUint64(456789).MapToInt32(func(uint64) (int32, error) {
			return 333, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 333)

		test2, err2 := NewEmptyUint64().MapToInt32(func(uint64) (int32, error) {
			return 333, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint64(456789).MapToInt32(func(uint64) (int32, error) {
			return 333, errors.New("int32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint64_ForceMapToInt64(t *testing.T) {
	Convey("Uint64.ForceMapToInt64", t, func() {
		So(NewUint64(456789).ForceMapToInt64(func(uint64) int64 { return 444 }).Get(),
			ShouldEqual, 444)

		So(NewEmptyUint64().ForceMapToInt64(func(uint64) int64 { return 444 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToNillableInt64(t *testing.T) {
	Convey("Uint64.MapToNillableInt64", t, func() {
		So(NewUint64(456789).MapToNillableInt64(func(uint64) *int64 {
			var tmp int64 = 444
			return &tmp
		}).Get(), ShouldEqual, 444)

		So(NewEmptyUint64().MapToNillableInt64(func(uint64) *int64 {
			var tmp int64 = 444
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint64(456789).MapToNillableInt64(func(uint64) *int64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToInt64(t *testing.T) {
	Convey("Uint64.MapToInt64", t, func() {
		test1, err1 := NewUint64(456789).MapToInt64(func(uint64) (int64, error) {
			return 444, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 444)

		test2, err2 := NewEmptyUint64().MapToInt64(func(uint64) (int64, error) {
			return 444, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint64(456789).MapToInt64(func(uint64) (int64, error) {
			return 444, errors.New("int64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint64_ForceMapToInt8(t *testing.T) {
	Convey("Uint64.ForceMapToInt8", t, func() {
		So(NewUint64(456789).ForceMapToInt8(func(uint64) int8 { return 111 }).Get(),
			ShouldEqual, 111)

		So(NewEmptyUint64().ForceMapToInt8(func(uint64) int8 { return 111 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToNillableInt8(t *testing.T) {
	Convey("Uint64.MapToNillableInt8", t, func() {
		So(NewUint64(456789).MapToNillableInt8(func(uint64) *int8 {
			var tmp int8 = 111
			return &tmp
		}).Get(), ShouldEqual, 111)

		So(NewEmptyUint64().MapToNillableInt8(func(uint64) *int8 {
			var tmp int8 = 111
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint64(456789).MapToNillableInt8(func(uint64) *int8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToInt8(t *testing.T) {
	Convey("Uint64.MapToInt8", t, func() {
		test1, err1 := NewUint64(456789).MapToInt8(func(uint64) (int8, error) {
			return 111, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 111)

		test2, err2 := NewEmptyUint64().MapToInt8(func(uint64) (int8, error) {
			return 111, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint64(456789).MapToInt8(func(uint64) (int8, error) {
			return 111, errors.New("int8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint64_ForceMapToUntyped(t *testing.T) {
	Convey("Uint64.ForceMapToUntyped", t, func() {
		So(NewUint64(456789).ForceMapToUntyped(func(uint64) interface{} { return "111111" }).Get(),
			ShouldEqual, "111111")

		So(NewEmptyUint64().ForceMapToUntyped(func(uint64) interface{} { return "111111" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToNillableUntyped(t *testing.T) {
	Convey("Uint64.MapToNillableUntyped", t, func() {
		So(NewUint64(456789).MapToNillableUntyped(func(uint64) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).Get(), ShouldEqual, "111111")

		So(NewEmptyUint64().MapToNillableUntyped(func(uint64) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint64(456789).MapToNillableUntyped(func(uint64) *interface{} { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToUntyped(t *testing.T) {
	Convey("Uint64.MapToUntyped", t, func() {
		test1, err1 := NewUint64(456789).MapToUntyped(func(uint64) (interface{}, error) {
			return "111111", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "111111")

		test2, err2 := NewEmptyUint64().MapToUntyped(func(uint64) (interface{}, error) {
			return "111111", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint64(456789).MapToUntyped(func(uint64) (interface{}, error) {
			return "111111", errors.New("interface error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "interface error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint64_ForceMapToRune(t *testing.T) {
	Convey("Uint64.ForceMapToRune", t, func() {
		So(NewUint64(456789).ForceMapToRune(func(uint64) rune { return 555 }).Get(),
			ShouldEqual, 555)

		So(NewEmptyUint64().ForceMapToRune(func(uint64) rune { return 555 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToNillableRune(t *testing.T) {
	Convey("Uint64.MapToNillableRune", t, func() {
		So(NewUint64(456789).MapToNillableRune(func(uint64) *rune {
			var tmp rune = 555
			return &tmp
		}).Get(), ShouldEqual, 555)

		So(NewEmptyUint64().MapToNillableRune(func(uint64) *rune {
			var tmp rune = 555
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint64(456789).MapToNillableRune(func(uint64) *rune { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToRune(t *testing.T) {
	Convey("Uint64.MapToRune", t, func() {
		test1, err1 := NewUint64(456789).MapToRune(func(uint64) (rune, error) {
			return 555, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 555)

		test2, err2 := NewEmptyUint64().MapToRune(func(uint64) (rune, error) {
			return 555, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint64(456789).MapToRune(func(uint64) (rune, error) {
			return 555, errors.New("rune error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "rune error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint64_ForceMapToString(t *testing.T) {
	Convey("Uint64.ForceMapToString", t, func() {
		So(NewUint64(456789).ForceMapToString(func(uint64) string { return "bonanza" }).Get(),
			ShouldEqual, "bonanza")

		So(NewEmptyUint64().ForceMapToString(func(uint64) string { return "bonanza" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToNillableString(t *testing.T) {
	Convey("Uint64.MapToNillableString", t, func() {
		So(NewUint64(456789).MapToNillableString(func(uint64) *string {
			var tmp string = "bonanza"
			return &tmp
		}).Get(), ShouldEqual, "bonanza")

		So(NewEmptyUint64().MapToNillableString(func(uint64) *string {
			var tmp string = "bonanza"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint64(456789).MapToNillableString(func(uint64) *string { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToString(t *testing.T) {
	Convey("Uint64.MapToString", t, func() {
		test1, err1 := NewUint64(456789).MapToString(func(uint64) (string, error) {
			return "bonanza", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "bonanza")

		test2, err2 := NewEmptyUint64().MapToString(func(uint64) (string, error) {
			return "bonanza", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint64(456789).MapToString(func(uint64) (string, error) {
			return "bonanza", errors.New("string error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "string error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint64_ForceMapToUint(t *testing.T) {
	Convey("Uint64.ForceMapToUint", t, func() {
		So(NewUint64(456789).ForceMapToUint(func(uint64) uint { return 666 }).Get(),
			ShouldEqual, 666)

		So(NewEmptyUint64().ForceMapToUint(func(uint64) uint { return 666 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToNillableUint(t *testing.T) {
	Convey("Uint64.MapToNillableUint", t, func() {
		So(NewUint64(456789).MapToNillableUint(func(uint64) *uint {
			var tmp uint = 666
			return &tmp
		}).Get(), ShouldEqual, 666)

		So(NewEmptyUint64().MapToNillableUint(func(uint64) *uint {
			var tmp uint = 666
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint64(456789).MapToNillableUint(func(uint64) *uint { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToUint(t *testing.T) {
	Convey("Uint64.MapToUint", t, func() {
		test1, err1 := NewUint64(456789).MapToUint(func(uint64) (uint, error) {
			return 666, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 666)

		test2, err2 := NewEmptyUint64().MapToUint(func(uint64) (uint, error) {
			return 666, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint64(456789).MapToUint(func(uint64) (uint, error) {
			return 666, errors.New("uint error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint64_ForceMapToUint16(t *testing.T) {
	Convey("Uint64.ForceMapToUint16", t, func() {
		So(NewUint64(456789).ForceMapToUint16(func(uint64) uint16 { return 888 }).Get(),
			ShouldEqual, 888)

		So(NewEmptyUint64().ForceMapToUint16(func(uint64) uint16 { return 888 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToNillableUint16(t *testing.T) {
	Convey("Uint64.MapToNillableUint16", t, func() {
		So(NewUint64(456789).MapToNillableUint16(func(uint64) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).Get(), ShouldEqual, 888)

		So(NewEmptyUint64().MapToNillableUint16(func(uint64) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint64(456789).MapToNillableUint16(func(uint64) *uint16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToUint16(t *testing.T) {
	Convey("Uint64.MapToUint16", t, func() {
		test1, err1 := NewUint64(456789).MapToUint16(func(uint64) (uint16, error) {
			return 888, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 888)

		test2, err2 := NewEmptyUint64().MapToUint16(func(uint64) (uint16, error) {
			return 888, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint64(456789).MapToUint16(func(uint64) (uint16, error) {
			return 888, errors.New("uint16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint64_ForceMapToUint32(t *testing.T) {
	Convey("Uint64.ForceMapToUint32", t, func() {
		So(NewUint64(456789).ForceMapToUint32(func(uint64) uint32 { return 999 }).Get(),
			ShouldEqual, 999)

		So(NewEmptyUint64().ForceMapToUint32(func(uint64) uint32 { return 999 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToNillableUint32(t *testing.T) {
	Convey("Uint64.MapToNillableUint32", t, func() {
		So(NewUint64(456789).MapToNillableUint32(func(uint64) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).Get(), ShouldEqual, 999)

		So(NewEmptyUint64().MapToNillableUint32(func(uint64) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint64(456789).MapToNillableUint32(func(uint64) *uint32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToUint32(t *testing.T) {
	Convey("Uint64.MapToUint32", t, func() {
		test1, err1 := NewUint64(456789).MapToUint32(func(uint64) (uint32, error) {
			return 999, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 999)

		test2, err2 := NewEmptyUint64().MapToUint32(func(uint64) (uint32, error) {
			return 999, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint64(456789).MapToUint32(func(uint64) (uint32, error) {
			return 999, errors.New("uint32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint64_ForceMapToUint8(t *testing.T) {
	Convey("Uint64.ForceMapToUint8", t, func() {
		So(NewUint64(456789).ForceMapToUint8(func(uint64) uint8 { return 77 }).Get(),
			ShouldEqual, 77)

		So(NewEmptyUint64().ForceMapToUint8(func(uint64) uint8 { return 77 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToNillableUint8(t *testing.T) {
	Convey("Uint64.MapToNillableUint8", t, func() {
		So(NewUint64(456789).MapToNillableUint8(func(uint64) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).Get(), ShouldEqual, 77)

		So(NewEmptyUint64().MapToNillableUint8(func(uint64) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint64(456789).MapToNillableUint8(func(uint64) *uint8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint64_MapToUint8(t *testing.T) {
	Convey("Uint64.MapToUint8", t, func() {
		test1, err1 := NewUint64(456789).MapToUint8(func(uint64) (uint8, error) {
			return 77, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 77)

		test2, err2 := NewEmptyUint64().MapToUint8(func(uint64) (uint8, error) {
			return 77, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint64(456789).MapToUint8(func(uint64) (uint8, error) {
			return 77, errors.New("uint8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}
