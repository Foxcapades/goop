package option

/* Generated @ 2020-05-18T16:16:44.814634002-04:00 */

// NewInt8 creates a new int8 option wrapping the given value.
func NewInt8(val int8) Int8 {
	return Int8{&val}
}

// NewMaybeInt8 creates a new option from the given
// pointer.
//
// If the pointer is nil, the resulting option will be
// empty, else the option will contain the value the pointer
// points to.
//
// Note, the value will be copied when wrapped.
func NewMaybeInt8(val *int8) Int8 {
	tmp := Int8{}

	if !isNil(val) {
		t2 := *val
		tmp.value = &t2
	}

	return tmp
}

// NewEmptyInt8 creates a new empty int8 option.
func NewEmptyInt8() Int8 {
	return Int8{}
}

// Int8 wraps an int8 value that may or may not exist.
//
// The value inside this struct cannot be modified (except
// maybe through reflection).
type Int8 struct {
	value *int8
}

// IsNil returns whether or not this option is empty.
func (o Int8) IsNil() bool {
	return isNil(o.value)
}

// IsPresent returns whether or not this option has a value.
func (o Int8) IsPresent() bool {
	return !isNil(o.value)
}

// Or returns either the value stored in this option or the
// given default value.
func (o Int8) Or(def int8) int8 {
	if o.value == nil {
		return def
	}

	return *o.value
}

// OrGet returns either the value stored in this option or
// result of calling the given function.
func (o Int8) OrGet(fn func() int8) int8 {
	if o.value == nil {
		return fn()
	}

	return *o.value
}

// OrPanicWith returns the value stored in this option or
// panics with the value returned by the given function.
func (o Int8) OrPanicWith(fn func() interface{}) int8 {
	if o.value == nil {
		panic(fn())
	}

	return *o.value
}

// OrPanicWithVal returns the value stored in this option
// or panics with the given value.
func (o Int8) OrPanicWithVal(err interface{}) int8 {
	if o.value == nil {
		panic(err)
	}

	return *o.value
}

// Get returns the value stored in this option or panics
// with a default empty unwrap error message.
func (o Int8) Get() int8 {
	return o.OrPanicWithVal("attempted to unwrap an empty option of int8")
}

// MapToBool either returns an empty option of bool if
// this option is empty, or calls the given function and
// returns an option wrapping the result.
//
// If the given function returns an error, the returned
// bool option will be empty.
func (o Int8) MapToBool(fn func(int8) (bool, error)) (Bool, error) {
	if o.value == nil {
		return NewEmptyBool(), nil
	}

	val, err := fn(*o.value)
	if err != nil {
		return NewEmptyBool(), err
	}

	return NewBool(val), nil
}

// ForceMapToBool either returns an empty option of bool
// if this option is empty, or returns a new option of
// bool wrapping the result of the given function.
func (o Int8) ForceMapToBool(fn func(int8) bool) Bool {
	if o.value == nil {
		return NewEmptyBool()
	}

	return NewBool(fn(*o.value))
}

// MapToNillableBool either returns an empty option of bool
// if this option is empty, or returns a new bool option of either
// some or none based on the value returned by the function.
func (o Int8) MapToNillableBool(fn func(int8) *bool) Bool {
	if o.value == nil {
		return NewEmptyBool()
	}

	return NewMaybeBool(fn(*o.value))
}

// MapToByte either returns an empty option of byte if
// this option is empty, or calls the given function and
// returns an option wrapping the result.
//
// If the given function returns an error, the returned
// byte option will be empty.
func (o Int8) MapToByte(fn func(int8) (byte, error)) (Byte, error) {
	if o.value == nil {
		return NewEmptyByte(), nil
	}

	val, err := fn(*o.value)
	if err != nil {
		return NewEmptyByte(), err
	}

	return NewByte(val), nil
}

// ForceMapToByte either returns an empty option of byte
// if this option is empty, or returns a new option of
// byte wrapping the result of the given function.
func (o Int8) ForceMapToByte(fn func(int8) byte) Byte {
	if o.value == nil {
		return NewEmptyByte()
	}

	return NewByte(fn(*o.value))
}

