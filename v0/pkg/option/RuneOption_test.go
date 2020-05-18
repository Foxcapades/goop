// Package option_test generated @ 2020-05-17 22:19:55.164878293 -0400 EDT m=+0.001439508
package option_test

import (
	"errors"
	"testing"

	. "github.com/Foxcapades/goop/v0/pkg/option"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewRune(t *testing.T) {
	Convey("NewRune", t, func() {
		test := NewRune(454545)
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyRune(t *testing.T) {
	Convey("NewEmptyRune", t, func() {
		test := NewEmptyRune()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeRune(t *testing.T) {
	Convey("NewMaybeRune", t, func() {
		var val1 rune

		test1 := NewMaybeRune(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeRune(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestRune_Get(t *testing.T) {
	Convey("Rune.Get", t, func() {
		test1 := NewRune(454545)
		So(test1.Get(), ShouldEqual, 454545)

		test2 := NewEmptyRune()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestRune_Or(t *testing.T) {
	Convey("Rune.Or", t, func() {
		test1 := NewRune(454545)
		So(test1.Or(898989), ShouldEqual, 454545)

		test2 := NewEmptyRune()
		So(test2.Or(454545), ShouldEqual, 454545)
	})
}

func TestRune_OrGet(t *testing.T) {
	Convey("Rune.OrGet", t, func() {
		test1 := NewRune(454545)
		So(test1.OrGet(func() rune { return 898989 }), ShouldEqual, 454545)

		test2 := NewEmptyRune()
		So(test2.OrGet(func() rune { return 454545 }), ShouldEqual, 454545)
	})
}

func TestRune_OrPanicWith(t *testing.T) {
	Convey("Rune.OrPanicWith", t, func() {
		test1 := NewRune(454545)
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, 454545)

		test2 := NewEmptyRune()
		So(func() {
			test2.OrPanicWith(func() interface{} { return "test value" })
		}, ShouldPanicWith, "test value")
	})
}

func TestRune_MapToNillable(t *testing.T) {
	Convey("Rune.MapToNillable", t, func() {
		test1 := NewRune(454545)
		So(test1.MapToNillable(func(b rune) *rune { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewRune(454545)
		So(test2.MapToNillable(func(b rune) *rune { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyRune()
		So(func() {
			test3.MapToNillable(func(b rune) *rune { panic("foo") }).IsNil()
		}, ShouldNotPanic)
	})
}

func TestRune_ForceMapToBool(t *testing.T) {
	Convey("Rune.ForceMapToBool", t, func() {
		So(NewRune(454545).ForceMapToBool(func(rune) bool { return true }).Get(),
			ShouldEqual, true)

		So(NewEmptyRune().ForceMapToBool(func(rune) bool { return true }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToNillableBool(t *testing.T) {
	Convey("Rune.MapToNillableBool", t, func() {
		So(NewRune(454545).MapToNillableBool(func(rune) *bool {
			var tmp bool = true
			return &tmp
		}).Get(), ShouldEqual, true)

		So(NewEmptyRune().MapToNillableBool(func(rune) *bool {
			var tmp bool = true
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewRune(454545).MapToNillableBool(func(rune) *bool { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToBool(t *testing.T) {
	Convey("Rune.MapToBool", t, func() {
		test1, err1 := NewRune(454545).MapToBool(func(rune) (bool, error) {
			return true, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, true)

		test2, err2 := NewEmptyRune().MapToBool(func(rune) (bool, error) {
			return true, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewRune(454545).MapToBool(func(rune) (bool, error) {
			return true, errors.New("bool error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "bool error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestRune_ForceMapToByte(t *testing.T) {
	Convey("Rune.ForceMapToByte", t, func() {
		So(NewRune(454545).ForceMapToByte(func(rune) byte { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyRune().ForceMapToByte(func(rune) byte { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToNillableByte(t *testing.T) {
	Convey("Rune.MapToNillableByte", t, func() {
		So(NewRune(454545).MapToNillableByte(func(rune) *byte {
			var tmp byte = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyRune().MapToNillableByte(func(rune) *byte {
			var tmp byte = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewRune(454545).MapToNillableByte(func(rune) *byte { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToByte(t *testing.T) {
	Convey("Rune.MapToByte", t, func() {
		test1, err1 := NewRune(454545).MapToByte(func(rune) (byte, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyRune().MapToByte(func(rune) (byte, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewRune(454545).MapToByte(func(rune) (byte, error) {
			return 222, errors.New("byte error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "byte error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestRune_ForceMapToComplex128(t *testing.T) {
	Convey("Rune.ForceMapToComplex128", t, func() {
		So(NewRune(454545).ForceMapToComplex128(func(rune) complex128 { return 987654 }).Get(),
			ShouldEqual, 987654)

		So(NewEmptyRune().ForceMapToComplex128(func(rune) complex128 { return 987654 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToNillableComplex128(t *testing.T) {
	Convey("Rune.MapToNillableComplex128", t, func() {
		So(NewRune(454545).MapToNillableComplex128(func(rune) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).Get(), ShouldEqual, 987654)

		So(NewEmptyRune().MapToNillableComplex128(func(rune) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewRune(454545).MapToNillableComplex128(func(rune) *complex128 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToComplex128(t *testing.T) {
	Convey("Rune.MapToComplex128", t, func() {
		test1, err1 := NewRune(454545).MapToComplex128(func(rune) (complex128, error) {
			return 987654, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 987654)

		test2, err2 := NewEmptyRune().MapToComplex128(func(rune) (complex128, error) {
			return 987654, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewRune(454545).MapToComplex128(func(rune) (complex128, error) {
			return 987654, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestRune_ForceMapToComplex64(t *testing.T) {
	Convey("Rune.ForceMapToComplex64", t, func() {
		So(NewRune(454545).ForceMapToComplex64(func(rune) complex64 { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyRune().ForceMapToComplex64(func(rune) complex64 { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToNillableComplex64(t *testing.T) {
	Convey("Rune.MapToNillableComplex64", t, func() {
		So(NewRune(454545).MapToNillableComplex64(func(rune) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyRune().MapToNillableComplex64(func(rune) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewRune(454545).MapToNillableComplex64(func(rune) *complex64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToComplex64(t *testing.T) {
	Convey("Rune.MapToComplex64", t, func() {
		test1, err1 := NewRune(454545).MapToComplex64(func(rune) (complex64, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyRune().MapToComplex64(func(rune) (complex64, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewRune(454545).MapToComplex64(func(rune) (complex64, error) {
			return 123456, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestRune_ForceMapToFloat32(t *testing.T) {
	Convey("Rune.ForceMapToFloat32", t, func() {
		So(NewRune(454545).ForceMapToFloat32(func(rune) float32 { return 32.3 }).Get(),
			ShouldEqual, 32.3)

		So(NewEmptyRune().ForceMapToFloat32(func(rune) float32 { return 32.3 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToNillableFloat32(t *testing.T) {
	Convey("Rune.MapToNillableFloat32", t, func() {
		So(NewRune(454545).MapToNillableFloat32(func(rune) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).Get(), ShouldEqual, 32.3)

		So(NewEmptyRune().MapToNillableFloat32(func(rune) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewRune(454545).MapToNillableFloat32(func(rune) *float32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToFloat32(t *testing.T) {
	Convey("Rune.MapToFloat32", t, func() {
		test1, err1 := NewRune(454545).MapToFloat32(func(rune) (float32, error) {
			return 32.3, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 32.3)

		test2, err2 := NewEmptyRune().MapToFloat32(func(rune) (float32, error) {
			return 32.3, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewRune(454545).MapToFloat32(func(rune) (float32, error) {
			return 32.3, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestRune_ForceMapToFloat64(t *testing.T) {
	Convey("Rune.ForceMapToFloat64", t, func() {
		So(NewRune(454545).ForceMapToFloat64(func(rune) float64 { return 3.2455 }).Get(),
			ShouldEqual, 3.2455)

		So(NewEmptyRune().ForceMapToFloat64(func(rune) float64 { return 3.2455 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToNillableFloat64(t *testing.T) {
	Convey("Rune.MapToNillableFloat64", t, func() {
		So(NewRune(454545).MapToNillableFloat64(func(rune) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).Get(), ShouldEqual, 3.2455)

		So(NewEmptyRune().MapToNillableFloat64(func(rune) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewRune(454545).MapToNillableFloat64(func(rune) *float64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToFloat64(t *testing.T) {
	Convey("Rune.MapToFloat64", t, func() {
		test1, err1 := NewRune(454545).MapToFloat64(func(rune) (float64, error) {
			return 3.2455, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 3.2455)

		test2, err2 := NewEmptyRune().MapToFloat64(func(rune) (float64, error) {
			return 3.2455, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewRune(454545).MapToFloat64(func(rune) (float64, error) {
			return 3.2455, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestRune_ForceMapToInt(t *testing.T) {
	Convey("Rune.ForceMapToInt", t, func() {
		So(NewRune(454545).ForceMapToInt(func(rune) int { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyRune().ForceMapToInt(func(rune) int { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToNillableInt(t *testing.T) {
	Convey("Rune.MapToNillableInt", t, func() {
		So(NewRune(454545).MapToNillableInt(func(rune) *int {
			var tmp int = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyRune().MapToNillableInt(func(rune) *int {
			var tmp int = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewRune(454545).MapToNillableInt(func(rune) *int { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToInt(t *testing.T) {
	Convey("Rune.MapToInt", t, func() {
		test1, err1 := NewRune(454545).MapToInt(func(rune) (int, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyRune().MapToInt(func(rune) (int, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewRune(454545).MapToInt(func(rune) (int, error) {
			return 123456, errors.New("int error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestRune_ForceMapToInt16(t *testing.T) {
	Convey("Rune.ForceMapToInt16", t, func() {
		So(NewRune(454545).ForceMapToInt16(func(rune) int16 { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyRune().ForceMapToInt16(func(rune) int16 { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToNillableInt16(t *testing.T) {
	Convey("Rune.MapToNillableInt16", t, func() {
		So(NewRune(454545).MapToNillableInt16(func(rune) *int16 {
			var tmp int16 = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyRune().MapToNillableInt16(func(rune) *int16 {
			var tmp int16 = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewRune(454545).MapToNillableInt16(func(rune) *int16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToInt16(t *testing.T) {
	Convey("Rune.MapToInt16", t, func() {
		test1, err1 := NewRune(454545).MapToInt16(func(rune) (int16, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyRune().MapToInt16(func(rune) (int16, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewRune(454545).MapToInt16(func(rune) (int16, error) {
			return 222, errors.New("int16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestRune_ForceMapToInt32(t *testing.T) {
	Convey("Rune.ForceMapToInt32", t, func() {
		So(NewRune(454545).ForceMapToInt32(func(rune) int32 { return 333 }).Get(),
			ShouldEqual, 333)

		So(NewEmptyRune().ForceMapToInt32(func(rune) int32 { return 333 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToNillableInt32(t *testing.T) {
	Convey("Rune.MapToNillableInt32", t, func() {
		So(NewRune(454545).MapToNillableInt32(func(rune) *int32 {
			var tmp int32 = 333
			return &tmp
		}).Get(), ShouldEqual, 333)

		So(NewEmptyRune().MapToNillableInt32(func(rune) *int32 {
			var tmp int32 = 333
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewRune(454545).MapToNillableInt32(func(rune) *int32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToInt32(t *testing.T) {
	Convey("Rune.MapToInt32", t, func() {
		test1, err1 := NewRune(454545).MapToInt32(func(rune) (int32, error) {
			return 333, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 333)

		test2, err2 := NewEmptyRune().MapToInt32(func(rune) (int32, error) {
			return 333, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewRune(454545).MapToInt32(func(rune) (int32, error) {
			return 333, errors.New("int32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestRune_ForceMapToInt64(t *testing.T) {
	Convey("Rune.ForceMapToInt64", t, func() {
		So(NewRune(454545).ForceMapToInt64(func(rune) int64 { return 444 }).Get(),
			ShouldEqual, 444)

		So(NewEmptyRune().ForceMapToInt64(func(rune) int64 { return 444 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToNillableInt64(t *testing.T) {
	Convey("Rune.MapToNillableInt64", t, func() {
		So(NewRune(454545).MapToNillableInt64(func(rune) *int64 {
			var tmp int64 = 444
			return &tmp
		}).Get(), ShouldEqual, 444)

		So(NewEmptyRune().MapToNillableInt64(func(rune) *int64 {
			var tmp int64 = 444
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewRune(454545).MapToNillableInt64(func(rune) *int64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToInt64(t *testing.T) {
	Convey("Rune.MapToInt64", t, func() {
		test1, err1 := NewRune(454545).MapToInt64(func(rune) (int64, error) {
			return 444, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 444)

		test2, err2 := NewEmptyRune().MapToInt64(func(rune) (int64, error) {
			return 444, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewRune(454545).MapToInt64(func(rune) (int64, error) {
			return 444, errors.New("int64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestRune_ForceMapToInt8(t *testing.T) {
	Convey("Rune.ForceMapToInt8", t, func() {
		So(NewRune(454545).ForceMapToInt8(func(rune) int8 { return 111 }).Get(),
			ShouldEqual, 111)

		So(NewEmptyRune().ForceMapToInt8(func(rune) int8 { return 111 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToNillableInt8(t *testing.T) {
	Convey("Rune.MapToNillableInt8", t, func() {
		So(NewRune(454545).MapToNillableInt8(func(rune) *int8 {
			var tmp int8 = 111
			return &tmp
		}).Get(), ShouldEqual, 111)

		So(NewEmptyRune().MapToNillableInt8(func(rune) *int8 {
			var tmp int8 = 111
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewRune(454545).MapToNillableInt8(func(rune) *int8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToInt8(t *testing.T) {
	Convey("Rune.MapToInt8", t, func() {
		test1, err1 := NewRune(454545).MapToInt8(func(rune) (int8, error) {
			return 111, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 111)

		test2, err2 := NewEmptyRune().MapToInt8(func(rune) (int8, error) {
			return 111, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewRune(454545).MapToInt8(func(rune) (int8, error) {
			return 111, errors.New("int8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestRune_ForceMapToUntyped(t *testing.T) {
	Convey("Rune.ForceMapToUntyped", t, func() {
		So(NewRune(454545).ForceMapToUntyped(func(rune) interface{} { return "111111" }).Get(),
			ShouldEqual, "111111")

		So(NewEmptyRune().ForceMapToUntyped(func(rune) interface{} { return "111111" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToNillableUntyped(t *testing.T) {
	Convey("Rune.MapToNillableUntyped", t, func() {
		So(NewRune(454545).MapToNillableUntyped(func(rune) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).Get(), ShouldEqual, "111111")

		So(NewEmptyRune().MapToNillableUntyped(func(rune) *interface{} {
			var tmp interface{} = "111111"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewRune(454545).MapToNillableUntyped(func(rune) *interface{} { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToUntyped(t *testing.T) {
	Convey("Rune.MapToUntyped", t, func() {
		test1, err1 := NewRune(454545).MapToUntyped(func(rune) (interface{}, error) {
			return "111111", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "111111")

		test2, err2 := NewEmptyRune().MapToUntyped(func(rune) (interface{}, error) {
			return "111111", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewRune(454545).MapToUntyped(func(rune) (interface{}, error) {
			return "111111", errors.New("interface error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "interface error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestRune_ForceMapToString(t *testing.T) {
	Convey("Rune.ForceMapToString", t, func() {
		So(NewRune(454545).ForceMapToString(func(rune) string { return "bonanza" }).Get(),
			ShouldEqual, "bonanza")

		So(NewEmptyRune().ForceMapToString(func(rune) string { return "bonanza" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToNillableString(t *testing.T) {
	Convey("Rune.MapToNillableString", t, func() {
		So(NewRune(454545).MapToNillableString(func(rune) *string {
			var tmp string = "bonanza"
			return &tmp
		}).Get(), ShouldEqual, "bonanza")

		So(NewEmptyRune().MapToNillableString(func(rune) *string {
			var tmp string = "bonanza"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewRune(454545).MapToNillableString(func(rune) *string { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToString(t *testing.T) {
	Convey("Rune.MapToString", t, func() {
		test1, err1 := NewRune(454545).MapToString(func(rune) (string, error) {
			return "bonanza", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "bonanza")

		test2, err2 := NewEmptyRune().MapToString(func(rune) (string, error) {
			return "bonanza", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewRune(454545).MapToString(func(rune) (string, error) {
			return "bonanza", errors.New("string error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "string error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestRune_ForceMapToUint(t *testing.T) {
	Convey("Rune.ForceMapToUint", t, func() {
		So(NewRune(454545).ForceMapToUint(func(rune) uint { return 666 }).Get(),
			ShouldEqual, 666)

		So(NewEmptyRune().ForceMapToUint(func(rune) uint { return 666 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToNillableUint(t *testing.T) {
	Convey("Rune.MapToNillableUint", t, func() {
		So(NewRune(454545).MapToNillableUint(func(rune) *uint {
			var tmp uint = 666
			return &tmp
		}).Get(), ShouldEqual, 666)

		So(NewEmptyRune().MapToNillableUint(func(rune) *uint {
			var tmp uint = 666
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewRune(454545).MapToNillableUint(func(rune) *uint { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToUint(t *testing.T) {
	Convey("Rune.MapToUint", t, func() {
		test1, err1 := NewRune(454545).MapToUint(func(rune) (uint, error) {
			return 666, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 666)

		test2, err2 := NewEmptyRune().MapToUint(func(rune) (uint, error) {
			return 666, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewRune(454545).MapToUint(func(rune) (uint, error) {
			return 666, errors.New("uint error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestRune_ForceMapToUint16(t *testing.T) {
	Convey("Rune.ForceMapToUint16", t, func() {
		So(NewRune(454545).ForceMapToUint16(func(rune) uint16 { return 888 }).Get(),
			ShouldEqual, 888)

		So(NewEmptyRune().ForceMapToUint16(func(rune) uint16 { return 888 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToNillableUint16(t *testing.T) {
	Convey("Rune.MapToNillableUint16", t, func() {
		So(NewRune(454545).MapToNillableUint16(func(rune) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).Get(), ShouldEqual, 888)

		So(NewEmptyRune().MapToNillableUint16(func(rune) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewRune(454545).MapToNillableUint16(func(rune) *uint16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToUint16(t *testing.T) {
	Convey("Rune.MapToUint16", t, func() {
		test1, err1 := NewRune(454545).MapToUint16(func(rune) (uint16, error) {
			return 888, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 888)

		test2, err2 := NewEmptyRune().MapToUint16(func(rune) (uint16, error) {
			return 888, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewRune(454545).MapToUint16(func(rune) (uint16, error) {
			return 888, errors.New("uint16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestRune_ForceMapToUint32(t *testing.T) {
	Convey("Rune.ForceMapToUint32", t, func() {
		So(NewRune(454545).ForceMapToUint32(func(rune) uint32 { return 999 }).Get(),
			ShouldEqual, 999)

		So(NewEmptyRune().ForceMapToUint32(func(rune) uint32 { return 999 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToNillableUint32(t *testing.T) {
	Convey("Rune.MapToNillableUint32", t, func() {
		So(NewRune(454545).MapToNillableUint32(func(rune) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).Get(), ShouldEqual, 999)

		So(NewEmptyRune().MapToNillableUint32(func(rune) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewRune(454545).MapToNillableUint32(func(rune) *uint32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToUint32(t *testing.T) {
	Convey("Rune.MapToUint32", t, func() {
		test1, err1 := NewRune(454545).MapToUint32(func(rune) (uint32, error) {
			return 999, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 999)

		test2, err2 := NewEmptyRune().MapToUint32(func(rune) (uint32, error) {
			return 999, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewRune(454545).MapToUint32(func(rune) (uint32, error) {
			return 999, errors.New("uint32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestRune_ForceMapToUint64(t *testing.T) {
	Convey("Rune.ForceMapToUint64", t, func() {
		So(NewRune(454545).ForceMapToUint64(func(rune) uint64 { return 101010 }).Get(),
			ShouldEqual, 101010)

		So(NewEmptyRune().ForceMapToUint64(func(rune) uint64 { return 101010 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToNillableUint64(t *testing.T) {
	Convey("Rune.MapToNillableUint64", t, func() {
		So(NewRune(454545).MapToNillableUint64(func(rune) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).Get(), ShouldEqual, 101010)

		So(NewEmptyRune().MapToNillableUint64(func(rune) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewRune(454545).MapToNillableUint64(func(rune) *uint64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToUint64(t *testing.T) {
	Convey("Rune.MapToUint64", t, func() {
		test1, err1 := NewRune(454545).MapToUint64(func(rune) (uint64, error) {
			return 101010, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 101010)

		test2, err2 := NewEmptyRune().MapToUint64(func(rune) (uint64, error) {
			return 101010, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewRune(454545).MapToUint64(func(rune) (uint64, error) {
			return 101010, errors.New("uint64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestRune_ForceMapToUint8(t *testing.T) {
	Convey("Rune.ForceMapToUint8", t, func() {
		So(NewRune(454545).ForceMapToUint8(func(rune) uint8 { return 77 }).Get(),
			ShouldEqual, 77)

		So(NewEmptyRune().ForceMapToUint8(func(rune) uint8 { return 77 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToNillableUint8(t *testing.T) {
	Convey("Rune.MapToNillableUint8", t, func() {
		So(NewRune(454545).MapToNillableUint8(func(rune) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).Get(), ShouldEqual, 77)

		So(NewEmptyRune().MapToNillableUint8(func(rune) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewRune(454545).MapToNillableUint8(func(rune) *uint8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestRune_MapToUint8(t *testing.T) {
	Convey("Rune.MapToUint8", t, func() {
		test1, err1 := NewRune(454545).MapToUint8(func(rune) (uint8, error) {
			return 77, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 77)

		test2, err2 := NewEmptyRune().MapToUint8(func(rune) (uint8, error) {
			return 77, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewRune(454545).MapToUint8(func(rune) (uint8, error) {
			return 77, errors.New("uint8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}
