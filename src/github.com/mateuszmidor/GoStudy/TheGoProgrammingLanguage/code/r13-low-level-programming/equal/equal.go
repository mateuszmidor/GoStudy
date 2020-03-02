package equal

import (
	"reflect"
	"unsafe"
)

type comparison struct {
	x, y unsafe.Pointer
	t    reflect.Type
}

// Equal tests x and y for deep equality
func Equal(x, y interface{}) bool {
	seen := make(map[comparison]bool)
	return equal(reflect.ValueOf(x), reflect.ValueOf(y), seen)
}

func equal(x, y reflect.Value, seen map[comparison]bool) bool {
	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}

	if x.Type() != y.Type() {
		return false
	}

	// check for cycles
	if x.CanAddr() && y.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		yptr := unsafe.Pointer(y.UnsafeAddr())

		if xptr == yptr {
			return true // same addresses so same objects
		}

		c := comparison{xptr, yptr, x.Type()}
		if seen[c] {
			return true
		}
		seen[c] = true
	}

	switch x.Kind() {
	case reflect.Bool:
		return x.Bool() == y.Bool()

	case reflect.String:
		return x.String() == y.String()

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return x.Uint() == y.Uint()

	// case signed numbers...

	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return x.Pointer() == y.Pointer()

	case reflect.Ptr, reflect.Interface:
		return equal(x.Elem(), y.Elem(), seen)

	case reflect.Array, reflect.Slice:
		if x.Len() != y.Len() {
			return false
		}
		for i := 0; i < x.Len(); i++ {
			if !equal(x.Index(i), y.Index(i), seen) {
				return false
			}
		}
		return true

		// case struct

		// case map
	case reflect.Map:
		xEmpty := x.IsZero() || x.IsNil()
		yEmpty := y.IsZero() || y.IsNil()
		if xEmpty && yEmpty {
			return true
		}

		if xEmpty != yEmpty {
			return false
		}
		if len(x.MapKeys()) != len(y.MapKeys()) {
			return false
		}

		for _, key := range x.MapKeys() {
			if !equal(x.MapIndex(key), y.MapIndex(key), seen) {
				return false
			}
			// display(fmt.Sprintf("%s[%s]", path, formatAtom(key)), v.MapIndex(key))
		}

		return true
	}
	panic("unreachable")
}