// MapToNillableByte either returns an empty option of byte
// if this option is empty, or returns a new byte option of either
// some or none based on the value returned by the function.
func (o Int8) MapToNillableByte(fn func(int8) *byte) Byte {
	if o.value == nil {
		return NewEmptyByte()
	}

	return NewMaybeByte(fn(*o.value))
}

// MapToComplex128 either returns an empty option of complex128 if
// this option is empty, or calls the given function and
// returns an option wrapping the result.
//
// If the given function returns an error, the returned
// complex128 option will be empty.
func (o Int8) MapToComplex128(fn func(int8) (complex128, error)) (Complex128, error) {
	if o.value == nil {
		return NewEmptyComplex128(), nil
	}

	val, err := fn(*o.value)
	if err != nil {
		return NewEmptyComplex128(), err
	}

	return NewComplex128(val), nil
}

// ForceMapToComplex128 either returns an empty option of complex128
// if this option is empty, or returns a new option of
// complex128 wrapping the result of the given function.
func (o Int8) ForceMapToComplex128(fn func(int8) complex128) Complex128 {
	if o.value == nil {
		return NewEmptyComplex128()
	}

	return NewComplex128(fn(*o.value))
}

// MapToNillableComplex128 either returns an empty option of complex128
// if this option is empty, or returns a new complex128 option of either
// some or none based on the value returned by the function.
func (o Int8) MapToNillableComplex128(fn func(int8) *complex128) Complex128 {
	if o.value == nil {
		return NewEmptyComplex128()
	}

	return NewMaybeComplex128(fn(*o.value))
}

// MapToComplex64 either returns an empty option of complex64 if
// this option is empty, or calls the given function and
// returns an option wrapping the result.
//
// If the given function returns an error, the returned
// complex64 option will be empty.
func (o Int8) MapToComplex64(fn func(int8) (complex64, error)) (Complex64, error) {
	if o.value == nil {
		return NewEmptyComplex64(), nil
	}

	val, err := fn(*o.value)
	if err != nil {
		return NewEmptyComplex64(), err
	}

	return NewComplex64(val), nil
}

// ForceMapToComplex64 either returns an empty option of complex64
// if this option is empty, or returns a new option of
// complex64 wrapping the result of the given function.
func (o Int8) ForceMapToComplex64(fn func(int8) complex64) Complex64 {
	if o.value == nil {
		return NewEmptyComplex64()
	}

	return NewComplex64(fn(*o.value))
}

// MapToNillableComplex64 either returns an empty option of complex64
// if this option is empty, or returns a new complex64 option of either
// some or none based on the value returned by the function.
func (o Int8) MapToNillableComplex64(fn func(int8) *complex64) Complex64 {
	if o.value == nil {
		return NewEmptyComplex64()
	}

	return NewMaybeComplex64(fn(*o.value))
}

// MapToFloat32 either returns an empty option of float32 if
// this option is empty, or calls the given function and
// returns an option wrapping the result.
//
// If the given function returns an error, the returned
// float32 option will be empty.
func (o Int8) MapToFloat32(fn func(int8) (float32, error)) (Float32, error) {
	if o.value == nil {
		return NewEmptyFloat32(), nil
	}

	val, err := fn(*o.value)
	if err != nil {
		return NewEmptyFloat32(), err
	}

	return NewFloat32(val), nil
}

// ForceMapToFloat32 either returns an empty option of float32
// if this option is empty, or returns a new option of
// float32 wrapping the result of the given function.
func (o Int8) ForceMapToFloat32(fn func(int8) float32) Float32 {
	if o.value == nil {
		return NewEmptyFloat32()
	}

	return NewFloat32(fn(*o.value))
}

// MapToNillableFloat32 either returns an empty option of float32
// if this option is empty, or returns a new float32 option of either
// some or none based on the value returned by the function.
func (o Int8) MapToNillableFloat32(fn func(int8) *float32) Float32 {
	if o.value == nil {
		return NewEmptyFloat32()
	}

	return NewMaybeFloat32(fn(*o.value))
}

