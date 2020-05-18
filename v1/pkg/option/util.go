package option

import "reflect"

func isNil(v interface{}) bool {
	if v == nil {
		return true
	}

	ref := reflect.ValueOf(v)

	switch ref.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Array, reflect.Chan:
		return ref.IsNil()
	case reflect.Invalid:
		return true
	}

	return false
}
