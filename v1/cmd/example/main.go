package main

import (
	"fmt"
	"strconv"

	"github.com/Foxcapades/goop/v0/pkg/option"
)

// NewTest creates a new Test struct with the given string.
func NewTest(val string) Test {
	return Test{val}
}

// Test struct.
type Test struct {
	Value string
}

func (t *Test) String() *string {
	return &t.Value
}

func main() {
	fmt.Println(option.NewBool(true).
		ForceMapToInt(func(b bool) int {
			if b {
				return 5
			}
			return 10
		}).
		ForceMapToString(strconv.Itoa).
		ForceMapToUntyped(func(s string) interface{} { return NewTest(s) }).
		MapToNillableString(func(t interface{}) *string { return nil }).
		Or("15"))
}