// MapToFloat64 either returns an empty option of float64 if
// this option is empty, or calls the given function and
// returns an option wrapping the result.
//
// If the given function returns an error, the returned
// float64 option will be empty.
func (o Int8) MapToFloat64(fn func(int8) (float64, error)) (Float64, error) {
	if o.value == nil {
		return NewEmptyFloat64(), nil
	}

	val, err := fn(*o.value)
	if err != nil {
		return NewEmptyFloat64(), err
	}

	return NewFloat64(val), nil
}

// ForceMapToFloat64 either returns an empty option of float64
// if this option is empty, or returns a new option of
// float64 wrapping the result of the given function.
func (o Int8) ForceMapToFloat64(fn func(int8) float64) Float64 {
	if o.value == nil {
		return NewEmptyFloat64()
	}

	return NewFloat64(fn(*o.value))
}

// MapToNillableFloat64 either returns an empty option of float64
// if this option is empty, or returns a new float64 option of either
// some or none based on the value returned by the function.
func (o Int8) MapToNillableFloat64(fn func(int8) *float64) Float64 {
	if o.value == nil {
		return NewEmptyFloat64()
	}

	return NewMaybeFloat64(fn(*o.value))
}

// MapToInt either returns an empty option of int if
// this option is empty, or calls the given function and
// returns an option wrapping the result.
//
// If the given function returns an error, the returned
// int option will be empty.
func (o Int8) MapToInt(fn func(int8) (int, error)) (Int, error) {
	if o.value == nil {
		return NewEmptyInt(), nil
	}

	val, err := fn(*o.value)
	if err != nil {
		return NewEmptyInt(), err
	}

	return NewInt(val), nil
}

// ForceMapToInt either returns an empty option of int
// if this option is empty, or returns a new option of
// int wrapping the result of the given function.
func (o Int8) ForceMapToInt(fn func(int8) int) Int {
	if o.value == nil {
		return NewEmptyInt()
	}

	return NewInt(fn(*o.value))
}

// MapToNillableInt either returns an empty option of int
// if this option is empty, or returns a new int option of either
// some or none based on the value returned by the function.
func (o Int8) MapToNillableInt(fn func(int8) *int) Int {
	if o.value == nil {
		return NewEmptyInt()
	}

	return NewMaybeInt(fn(*o.value))
}

// MapToInt16 either returns an empty option of int16 if
// this option is empty, or calls the given function and
// returns an option wrapping the result.
//
// If the given function returns an error, the returned
// int16 option will be empty.
func (o Int8) MapToInt16(fn func(int8) (int16, error)) (Int16, error) {
	if o.value == nil {
		return NewEmptyInt16(), nil
	}

	val, err := fn(*o.value)
	if err != nil {
		return NewEmptyInt16(), err
	}

	return NewInt16(val), nil
}

// ForceMapToInt16 either returns an empty option of int16
// if this option is empty, or returns a new option of
// int16 wrapping the result of the given function.
func (o Int8) ForceMapToInt16(fn func(int8) int16) Int16 {
	if o.value == nil {
		return NewEmptyInt16()
	}

	return NewInt16(fn(*o.value))
}

// MapToNillableInt16 either returns an empty option of int16
// if this option is empty, or returns a new int16 option of either
// some or none based on the value returned by the function.
func (o Int8) MapToNillableInt16(fn func(int8) *int16) Int16 {
	if o.value == nil {
		return NewEmptyInt16()
	}

	return NewMaybeInt16(fn(*o.value))
}

// MapToInt32 either returns an empty option of int32 if
// this option is empty, or calls the given function and
// returns an option wrapping the result.
//
// If the given function returns an error, the returned
// int32 option will be empty.
func (o Int8) MapToInt32(fn func(int8) (int32, error)) (Int32, error) {
	if o.value == nil {
		return NewEmptyInt32(), nil
	}

	val, err := fn(*o.value)
	if err != nil {
		return NewEmptyInt32(), err
	}

	return NewInt32(val), nil
}

