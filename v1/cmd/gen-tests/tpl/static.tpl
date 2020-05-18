{{ define "static" }}
func TestNew{{ .Name }}(t *testing.T) {
	Convey("New{{ .Name }}", t, func() {
		test := New{{ .Name }}({{ .Static.DummyValue }})
		So(test.IsPresent(), ShouldBeTrue)
		So(test.IsNil(), ShouldBeFalse)
	})
}

func TestNewEmpty{{ .Name }}(t *testing.T) {
	Convey("NewEmpty{{ .Name }}", t, func() {
		test := NewEmpty{{ .Name }}()
		So(test.IsPresent(), ShouldBeFalse)
		So(test.IsNil(), ShouldBeTrue)
	})
}

func TestNewMaybe{{ .Name }}(t *testing.T) {
	Convey("NewMaybe{{ .Name }}", t, func() {
		var val1 {{ .Type }}

		test1 := NewMaybe{{ .Name }}(&val1)
		So(test1.IsPresent(), ShouldBeTrue)
		So(test1.IsNil(), ShouldBeFalse)

		test2 := NewMaybe{{ .Name }}(nil)
		So(test2.IsPresent(), ShouldBeFalse)
		So(test2.IsNil(), ShouldBeTrue)
	})
}

func Test{{ .Name }}_Get(t *testing.T) {
	Convey("{{ .Name }}.Get", t, func() {
		test1 := New{{ .Name }}({{ .Static.DummyValue }})
		So(test1.Get(), ShouldEqual, {{ .Static.DummyValue }})

		test2 := NewEmpty{{ .Name }}()
		So(func() { test2.Get() }, ShouldPanic)
	})
}

func Test{{ .Name }}_Or(t *testing.T) {
	Convey("{{ .Name }}.Or", t, func() {
		test1 := New{{ .Name }}({{ .Static.DummyValue }})
		So(test1.Or({{ .Static.DummyAlt }}), ShouldEqual, {{ .Static.DummyValue }})

		test2 := NewEmpty{{ .Name }}()
		So(test2.Or({{ .Static.DummyValue }}), ShouldEqual, {{ .Static.DummyValue }})
	})
}

func Test{{ .Name }}_OrGet(t *testing.T) {
	Convey("{{ .Name }}.OrGet", t, func() {
		test1 := New{{ .Name }}({{ .Static.DummyValue }})
		So(test1.OrGet(func() {{ .Type }} { return {{ .Static.DummyAlt }} }), ShouldEqual, {{ .Static.DummyValue }})

		test2 := NewEmpty{{ .Name }}()
		So(test2.OrGet(func() {{ .Type }} { return {{ .Static.DummyValue }} }), ShouldEqual, {{ .Static.DummyValue }})
	})
}

func Test{{ .Name }}_OrPanicWith(t *testing.T) {
	Convey("{{ .Name }}.OrPanicWith", t, func() {
		test1 := New{{ .Name }}({{ .Static.DummyValue }})
		So(test1.OrPanicWith(func() interface{} { return "panic!" }), ShouldEqual, {{ .Static.DummyValue }})

		test2 := NewEmpty{{ .Name }}()
		So(func() {
			test2.OrPanicWith(func() interface{} { return "test value" })
		}, ShouldPanicWith, "test value")
	})
}

func Test{{ .Name }}_MapToNillable(t *testing.T) {
	Convey("{{ .Name }}.MapToNillable", t, func() {
		test1 := New{{ .Name }}({{ .Static.DummyValue }})
		So(test1.MapToNillable(func(b {{ .Type }}) *{{ .Type }} { return &b }).IsPresent(), ShouldBeTrue)

		test2 := New{{ .Name }}({{ .Static.DummyValue }})
		So(test2.MapToNillable(func(b {{ .Type }}) *{{ .Type }} { return nil }).IsNil(), ShouldBeTrue)

		test3 := NewEmpty{{ .Name }}()
		So(func() {
			test3.MapToNillable(func(b {{ .Type }}) *{{ .Type }} { panic("foo") }).IsNil()
		}, ShouldNotPanic)
	})
}{{ end }}