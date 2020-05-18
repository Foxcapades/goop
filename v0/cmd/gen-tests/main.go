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
	Loop   map[string]loopData
}

type loopData struct {
	Name           string
	MapReturnValue interface{}
	MapReturnError string
}

func main() {

	tpl := template.Must(template.ParseGlob("v0/cmd/gen-tests/tpl/*"))
	now := time.Now()

	// list of types to generate for
	types := []tplData{
		{"bool", "Bool", now, staticData{true, false}, nil},
		{"float32", "Float32", now, staticData{28.3, 69.9}, nil},
		{"float64", "Float64", now, staticData{280.4, 325.1118}, nil},
		{"complex64", "Complex64", now, staticData{2548, 77.777}, nil},
		{"complex128", "Complex128", now, staticData{58.98, 666.666}, nil},
		{"string", "String", now, staticData{`"hi"`, `"bye"`}, nil},
		{"int", "Int", now, staticData{123, 321}, nil},
		{"int8", "Int8", now, staticData{127, -128}, nil},
		{"int16", "Int16", now, staticData{256, 585}, nil},
		{"int32", "Int32", now, staticData{1245, 1234}, nil},
		{"int64", "Int64", now, staticData{55555, 99999}, nil},
		{"rune", "Rune", now, staticData{454545, 898989}, nil},
		{"byte", "Byte", now, staticData{255, 56}, nil},
		{"uint", "Uint", now, staticData{5898, 36985}, nil},
		{"uint8", "Uint8", now, staticData{123, 245}, nil},
		{"uint16", "Uint16", now, staticData{4562, 7894}, nil},
		{"uint32", "Uint32", now, staticData{123456, 654321}, nil},
		{"uint64", "Uint64", now, staticData{456789, 9874654}, nil},
		{"interface{}", "Untyped", now, staticData{`"clam"`, 69}, nil},
	}
	dyn := map[string]loopData{
		"bool":        {"Bool", true, "bool error"},
		"float32":     {"Float32", 32.3, "float error"},
		"float64":     {"Float64", 3.2455, "float error"},
		"complex64":   {"Complex64", 123456, "complex error"},
		"complex128":  {"Complex128", 987654, "complex error"},
		"string":      {"String", `"bonanza"`, "string error"},
		"int":         {"Int", 123456, "int error"},
		"int8":        {"Int8", 111, "int8 error"},
		"int16":       {"Int16", 222, "int16 error"},
		"int32":       {"Int32", 333, "int32 error"},
		"int64":       {"Int64", 444, "int64 error"},
		"rune":        {"Rune", 555, "rune error"},
		"byte":        {"Byte", 222, "byte error"},
		"uint":        {"Uint", 666, "uint error"},
		"uint8":       {"Uint8", 77, "uint8 error"},
		"uint16":      {"Uint16", 888, "uint16 error"},
		"uint32":      {"Uint32", 999, "uint32 error"},
		"uint64":      {"Uint64", 101010, "uint64 error"},
		"interface{}": {"Untyped", `"111111"`, "interface error"},
	}

	for _, v := range types {

		file, err := os.Create(filePrefix + v.Name + fileSuffix)
		check(err)

		v.Loop = dyn

		check(tpl.ExecuteTemplate(file, "tests", v))

		_ = file.Close()
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