// ForceMapToInt32 either returns an empty option of int32
// if this option is empty, or returns a new option of
// int32 wrapping the result of the given function.
func (o Int8) ForceMapToInt32(fn func(int8) int32) Int32 {
	if o.value == nil {
		return NewEmptyInt32()
	}

	return NewInt32(fn(*o.value))
}

// MapToNillableInt32 either returns an empty option of int32
// if this option is empty, or returns a new int32 option of either
// some or none based on the value returned by the function.
func (o Int8) MapToNillableInt32(fn func(int8) *int32) Int32 {
	if o.value == nil {
		return NewEmptyInt32()
	}

	return NewMaybeInt32(fn(*o.value))
}

// MapToInt64 either returns an empty option of int64 if
// this option is empty, or calls the given function and
// returns an option wrapping the result.
//
// If the given function returns an error, the returned
// int64 option will be empty.
func (o Int8) MapToInt64(fn func(int8) (int64, error)) (Int64, error) {
	if o.value == nil {
		return NewEmptyInt64(), nil
	}

	val, err := fn(*o.value)
	if err != nil {
		return NewEmptyInt64(), err
	}

	return NewInt64(val), nil
}

// ForceMapToInt64 either returns an empty option of int64
// if this option is empty, or returns a new option of
// int64 wrapping the result of the given function.
func (o Int8) ForceMapToInt64(fn func(int8) int64) Int64 {
	if o.value == nil {
		return NewEmptyInt64()
	}

	return NewInt64(fn(*o.value))
}

// MapToNillableInt64 either returns an empty option of int64
// if this option is empty, or returns a new int64 option of either
// some or none based on the value returned by the function.
func (o Int8) MapToNillableInt64(fn func(int8) *int64) Int64 {
	if o.value == nil {
		return NewEmptyInt64()
	}

	return NewMaybeInt64(fn(*o.value))
}

// MapToUntyped either returns an empty option of interface{} if
// this option is empty, or calls the given function and
// returns an option wrapping the result.
//
// If the given function returns an error, the returned
// interface{} option will be empty.
//
// Since "interface{}" is a nillable type, the resulting
// will be empty if the return value was nil.
func (o Int8) MapToUntyped(fn func(int8) (interface{}, error)) (Untyped, error) {
	if o.value == nil {
		return NewEmptyUntyped(), nil
	}

	val, err := fn(*o.value)
	if err != nil {
		return NewEmptyUntyped(), err
	}

	if isNil(val) {
		return NewEmptyUntyped(), nil
	}

	return NewUntyped(val), nil
}

// ForceMapToUntyped either returns an empty option of interface{}
// if this option is empty, or returns a new option of
// interface{} wrapping the result of the given function.
//
// Since "interface{}" is a nillable type, the resulting
// option will be empty if the return value was nil.
func (o Int8) ForceMapToUntyped(fn func(int8) interface{}) Untyped {
	if o.value == nil {
		return NewEmptyUntyped()
	}

	val := fn(*o.value)

	if isNil(val) {
		return NewEmptyUntyped()
	}

	return NewUntyped(val)
}

// MapToNillableUntyped either returns an empty option of interface{}
// if this option is empty, or returns a new interface{} option of either
// some or none based on the value returned by the function.
func (o Int8) MapToNillableUntyped(fn func(int8) *interface{}) Untyped {
	if o.value == nil {
		return NewEmptyUntyped()
	}

	return NewMaybeUntyped(fn(*o.value))
}

// MapToRune either returns an empty option of rune if
// this option is empty, or calls the given function and
// returns an option wrapping the result.
//
// If the given function returns an error, the returned
// rune option will be empty.
func (o Int8) MapToRune(fn func(int8) (rune, error)) (Rune, error) {
	if o.value == nil {
		return NewEmptyRune(), nil
	}

	val, err := fn(*o.value)
	if err != nil {
		return NewEmptyRune(), err
	}

	return NewRune(val), nil
}

