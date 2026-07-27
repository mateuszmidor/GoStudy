package defaults

import (
	"fmt"
	"reflect"
	"strconv"
)

// SetEmptyFieldsToDefault updates empty fields to their defaults provided in `default` tag, eg. `default:"12"`
func SetEmptyFieldsToDefault(out interface{}) {
	v := reflect.ValueOf(out).Elem()

	// non recursive so only flat structures are supported
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if !field.IsZero() {
			continue // this field is set
		}

		defval := v.Type().Field(i).Tag.Get("default")
		if defval == "" {
			continue // no default value for this field
		}

		// set default value
		parseInto(defval, field)
	}
}

func parseInto(strval string, v reflect.Value) {
	switch v.Kind() {

	case reflect.String:
		v.SetString(strval)

	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8: // only unsigned ints
		val, _ := strconv.ParseUint(strval, 10, 64)
		v.SetUint(val)

	default:
		panic(fmt.Sprintf("Unsupported type %v", v.Kind()))
	}
}
