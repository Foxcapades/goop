// Package option_test generated @ 2020-05-17 22:19:55.164878293 -0400 EDT m=+0.001439508
package option_test

import (
	"errors"
	"testing"

	. "github.com/Foxcapades/goop/v1/pkg/option"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewUntyped(t *testing.T) {
	Convey("NewUntyped", t, func() {
		test := NewUntyped("clam")
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmptyUntyped(t *testing.T) {
	Convey("NewEmptyUntyped", t, func() {
		test := NewEmptyUntyped()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybeUntyped(t *testing.T) {
	Convey("NewMaybeUntyped", t, func() {
		var val1 interface{}

		test1 := NewMaybeUntyped(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybeUntyped(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func TestUntyped_Get(t *testing.T) {
	Convey("Untyped.Get", t, func() {
		test1 := NewUntyped("clam")
		So(test1.Get(), ShouldEqual, "clam")

		test2 := NewEmptyUntyped()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func TestUntyped_Or(t *testing.T) {
	Convey("Untyped.Or", t, func() {
		test1 := NewUntyped("clam")
		So(test1.Or(69), ShouldEqual, "clam")

		test2 := NewEmptyUntyped()
		So(test2.Or("clam"), ShouldEqual, "clam")
	})
}

func TestUntyped_OrGet(t *testing.T) {
	Convey("Untyped.OrGet", t, func() {
		test1 := NewUntyped("clam")
		So(test1.OrGet(func() interface{} { return 69 }), ShouldEqual, "clam")

		test2 := NewEmptyUntyped()
		So(test2.OrGet(func() interface{} { return "clam" }), ShouldEqual, "clam")
	})
}

func TestUntyped_OrPanicWith(t *testing.T) {
	Convey("Untyped.OrPanicWith", t, func() {
		test1 := NewUntyped("clam")
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, "clam")

		test2 := NewEmptyUntyped()
		So(func() {
			test2.OrPanicWith(func() interface{} { return "test value" })
		}, ShouldPanicWith, "test value")
	})
}

func TestUntyped_MapToNillable(t *testing.T) {
	Convey("Untyped.MapToNillable", t, func() {
		test1 := NewUntyped("clam")
		So(test1.MapToNillable(func(b interface{}) *interface{} { return &b }).IsPresent(), ShouldBeTrue)

		test2 := NewUntyped("clam")
		So(test2.MapToNillable(func(b interface{}) *interface{} { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmptyUntyped()
		So(func() {
			test3.MapToNillable(func(b interface{}) *interface{} { panic("foo") }).IsNil()
		}, ShouldNotPanic)
	})
}

func TestUntyped_ForceMapToBool(t *testing.T) {
	Convey("Untyped.ForceMapToBool", t, func() {
		So(NewUntyped("clam").ForceMapToBool(func(interface{}) bool { return true }).Get(),
			ShouldEqual, true)

		So(NewEmptyUntyped().ForceMapToBool(func(interface{}) bool { return true }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToNillableBool(t *testing.T) {
	Convey("Untyped.MapToNillableBool", t, func() {
		So(NewUntyped("clam").MapToNillableBool(func(interface{}) *bool {
			var tmp bool = true
			return &tmp
		}).Get(), ShouldEqual, true)

		So(NewEmptyUntyped().MapToNillableBool(func(interface{}) *bool {
			var tmp bool = true
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUntyped("clam").MapToNillableBool(func(interface{}) *bool { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToBool(t *testing.T) {
	Convey("Untyped.MapToBool", t, func() {
		test1, err1 := NewUntyped("clam").MapToBool(func(interface{}) (bool, error) {
			return true, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, true)

		test2, err2 := NewEmptyUntyped().MapToBool(func(interface{}) (bool, error) {
			return true, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUntyped("clam").MapToBool(func(interface{}) (bool, error) {
			return true, errors.New("bool error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "bool error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUntyped_ForceMapToByte(t *testing.T) {
	Convey("Untyped.ForceMapToByte", t, func() {
		So(NewUntyped("clam").ForceMapToByte(func(interface{}) byte { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyUntyped().ForceMapToByte(func(interface{}) byte { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToNillableByte(t *testing.T) {
	Convey("Untyped.MapToNillableByte", t, func() {
		So(NewUntyped("clam").MapToNillableByte(func(interface{}) *byte {
			var tmp byte = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyUntyped().MapToNillableByte(func(interface{}) *byte {
			var tmp byte = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUntyped("clam").MapToNillableByte(func(interface{}) *byte { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToByte(t *testing.T) {
	Convey("Untyped.MapToByte", t, func() {
		test1, err1 := NewUntyped("clam").MapToByte(func(interface{}) (byte, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyUntyped().MapToByte(func(interface{}) (byte, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUntyped("clam").MapToByte(func(interface{}) (byte, error) {
			return 222, errors.New("byte error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "byte error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUntyped_ForceMapToComplex128(t *testing.T) {
	Convey("Untyped.ForceMapToComplex128", t, func() {
		So(NewUntyped("clam").ForceMapToComplex128(func(interface{}) complex128 { return 987654 }).Get(),
			ShouldEqual, 987654)

		So(NewEmptyUntyped().ForceMapToComplex128(func(interface{}) complex128 { return 987654 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToNillableComplex128(t *testing.T) {
	Convey("Untyped.MapToNillableComplex128", t, func() {
		So(NewUntyped("clam").MapToNillableComplex128(func(interface{}) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).Get(), ShouldEqual, 987654)

		So(NewEmptyUntyped().MapToNillableComplex128(func(interface{}) *complex128 {
			var tmp complex128 = 987654
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUntyped("clam").MapToNillableComplex128(func(interface{}) *complex128 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToComplex128(t *testing.T) {
	Convey("Untyped.MapToComplex128", t, func() {
		test1, err1 := NewUntyped("clam").MapToComplex128(func(interface{}) (complex128, error) {
			return 987654, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 987654)

		test2, err2 := NewEmptyUntyped().MapToComplex128(func(interface{}) (complex128, error) {
			return 987654, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUntyped("clam").MapToComplex128(func(interface{}) (complex128, error) {
			return 987654, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUntyped_ForceMapToComplex64(t *testing.T) {
	Convey("Untyped.ForceMapToComplex64", t, func() {
		So(NewUntyped("clam").ForceMapToComplex64(func(interface{}) complex64 { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyUntyped().ForceMapToComplex64(func(interface{}) complex64 { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToNillableComplex64(t *testing.T) {
	Convey("Untyped.MapToNillableComplex64", t, func() {
		So(NewUntyped("clam").MapToNillableComplex64(func(interface{}) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyUntyped().MapToNillableComplex64(func(interface{}) *complex64 {
			var tmp complex64 = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUntyped("clam").MapToNillableComplex64(func(interface{}) *complex64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToComplex64(t *testing.T) {
	Convey("Untyped.MapToComplex64", t, func() {
		test1, err1 := NewUntyped("clam").MapToComplex64(func(interface{}) (complex64, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyUntyped().MapToComplex64(func(interface{}) (complex64, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUntyped("clam").MapToComplex64(func(interface{}) (complex64, error) {
			return 123456, errors.New("complex error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "complex error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUntyped_ForceMapToFloat32(t *testing.T) {
	Convey("Untyped.ForceMapToFloat32", t, func() {
		So(NewUntyped("clam").ForceMapToFloat32(func(interface{}) float32 { return 32.3 }).Get(),
			ShouldEqual, 32.3)

		So(NewEmptyUntyped().ForceMapToFloat32(func(interface{}) float32 { return 32.3 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToNillableFloat32(t *testing.T) {
	Convey("Untyped.MapToNillableFloat32", t, func() {
		So(NewUntyped("clam").MapToNillableFloat32(func(interface{}) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).Get(), ShouldEqual, 32.3)

		So(NewEmptyUntyped().MapToNillableFloat32(func(interface{}) *float32 {
			var tmp float32 = 32.3
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUntyped("clam").MapToNillableFloat32(func(interface{}) *float32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToFloat32(t *testing.T) {
	Convey("Untyped.MapToFloat32", t, func() {
		test1, err1 := NewUntyped("clam").MapToFloat32(func(interface{}) (float32, error) {
			return 32.3, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 32.3)

		test2, err2 := NewEmptyUntyped().MapToFloat32(func(interface{}) (float32, error) {
			return 32.3, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUntyped("clam").MapToFloat32(func(interface{}) (float32, error) {
			return 32.3, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUntyped_ForceMapToFloat64(t *testing.T) {
	Convey("Untyped.ForceMapToFloat64", t, func() {
		So(NewUntyped("clam").ForceMapToFloat64(func(interface{}) float64 { return 3.2455 }).Get(),
			ShouldEqual, 3.2455)

		So(NewEmptyUntyped().ForceMapToFloat64(func(interface{}) float64 { return 3.2455 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToNillableFloat64(t *testing.T) {
	Convey("Untyped.MapToNillableFloat64", t, func() {
		So(NewUntyped("clam").MapToNillableFloat64(func(interface{}) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).Get(), ShouldEqual, 3.2455)

		So(NewEmptyUntyped().MapToNillableFloat64(func(interface{}) *float64 {
			var tmp float64 = 3.2455
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUntyped("clam").MapToNillableFloat64(func(interface{}) *float64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToFloat64(t *testing.T) {
	Convey("Untyped.MapToFloat64", t, func() {
		test1, err1 := NewUntyped("clam").MapToFloat64(func(interface{}) (float64, error) {
			return 3.2455, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 3.2455)

		test2, err2 := NewEmptyUntyped().MapToFloat64(func(interface{}) (float64, error) {
			return 3.2455, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUntyped("clam").MapToFloat64(func(interface{}) (float64, error) {
			return 3.2455, errors.New("float error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "float error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUntyped_ForceMapToInt(t *testing.T) {
	Convey("Untyped.ForceMapToInt", t, func() {
		So(NewUntyped("clam").ForceMapToInt(func(interface{}) int { return 123456 }).Get(),
			ShouldEqual, 123456)

		So(NewEmptyUntyped().ForceMapToInt(func(interface{}) int { return 123456 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToNillableInt(t *testing.T) {
	Convey("Untyped.MapToNillableInt", t, func() {
		So(NewUntyped("clam").MapToNillableInt(func(interface{}) *int {
			var tmp int = 123456
			return &tmp
		}).Get(), ShouldEqual, 123456)

		So(NewEmptyUntyped().MapToNillableInt(func(interface{}) *int {
			var tmp int = 123456
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUntyped("clam").MapToNillableInt(func(interface{}) *int { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToInt(t *testing.T) {
	Convey("Untyped.MapToInt", t, func() {
		test1, err1 := NewUntyped("clam").MapToInt(func(interface{}) (int, error) {
			return 123456, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 123456)

		test2, err2 := NewEmptyUntyped().MapToInt(func(interface{}) (int, error) {
			return 123456, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUntyped("clam").MapToInt(func(interface{}) (int, error) {
			return 123456, errors.New("int error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUntyped_ForceMapToInt16(t *testing.T) {
	Convey("Untyped.ForceMapToInt16", t, func() {
		So(NewUntyped("clam").ForceMapToInt16(func(interface{}) int16 { return 222 }).Get(),
			ShouldEqual, 222)

		So(NewEmptyUntyped().ForceMapToInt16(func(interface{}) int16 { return 222 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToNillableInt16(t *testing.T) {
	Convey("Untyped.MapToNillableInt16", t, func() {
		So(NewUntyped("clam").MapToNillableInt16(func(interface{}) *int16 {
			var tmp int16 = 222
			return &tmp
		}).Get(), ShouldEqual, 222)

		So(NewEmptyUntyped().MapToNillableInt16(func(interface{}) *int16 {
			var tmp int16 = 222
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUntyped("clam").MapToNillableInt16(func(interface{}) *int16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToInt16(t *testing.T) {
	Convey("Untyped.MapToInt16", t, func() {
		test1, err1 := NewUntyped("clam").MapToInt16(func(interface{}) (int16, error) {
			return 222, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 222)

		test2, err2 := NewEmptyUntyped().MapToInt16(func(interface{}) (int16, error) {
			return 222, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUntyped("clam").MapToInt16(func(interface{}) (int16, error) {
			return 222, errors.New("int16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUntyped_ForceMapToInt32(t *testing.T) {
	Convey("Untyped.ForceMapToInt32", t, func() {
		So(NewUntyped("clam").ForceMapToInt32(func(interface{}) int32 { return 333 }).Get(),
			ShouldEqual, 333)

		So(NewEmptyUntyped().ForceMapToInt32(func(interface{}) int32 { return 333 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToNillableInt32(t *testing.T) {
	Convey("Untyped.MapToNillableInt32", t, func() {
		So(NewUntyped("clam").MapToNillableInt32(func(interface{}) *int32 {
			var tmp int32 = 333
			return &tmp
		}).Get(), ShouldEqual, 333)

		So(NewEmptyUntyped().MapToNillableInt32(func(interface{}) *int32 {
			var tmp int32 = 333
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUntyped("clam").MapToNillableInt32(func(interface{}) *int32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToInt32(t *testing.T) {
	Convey("Untyped.MapToInt32", t, func() {
		test1, err1 := NewUntyped("clam").MapToInt32(func(interface{}) (int32, error) {
			return 333, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 333)

		test2, err2 := NewEmptyUntyped().MapToInt32(func(interface{}) (int32, error) {
			return 333, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUntyped("clam").MapToInt32(func(interface{}) (int32, error) {
			return 333, errors.New("int32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUntyped_ForceMapToInt64(t *testing.T) {
	Convey("Untyped.ForceMapToInt64", t, func() {
		So(NewUntyped("clam").ForceMapToInt64(func(interface{}) int64 { return 444 }).Get(),
			ShouldEqual, 444)

		So(NewEmptyUntyped().ForceMapToInt64(func(interface{}) int64 { return 444 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToNillableInt64(t *testing.T) {
	Convey("Untyped.MapToNillableInt64", t, func() {
		So(NewUntyped("clam").MapToNillableInt64(func(interface{}) *int64 {
			var tmp int64 = 444
			return &tmp
		}).Get(), ShouldEqual, 444)

		So(NewEmptyUntyped().MapToNillableInt64(func(interface{}) *int64 {
			var tmp int64 = 444
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUntyped("clam").MapToNillableInt64(func(interface{}) *int64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToInt64(t *testing.T) {
	Convey("Untyped.MapToInt64", t, func() {
		test1, err1 := NewUntyped("clam").MapToInt64(func(interface{}) (int64, error) {
			return 444, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 444)

		test2, err2 := NewEmptyUntyped().MapToInt64(func(interface{}) (int64, error) {
			return 444, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUntyped("clam").MapToInt64(func(interface{}) (int64, error) {
			return 444, errors.New("int64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUntyped_ForceMapToInt8(t *testing.T) {
	Convey("Untyped.ForceMapToInt8", t, func() {
		So(NewUntyped("clam").ForceMapToInt8(func(interface{}) int8 { return 111 }).Get(),
			ShouldEqual, 111)

		So(NewEmptyUntyped().ForceMapToInt8(func(interface{}) int8 { return 111 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToNillableInt8(t *testing.T) {
	Convey("Untyped.MapToNillableInt8", t, func() {
		So(NewUntyped("clam").MapToNillableInt8(func(interface{}) *int8 {
			var tmp int8 = 111
			return &tmp
		}).Get(), ShouldEqual, 111)

		So(NewEmptyUntyped().MapToNillableInt8(func(interface{}) *int8 {
			var tmp int8 = 111
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUntyped("clam").MapToNillableInt8(func(interface{}) *int8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToInt8(t *testing.T) {
	Convey("Untyped.MapToInt8", t, func() {
		test1, err1 := NewUntyped("clam").MapToInt8(func(interface{}) (int8, error) {
			return 111, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 111)

		test2, err2 := NewEmptyUntyped().MapToInt8(func(interface{}) (int8, error) {
			return 111, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUntyped("clam").MapToInt8(func(interface{}) (int8, error) {
			return 111, errors.New("int8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "int8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUntyped_ForceMapToRune(t *testing.T) {
	Convey("Untyped.ForceMapToRune", t, func() {
		So(NewUntyped("clam").ForceMapToRune(func(interface{}) rune { return 555 }).Get(),
			ShouldEqual, 555)

		So(NewEmptyUntyped().ForceMapToRune(func(interface{}) rune { return 555 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToNillableRune(t *testing.T) {
	Convey("Untyped.MapToNillableRune", t, func() {
		So(NewUntyped("clam").MapToNillableRune(func(interface{}) *rune {
			var tmp rune = 555
			return &tmp
		}).Get(), ShouldEqual, 555)

		So(NewEmptyUntyped().MapToNillableRune(func(interface{}) *rune {
			var tmp rune = 555
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUntyped("clam").MapToNillableRune(func(interface{}) *rune { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToRune(t *testing.T) {
	Convey("Untyped.MapToRune", t, func() {
		test1, err1 := NewUntyped("clam").MapToRune(func(interface{}) (rune, error) {
			return 555, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 555)

		test2, err2 := NewEmptyUntyped().MapToRune(func(interface{}) (rune, error) {
			return 555, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUntyped("clam").MapToRune(func(interface{}) (rune, error) {
			return 555, errors.New("rune error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "rune error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUntyped_ForceMapToString(t *testing.T) {
	Convey("Untyped.ForceMapToString", t, func() {
		So(NewUntyped("clam").ForceMapToString(func(interface{}) string { return "bonanza" }).Get(),
			ShouldEqual, "bonanza")

		So(NewEmptyUntyped().ForceMapToString(func(interface{}) string { return "bonanza" }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToNillableString(t *testing.T) {
	Convey("Untyped.MapToNillableString", t, func() {
		So(NewUntyped("clam").MapToNillableString(func(interface{}) *string {
			var tmp string = "bonanza"
			return &tmp
		}).Get(), ShouldEqual, "bonanza")

		So(NewEmptyUntyped().MapToNillableString(func(interface{}) *string {
			var tmp string = "bonanza"
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUntyped("clam").MapToNillableString(func(interface{}) *string { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToString(t *testing.T) {
	Convey("Untyped.MapToString", t, func() {
		test1, err1 := NewUntyped("clam").MapToString(func(interface{}) (string, error) {
			return "bonanza", nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, "bonanza")

		test2, err2 := NewEmptyUntyped().MapToString(func(interface{}) (string, error) {
			return "bonanza", nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUntyped("clam").MapToString(func(interface{}) (string, error) {
			return "bonanza", errors.New("string error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "string error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUntyped_ForceMapToUint(t *testing.T) {
	Convey("Untyped.ForceMapToUint", t, func() {
		So(NewUntyped("clam").ForceMapToUint(func(interface{}) uint { return 666 }).Get(),
			ShouldEqual, 666)

		So(NewEmptyUntyped().ForceMapToUint(func(interface{}) uint { return 666 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToNillableUint(t *testing.T) {
	Convey("Untyped.MapToNillableUint", t, func() {
		So(NewUntyped("clam").MapToNillableUint(func(interface{}) *uint {
			var tmp uint = 666
			return &tmp
		}).Get(), ShouldEqual, 666)

		So(NewEmptyUntyped().MapToNillableUint(func(interface{}) *uint {
			var tmp uint = 666
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUntyped("clam").MapToNillableUint(func(interface{}) *uint { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToUint(t *testing.T) {
	Convey("Untyped.MapToUint", t, func() {
		test1, err1 := NewUntyped("clam").MapToUint(func(interface{}) (uint, error) {
			return 666, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 666)

		test2, err2 := NewEmptyUntyped().MapToUint(func(interface{}) (uint, error) {
			return 666, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUntyped("clam").MapToUint(func(interface{}) (uint, error) {
			return 666, errors.New("uint error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUntyped_ForceMapToUint16(t *testing.T) {
	Convey("Untyped.ForceMapToUint16", t, func() {
		So(NewUntyped("clam").ForceMapToUint16(func(interface{}) uint16 { return 888 }).Get(),
			ShouldEqual, 888)

		So(NewEmptyUntyped().ForceMapToUint16(func(interface{}) uint16 { return 888 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToNillableUint16(t *testing.T) {
	Convey("Untyped.MapToNillableUint16", t, func() {
		So(NewUntyped("clam").MapToNillableUint16(func(interface{}) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).Get(), ShouldEqual, 888)

		So(NewEmptyUntyped().MapToNillableUint16(func(interface{}) *uint16 {
			var tmp uint16 = 888
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUntyped("clam").MapToNillableUint16(func(interface{}) *uint16 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToUint16(t *testing.T) {
	Convey("Untyped.MapToUint16", t, func() {
		test1, err1 := NewUntyped("clam").MapToUint16(func(interface{}) (uint16, error) {
			return 888, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 888)

		test2, err2 := NewEmptyUntyped().MapToUint16(func(interface{}) (uint16, error) {
			return 888, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUntyped("clam").MapToUint16(func(interface{}) (uint16, error) {
			return 888, errors.New("uint16 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint16 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUntyped_ForceMapToUint32(t *testing.T) {
	Convey("Untyped.ForceMapToUint32", t, func() {
		So(NewUntyped("clam").ForceMapToUint32(func(interface{}) uint32 { return 999 }).Get(),
			ShouldEqual, 999)

		So(NewEmptyUntyped().ForceMapToUint32(func(interface{}) uint32 { return 999 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToNillableUint32(t *testing.T) {
	Convey("Untyped.MapToNillableUint32", t, func() {
		So(NewUntyped("clam").MapToNillableUint32(func(interface{}) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).Get(), ShouldEqual, 999)

		So(NewEmptyUntyped().MapToNillableUint32(func(interface{}) *uint32 {
			var tmp uint32 = 999
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUntyped("clam").MapToNillableUint32(func(interface{}) *uint32 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToUint32(t *testing.T) {
	Convey("Untyped.MapToUint32", t, func() {
		test1, err1 := NewUntyped("clam").MapToUint32(func(interface{}) (uint32, error) {
			return 999, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 999)

		test2, err2 := NewEmptyUntyped().MapToUint32(func(interface{}) (uint32, error) {
			return 999, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUntyped("clam").MapToUint32(func(interface{}) (uint32, error) {
			return 999, errors.New("uint32 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint32 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUntyped_ForceMapToUint64(t *testing.T) {
	Convey("Untyped.ForceMapToUint64", t, func() {
		So(NewUntyped("clam").ForceMapToUint64(func(interface{}) uint64 { return 101010 }).Get(),
			ShouldEqual, 101010)

		So(NewEmptyUntyped().ForceMapToUint64(func(interface{}) uint64 { return 101010 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToNillableUint64(t *testing.T) {
	Convey("Untyped.MapToNillableUint64", t, func() {
		So(NewUntyped("clam").MapToNillableUint64(func(interface{}) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).Get(), ShouldEqual, 101010)

		So(NewEmptyUntyped().MapToNillableUint64(func(interface{}) *uint64 {
			var tmp uint64 = 101010
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUntyped("clam").MapToNillableUint64(func(interface{}) *uint64 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToUint64(t *testing.T) {
	Convey("Untyped.MapToUint64", t, func() {
		test1, err1 := NewUntyped("clam").MapToUint64(func(interface{}) (uint64, error) {
			return 101010, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 101010)

		test2, err2 := NewEmptyUntyped().MapToUint64(func(interface{}) (uint64, error) {
			return 101010, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUntyped("clam").MapToUint64(func(interface{}) (uint64, error) {
			return 101010, errors.New("uint64 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint64 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}

func TestUntyped_ForceMapToUint8(t *testing.T) {
	Convey("Untyped.ForceMapToUint8", t, func() {
		So(NewUntyped("clam").ForceMapToUint8(func(interface{}) uint8 { return 77 }).Get(),
			ShouldEqual, 77)

		So(NewEmptyUntyped().ForceMapToUint8(func(interface{}) uint8 { return 77 }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToNillableUint8(t *testing.T) {
	Convey("Untyped.MapToNillableUint8", t, func() {
		So(NewUntyped("clam").MapToNillableUint8(func(interface{}) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).Get(), ShouldEqual, 77)

		So(NewEmptyUntyped().MapToNillableUint8(func(interface{}) *uint8 {
			var tmp uint8 = 77
			return &tmp
		}).IsNil(), ShouldBeTrue)

		So(NewUntyped("clam").MapToNillableUint8(func(interface{}) *uint8 { return nil }).IsNil(),
			ShouldBeTrue)
	})
}

func TestUntyped_MapToUint8(t *testing.T) {
	Convey("Untyped.MapToUint8", t, func() {
		test1, err1 := NewUntyped("clam").MapToUint8(func(interface{}) (uint8, error) {
			return 77, nil
		})

		So(err1, ShouldBeNil)
		So(test1.Get(), ShouldEqual, 77)

		test2, err2 := NewEmptyUntyped().MapToUint8(func(interface{}) (uint8, error) {
			return 77, nil
		})

		So(err2, ShouldBeNil)
		So(test2.IsNil(), ShouldBeTrue)

		test3, err3 := NewUntyped("clam").MapToUint8(func(interface{}) (uint8, error) {
			return 77, errors.New("uint8 error")
		})

		So(err3, ShouldNotBeNil)
		So(err3.Error(), ShouldEqual, "uint8 error")
		So(test3.IsNil(), ShouldBeTrue)
	})
}