// ForceMapToRune either returns an empty option of rune
// if this option is empty, or returns a new option of
// rune wrapping the result of the given function.
func (o Int8) ForceMapToRune(fn func(int8) rune) Rune {
	if o.value == nil {
		return NewEmptyRune()
	}

	return NewRune(fn(*o.value))
}

// MapToNillableRune either returns an empty option of rune
// if this option is empty, or returns a new rune option of either
// some or none based on the value returned by the function.
func (o Int8) MapToNillableRune(fn func(int8) *rune) Rune {
	if o.value == nil {
		return NewEmptyRune()
	}

	return NewMaybeRune(fn(*o.value))
}

// MapToString either returns an empty option of string if
// this option is empty, or calls the given function and
// returns an option wrapping the result.
//
// If the given function returns an error, the returned
// string option will be empty.
func (o Int8) MapToString(fn func(int8) (string, error)) (String, error) {
	if o.value == nil {
		return NewEmptyString(), nil
	}

	val, err := fn(*o.value)
	if err != nil {
		return NewEmptyString(), err
	}

	return NewString(val), nil
}

// ForceMapToString either returns an empty option of string
// if this option is empty, or returns a new option of
// string wrapping the result of the given function.
func (o Int8) ForceMapToString(fn func(int8) string) String {
	if o.value == nil {
		return NewEmptyString()
	}

	return NewString(fn(*o.value))
}

// MapToNillableString either returns an empty option of string
// if this option is empty, or returns a new string option of either
// some or none based on the value returned by the function.
func (o Int8) MapToNillableString(fn func(int8) *string) String {
	if o.value == nil {
		return NewEmptyString()
	}

	return NewMaybeString(fn(*o.value))
}

// MapToUint either returns an empty option of uint if
// this option is empty, or calls the given function and
// returns an option wrapping the result.
//
// If the given function returns an error, the returned
// uint option will be empty.
func (o Int8) MapToUint(fn func(int8) (uint, error)) (Uint, error) {
	if o.value == nil {
		return NewEmptyUint(), nil
	}

	val, err := fn(*o.value)
	if err != nil {
		return NewEmptyUint(), err
	}

	return NewUint(val), nil
}

// ForceMapToUint either returns an empty option of uint
// if this option is empty, or returns a new option of
// uint wrapping the result of the given function.
func (o Int8) ForceMapToUint(fn func(int8) uint) Uint {
	if o.value == nil {
		return NewEmptyUint()
	}

	return NewUint(fn(*o.value))
}

// MapToNillableUint either returns an empty option of uint
// if this option is empty, or returns a new uint option of either
// some or none based on the value returned by the function.
func (o Int8) MapToNillableUint(fn func(int8) *uint) Uint {
	if o.value == nil {
		return NewEmptyUint()
	}

	return NewMaybeUint(fn(*o.value))
}

// MapToUint16 either returns an empty option of uint16 if
// this option is empty, or calls the given function and
// returns an option wrapping the result.
//
// If the given function returns an error, the returned
// uint16 option will be empty.
func (o Int8) MapToUint16(fn func(int8) (uint16, error)) (Uint16, error) {
	if o.value == nil {
		return NewEmptyUint16(), nil
	}

	val, err := fn(*o.value)
	if err != nil {
		return NewEmptyUint16(), err
	}

	return NewUint16(val), nil
}

// ForceMapToUint16 either returns an empty option of uint16
// if this option is empty, or returns a new option of
// uint16 wrapping the result of the given function.
func (o Int8) ForceMapToUint16(fn func(int8) uint16) Uint16 {
	if o.value == nil {
		return NewEmptyUint16()
	}

	return NewUint16(fn(*o.value))
}

// MapToNillableUint16 either returns an empty option of uint16
// if this option is empty, or returns a new uint16 option of either
// some or none based on the value returned by the function.
func (o Int8) MapToNillableUint16(fn func(int8) *uint16) Uint16 {
	if o.value == nil {
		return NewEmptyUint16()
	}

	return NewMaybeUint16(fn(*o.value))
}

