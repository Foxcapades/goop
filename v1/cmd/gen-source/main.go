//+build ignore

package main

import (
	"github.com/Foxcapades/goop/v1/internal/gen"
	"html/template"
	"os"
	"time"
)

const (
	filePrefix = "v1/pkg/option/"
	fileSuffix = "Option.go"
)

func main() {

	// list of types to generate for
	types := map[string]string{
		"bool":        "Bool",
		"float32":     "Float32",
		"float64":     "Float64",
		"complex64":   "Complex64",
		"complex128":  "Complex128",
		"string":      "String",
		"int":         "Int",
		"int8":        "Int8",
		"int16":       "Int16",
		"int32":       "Int32",
		"int64":       "Int64",
		"rune":        "Rune",
		"byte":        "Byte",
		"uint":        "Uint",
		"uint8":       "Uint8",
		"uint16":      "Uint16",
		"uint32":      "Uint32",
		"uint64":      "Uint64",
		"interface{}": "Untyped",
	}

	tpl := template.Must(template.New("").Parse(gen.Template))
	now := time.Now()

	for Type, Name := range types {

		file, err := os.Create(filePrefix + Name + fileSuffix)
		check(err)

		alts := make(map[string]string, len(types)-1)
		for AltType, AltName := range types {
			if Type != AltType {
				alts[AltType] = AltName
			}
		}

		check(tpl.Execute(file, struct {
			Time       time.Time
			Type       string
			Name       string
			Alternates map[string]string
		}{
			Time:       now,
			Type:       Type,
			Name:       Name,
			Alternates: alts,
		}))

		_ = file.Close()
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
