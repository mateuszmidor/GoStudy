package filter

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// DigAttribute returns value of complex structure found at given path, the valid path format is: "persons.1.address.streetName"
func DigAttribute(resource interface{}, path string) interface{} {
	paths := strings.Split(path, ".")
	currentResource := reflect.ValueOf(resource)

	for _, p := range paths {
		switch currentResource.Kind() {
		case reflect.Struct:
			field := currentResource.FieldByName(p)
			if !field.IsValid() {
				return nil
			}
			currentResource = field

		case reflect.Interface:
			if currentResource.IsNil() {
				return nil
			}
			key := reflect.ValueOf(p)
			value := currentResource.Elem().MapIndex(key)
			if !value.IsValid() {
				return nil
			}
			currentResource = value

		case reflect.Map:
			if currentResource.IsNil() {
				return nil
			}
			key := reflect.ValueOf(p)
			value := currentResource.MapIndex(key)
			if !value.IsValid() {
				return nil
			}
			currentResource = value

		case reflect.Slice:
			index, err := strconv.Atoi(p)
			if err != nil {
				return nil
			}
			if index < 0 || index >= currentResource.Len() {
				return nil
			}
			currentResource = currentResource.Index(index)

		default:
			fmt.Printf("could not dig attribute, path: %s, type: %s\n", p, currentResource.Kind().String())
			return nil
		}
	}

	return currentResource.Interface()
}
