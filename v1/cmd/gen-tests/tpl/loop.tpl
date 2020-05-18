{{ define "loop" }}
    {{ range $type, $val := .Loop }}
        {{ if ne $.Type $type }}
					func Test{{ $.Name }}_ForceMapTo{{ $val.Name }}(t *testing.T) {
						Convey("{{ $.Name }}.ForceMapTo{{ $val.Name }}", t, func() {
							So(New{{ $.Name }}({{ $.Static.DummyValue }}).ForceMapTo{{ $val.Name }}(func({{ $.Type }}) {{ $type }} { return {{ $val.MapReturnValue }} }).Get(),
							ShouldEqual, {{ $val.MapReturnValue }})

							So(NewEmpty{{ $.Name }}().ForceMapTo{{ $val.Name }}(func({{ $.Type }}) {{ $type }} { return {{ $val.MapReturnValue }} }).IsNil(),
							ShouldBeTrue)
						})
					}

					func Test{{ $.Name }}_MapToNillable{{ $val.Name }}(t *testing.T) {
						Convey("{{ $.Name }}.MapToNillable{{ $val.Name }}", t, func() {
							So(New{{ $.Name }}({{ $.Static.DummyValue }}).MapToNillable{{ $val.Name }}(func({{ $.Type }}) *{{ $type }} {
							var tmp {{ $type }} = {{ $val.MapReturnValue }}
							return &tmp
							}).Get(), ShouldEqual, {{ $val.MapReturnValue }})

							So(NewEmpty{{ $.Name }}().MapToNillable{{ $val.Name }}(func({{ $.Type }}) *{{ $type }} {
							var tmp {{ $type }} = {{ $val.MapReturnValue }}
							return &tmp
							}).IsNil(), ShouldBeTrue)

							So(New{{ $.Name }}({{ $.Static.DummyValue }}).MapToNillable{{ $val.Name }}(func({{ $.Type }}) *{{ $type }} { return nil }).IsNil(),
							ShouldBeTrue)
						})
					}

					func Test{{ $.Name }}_MapTo{{ $val.Name }}(t *testing.T) {
						Convey("{{ $.Name }}.MapTo{{ $val.Name }}", t, func() {
							test1, err1 := New{{ $.Name }}({{ $.Static.DummyValue }}).MapTo{{ $val.Name }}(func({{ $.Type }}) ({{ $type }}, error) {
								return {{ $val.MapReturnValue }}, nil
							})

							So(err1, ShouldBeNil)
							So(test1.Get(), ShouldEqual, {{ $val.MapReturnValue }})

							test2, err2 := NewEmpty{{ $.Name }}().MapTo{{ $val.Name }}(func({{ $.Type }}) ({{ $type }}, error) {
								return {{ $val.MapReturnValue }}, nil
							})

							So(err2, ShouldBeNil)
							So(test2.IsNil(), ShouldBeTrue)

							test3, err3 := New{{ $.Name }}({{ $.Static.DummyValue }}).MapTo{{ $val.Name }}(func({{ $.Type }}) ({{ $type }}, error) {
								return {{ $val.MapReturnValue }}, errors.New("{{ $val.MapReturnError }}")
							})

							So(err3, ShouldNotBeNil)
							So(err3.Error(), ShouldEqual, "{{ $val.MapReturnError }}")
							So(test3.IsNil(), ShouldBeTrue)
						})
					}
        {{ end }}
    {{ end }}
{{ end }}