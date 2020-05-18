//+build ignore

package main

import (
	"os"
	"text/template"
	"time"
)

const (
	filePrefix = "v0/pkg/option/"
	fileSuffix = "Option_test.go"
)

type staticData struct {
	DummyValue interface{}
	DummyAlt   interface{}
}

type tplData struct {
	Type   string
	Name   string
	Time   time.Time
	Static staticData
}

func main() {

	tpl := template.Must(template.ParseGlob("v0/cmd/gen-tests/tpl/*"))
	now := time.Now()

	// list of types to generate for
	types := []tplData{
		{"bool",        "Bool",       now, staticData{true,   false}},
		{"float32",     "Float32",    now, staticData{28.3,   69.9}},
		{"float64",     "Float64",    now, staticData{280.4,  325.1118}},
		{"complex64",   "Complex64",  now, staticData{2548,   77.777}},
		{"complex128",  "Complex128", now, staticData{58.98,  666.666}},
		{"string",      "String",     now, staticData{`"hi"`,   `"bye"`}},
		{"int",         "Int",        now, staticData{123,    321}},
		{"int8",        "Int8",       now, staticData{127,    -128}},
		{"int16",       "Int16",      now, staticData{256,    585}},
		{"int32",       "Int32",      now, staticData{1245,   1234}},
		{"int64",       "Int64",      now, staticData{55555,  99999}},
		{"rune",        "Rune",       now, staticData{454545, 898989}},
		{"byte",        "Byte",       now, staticData{255,    56}},
		{"uint",        "Uint",       now, staticData{5898,   36985}},
		{"uint8",       "Uint8",      now, staticData{123,    245}},
		{"uint16",      "Uint16",     now, staticData{4562,   7894}},
		{"uint32",      "Uint32",     now, staticData{123456, 654321}},
		{"uint64",      "Uint64",     now, staticData{456789, 9874654}},
		{"interface{}", "Untyped",    now, staticData{`"clam"`, 69}},
	}

	for _, v := range types {

		file, err := os.Create(filePrefix + v.Name + fileSuffix)
		check(err)

		//alts := make(map[string]string, len(types)-1)
		//for AltType, AltName := range types {
		//	if Type != AltType {
		//		alts[AltType] = AltName
		//	}
		//}

		check(tpl.ExecuteTemplate(file, "tests", v))

		_ = file.Close()
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
