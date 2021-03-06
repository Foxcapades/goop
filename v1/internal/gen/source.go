package gen

// Template defines the structure for the generated option type files.
const Template = `package option

/* Generated @ {{ .Time }} */

// New{{ .Name }} creates a new {{ .Type }} option wrapping the given value.
{{ if eq .Type "interface{}" }}//
// If the value given to this function is nil, this function will panic.
{{end}}func New{{ .Name }}(val {{ .Type }}) {{ .Name }} {
{{ if eq .Type "interface{}" }}	if isNil(val) {
		panic("attempted to create a some option from a nil value")
	}

{{end}}	return {{ .Name }}{&val}
}

// NewMaybe{{ .Name }} creates a new option from the given
// pointer.
//
// If the pointer is nil, the resulting option will be
// empty, else the option will contain the value the pointer
// points to.
//
// Note, the value will be copied when wrapped.
func NewMaybe{{ .Name }}(val *{{.Type}}) {{ .Name }} {
	tmp := {{ .Name }}{}

	if !isNil(val) {
		t2 := *val
		tmp.value = &t2
	}

	return tmp
}

// NewEmpty{{ .Name }} creates a new empty {{ .Type }} option.
func NewEmpty{{ .Name }}() {{ .Name }} {
	return {{ .Name }}{}
}

// {{ .Name }} wraps an {{ .Type }} value that may or may not exist.
//
// The value inside this struct cannot be modified (except
// maybe through reflection).
type {{ .Name }} struct {
	value *{{ .Type }}
}

// IsNil returns whether or not this option is empty.
func (o {{ .Name }}) IsNil() bool {
	return isNil(o.value)
}

// IsPresent returns whether or not this option has a value.
func (o {{ .Name }}) IsPresent() bool {
	return !isNil(o.value)
}

// Or returns either the value stored in this option or the
// given default value.
func (o {{ .Name }}) Or(def {{ .Type }}) {{ .Type }} {
	if o.value == nil {
		return def
	}

	return *o.value
}

// OrGet returns either the value stored in this option or
// result of calling the given function.
func (o {{ .Name }}) OrGet(fn func() {{ .Type }}) {{ .Type }} {
	if o.value == nil {
		return fn()
	}

	return *o.value
}

// OrPanicWith returns the value stored in this option or
// panics with the value returned by the given function.
func (o {{ .Name }}) OrPanicWith(fn func() interface{}) {{ .Type }} {
	if o.value == nil {
		panic(fn())
	}

	return *o.value
}

// OrPanicWithVal returns the value stored in this option
// or panics with the given value.
func (o {{ .Name }}) OrPanicWithVal(err interface{}) {{ .Type }} {
	if o.value == nil {
		panic(err)
	}

	return *o.value
}

// Get returns the value stored in this option or panics
// with a default empty unwrap error message.
func (o {{ .Name }}) Get() {{ .Type }} {
	return o.OrPanicWithVal("attempted to unwrap an empty option of {{ .Type }}")
}
{{ range $AltType, $AltName := .Alternates }}
// MapTo{{ $AltName }} either returns an empty option of {{ $AltType }} if
// this option is empty, or calls the given function and
// returns an option wrapping the result.
//
// If the given function returns an error, the returned
// {{ $AltType }} option will be empty.{{ if eq $AltType "interface{}" }}
//
// Since "interface{}" is a nillable type, the resulting
// will be empty if the return value was nil.{{ end }}
func (o {{ $.Name }}) MapTo{{ $AltName }}(fn func({{ $.Type }}) ({{ $AltType }}, error)) ({{ $AltName }}, error) {
	if o.value == nil {
		return NewEmpty{{ $AltName }}(), nil
	}

	val, err := fn(*o.value)
	if err != nil {
		return NewEmpty{{ $AltName }}(), err
	}
{{ if eq $AltType "interface{}" }}
	if isNil(val) {
		return NewEmpty{{ $AltName }}(), nil
	}
{{ end }}
	return New{{ $AltName }}(val), nil
}

// ForceMapTo{{ $AltName }} either returns an empty option of {{ $AltType }}
// if this option is empty, or returns a new option of
// {{ $AltType }} wrapping the result of the given function.{{ if eq $AltType "interface{}" }}
//
// Since "interface{}" is a nillable type, the resulting
// option will be empty if the return value was nil.{{ end }}
func (o {{ $.Name }}) ForceMapTo{{ $AltName }}(fn func({{ $.Type }}) {{ $AltType }}) {{ $AltName }} {
	if o.value == nil {
		return NewEmpty{{ $AltName }}()
	}
{{ if eq $AltType "interface{}" }}
	val := fn(*o.value)

	if isNil(val) {
		return NewEmpty{{ $AltName }}()
	}

	return New{{ $AltName }}(val){{ else }}
	return New{{ $AltName }}(fn(*o.value)){{end}}
}

// MapToNillable{{ $AltName }} either returns an empty option of {{ $AltType }}
// if this option is empty, or returns a new {{ $AltType }} option of either
// some or none based on the value returned by the function.
func (o {{ $.Name }}) MapToNillable{{ $AltName }}(fn func({{ $.Type }}) *{{ $AltType }}) {{ $AltName }} {
	if o.value == nil {
		return NewEmpty{{ $AltName }}()
	}

	return NewMaybe{{ $AltName }}(fn(*o.value))
}
{{ end }}
// MapToNillable either returns an empty option of {{ .Type }}
// if this option is empty, or returns a new {{ .Type }} option of either
// some or none based on the value returned by the function.
func (o {{ $.Name }}) MapToNillable(fn func({{ .Type }}) *{{ .Type }}) {{ .Name }} {
	if o.value == nil {
		return NewEmpty{{ .Name }}()
	}

	return NewMaybe{{ .Name }}(fn(*o.value))
}
`
