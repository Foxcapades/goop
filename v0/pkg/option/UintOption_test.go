// Package option_test generated @ 2020-05-17 22:19:55.164878293 -0400 EDT m=+0.001439508
package option_test

import (
	"errors"
	"testing"

	. "github.com/Foxcapades/goop/v0/pkg/option"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewUint(t *testing.T) {
	Convey("NewUint", t, func() {
		test := NewUint(5898)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyUint(t *testing.T) {
	Convey("NewEmptyUint", t, func() {
		test := NewEmptyUint()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeUint(t *testing.T) {
	Convey("NewMaybeUint", t, func() {
		var val1 uint

		test1 := NewMaybeUint(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeUint(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestUint_Get(t *testing.T) {
	Convey("Uint.Get", t, func() {
		test1 := NewUint(5898)
		So(test1.Get(), ShouldEqual, 5898)

		test2 := NewEmptyUint()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestUint_Or(t *testing.T) {
	Convey("Uint.Or", t, func() {
		test1 := NewUint(5898)
		So(test1.Or(36985), ShouldEqual, 5898)

		test2 := NewEmptyUint()
		So(test2.Or(5898), ShouldEqual, 5898)
	})
}

func TestUint_OrGet(t *testing.T) {
	Convey("Uint.OrGet", t, func() {
		test1 := NewUint(5898)
		So(test1.OrGet(func() uint { return 36985 }), ShouldEqual, 5898)

		test2 := NewEmptyUint()
		So(test2.OrGet(func() uint { return 5898 }), ShouldEqual, 5898)
	})
}

func TestUint_OrPanicWith(t *testing.T) {
	Convey("Uint.OrPanicWith", t, func() {
		test1 := NewUint(5898)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 5898)

		test2 := NewEmptyUint()
		So(func() {
			test2.OrPanicWith(func() interface{} { return "test value" })
		}, ShouldPanicWith, "test value")
	})
}

func TestUint_MapToNillable(t *testing.T) {
	Convey("Uint.MapToNillable", t, func() {
		test1 := NewUint(5898)
		So(test1.MapToNillable(func(b uint) *uint { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewUint(5898)
		So(test2.MapToNillable(func(b uint) *uint { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyUint()
		So(func() {
			test3.MapToNillable(func(b uint) *uint { panic("foo") }).IsNil()
		}, ShouldNotPanic)
	})
}

func TestUint_ForceMapToBool(t *testing.T) {
	Convey("Uint.ForceMapToBool", t, func() {
		So(NewUint(5898).ForceMapToBool(func(uint) bool { return true }).Get(),
			ShouldEqual, true)

		So(NewEmptyUint().ForceMapToBool(func(uint) bool { return true }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToNillableBool(t *testing.T) {
	Convey("Uint.MapToNillableBool", t, func() {
		So(NewUint(5898).MapToNillableBool(func(uint) *bool {
			var tmp bool = true
			return &tmp
		}).Get(), ShouldEqual, true)

		So(NewEmptyUint().MapToNillableBool(func(uint) *bool {
			var tmp bool = true
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint(5898).MapToNillableBool(func(uint) *bool { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToBool(t *testing.T) {
	Convey("Uint.MapToBool", t, func() {
		test1, err1 := NewUint(5898).MapToBool(func(uint) (bool, error) {
			return true, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, true)

		test2, err2 := NewEmptyUint().MapToBool(func(uint) (bool, error) {
			return true, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint(5898).MapToBool(func(uint) (bool, error) {
			return true, errors.New("bool error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "bool error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint_ForceMapToByte(t *testing.T) {
	Convey("Uint.ForceMapToByte", t, func() {
		So(NewUint(5898).ForceMapToByte(func(uint) byte { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyUint().ForceMapToByte(func(uint) byte { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToNillableByte(t *testing.T) {
	Convey("Uint.MapToNillableByte", t, func() {
		So(NewUint(5898).MapToNillableByte(func(uint) *byte {
			var tmp byte = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyUint().MapToNillableByte(func(uint) *byte {
			var tmp byte = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint(5898).MapToNillableByte(func(uint) *byte { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToByte(t *testing.T) {
	Convey("Uint.MapToByte", t, func() {
		test1, err1 := NewUint(5898).MapToByte(func(uint) (byte, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyUint().MapToByte(func(uint) (byte, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint(5898).MapToByte(func(uint) (byte, error) {
			return 222, errors.New("byte error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "byte error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint_ForceMapToComplex128(t *testing.T) {
	Convey("Uint.ForceMapToComplex128", t, func() {
		So(NewUint(5898).ForceMapToComplex128(func(uint) complex128 { return 987654 }).Get(),
			ShouldEqual, 987654)

		So(NewEmptyUint().ForceMapToComplex128(func(uint) complex128 { return 987654 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToNillableComplex128(t *testing.T) {
	Convey("Uint.MapToNillableComplex128", t, func() {
		So(NewUint(5898).MapToNillableComplex128(func(uint) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).Get(), ShouldEqual, 987654)

		So(NewEmptyUint().MapToNillableComplex128(func(uint) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint(5898).MapToNillableComplex128(func(uint) *complex128 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToComplex128(t *testing.T) {
	Convey("Uint.MapToComplex128", t, func() {
		test1, err1 := NewUint(5898).MapToComplex128(func(uint) (complex128, error) {
			return 987654, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 987654)

		test2, err2 := NewEmptyUint().MapToComplex128(func(uint) (complex128, error) {
			return 987654, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint(5898).MapToComplex128(func(uint) (complex128, error) {
			return 987654, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint_ForceMapToComplex64(t *testing.T) {
	Convey("Uint.ForceMapToComplex64", t, func() {
		So(NewUint(5898).ForceMapToComplex64(func(uint) complex64 { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyUint().ForceMapToComplex64(func(uint) complex64 { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToNillableComplex64(t *testing.T) {
	Convey("Uint.MapToNillableComplex64", t, func() {
		So(NewUint(5898).MapToNillableComplex64(func(uint) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyUint().MapToNillableComplex64(func(uint) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint(5898).MapToNillableComplex64(func(uint) *complex64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToComplex64(t *testing.T) {
	Convey("Uint.MapToComplex64", t, func() {
		test1, err1 := NewUint(5898).MapToComplex64(func(uint) (complex64, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyUint().MapToComplex64(func(uint) (complex64, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint(5898).MapToComplex64(func(uint) (complex64, error) {
			return 123456, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint_ForceMapToFloat32(t *testing.T) {
	Convey("Uint.ForceMapToFloat32", t, func() {
		So(NewUint(5898).ForceMapToFloat32(func(uint) float32 { return 32.3 }).Get(),
			ShouldEqual, 32.3)

		So(NewEmptyUint().ForceMapToFloat32(func(uint) float32 { return 32.3 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToNillableFloat32(t *testing.T) {
	Convey("Uint.MapToNillableFloat32", t, func() {
		So(NewUint(5898).MapToNillableFloat32(func(uint) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).Get(), ShouldEqual, 32.3)

		So(NewEmptyUint().MapToNillableFloat32(func(uint) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint(5898).MapToNillableFloat32(func(uint) *float32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToFloat32(t *testing.T) {
	Convey("Uint.MapToFloat32", t, func() {
		test1, err1 := NewUint(5898).MapToFloat32(func(uint) (float32, error) {
			return 32.3, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 32.3)

		test2, err2 := NewEmptyUint().MapToFloat32(func(uint) (float32, error) {
			return 32.3, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint(5898).MapToFloat32(func(uint) (float32, error) {
			return 32.3, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint_ForceMapToFloat64(t *testing.T) {
	Convey("Uint.ForceMapToFloat64", t, func() {
		So(NewUint(5898).ForceMapToFloat64(func(uint) float64 { return 3.2455 }).Get(),
			ShouldEqual, 3.2455)

		So(NewEmptyUint().ForceMapToFloat64(func(uint) float64 { return 3.2455 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToNillableFloat64(t *testing.T) {
	Convey("Uint.MapToNillableFloat64", t, func() {
		So(NewUint(5898).MapToNillableFloat64(func(uint) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).Get(), ShouldEqual, 3.2455)

		So(NewEmptyUint().MapToNillableFloat64(func(uint) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint(5898).MapToNillableFloat64(func(uint) *float64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToFloat64(t *testing.T) {
	Convey("Uint.MapToFloat64", t, func() {
		test1, err1 := NewUint(5898).MapToFloat64(func(uint) (float64, error) {
			return 3.2455, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 3.2455)

		test2, err2 := NewEmptyUint().MapToFloat64(func(uint) (float64, error) {
			return 3.2455, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint(5898).MapToFloat64(func(uint) (float64, error) {
			return 3.2455, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint_ForceMapToInt(t *testing.T) {
	Convey("Uint.ForceMapToInt", t, func() {
		So(NewUint(5898).ForceMapToInt(func(uint) int { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyUint().ForceMapToInt(func(uint) int { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToNillableInt(t *testing.T) {
	Convey("Uint.MapToNillableInt", t, func() {
		So(NewUint(5898).MapToNillableInt(func(uint) *int {
			var tmp int = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyUint().MapToNillableInt(func(uint) *int {
			var tmp int = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint(5898).MapToNillableInt(func(uint) *int { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToInt(t *testing.T) {
	Convey("Uint.MapToInt", t, func() {
		test1, err1 := NewUint(5898).MapToInt(func(uint) (int, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyUint().MapToInt(func(uint) (int, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint(5898).MapToInt(func(uint) (int, error) {
			return 123456, errors.New("int error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint_ForceMapToInt16(t *testing.T) {
	Convey("Uint.ForceMapToInt16", t, func() {
		So(NewUint(5898).ForceMapToInt16(func(uint) int16 { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyUint().ForceMapToInt16(func(uint) int16 { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToNillableInt16(t *testing.T) {
	Convey("Uint.MapToNillableInt16", t, func() {
		So(NewUint(5898).MapToNillableInt16(func(uint) *int16 {
			var tmp int16 = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyUint().MapToNillableInt16(func(uint) *int16 {
			var tmp int16 = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint(5898).MapToNillableInt16(func(uint) *int16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToInt16(t *testing.T) {
	Convey("Uint.MapToInt16", t, func() {
		test1, err1 := NewUint(5898).MapToInt16(func(uint) (int16, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyUint().MapToInt16(func(uint) (int16, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint(5898).MapToInt16(func(uint) (int16, error) {
			return 222, errors.New("int16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint_ForceMapToInt32(t *testing.T) {
	Convey("Uint.ForceMapToInt32", t, func() {
		So(NewUint(5898).ForceMapToInt32(func(uint) int32 { return 333 }).Get(),
			ShouldEqual, 333)

		So(NewEmptyUint().ForceMapToInt32(func(uint) int32 { return 333 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToNillableInt32(t *testing.T) {
	Convey("Uint.MapToNillableInt32", t, func() {
		So(NewUint(5898).MapToNillableInt32(func(uint) *int32 {
			var tmp int32 = 333
			return &tmp
		}).Get(), ShouldEqual, 333)

		So(NewEmptyUint().MapToNillableInt32(func(uint) *int32 {
			var tmp int32 = 333
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint(5898).MapToNillableInt32(func(uint) *int32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToInt32(t *testing.T) {
	Convey("Uint.MapToInt32", t, func() {
		test1, err1 := NewUint(5898).MapToInt32(func(uint) (int32, error) {
			return 333, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 333)

		test2, err2 := NewEmptyUint().MapToInt32(func(uint) (int32, error) {
			return 333, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint(5898).MapToInt32(func(uint) (int32, error) {
			return 333, errors.New("int32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint_ForceMapToInt64(t *testing.T) {
	Convey("Uint.ForceMapToInt64", t, func() {
		So(NewUint(5898).ForceMapToInt64(func(uint) int64 { return 444 }).Get(),
			ShouldEqual, 444)

		So(NewEmptyUint().ForceMapToInt64(func(uint) int64 { return 444 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToNillableInt64(t *testing.T) {
	Convey("Uint.MapToNillableInt64", t, func() {
		So(NewUint(5898).MapToNillableInt64(func(uint) *int64 {
			var tmp int64 = 444
			return &tmp
		}).Get(), ShouldEqual, 444)

		So(NewEmptyUint().MapToNillableInt64(func(uint) *int64 {
			var tmp int64 = 444
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint(5898).MapToNillableInt64(func(uint) *int64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToInt64(t *testing.T) {
	Convey("Uint.MapToInt64", t, func() {
		test1, err1 := NewUint(5898).MapToInt64(func(uint) (int64, error) {
			return 444, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 444)

		test2, err2 := NewEmptyUint().MapToInt64(func(uint) (int64, error) {
			return 444, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint(5898).MapToInt64(func(uint) (int64, error) {
			return 444, errors.New("int64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint_ForceMapToInt8(t *testing.T) {
	Convey("Uint.ForceMapToInt8", t, func() {
		So(NewUint(5898).ForceMapToInt8(func(uint) int8 { return 111 }).Get(),
			ShouldEqual, 111)

		So(NewEmptyUint().ForceMapToInt8(func(uint) int8 { return 111 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToNillableInt8(t *testing.T) {
	Convey("Uint.MapToNillableInt8", t, func() {
		So(NewUint(5898).MapToNillableInt8(func(uint) *int8 {
			var tmp int8 = 111
			return &tmp
		}).Get(), ShouldEqual, 111)

		So(NewEmptyUint().MapToNillableInt8(func(uint) *int8 {
			var tmp int8 = 111
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint(5898).MapToNillableInt8(func(uint) *int8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToInt8(t *testing.T) {
	Convey("Uint.MapToInt8", t, func() {
		test1, err1 := NewUint(5898).MapToInt8(func(uint) (int8, error) {
			return 111, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 111)

		test2, err2 := NewEmptyUint().MapToInt8(func(uint) (int8, error) {
			return 111, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint(5898).MapToInt8(func(uint) (int8, error) {
			return 111, errors.New("int8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint_ForceMapToUntyped(t *testing.T) {
	Convey("Uint.ForceMapToUntyped", t, func() {
		So(NewUint(5898).ForceMapToUntyped(func(uint) interface{} { return "111111" }).Get(),
			ShouldEqual, "111111")

		So(NewEmptyUint().ForceMapToUntyped(func(uint) interface{} { return "111111" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToNillableUntyped(t *testing.T) {
	Convey("Uint.MapToNillableUntyped", t, func() {
		So(NewUint(5898).MapToNillableUntyped(func(uint) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).Get(), ShouldEqual, "111111")

		So(NewEmptyUint().MapToNillableUntyped(func(uint) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint(5898).MapToNillableUntyped(func(uint) *interface{} { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToUntyped(t *testing.T) {
	Convey("Uint.MapToUntyped", t, func() {
		test1, err1 := NewUint(5898).MapToUntyped(func(uint) (interface{}, error) {
			return "111111", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "111111")

		test2, err2 := NewEmptyUint().MapToUntyped(func(uint) (interface{}, error) {
			return "111111", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint(5898).MapToUntyped(func(uint) (interface{}, error) {
			return "111111", errors.New("interface error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "interface error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint_ForceMapToRune(t *testing.T) {
	Convey("Uint.ForceMapToRune", t, func() {
		So(NewUint(5898).ForceMapToRune(func(uint) rune { return 555 }).Get(),
			ShouldEqual, 555)

		So(NewEmptyUint().ForceMapToRune(func(uint) rune { return 555 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToNillableRune(t *testing.T) {
	Convey("Uint.MapToNillableRune", t, func() {
		So(NewUint(5898).MapToNillableRune(func(uint) *rune {
			var tmp rune = 555
			return &tmp
		}).Get(), ShouldEqual, 555)

		So(NewEmptyUint().MapToNillableRune(func(uint) *rune {
			var tmp rune = 555
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint(5898).MapToNillableRune(func(uint) *rune { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToRune(t *testing.T) {
	Convey("Uint.MapToRune", t, func() {
		test1, err1 := NewUint(5898).MapToRune(func(uint) (rune, error) {
			return 555, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 555)

		test2, err2 := NewEmptyUint().MapToRune(func(uint) (rune, error) {
			return 555, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint(5898).MapToRune(func(uint) (rune, error) {
			return 555, errors.New("rune error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "rune error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint_ForceMapToString(t *testing.T) {
	Convey("Uint.ForceMapToString", t, func() {
		So(NewUint(5898).ForceMapToString(func(uint) string { return "bonanza" }).Get(),
			ShouldEqual, "bonanza")

		So(NewEmptyUint().ForceMapToString(func(uint) string { return "bonanza" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToNillableString(t *testing.T) {
	Convey("Uint.MapToNillableString", t, func() {
		So(NewUint(5898).MapToNillableString(func(uint) *string {
			var tmp string = "bonanza"
			return &tmp
		}).Get(), ShouldEqual, "bonanza")

		So(NewEmptyUint().MapToNillableString(func(uint) *string {
			var tmp string = "bonanza"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint(5898).MapToNillableString(func(uint) *string { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToString(t *testing.T) {
	Convey("Uint.MapToString", t, func() {
		test1, err1 := NewUint(5898).MapToString(func(uint) (string, error) {
			return "bonanza", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "bonanza")

		test2, err2 := NewEmptyUint().MapToString(func(uint) (string, error) {
			return "bonanza", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint(5898).MapToString(func(uint) (string, error) {
			return "bonanza", errors.New("string error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "string error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint_ForceMapToUint16(t *testing.T) {
	Convey("Uint.ForceMapToUint16", t, func() {
		So(NewUint(5898).ForceMapToUint16(func(uint) uint16 { return 888 }).Get(),
			ShouldEqual, 888)

		So(NewEmptyUint().ForceMapToUint16(func(uint) uint16 { return 888 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToNillableUint16(t *testing.T) {
	Convey("Uint.MapToNillableUint16", t, func() {
		So(NewUint(5898).MapToNillableUint16(func(uint) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).Get(), ShouldEqual, 888)

		So(NewEmptyUint().MapToNillableUint16(func(uint) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint(5898).MapToNillableUint16(func(uint) *uint16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToUint16(t *testing.T) {
	Convey("Uint.MapToUint16", t, func() {
		test1, err1 := NewUint(5898).MapToUint16(func(uint) (uint16, error) {
			return 888, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 888)

		test2, err2 := NewEmptyUint().MapToUint16(func(uint) (uint16, error) {
			return 888, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint(5898).MapToUint16(func(uint) (uint16, error) {
			return 888, errors.New("uint16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint_ForceMapToUint32(t *testing.T) {
	Convey("Uint.ForceMapToUint32", t, func() {
		So(NewUint(5898).ForceMapToUint32(func(uint) uint32 { return 999 }).Get(),
			ShouldEqual, 999)

		So(NewEmptyUint().ForceMapToUint32(func(uint) uint32 { return 999 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToNillableUint32(t *testing.T) {
	Convey("Uint.MapToNillableUint32", t, func() {
		So(NewUint(5898).MapToNillableUint32(func(uint) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).Get(), ShouldEqual, 999)

		So(NewEmptyUint().MapToNillableUint32(func(uint) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint(5898).MapToNillableUint32(func(uint) *uint32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToUint32(t *testing.T) {
	Convey("Uint.MapToUint32", t, func() {
		test1, err1 := NewUint(5898).MapToUint32(func(uint) (uint32, error) {
			return 999, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 999)

		test2, err2 := NewEmptyUint().MapToUint32(func(uint) (uint32, error) {
			return 999, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint(5898).MapToUint32(func(uint) (uint32, error) {
			return 999, errors.New("uint32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint_ForceMapToUint64(t *testing.T) {
	Convey("Uint.ForceMapToUint64", t, func() {
		So(NewUint(5898).ForceMapToUint64(func(uint) uint64 { return 101010 }).Get(),
			ShouldEqual, 101010)

		So(NewEmptyUint().ForceMapToUint64(func(uint) uint64 { return 101010 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToNillableUint64(t *testing.T) {
	Convey("Uint.MapToNillableUint64", t, func() {
		So(NewUint(5898).MapToNillableUint64(func(uint) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).Get(), ShouldEqual, 101010)

		So(NewEmptyUint().MapToNillableUint64(func(uint) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint(5898).MapToNillableUint64(func(uint) *uint64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToUint64(t *testing.T) {
	Convey("Uint.MapToUint64", t, func() {
		test1, err1 := NewUint(5898).MapToUint64(func(uint) (uint64, error) {
			return 101010, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 101010)

		test2, err2 := NewEmptyUint().MapToUint64(func(uint) (uint64, error) {
			return 101010, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint(5898).MapToUint64(func(uint) (uint64, error) {
			return 101010, errors.New("uint64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUint_ForceMapToUint8(t *testing.T) {
	Convey("Uint.ForceMapToUint8", t, func() {
		So(NewUint(5898).ForceMapToUint8(func(uint) uint8 { return 77 }).Get(),
			ShouldEqual, 77)

		So(NewEmptyUint().ForceMapToUint8(func(uint) uint8 { return 77 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToNillableUint8(t *testing.T) {
	Convey("Uint.MapToNillableUint8", t, func() {
		So(NewUint(5898).MapToNillableUint8(func(uint) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).Get(), ShouldEqual, 77)

		So(NewEmptyUint().MapToNillableUint8(func(uint) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUint(5898).MapToNillableUint8(func(uint) *uint8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUint_MapToUint8(t *testing.T) {
	Convey("Uint.MapToUint8", t, func() {
		test1, err1 := NewUint(5898).MapToUint8(func(uint) (uint8, error) {
			return 77, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 77)

		test2, err2 := NewEmptyUint().MapToUint8(func(uint) (uint8, error) {
			return 77, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUint(5898).MapToUint8(func(uint) (uint8, error) {
			return 77, errors.New("uint8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}
