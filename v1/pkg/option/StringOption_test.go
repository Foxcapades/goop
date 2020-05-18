// Package option_test generated @ 2020-05-17 22:19:55.164878293 -0400 EDT m=+0.001439508
package option_test

import (
	"errors"
	"testing"

	. "github.com/Foxcapades/goop/v1/pkg/option"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewString(t *testing.T) {
	Convey("NewString", t, func() {
		test := NewString("hi")
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyString(t *testing.T) {
	Convey("NewEmptyString", t, func() {
		test := NewEmptyString()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeString(t *testing.T) {
	Convey("NewMaybeString", t, func() {
		var val1 string

		test1 := NewMaybeString(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeString(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestString_Get(t *testing.T) {
	Convey("String.Get", t, func() {
		test1 := NewString("hi")
		So(test1.Get(), ShouldEqual, "hi")

		test2 := NewEmptyString()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestString_Or(t *testing.T) {
	Convey("String.Or", t, func() {
		test1 := NewString("hi")
		So(test1.Or("bye"), ShouldEqual, "hi")

		test2 := NewEmptyString()
		So(test2.Or("hi"), ShouldEqual, "hi")
	})
}

func TestString_OrGet(t *testing.T) {
	Convey("String.OrGet", t, func() {
		test1 := NewString("hi")
		So(test1.OrGet(func() string { return "bye" }), ShouldEqual, "hi")

		test2 := NewEmptyString()
		So(test2.OrGet(func() string { return "hi" }), ShouldEqual, "hi")
	})
}

func TestString_OrPanicWith(t *testing.T) {
	Convey("String.OrPanicWith", t, func() {
		test1 := NewString("hi")
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, "hi")

		test2 := NewEmptyString()
		So(func() {
			test2.OrPanicWith(func() interface{} { return "test value" })
		}, ShouldPanicWith, "test value")
	})
}

func TestString_MapToNillable(t *testing.T) {
	Convey("String.MapToNillable", t, func() {
		test1 := NewString("hi")
		So(test1.MapToNillable(func(b string) *string { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewString("hi")
		So(test2.MapToNillable(func(b string) *string { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyString()
		So(func() {
			test3.MapToNillable(func(b string) *string { panic("foo") }).IsNil()
		}, ShouldNotPanic)
	})
}

func TestString_ForceMapToBool(t *testing.T) {
	Convey("String.ForceMapToBool", t, func() {
		So(NewString("hi").ForceMapToBool(func(string) bool { return true }).Get(),
			ShouldEqual, true)

		So(NewEmptyString().ForceMapToBool(func(string) bool { return true }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToNillableBool(t *testing.T) {
	Convey("String.MapToNillableBool", t, func() {
		So(NewString("hi").MapToNillableBool(func(string) *bool {
			var tmp bool = true
			return &tmp
		}).Get(), ShouldEqual, true)

		So(NewEmptyString().MapToNillableBool(func(string) *bool {
			var tmp bool = true
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewString("hi").MapToNillableBool(func(string) *bool { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToBool(t *testing.T) {
	Convey("String.MapToBool", t, func() {
		test1, err1 := NewString("hi").MapToBool(func(string) (bool, error) {
			return true, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, true)

		test2, err2 := NewEmptyString().MapToBool(func(string) (bool, error) {
			return true, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewString("hi").MapToBool(func(string) (bool, error) {
			return true, errors.New("bool error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "bool error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestString_ForceMapToByte(t *testing.T) {
	Convey("String.ForceMapToByte", t, func() {
		So(NewString("hi").ForceMapToByte(func(string) byte { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyString().ForceMapToByte(func(string) byte { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToNillableByte(t *testing.T) {
	Convey("String.MapToNillableByte", t, func() {
		So(NewString("hi").MapToNillableByte(func(string) *byte {
			var tmp byte = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyString().MapToNillableByte(func(string) *byte {
			var tmp byte = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewString("hi").MapToNillableByte(func(string) *byte { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToByte(t *testing.T) {
	Convey("String.MapToByte", t, func() {
		test1, err1 := NewString("hi").MapToByte(func(string) (byte, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyString().MapToByte(func(string) (byte, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewString("hi").MapToByte(func(string) (byte, error) {
			return 222, errors.New("byte error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "byte error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestString_ForceMapToComplex128(t *testing.T) {
	Convey("String.ForceMapToComplex128", t, func() {
		So(NewString("hi").ForceMapToComplex128(func(string) complex128 { return 987654 }).Get(),
			ShouldEqual, 987654)

		So(NewEmptyString().ForceMapToComplex128(func(string) complex128 { return 987654 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToNillableComplex128(t *testing.T) {
	Convey("String.MapToNillableComplex128", t, func() {
		So(NewString("hi").MapToNillableComplex128(func(string) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).Get(), ShouldEqual, 987654)

		So(NewEmptyString().MapToNillableComplex128(func(string) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewString("hi").MapToNillableComplex128(func(string) *complex128 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToComplex128(t *testing.T) {
	Convey("String.MapToComplex128", t, func() {
		test1, err1 := NewString("hi").MapToComplex128(func(string) (complex128, error) {
			return 987654, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 987654)

		test2, err2 := NewEmptyString().MapToComplex128(func(string) (complex128, error) {
			return 987654, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewString("hi").MapToComplex128(func(string) (complex128, error) {
			return 987654, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestString_ForceMapToComplex64(t *testing.T) {
	Convey("String.ForceMapToComplex64", t, func() {
		So(NewString("hi").ForceMapToComplex64(func(string) complex64 { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyString().ForceMapToComplex64(func(string) complex64 { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToNillableComplex64(t *testing.T) {
	Convey("String.MapToNillableComplex64", t, func() {
		So(NewString("hi").MapToNillableComplex64(func(string) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyString().MapToNillableComplex64(func(string) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewString("hi").MapToNillableComplex64(func(string) *complex64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToComplex64(t *testing.T) {
	Convey("String.MapToComplex64", t, func() {
		test1, err1 := NewString("hi").MapToComplex64(func(string) (complex64, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyString().MapToComplex64(func(string) (complex64, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewString("hi").MapToComplex64(func(string) (complex64, error) {
			return 123456, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestString_ForceMapToFloat32(t *testing.T) {
	Convey("String.ForceMapToFloat32", t, func() {
		So(NewString("hi").ForceMapToFloat32(func(string) float32 { return 32.3 }).Get(),
			ShouldEqual, 32.3)

		So(NewEmptyString().ForceMapToFloat32(func(string) float32 { return 32.3 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToNillableFloat32(t *testing.T) {
	Convey("String.MapToNillableFloat32", t, func() {
		So(NewString("hi").MapToNillableFloat32(func(string) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).Get(), ShouldEqual, 32.3)

		So(NewEmptyString().MapToNillableFloat32(func(string) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewString("hi").MapToNillableFloat32(func(string) *float32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToFloat32(t *testing.T) {
	Convey("String.MapToFloat32", t, func() {
		test1, err1 := NewString("hi").MapToFloat32(func(string) (float32, error) {
			return 32.3, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 32.3)

		test2, err2 := NewEmptyString().MapToFloat32(func(string) (float32, error) {
			return 32.3, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewString("hi").MapToFloat32(func(string) (float32, error) {
			return 32.3, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestString_ForceMapToFloat64(t *testing.T) {
	Convey("String.ForceMapToFloat64", t, func() {
		So(NewString("hi").ForceMapToFloat64(func(string) float64 { return 3.2455 }).Get(),
			ShouldEqual, 3.2455)

		So(NewEmptyString().ForceMapToFloat64(func(string) float64 { return 3.2455 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToNillableFloat64(t *testing.T) {
	Convey("String.MapToNillableFloat64", t, func() {
		So(NewString("hi").MapToNillableFloat64(func(string) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).Get(), ShouldEqual, 3.2455)

		So(NewEmptyString().MapToNillableFloat64(func(string) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewString("hi").MapToNillableFloat64(func(string) *float64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToFloat64(t *testing.T) {
	Convey("String.MapToFloat64", t, func() {
		test1, err1 := NewString("hi").MapToFloat64(func(string) (float64, error) {
			return 3.2455, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 3.2455)

		test2, err2 := NewEmptyString().MapToFloat64(func(string) (float64, error) {
			return 3.2455, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewString("hi").MapToFloat64(func(string) (float64, error) {
			return 3.2455, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestString_ForceMapToInt(t *testing.T) {
	Convey("String.ForceMapToInt", t, func() {
		So(NewString("hi").ForceMapToInt(func(string) int { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyString().ForceMapToInt(func(string) int { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToNillableInt(t *testing.T) {
	Convey("String.MapToNillableInt", t, func() {
		So(NewString("hi").MapToNillableInt(func(string) *int {
			var tmp int = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyString().MapToNillableInt(func(string) *int {
			var tmp int = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewString("hi").MapToNillableInt(func(string) *int { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToInt(t *testing.T) {
	Convey("String.MapToInt", t, func() {
		test1, err1 := NewString("hi").MapToInt(func(string) (int, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyString().MapToInt(func(string) (int, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewString("hi").MapToInt(func(string) (int, error) {
			return 123456, errors.New("int error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestString_ForceMapToInt16(t *testing.T) {
	Convey("String.ForceMapToInt16", t, func() {
		So(NewString("hi").ForceMapToInt16(func(string) int16 { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyString().ForceMapToInt16(func(string) int16 { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToNillableInt16(t *testing.T) {
	Convey("String.MapToNillableInt16", t, func() {
		So(NewString("hi").MapToNillableInt16(func(string) *int16 {
			var tmp int16 = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyString().MapToNillableInt16(func(string) *int16 {
			var tmp int16 = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewString("hi").MapToNillableInt16(func(string) *int16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToInt16(t *testing.T) {
	Convey("String.MapToInt16", t, func() {
		test1, err1 := NewString("hi").MapToInt16(func(string) (int16, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyString().MapToInt16(func(string) (int16, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewString("hi").MapToInt16(func(string) (int16, error) {
			return 222, errors.New("int16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestString_ForceMapToInt32(t *testing.T) {
	Convey("String.ForceMapToInt32", t, func() {
		So(NewString("hi").ForceMapToInt32(func(string) int32 { return 333 }).Get(),
			ShouldEqual, 333)

		So(NewEmptyString().ForceMapToInt32(func(string) int32 { return 333 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToNillableInt32(t *testing.T) {
	Convey("String.MapToNillableInt32", t, func() {
		So(NewString("hi").MapToNillableInt32(func(string) *int32 {
			var tmp int32 = 333
			return &tmp
		}).Get(), ShouldEqual, 333)

		So(NewEmptyString().MapToNillableInt32(func(string) *int32 {
			var tmp int32 = 333
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewString("hi").MapToNillableInt32(func(string) *int32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToInt32(t *testing.T) {
	Convey("String.MapToInt32", t, func() {
		test1, err1 := NewString("hi").MapToInt32(func(string) (int32, error) {
			return 333, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 333)

		test2, err2 := NewEmptyString().MapToInt32(func(string) (int32, error) {
			return 333, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewString("hi").MapToInt32(func(string) (int32, error) {
			return 333, errors.New("int32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestString_ForceMapToInt64(t *testing.T) {
	Convey("String.ForceMapToInt64", t, func() {
		So(NewString("hi").ForceMapToInt64(func(string) int64 { return 444 }).Get(),
			ShouldEqual, 444)

		So(NewEmptyString().ForceMapToInt64(func(string) int64 { return 444 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToNillableInt64(t *testing.T) {
	Convey("String.MapToNillableInt64", t, func() {
		So(NewString("hi").MapToNillableInt64(func(string) *int64 {
			var tmp int64 = 444
			return &tmp
		}).Get(), ShouldEqual, 444)

		So(NewEmptyString().MapToNillableInt64(func(string) *int64 {
			var tmp int64 = 444
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewString("hi").MapToNillableInt64(func(string) *int64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToInt64(t *testing.T) {
	Convey("String.MapToInt64", t, func() {
		test1, err1 := NewString("hi").MapToInt64(func(string) (int64, error) {
			return 444, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 444)

		test2, err2 := NewEmptyString().MapToInt64(func(string) (int64, error) {
			return 444, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewString("hi").MapToInt64(func(string) (int64, error) {
			return 444, errors.New("int64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestString_ForceMapToInt8(t *testing.T) {
	Convey("String.ForceMapToInt8", t, func() {
		So(NewString("hi").ForceMapToInt8(func(string) int8 { return 111 }).Get(),
			ShouldEqual, 111)

		So(NewEmptyString().ForceMapToInt8(func(string) int8 { return 111 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToNillableInt8(t *testing.T) {
	Convey("String.MapToNillableInt8", t, func() {
		So(NewString("hi").MapToNillableInt8(func(string) *int8 {
			var tmp int8 = 111
			return &tmp
		}).Get(), ShouldEqual, 111)

		So(NewEmptyString().MapToNillableInt8(func(string) *int8 {
			var tmp int8 = 111
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewString("hi").MapToNillableInt8(func(string) *int8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToInt8(t *testing.T) {
	Convey("String.MapToInt8", t, func() {
		test1, err1 := NewString("hi").MapToInt8(func(string) (int8, error) {
			return 111, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 111)

		test2, err2 := NewEmptyString().MapToInt8(func(string) (int8, error) {
			return 111, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewString("hi").MapToInt8(func(string) (int8, error) {
			return 111, errors.New("int8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestString_ForceMapToUntyped(t *testing.T) {
	Convey("String.ForceMapToUntyped", t, func() {
		So(NewString("hi").ForceMapToUntyped(func(string) interface{} { return "111111" }).Get(),
			ShouldEqual, "111111")

		So(NewEmptyString().ForceMapToUntyped(func(string) interface{} { return "111111" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToNillableUntyped(t *testing.T) {
	Convey("String.MapToNillableUntyped", t, func() {
		So(NewString("hi").MapToNillableUntyped(func(string) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).Get(), ShouldEqual, "111111")

		So(NewEmptyString().MapToNillableUntyped(func(string) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewString("hi").MapToNillableUntyped(func(string) *interface{} { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToUntyped(t *testing.T) {
	Convey("String.MapToUntyped", t, func() {
		test1, err1 := NewString("hi").MapToUntyped(func(string) (interface{}, error) {
			return "111111", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "111111")

		test2, err2 := NewEmptyString().MapToUntyped(func(string) (interface{}, error) {
			return "111111", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewString("hi").MapToUntyped(func(string) (interface{}, error) {
			return "111111", errors.New("interface error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "interface error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestString_ForceMapToRune(t *testing.T) {
	Convey("String.ForceMapToRune", t, func() {
		So(NewString("hi").ForceMapToRune(func(string) rune { return 555 }).Get(),
			ShouldEqual, 555)

		So(NewEmptyString().ForceMapToRune(func(string) rune { return 555 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToNillableRune(t *testing.T) {
	Convey("String.MapToNillableRune", t, func() {
		So(NewString("hi").MapToNillableRune(func(string) *rune {
			var tmp rune = 555
			return &tmp
		}).Get(), ShouldEqual, 555)

		So(NewEmptyString().MapToNillableRune(func(string) *rune {
			var tmp rune = 555
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewString("hi").MapToNillableRune(func(string) *rune { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToRune(t *testing.T) {
	Convey("String.MapToRune", t, func() {
		test1, err1 := NewString("hi").MapToRune(func(string) (rune, error) {
			return 555, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 555)

		test2, err2 := NewEmptyString().MapToRune(func(string) (rune, error) {
			return 555, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewString("hi").MapToRune(func(string) (rune, error) {
			return 555, errors.New("rune error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "rune error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestString_ForceMapToUint(t *testing.T) {
	Convey("String.ForceMapToUint", t, func() {
		So(NewString("hi").ForceMapToUint(func(string) uint { return 666 }).Get(),
			ShouldEqual, 666)

		So(NewEmptyString().ForceMapToUint(func(string) uint { return 666 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToNillableUint(t *testing.T) {
	Convey("String.MapToNillableUint", t, func() {
		So(NewString("hi").MapToNillableUint(func(string) *uint {
			var tmp uint = 666
			return &tmp
		}).Get(), ShouldEqual, 666)

		So(NewEmptyString().MapToNillableUint(func(string) *uint {
			var tmp uint = 666
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewString("hi").MapToNillableUint(func(string) *uint { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToUint(t *testing.T) {
	Convey("String.MapToUint", t, func() {
		test1, err1 := NewString("hi").MapToUint(func(string) (uint, error) {
			return 666, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 666)

		test2, err2 := NewEmptyString().MapToUint(func(string) (uint, error) {
			return 666, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewString("hi").MapToUint(func(string) (uint, error) {
			return 666, errors.New("uint error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestString_ForceMapToUint16(t *testing.T) {
	Convey("String.ForceMapToUint16", t, func() {
		So(NewString("hi").ForceMapToUint16(func(string) uint16 { return 888 }).Get(),
			ShouldEqual, 888)

		So(NewEmptyString().ForceMapToUint16(func(string) uint16 { return 888 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToNillableUint16(t *testing.T) {
	Convey("String.MapToNillableUint16", t, func() {
		So(NewString("hi").MapToNillableUint16(func(string) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).Get(), ShouldEqual, 888)

		So(NewEmptyString().MapToNillableUint16(func(string) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewString("hi").MapToNillableUint16(func(string) *uint16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToUint16(t *testing.T) {
	Convey("String.MapToUint16", t, func() {
		test1, err1 := NewString("hi").MapToUint16(func(string) (uint16, error) {
			return 888, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 888)

		test2, err2 := NewEmptyString().MapToUint16(func(string) (uint16, error) {
			return 888, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewString("hi").MapToUint16(func(string) (uint16, error) {
			return 888, errors.New("uint16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestString_ForceMapToUint32(t *testing.T) {
	Convey("String.ForceMapToUint32", t, func() {
		So(NewString("hi").ForceMapToUint32(func(string) uint32 { return 999 }).Get(),
			ShouldEqual, 999)

		So(NewEmptyString().ForceMapToUint32(func(string) uint32 { return 999 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToNillableUint32(t *testing.T) {
	Convey("String.MapToNillableUint32", t, func() {
		So(NewString("hi").MapToNillableUint32(func(string) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).Get(), ShouldEqual, 999)

		So(NewEmptyString().MapToNillableUint32(func(string) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewString("hi").MapToNillableUint32(func(string) *uint32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToUint32(t *testing.T) {
	Convey("String.MapToUint32", t, func() {
		test1, err1 := NewString("hi").MapToUint32(func(string) (uint32, error) {
			return 999, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 999)

		test2, err2 := NewEmptyString().MapToUint32(func(string) (uint32, error) {
			return 999, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewString("hi").MapToUint32(func(string) (uint32, error) {
			return 999, errors.New("uint32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestString_ForceMapToUint64(t *testing.T) {
	Convey("String.ForceMapToUint64", t, func() {
		So(NewString("hi").ForceMapToUint64(func(string) uint64 { return 101010 }).Get(),
			ShouldEqual, 101010)

		So(NewEmptyString().ForceMapToUint64(func(string) uint64 { return 101010 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToNillableUint64(t *testing.T) {
	Convey("String.MapToNillableUint64", t, func() {
		So(NewString("hi").MapToNillableUint64(func(string) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).Get(), ShouldEqual, 101010)

		So(NewEmptyString().MapToNillableUint64(func(string) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewString("hi").MapToNillableUint64(func(string) *uint64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToUint64(t *testing.T) {
	Convey("String.MapToUint64", t, func() {
		test1, err1 := NewString("hi").MapToUint64(func(string) (uint64, error) {
			return 101010, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 101010)

		test2, err2 := NewEmptyString().MapToUint64(func(string) (uint64, error) {
			return 101010, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewString("hi").MapToUint64(func(string) (uint64, error) {
			return 101010, errors.New("uint64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestString_ForceMapToUint8(t *testing.T) {
	Convey("String.ForceMapToUint8", t, func() {
		So(NewString("hi").ForceMapToUint8(func(string) uint8 { return 77 }).Get(),
			ShouldEqual, 77)

		So(NewEmptyString().ForceMapToUint8(func(string) uint8 { return 77 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToNillableUint8(t *testing.T) {
	Convey("String.MapToNillableUint8", t, func() {
		So(NewString("hi").MapToNillableUint8(func(string) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).Get(), ShouldEqual, 77)

		So(NewEmptyString().MapToNillableUint8(func(string) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewString("hi").MapToNillableUint8(func(string) *uint8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestString_MapToUint8(t *testing.T) {
	Convey("String.MapToUint8", t, func() {
		test1, err1 := NewString("hi").MapToUint8(func(string) (uint8, error) {
			return 77, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 77)

		test2, err2 := NewEmptyString().MapToUint8(func(string) (uint8, error) {
			return 77, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewString("hi").MapToUint8(func(string) (uint8, error) {
			return 77, errors.New("uint8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}
