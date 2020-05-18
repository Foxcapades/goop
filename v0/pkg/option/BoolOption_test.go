// Package option_test generated @ 2020-05-17 22:19:55.164878293 -0400 EDT m=+0.001439508
package option_test

import (
	"errors"
	"testing"

	. "github.com/Foxcapades/goop/v0/pkg/option"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewBool(t *testing.T) {
	Convey("NewBool", t, func() {
		test := NewBool(true)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyBool(t *testing.T) {
	Convey("NewEmptyBool", t, func() {
		test := NewEmptyBool()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeBool(t *testing.T) {
	Convey("NewMaybeBool", t, func() {
		var val1 bool

		test1 := NewMaybeBool(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeBool(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestBool_Get(t *testing.T) {
	Convey("Bool.Get", t, func() {
		test1 := NewBool(true)
		So(test1.Get(), ShouldEqual, true)

		test2 := NewEmptyBool()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestBool_Or(t *testing.T) {
	Convey("Bool.Or", t, func() {
		test1 := NewBool(true)
		So(test1.Or(false), ShouldEqual, true)

		test2 := NewEmptyBool()
		So(test2.Or(true), ShouldEqual, true)
	})
}

func TestBool_OrGet(t *testing.T) {
	Convey("Bool.OrGet", t, func() {
		test1 := NewBool(true)
		So(test1.OrGet(func() bool { return false }), ShouldEqual, true)

		test2 := NewEmptyBool()
		So(test2.OrGet(func() bool { return true }), ShouldEqual, true)
	})
}

func TestBool_OrPanicWith(t *testing.T) {
	Convey("Bool.OrPanicWith", t, func() {
		test1 := NewBool(true)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, true)

		test2 := NewEmptyBool()
		So(func() {
			test2.OrPanicWith(func() interface{} { return "test value" })
		}, ShouldPanicWith, "test value")
	})
}

func TestBool_MapToNillable(t *testing.T) {
	Convey("Bool.MapToNillable", t, func() {
		test1 := NewBool(true)
		So(test1.MapToNillable(func(b bool) *bool { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewBool(true)
		So(test2.MapToNillable(func(b bool) *bool { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyBool()
		So(func() {
			test3.MapToNillable(func(b bool) *bool { panic("foo") }).IsNil()
		}, ShouldNotPanic)
	})
}

func TestBool_ForceMapToByte(t *testing.T) {
	Convey("Bool.ForceMapToByte", t, func() {
		So(NewBool(true).ForceMapToByte(func(bool) byte { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyBool().ForceMapToByte(func(bool) byte { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToNillableByte(t *testing.T) {
	Convey("Bool.MapToNillableByte", t, func() {
		So(NewBool(true).MapToNillableByte(func(bool) *byte {
			var tmp byte = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyBool().MapToNillableByte(func(bool) *byte {
			var tmp byte = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewBool(true).MapToNillableByte(func(bool) *byte { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToByte(t *testing.T) {
	Convey("Bool.MapToByte", t, func() {
		test1, err1 := NewBool(true).MapToByte(func(bool) (byte, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyBool().MapToByte(func(bool) (byte, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewBool(true).MapToByte(func(bool) (byte, error) {
			return 222, errors.New("byte error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "byte error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestBool_ForceMapToComplex128(t *testing.T) {
	Convey("Bool.ForceMapToComplex128", t, func() {
		So(NewBool(true).ForceMapToComplex128(func(bool) complex128 { return 987654 }).Get(),
			ShouldEqual, 987654)

		So(NewEmptyBool().ForceMapToComplex128(func(bool) complex128 { return 987654 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToNillableComplex128(t *testing.T) {
	Convey("Bool.MapToNillableComplex128", t, func() {
		So(NewBool(true).MapToNillableComplex128(func(bool) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).Get(), ShouldEqual, 987654)

		So(NewEmptyBool().MapToNillableComplex128(func(bool) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewBool(true).MapToNillableComplex128(func(bool) *complex128 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToComplex128(t *testing.T) {
	Convey("Bool.MapToComplex128", t, func() {
		test1, err1 := NewBool(true).MapToComplex128(func(bool) (complex128, error) {
			return 987654, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 987654)

		test2, err2 := NewEmptyBool().MapToComplex128(func(bool) (complex128, error) {
			return 987654, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewBool(true).MapToComplex128(func(bool) (complex128, error) {
			return 987654, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestBool_ForceMapToComplex64(t *testing.T) {
	Convey("Bool.ForceMapToComplex64", t, func() {
		So(NewBool(true).ForceMapToComplex64(func(bool) complex64 { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyBool().ForceMapToComplex64(func(bool) complex64 { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToNillableComplex64(t *testing.T) {
	Convey("Bool.MapToNillableComplex64", t, func() {
		So(NewBool(true).MapToNillableComplex64(func(bool) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyBool().MapToNillableComplex64(func(bool) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewBool(true).MapToNillableComplex64(func(bool) *complex64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToComplex64(t *testing.T) {
	Convey("Bool.MapToComplex64", t, func() {
		test1, err1 := NewBool(true).MapToComplex64(func(bool) (complex64, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyBool().MapToComplex64(func(bool) (complex64, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewBool(true).MapToComplex64(func(bool) (complex64, error) {
			return 123456, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestBool_ForceMapToFloat32(t *testing.T) {
	Convey("Bool.ForceMapToFloat32", t, func() {
		So(NewBool(true).ForceMapToFloat32(func(bool) float32 { return 32.3 }).Get(),
			ShouldEqual, 32.3)

		So(NewEmptyBool().ForceMapToFloat32(func(bool) float32 { return 32.3 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToNillableFloat32(t *testing.T) {
	Convey("Bool.MapToNillableFloat32", t, func() {
		So(NewBool(true).MapToNillableFloat32(func(bool) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).Get(), ShouldEqual, 32.3)

		So(NewEmptyBool().MapToNillableFloat32(func(bool) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewBool(true).MapToNillableFloat32(func(bool) *float32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToFloat32(t *testing.T) {
	Convey("Bool.MapToFloat32", t, func() {
		test1, err1 := NewBool(true).MapToFloat32(func(bool) (float32, error) {
			return 32.3, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 32.3)

		test2, err2 := NewEmptyBool().MapToFloat32(func(bool) (float32, error) {
			return 32.3, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewBool(true).MapToFloat32(func(bool) (float32, error) {
			return 32.3, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestBool_ForceMapToFloat64(t *testing.T) {
	Convey("Bool.ForceMapToFloat64", t, func() {
		So(NewBool(true).ForceMapToFloat64(func(bool) float64 { return 3.2455 }).Get(),
			ShouldEqual, 3.2455)

		So(NewEmptyBool().ForceMapToFloat64(func(bool) float64 { return 3.2455 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToNillableFloat64(t *testing.T) {
	Convey("Bool.MapToNillableFloat64", t, func() {
		So(NewBool(true).MapToNillableFloat64(func(bool) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).Get(), ShouldEqual, 3.2455)

		So(NewEmptyBool().MapToNillableFloat64(func(bool) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewBool(true).MapToNillableFloat64(func(bool) *float64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToFloat64(t *testing.T) {
	Convey("Bool.MapToFloat64", t, func() {
		test1, err1 := NewBool(true).MapToFloat64(func(bool) (float64, error) {
			return 3.2455, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 3.2455)

		test2, err2 := NewEmptyBool().MapToFloat64(func(bool) (float64, error) {
			return 3.2455, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewBool(true).MapToFloat64(func(bool) (float64, error) {
			return 3.2455, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestBool_ForceMapToInt(t *testing.T) {
	Convey("Bool.ForceMapToInt", t, func() {
		So(NewBool(true).ForceMapToInt(func(bool) int { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyBool().ForceMapToInt(func(bool) int { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToNillableInt(t *testing.T) {
	Convey("Bool.MapToNillableInt", t, func() {
		So(NewBool(true).MapToNillableInt(func(bool) *int {
			var tmp int = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyBool().MapToNillableInt(func(bool) *int {
			var tmp int = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewBool(true).MapToNillableInt(func(bool) *int { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToInt(t *testing.T) {
	Convey("Bool.MapToInt", t, func() {
		test1, err1 := NewBool(true).MapToInt(func(bool) (int, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyBool().MapToInt(func(bool) (int, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewBool(true).MapToInt(func(bool) (int, error) {
			return 123456, errors.New("int error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestBool_ForceMapToInt16(t *testing.T) {
	Convey("Bool.ForceMapToInt16", t, func() {
		So(NewBool(true).ForceMapToInt16(func(bool) int16 { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyBool().ForceMapToInt16(func(bool) int16 { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToNillableInt16(t *testing.T) {
	Convey("Bool.MapToNillableInt16", t, func() {
		So(NewBool(true).MapToNillableInt16(func(bool) *int16 {
			var tmp int16 = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyBool().MapToNillableInt16(func(bool) *int16 {
			var tmp int16 = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewBool(true).MapToNillableInt16(func(bool) *int16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToInt16(t *testing.T) {
	Convey("Bool.MapToInt16", t, func() {
		test1, err1 := NewBool(true).MapToInt16(func(bool) (int16, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyBool().MapToInt16(func(bool) (int16, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewBool(true).MapToInt16(func(bool) (int16, error) {
			return 222, errors.New("int16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestBool_ForceMapToInt32(t *testing.T) {
	Convey("Bool.ForceMapToInt32", t, func() {
		So(NewBool(true).ForceMapToInt32(func(bool) int32 { return 333 }).Get(),
			ShouldEqual, 333)

		So(NewEmptyBool().ForceMapToInt32(func(bool) int32 { return 333 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToNillableInt32(t *testing.T) {
	Convey("Bool.MapToNillableInt32", t, func() {
		So(NewBool(true).MapToNillableInt32(func(bool) *int32 {
			var tmp int32 = 333
			return &tmp
		}).Get(), ShouldEqual, 333)

		So(NewEmptyBool().MapToNillableInt32(func(bool) *int32 {
			var tmp int32 = 333
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewBool(true).MapToNillableInt32(func(bool) *int32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToInt32(t *testing.T) {
	Convey("Bool.MapToInt32", t, func() {
		test1, err1 := NewBool(true).MapToInt32(func(bool) (int32, error) {
			return 333, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 333)

		test2, err2 := NewEmptyBool().MapToInt32(func(bool) (int32, error) {
			return 333, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewBool(true).MapToInt32(func(bool) (int32, error) {
			return 333, errors.New("int32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestBool_ForceMapToInt64(t *testing.T) {
	Convey("Bool.ForceMapToInt64", t, func() {
		So(NewBool(true).ForceMapToInt64(func(bool) int64 { return 444 }).Get(),
			ShouldEqual, 444)

		So(NewEmptyBool().ForceMapToInt64(func(bool) int64 { return 444 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToNillableInt64(t *testing.T) {
	Convey("Bool.MapToNillableInt64", t, func() {
		So(NewBool(true).MapToNillableInt64(func(bool) *int64 {
			var tmp int64 = 444
			return &tmp
		}).Get(), ShouldEqual, 444)

		So(NewEmptyBool().MapToNillableInt64(func(bool) *int64 {
			var tmp int64 = 444
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewBool(true).MapToNillableInt64(func(bool) *int64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToInt64(t *testing.T) {
	Convey("Bool.MapToInt64", t, func() {
		test1, err1 := NewBool(true).MapToInt64(func(bool) (int64, error) {
			return 444, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 444)

		test2, err2 := NewEmptyBool().MapToInt64(func(bool) (int64, error) {
			return 444, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewBool(true).MapToInt64(func(bool) (int64, error) {
			return 444, errors.New("int64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestBool_ForceMapToInt8(t *testing.T) {
	Convey("Bool.ForceMapToInt8", t, func() {
		So(NewBool(true).ForceMapToInt8(func(bool) int8 { return 111 }).Get(),
			ShouldEqual, 111)

		So(NewEmptyBool().ForceMapToInt8(func(bool) int8 { return 111 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToNillableInt8(t *testing.T) {
	Convey("Bool.MapToNillableInt8", t, func() {
		So(NewBool(true).MapToNillableInt8(func(bool) *int8 {
			var tmp int8 = 111
			return &tmp
		}).Get(), ShouldEqual, 111)

		So(NewEmptyBool().MapToNillableInt8(func(bool) *int8 {
			var tmp int8 = 111
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewBool(true).MapToNillableInt8(func(bool) *int8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToInt8(t *testing.T) {
	Convey("Bool.MapToInt8", t, func() {
		test1, err1 := NewBool(true).MapToInt8(func(bool) (int8, error) {
			return 111, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 111)

		test2, err2 := NewEmptyBool().MapToInt8(func(bool) (int8, error) {
			return 111, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewBool(true).MapToInt8(func(bool) (int8, error) {
			return 111, errors.New("int8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestBool_ForceMapToUntyped(t *testing.T) {
	Convey("Bool.ForceMapToUntyped", t, func() {
		So(NewBool(true).ForceMapToUntyped(func(bool) interface{} { return "111111" }).Get(),
			ShouldEqual, "111111")

		So(NewEmptyBool().ForceMapToUntyped(func(bool) interface{} { return "111111" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToNillableUntyped(t *testing.T) {
	Convey("Bool.MapToNillableUntyped", t, func() {
		So(NewBool(true).MapToNillableUntyped(func(bool) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).Get(), ShouldEqual, "111111")

		So(NewEmptyBool().MapToNillableUntyped(func(bool) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewBool(true).MapToNillableUntyped(func(bool) *interface{} { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToUntyped(t *testing.T) {
	Convey("Bool.MapToUntyped", t, func() {
		test1, err1 := NewBool(true).MapToUntyped(func(bool) (interface{}, error) {
			return "111111", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "111111")

		test2, err2 := NewEmptyBool().MapToUntyped(func(bool) (interface{}, error) {
			return "111111", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewBool(true).MapToUntyped(func(bool) (interface{}, error) {
			return "111111", errors.New("interface error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "interface error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestBool_ForceMapToRune(t *testing.T) {
	Convey("Bool.ForceMapToRune", t, func() {
		So(NewBool(true).ForceMapToRune(func(bool) rune { return 555 }).Get(),
			ShouldEqual, 555)

		So(NewEmptyBool().ForceMapToRune(func(bool) rune { return 555 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToNillableRune(t *testing.T) {
	Convey("Bool.MapToNillableRune", t, func() {
		So(NewBool(true).MapToNillableRune(func(bool) *rune {
			var tmp rune = 555
			return &tmp
		}).Get(), ShouldEqual, 555)

		So(NewEmptyBool().MapToNillableRune(func(bool) *rune {
			var tmp rune = 555
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewBool(true).MapToNillableRune(func(bool) *rune { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToRune(t *testing.T) {
	Convey("Bool.MapToRune", t, func() {
		test1, err1 := NewBool(true).MapToRune(func(bool) (rune, error) {
			return 555, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 555)

		test2, err2 := NewEmptyBool().MapToRune(func(bool) (rune, error) {
			return 555, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewBool(true).MapToRune(func(bool) (rune, error) {
			return 555, errors.New("rune error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "rune error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestBool_ForceMapToString(t *testing.T) {
	Convey("Bool.ForceMapToString", t, func() {
		So(NewBool(true).ForceMapToString(func(bool) string { return "bonanza" }).Get(),
			ShouldEqual, "bonanza")

		So(NewEmptyBool().ForceMapToString(func(bool) string { return "bonanza" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToNillableString(t *testing.T) {
	Convey("Bool.MapToNillableString", t, func() {
		So(NewBool(true).MapToNillableString(func(bool) *string {
			var tmp string = "bonanza"
			return &tmp
		}).Get(), ShouldEqual, "bonanza")

		So(NewEmptyBool().MapToNillableString(func(bool) *string {
			var tmp string = "bonanza"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewBool(true).MapToNillableString(func(bool) *string { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToString(t *testing.T) {
	Convey("Bool.MapToString", t, func() {
		test1, err1 := NewBool(true).MapToString(func(bool) (string, error) {
			return "bonanza", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "bonanza")

		test2, err2 := NewEmptyBool().MapToString(func(bool) (string, error) {
			return "bonanza", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewBool(true).MapToString(func(bool) (string, error) {
			return "bonanza", errors.New("string error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "string error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestBool_ForceMapToUint(t *testing.T) {
	Convey("Bool.ForceMapToUint", t, func() {
		So(NewBool(true).ForceMapToUint(func(bool) uint { return 666 }).Get(),
			ShouldEqual, 666)

		So(NewEmptyBool().ForceMapToUint(func(bool) uint { return 666 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToNillableUint(t *testing.T) {
	Convey("Bool.MapToNillableUint", t, func() {
		So(NewBool(true).MapToNillableUint(func(bool) *uint {
			var tmp uint = 666
			return &tmp
		}).Get(), ShouldEqual, 666)

		So(NewEmptyBool().MapToNillableUint(func(bool) *uint {
			var tmp uint = 666
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewBool(true).MapToNillableUint(func(bool) *uint { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToUint(t *testing.T) {
	Convey("Bool.MapToUint", t, func() {
		test1, err1 := NewBool(true).MapToUint(func(bool) (uint, error) {
			return 666, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 666)

		test2, err2 := NewEmptyBool().MapToUint(func(bool) (uint, error) {
			return 666, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewBool(true).MapToUint(func(bool) (uint, error) {
			return 666, errors.New("uint error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestBool_ForceMapToUint16(t *testing.T) {
	Convey("Bool.ForceMapToUint16", t, func() {
		So(NewBool(true).ForceMapToUint16(func(bool) uint16 { return 888 }).Get(),
			ShouldEqual, 888)

		So(NewEmptyBool().ForceMapToUint16(func(bool) uint16 { return 888 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToNillableUint16(t *testing.T) {
	Convey("Bool.MapToNillableUint16", t, func() {
		So(NewBool(true).MapToNillableUint16(func(bool) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).Get(), ShouldEqual, 888)

		So(NewEmptyBool().MapToNillableUint16(func(bool) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewBool(true).MapToNillableUint16(func(bool) *uint16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToUint16(t *testing.T) {
	Convey("Bool.MapToUint16", t, func() {
		test1, err1 := NewBool(true).MapToUint16(func(bool) (uint16, error) {
			return 888, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 888)

		test2, err2 := NewEmptyBool().MapToUint16(func(bool) (uint16, error) {
			return 888, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewBool(true).MapToUint16(func(bool) (uint16, error) {
			return 888, errors.New("uint16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestBool_ForceMapToUint32(t *testing.T) {
	Convey("Bool.ForceMapToUint32", t, func() {
		So(NewBool(true).ForceMapToUint32(func(bool) uint32 { return 999 }).Get(),
			ShouldEqual, 999)

		So(NewEmptyBool().ForceMapToUint32(func(bool) uint32 { return 999 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToNillableUint32(t *testing.T) {
	Convey("Bool.MapToNillableUint32", t, func() {
		So(NewBool(true).MapToNillableUint32(func(bool) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).Get(), ShouldEqual, 999)

		So(NewEmptyBool().MapToNillableUint32(func(bool) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewBool(true).MapToNillableUint32(func(bool) *uint32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToUint32(t *testing.T) {
	Convey("Bool.MapToUint32", t, func() {
		test1, err1 := NewBool(true).MapToUint32(func(bool) (uint32, error) {
			return 999, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 999)

		test2, err2 := NewEmptyBool().MapToUint32(func(bool) (uint32, error) {
			return 999, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewBool(true).MapToUint32(func(bool) (uint32, error) {
			return 999, errors.New("uint32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestBool_ForceMapToUint64(t *testing.T) {
	Convey("Bool.ForceMapToUint64", t, func() {
		So(NewBool(true).ForceMapToUint64(func(bool) uint64 { return 101010 }).Get(),
			ShouldEqual, 101010)

		So(NewEmptyBool().ForceMapToUint64(func(bool) uint64 { return 101010 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToNillableUint64(t *testing.T) {
	Convey("Bool.MapToNillableUint64", t, func() {
		So(NewBool(true).MapToNillableUint64(func(bool) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).Get(), ShouldEqual, 101010)

		So(NewEmptyBool().MapToNillableUint64(func(bool) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewBool(true).MapToNillableUint64(func(bool) *uint64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToUint64(t *testing.T) {
	Convey("Bool.MapToUint64", t, func() {
		test1, err1 := NewBool(true).MapToUint64(func(bool) (uint64, error) {
			return 101010, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 101010)

		test2, err2 := NewEmptyBool().MapToUint64(func(bool) (uint64, error) {
			return 101010, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewBool(true).MapToUint64(func(bool) (uint64, error) {
			return 101010, errors.New("uint64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestBool_ForceMapToUint8(t *testing.T) {
	Convey("Bool.ForceMapToUint8", t, func() {
		So(NewBool(true).ForceMapToUint8(func(bool) uint8 { return 77 }).Get(),
			ShouldEqual, 77)

		So(NewEmptyBool().ForceMapToUint8(func(bool) uint8 { return 77 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToNillableUint8(t *testing.T) {
	Convey("Bool.MapToNillableUint8", t, func() {
		So(NewBool(true).MapToNillableUint8(func(bool) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).Get(), ShouldEqual, 77)

		So(NewEmptyBool().MapToNillableUint8(func(bool) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewBool(true).MapToNillableUint8(func(bool) *uint8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestBool_MapToUint8(t *testing.T) {
	Convey("Bool.MapToUint8", t, func() {
		test1, err1 := NewBool(true).MapToUint8(func(bool) (uint8, error) {
			return 77, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 77)

		test2, err2 := NewEmptyBool().MapToUint8(func(bool) (uint8, error) {
			return 77, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewBool(true).MapToUint8(func(bool) (uint8, error) {
			return 77, errors.New("uint8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}
