// Package option_test generated @ 2020-05-17 22:19:55.164878293 -0400 EDT m=+0.001439508
package option_test

import (
	"errors"
	"testing"

	. "github.com/Foxcapades/goop/v0/pkg/option"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewInt(t *testing.T) {
	Convey("NewInt", t, func() {
		test := NewInt(123)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyInt(t *testing.T) {
	Convey("NewEmptyInt", t, func() {
		test := NewEmptyInt()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeInt(t *testing.T) {
	Convey("NewMaybeInt", t, func() {
		var val1 int

		test1 := NewMaybeInt(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeInt(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestInt_Get(t *testing.T) {
	Convey("Int.Get", t, func() {
		test1 := NewInt(123)
		So(test1.Get(), ShouldEqual, 123)

		test2 := NewEmptyInt()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestInt_Or(t *testing.T) {
	Convey("Int.Or", t, func() {
		test1 := NewInt(123)
		So(test1.Or(321), ShouldEqual, 123)

		test2 := NewEmptyInt()
		So(test2.Or(123), ShouldEqual, 123)
	})
}

func TestInt_OrGet(t *testing.T) {
	Convey("Int.OrGet", t, func() {
		test1 := NewInt(123)
		So(test1.OrGet(func() int { return 321 }), ShouldEqual, 123)

		test2 := NewEmptyInt()
		So(test2.OrGet(func() int { return 123 }), ShouldEqual, 123)
	})
}

func TestInt_OrPanicWith(t *testing.T) {
	Convey("Int.OrPanicWith", t, func() {
		test1 := NewInt(123)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 123)

		test2 := NewEmptyInt()
		So(func() {
			test2.OrPanicWith(func() interface{} { return "test value" })
		}, ShouldPanicWith, "test value")
	})
}

func TestInt_MapToNillable(t *testing.T) {
	Convey("Int.MapToNillable", t, func() {
		test1 := NewInt(123)
		So(test1.MapToNillable(func(b int) *int { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewInt(123)
		So(test2.MapToNillable(func(b int) *int { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyInt()
		So(func() {
			test3.MapToNillable(func(b int) *int { panic("foo") }).IsNil()
		}, ShouldNotPanic)
	})
}

func TestInt_ForceMapToBool(t *testing.T) {
	Convey("Int.ForceMapToBool", t, func() {
		So(NewInt(123).ForceMapToBool(func(int) bool { return true }).Get(),
			ShouldEqual, true)

		So(NewEmptyInt().ForceMapToBool(func(int) bool { return true }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToNillableBool(t *testing.T) {
	Convey("Int.MapToNillableBool", t, func() {
		So(NewInt(123).MapToNillableBool(func(int) *bool {
			var tmp bool = true
			return &tmp
		}).Get(), ShouldEqual, true)

		So(NewEmptyInt().MapToNillableBool(func(int) *bool {
			var tmp bool = true
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt(123).MapToNillableBool(func(int) *bool { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToBool(t *testing.T) {
	Convey("Int.MapToBool", t, func() {
		test1, err1 := NewInt(123).MapToBool(func(int) (bool, error) {
			return true, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, true)

		test2, err2 := NewEmptyInt().MapToBool(func(int) (bool, error) {
			return true, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt(123).MapToBool(func(int) (bool, error) {
			return true, errors.New("bool error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "bool error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt_ForceMapToByte(t *testing.T) {
	Convey("Int.ForceMapToByte", t, func() {
		So(NewInt(123).ForceMapToByte(func(int) byte { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyInt().ForceMapToByte(func(int) byte { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToNillableByte(t *testing.T) {
	Convey("Int.MapToNillableByte", t, func() {
		So(NewInt(123).MapToNillableByte(func(int) *byte {
			var tmp byte = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyInt().MapToNillableByte(func(int) *byte {
			var tmp byte = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt(123).MapToNillableByte(func(int) *byte { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToByte(t *testing.T) {
	Convey("Int.MapToByte", t, func() {
		test1, err1 := NewInt(123).MapToByte(func(int) (byte, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyInt().MapToByte(func(int) (byte, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt(123).MapToByte(func(int) (byte, error) {
			return 222, errors.New("byte error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "byte error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt_ForceMapToComplex128(t *testing.T) {
	Convey("Int.ForceMapToComplex128", t, func() {
		So(NewInt(123).ForceMapToComplex128(func(int) complex128 { return 987654 }).Get(),
			ShouldEqual, 987654)

		So(NewEmptyInt().ForceMapToComplex128(func(int) complex128 { return 987654 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToNillableComplex128(t *testing.T) {
	Convey("Int.MapToNillableComplex128", t, func() {
		So(NewInt(123).MapToNillableComplex128(func(int) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).Get(), ShouldEqual, 987654)

		So(NewEmptyInt().MapToNillableComplex128(func(int) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt(123).MapToNillableComplex128(func(int) *complex128 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToComplex128(t *testing.T) {
	Convey("Int.MapToComplex128", t, func() {
		test1, err1 := NewInt(123).MapToComplex128(func(int) (complex128, error) {
			return 987654, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 987654)

		test2, err2 := NewEmptyInt().MapToComplex128(func(int) (complex128, error) {
			return 987654, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt(123).MapToComplex128(func(int) (complex128, error) {
			return 987654, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt_ForceMapToComplex64(t *testing.T) {
	Convey("Int.ForceMapToComplex64", t, func() {
		So(NewInt(123).ForceMapToComplex64(func(int) complex64 { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyInt().ForceMapToComplex64(func(int) complex64 { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToNillableComplex64(t *testing.T) {
	Convey("Int.MapToNillableComplex64", t, func() {
		So(NewInt(123).MapToNillableComplex64(func(int) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyInt().MapToNillableComplex64(func(int) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt(123).MapToNillableComplex64(func(int) *complex64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToComplex64(t *testing.T) {
	Convey("Int.MapToComplex64", t, func() {
		test1, err1 := NewInt(123).MapToComplex64(func(int) (complex64, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyInt().MapToComplex64(func(int) (complex64, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt(123).MapToComplex64(func(int) (complex64, error) {
			return 123456, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt_ForceMapToFloat32(t *testing.T) {
	Convey("Int.ForceMapToFloat32", t, func() {
		So(NewInt(123).ForceMapToFloat32(func(int) float32 { return 32.3 }).Get(),
			ShouldEqual, 32.3)

		So(NewEmptyInt().ForceMapToFloat32(func(int) float32 { return 32.3 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToNillableFloat32(t *testing.T) {
	Convey("Int.MapToNillableFloat32", t, func() {
		So(NewInt(123).MapToNillableFloat32(func(int) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).Get(), ShouldEqual, 32.3)

		So(NewEmptyInt().MapToNillableFloat32(func(int) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt(123).MapToNillableFloat32(func(int) *float32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToFloat32(t *testing.T) {
	Convey("Int.MapToFloat32", t, func() {
		test1, err1 := NewInt(123).MapToFloat32(func(int) (float32, error) {
			return 32.3, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 32.3)

		test2, err2 := NewEmptyInt().MapToFloat32(func(int) (float32, error) {
			return 32.3, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt(123).MapToFloat32(func(int) (float32, error) {
			return 32.3, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt_ForceMapToFloat64(t *testing.T) {
	Convey("Int.ForceMapToFloat64", t, func() {
		So(NewInt(123).ForceMapToFloat64(func(int) float64 { return 3.2455 }).Get(),
			ShouldEqual, 3.2455)

		So(NewEmptyInt().ForceMapToFloat64(func(int) float64 { return 3.2455 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToNillableFloat64(t *testing.T) {
	Convey("Int.MapToNillableFloat64", t, func() {
		So(NewInt(123).MapToNillableFloat64(func(int) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).Get(), ShouldEqual, 3.2455)

		So(NewEmptyInt().MapToNillableFloat64(func(int) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt(123).MapToNillableFloat64(func(int) *float64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToFloat64(t *testing.T) {
	Convey("Int.MapToFloat64", t, func() {
		test1, err1 := NewInt(123).MapToFloat64(func(int) (float64, error) {
			return 3.2455, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 3.2455)

		test2, err2 := NewEmptyInt().MapToFloat64(func(int) (float64, error) {
			return 3.2455, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt(123).MapToFloat64(func(int) (float64, error) {
			return 3.2455, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt_ForceMapToInt16(t *testing.T) {
	Convey("Int.ForceMapToInt16", t, func() {
		So(NewInt(123).ForceMapToInt16(func(int) int16 { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyInt().ForceMapToInt16(func(int) int16 { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToNillableInt16(t *testing.T) {
	Convey("Int.MapToNillableInt16", t, func() {
		So(NewInt(123).MapToNillableInt16(func(int) *int16 {
			var tmp int16 = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyInt().MapToNillableInt16(func(int) *int16 {
			var tmp int16 = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt(123).MapToNillableInt16(func(int) *int16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToInt16(t *testing.T) {
	Convey("Int.MapToInt16", t, func() {
		test1, err1 := NewInt(123).MapToInt16(func(int) (int16, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyInt().MapToInt16(func(int) (int16, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt(123).MapToInt16(func(int) (int16, error) {
			return 222, errors.New("int16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt_ForceMapToInt32(t *testing.T) {
	Convey("Int.ForceMapToInt32", t, func() {
		So(NewInt(123).ForceMapToInt32(func(int) int32 { return 333 }).Get(),
			ShouldEqual, 333)

		So(NewEmptyInt().ForceMapToInt32(func(int) int32 { return 333 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToNillableInt32(t *testing.T) {
	Convey("Int.MapToNillableInt32", t, func() {
		So(NewInt(123).MapToNillableInt32(func(int) *int32 {
			var tmp int32 = 333
			return &tmp
		}).Get(), ShouldEqual, 333)

		So(NewEmptyInt().MapToNillableInt32(func(int) *int32 {
			var tmp int32 = 333
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt(123).MapToNillableInt32(func(int) *int32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToInt32(t *testing.T) {
	Convey("Int.MapToInt32", t, func() {
		test1, err1 := NewInt(123).MapToInt32(func(int) (int32, error) {
			return 333, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 333)

		test2, err2 := NewEmptyInt().MapToInt32(func(int) (int32, error) {
			return 333, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt(123).MapToInt32(func(int) (int32, error) {
			return 333, errors.New("int32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt_ForceMapToInt64(t *testing.T) {
	Convey("Int.ForceMapToInt64", t, func() {
		So(NewInt(123).ForceMapToInt64(func(int) int64 { return 444 }).Get(),
			ShouldEqual, 444)

		So(NewEmptyInt().ForceMapToInt64(func(int) int64 { return 444 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToNillableInt64(t *testing.T) {
	Convey("Int.MapToNillableInt64", t, func() {
		So(NewInt(123).MapToNillableInt64(func(int) *int64 {
			var tmp int64 = 444
			return &tmp
		}).Get(), ShouldEqual, 444)

		So(NewEmptyInt().MapToNillableInt64(func(int) *int64 {
			var tmp int64 = 444
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt(123).MapToNillableInt64(func(int) *int64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToInt64(t *testing.T) {
	Convey("Int.MapToInt64", t, func() {
		test1, err1 := NewInt(123).MapToInt64(func(int) (int64, error) {
			return 444, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 444)

		test2, err2 := NewEmptyInt().MapToInt64(func(int) (int64, error) {
			return 444, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt(123).MapToInt64(func(int) (int64, error) {
			return 444, errors.New("int64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt_ForceMapToInt8(t *testing.T) {
	Convey("Int.ForceMapToInt8", t, func() {
		So(NewInt(123).ForceMapToInt8(func(int) int8 { return 111 }).Get(),
			ShouldEqual, 111)

		So(NewEmptyInt().ForceMapToInt8(func(int) int8 { return 111 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToNillableInt8(t *testing.T) {
	Convey("Int.MapToNillableInt8", t, func() {
		So(NewInt(123).MapToNillableInt8(func(int) *int8 {
			var tmp int8 = 111
			return &tmp
		}).Get(), ShouldEqual, 111)

		So(NewEmptyInt().MapToNillableInt8(func(int) *int8 {
			var tmp int8 = 111
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt(123).MapToNillableInt8(func(int) *int8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToInt8(t *testing.T) {
	Convey("Int.MapToInt8", t, func() {
		test1, err1 := NewInt(123).MapToInt8(func(int) (int8, error) {
			return 111, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 111)

		test2, err2 := NewEmptyInt().MapToInt8(func(int) (int8, error) {
			return 111, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt(123).MapToInt8(func(int) (int8, error) {
			return 111, errors.New("int8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt_ForceMapToUntyped(t *testing.T) {
	Convey("Int.ForceMapToUntyped", t, func() {
		So(NewInt(123).ForceMapToUntyped(func(int) interface{} { return "111111" }).Get(),
			ShouldEqual, "111111")

		So(NewEmptyInt().ForceMapToUntyped(func(int) interface{} { return "111111" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToNillableUntyped(t *testing.T) {
	Convey("Int.MapToNillableUntyped", t, func() {
		So(NewInt(123).MapToNillableUntyped(func(int) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).Get(), ShouldEqual, "111111")

		So(NewEmptyInt().MapToNillableUntyped(func(int) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt(123).MapToNillableUntyped(func(int) *interface{} { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToUntyped(t *testing.T) {
	Convey("Int.MapToUntyped", t, func() {
		test1, err1 := NewInt(123).MapToUntyped(func(int) (interface{}, error) {
			return "111111", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "111111")

		test2, err2 := NewEmptyInt().MapToUntyped(func(int) (interface{}, error) {
			return "111111", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt(123).MapToUntyped(func(int) (interface{}, error) {
			return "111111", errors.New("interface error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "interface error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt_ForceMapToRune(t *testing.T) {
	Convey("Int.ForceMapToRune", t, func() {
		So(NewInt(123).ForceMapToRune(func(int) rune { return 555 }).Get(),
			ShouldEqual, 555)

		So(NewEmptyInt().ForceMapToRune(func(int) rune { return 555 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToNillableRune(t *testing.T) {
	Convey("Int.MapToNillableRune", t, func() {
		So(NewInt(123).MapToNillableRune(func(int) *rune {
			var tmp rune = 555
			return &tmp
		}).Get(), ShouldEqual, 555)

		So(NewEmptyInt().MapToNillableRune(func(int) *rune {
			var tmp rune = 555
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt(123).MapToNillableRune(func(int) *rune { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToRune(t *testing.T) {
	Convey("Int.MapToRune", t, func() {
		test1, err1 := NewInt(123).MapToRune(func(int) (rune, error) {
			return 555, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 555)

		test2, err2 := NewEmptyInt().MapToRune(func(int) (rune, error) {
			return 555, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt(123).MapToRune(func(int) (rune, error) {
			return 555, errors.New("rune error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "rune error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt_ForceMapToString(t *testing.T) {
	Convey("Int.ForceMapToString", t, func() {
		So(NewInt(123).ForceMapToString(func(int) string { return "bonanza" }).Get(),
			ShouldEqual, "bonanza")

		So(NewEmptyInt().ForceMapToString(func(int) string { return "bonanza" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToNillableString(t *testing.T) {
	Convey("Int.MapToNillableString", t, func() {
		So(NewInt(123).MapToNillableString(func(int) *string {
			var tmp string = "bonanza"
			return &tmp
		}).Get(), ShouldEqual, "bonanza")

		So(NewEmptyInt().MapToNillableString(func(int) *string {
			var tmp string = "bonanza"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt(123).MapToNillableString(func(int) *string { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToString(t *testing.T) {
	Convey("Int.MapToString", t, func() {
		test1, err1 := NewInt(123).MapToString(func(int) (string, error) {
			return "bonanza", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "bonanza")

		test2, err2 := NewEmptyInt().MapToString(func(int) (string, error) {
			return "bonanza", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt(123).MapToString(func(int) (string, error) {
			return "bonanza", errors.New("string error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "string error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt_ForceMapToUint(t *testing.T) {
	Convey("Int.ForceMapToUint", t, func() {
		So(NewInt(123).ForceMapToUint(func(int) uint { return 666 }).Get(),
			ShouldEqual, 666)

		So(NewEmptyInt().ForceMapToUint(func(int) uint { return 666 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToNillableUint(t *testing.T) {
	Convey("Int.MapToNillableUint", t, func() {
		So(NewInt(123).MapToNillableUint(func(int) *uint {
			var tmp uint = 666
			return &tmp
		}).Get(), ShouldEqual, 666)

		So(NewEmptyInt().MapToNillableUint(func(int) *uint {
			var tmp uint = 666
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt(123).MapToNillableUint(func(int) *uint { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToUint(t *testing.T) {
	Convey("Int.MapToUint", t, func() {
		test1, err1 := NewInt(123).MapToUint(func(int) (uint, error) {
			return 666, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 666)

		test2, err2 := NewEmptyInt().MapToUint(func(int) (uint, error) {
			return 666, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt(123).MapToUint(func(int) (uint, error) {
			return 666, errors.New("uint error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt_ForceMapToUint16(t *testing.T) {
	Convey("Int.ForceMapToUint16", t, func() {
		So(NewInt(123).ForceMapToUint16(func(int) uint16 { return 888 }).Get(),
			ShouldEqual, 888)

		So(NewEmptyInt().ForceMapToUint16(func(int) uint16 { return 888 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToNillableUint16(t *testing.T) {
	Convey("Int.MapToNillableUint16", t, func() {
		So(NewInt(123).MapToNillableUint16(func(int) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).Get(), ShouldEqual, 888)

		So(NewEmptyInt().MapToNillableUint16(func(int) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt(123).MapToNillableUint16(func(int) *uint16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToUint16(t *testing.T) {
	Convey("Int.MapToUint16", t, func() {
		test1, err1 := NewInt(123).MapToUint16(func(int) (uint16, error) {
			return 888, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 888)

		test2, err2 := NewEmptyInt().MapToUint16(func(int) (uint16, error) {
			return 888, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt(123).MapToUint16(func(int) (uint16, error) {
			return 888, errors.New("uint16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt_ForceMapToUint32(t *testing.T) {
	Convey("Int.ForceMapToUint32", t, func() {
		So(NewInt(123).ForceMapToUint32(func(int) uint32 { return 999 }).Get(),
			ShouldEqual, 999)

		So(NewEmptyInt().ForceMapToUint32(func(int) uint32 { return 999 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToNillableUint32(t *testing.T) {
	Convey("Int.MapToNillableUint32", t, func() {
		So(NewInt(123).MapToNillableUint32(func(int) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).Get(), ShouldEqual, 999)

		So(NewEmptyInt().MapToNillableUint32(func(int) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt(123).MapToNillableUint32(func(int) *uint32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToUint32(t *testing.T) {
	Convey("Int.MapToUint32", t, func() {
		test1, err1 := NewInt(123).MapToUint32(func(int) (uint32, error) {
			return 999, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 999)

		test2, err2 := NewEmptyInt().MapToUint32(func(int) (uint32, error) {
			return 999, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt(123).MapToUint32(func(int) (uint32, error) {
			return 999, errors.New("uint32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt_ForceMapToUint64(t *testing.T) {
	Convey("Int.ForceMapToUint64", t, func() {
		So(NewInt(123).ForceMapToUint64(func(int) uint64 { return 101010 }).Get(),
			ShouldEqual, 101010)

		So(NewEmptyInt().ForceMapToUint64(func(int) uint64 { return 101010 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToNillableUint64(t *testing.T) {
	Convey("Int.MapToNillableUint64", t, func() {
		So(NewInt(123).MapToNillableUint64(func(int) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).Get(), ShouldEqual, 101010)

		So(NewEmptyInt().MapToNillableUint64(func(int) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt(123).MapToNillableUint64(func(int) *uint64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToUint64(t *testing.T) {
	Convey("Int.MapToUint64", t, func() {
		test1, err1 := NewInt(123).MapToUint64(func(int) (uint64, error) {
			return 101010, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 101010)

		test2, err2 := NewEmptyInt().MapToUint64(func(int) (uint64, error) {
			return 101010, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt(123).MapToUint64(func(int) (uint64, error) {
			return 101010, errors.New("uint64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestInt_ForceMapToUint8(t *testing.T) {
	Convey("Int.ForceMapToUint8", t, func() {
		So(NewInt(123).ForceMapToUint8(func(int) uint8 { return 77 }).Get(),
			ShouldEqual, 77)

		So(NewEmptyInt().ForceMapToUint8(func(int) uint8 { return 77 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToNillableUint8(t *testing.T) {
	Convey("Int.MapToNillableUint8", t, func() {
		So(NewInt(123).MapToNillableUint8(func(int) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).Get(), ShouldEqual, 77)

		So(NewEmptyInt().MapToNillableUint8(func(int) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewInt(123).MapToNillableUint8(func(int) *uint8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestInt_MapToUint8(t *testing.T) {
	Convey("Int.MapToUint8", t, func() {
		test1, err1 := NewInt(123).MapToUint8(func(int) (uint8, error) {
			return 77, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 77)

		test2, err2 := NewEmptyInt().MapToUint8(func(int) (uint8, error) {
			return 77, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewInt(123).MapToUint8(func(int) (uint8, error) {
			return 77, errors.New("uint8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}