// MapToUint32 either returns an empty option of uint32 if
// this option is empty, or calls the given function and
// returns an option wrapping the result.
//
// If the given function returns an error, the returned
// uint32 option will be empty.
func (o Int8) MapToUint32(fn func(int8) (uint32, error)) (Uint32, error) {
	if o.value == nil {
		return NewEmptyUint32(), nil
	}

	val, err := fn(*o.value)
	if err != nil {
		return NewEmptyUint32(), err
	}

	return NewUint32(val), nil
}

// ForceMapToUint32 either returns an empty option of uint32
// if this option is empty, or returns a new option of
// uint32 wrapping the result of the given function.
func (o Int8) ForceMapToUint32(fn func(int8) uint32) Uint32 {
	if o.value == nil {
		return NewEmptyUint32()
	}

	return NewUint32(fn(*o.value))
}

// MapToNillableUint32 either returns an empty option of uint32
// if this option is empty, or returns a new uint32 option of either
// some or none based on the value returned by the function.
func (o Int8) MapToNillableUint32(fn func(int8) *uint32) Uint32 {
	if o.value == nil {
		return NewEmptyUint32()
	}

	return NewMaybeUint32(fn(*o.value))
}

// MapToUint64 either returns an empty option of uint64 if
// this option is empty, or calls the given function and
// returns an option wrapping the result.
//
// If the given function returns an error, the returned
// uint64 option will be empty.
func (o Int8) MapToUint64(fn func(int8) (uint64, error)) (Uint64, error) {
	if o.value == nil {
		return NewEmptyUint64(), nil
	}

	val, err := fn(*o.value)
	if err != nil {
		return NewEmptyUint64(), err
	}

	return NewUint64(val), nil
}

// ForceMapToUint64 either returns an empty option of uint64
// if this option is empty, or returns a new option of
// uint64 wrapping the result of the given function.
func (o Int8) ForceMapToUint64(fn func(int8) uint64) Uint64 {
	if o.value == nil {
		return NewEmptyUint64()
	}

	return NewUint64(fn(*o.value))
}

// MapToNillableUint64 either returns an empty option of uint64
// if this option is empty, or returns a new uint64 option of either
// some or none based on the value returned by the function.
func (o Int8) MapToNillableUint64(fn func(int8) *uint64) Uint64 {
	if o.value == nil {
		return NewEmptyUint64()
	}

	return NewMaybeUint64(fn(*o.value))
}

// MapToUint8 either returns an empty option of uint8 if
// this option is empty, or calls the given function and
// returns an option wrapping the result.
//
// If the given function returns an error, the returned
// uint8 option will be empty.
func (o Int8) MapToUint8(fn func(int8) (uint8, error)) (Uint8, error) {
	if o.value == nil {
		return NewEmptyUint8(), nil
	}

	val, err := fn(*o.value)
	if err != nil {
		return NewEmptyUint8(), err
	}

	return NewUint8(val), nil
}

// ForceMapToUint8 either returns an empty option of uint8
// if this option is empty, or returns a new option of
// uint8 wrapping the result of the given function.
func (o Int8) ForceMapToUint8(fn func(int8) uint8) Uint8 {
	if o.value == nil {
		return NewEmptyUint8()
	}

	return NewUint8(fn(*o.value))
}

// MapToNillableUint8 either returns an empty option of uint8
// if this option is empty, or returns a new uint8 option of either
// some or none based on the value returned by the function.
func (o Int8) MapToNillableUint8(fn func(int8) *uint8) Uint8 {
	if o.value == nil {
		return NewEmptyUint8()
	}

	return NewMaybeUint8(fn(*o.value))
}

// MapToNillable either returns an empty option of int8
// if this option is empty, or returns a new int8 option of either
// some or none based on the value returned by the function.
func (o Int8) MapToNillable(fn func(int8) *int8) Int8 {
	if o.value == nil {
		return NewEmptyInt8()
	}

	return NewMaybeInt8(fn(*o.value))
}
