{{ define "tests" }}
	// Package option_test generated @ {{ .Time }}
	package option_test

	import (
		"errors"
		"testing"

		. "github.com/Foxcapades/goop/v0/pkg/option"
		. "github.com/smartystreets/goconvey/convey"
	)
	{{ template "static" $ }}
	{{ template "loop" $ }}
{{ end }}