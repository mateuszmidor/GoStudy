package main

import (
	"reflect"
	"strconv"
	"strings"
)

// path is like "persons.1.address.streetName"
func FieldAtPathMatches(resource interface{}, path string, value interface{}) bool {
	paths := strings.Split(path, ".")
	currentResource := reflect.ValueOf(resource)

	for _, p := range paths {
		switch currentResource.Kind() {
		case reflect.Struct:
			field := currentResource.FieldByName(p)
			if !field.IsValid() {
				return false
			}
			currentResource = field

		case reflect.Map:
			if currentResource.IsNil() {
				return false
			}
			key := reflect.ValueOf(p)
			value := currentResource.MapIndex(key)
			if !value.IsValid() {
				return false
			}
			currentResource = value

		case reflect.Slice:
			index, err := strconv.Atoi(p)
			if err != nil {
				return false
			}
			if index < 0 || index >= currentResource.Len() {
				return false
			}
			currentResource = currentResource.Index(index)

		default:
			return false
		}
	}

	return reflect.DeepEqual(currentResource.Interface(), value)
}
