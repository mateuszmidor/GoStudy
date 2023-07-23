package main

import (
	"reflect"
)

// Injector injects dependencies into given struct
type Injector struct {
	deps map[reflect.Type]reflect.Value
}

// NewInjector is constructor
func NewInjector() Injector {
	return Injector{deps: map[reflect.Type]reflect.Value{}}
}

// Set assigns value to a type
func (injector *Injector) Set(t reflect.Type, v interface{}) {
	injector.deps[t] = reflect.ValueOf(v)
}

// Inject actually injects the dependencies. Argument must be pointer
func (injector *Injector) Inject(objptr interface{}) {
	v := reflect.ValueOf(objptr).Elem()
	t := reflect.TypeOf(objptr).Elem()

	for i := 0; i < v.NumField(); i++ {
		vf := v.Field(i)
		vt := t.Field(i)

		if !canSet(vf) {
			continue
		}

		if !shouldSet(vt) {
			continue
		}

		if value, ok := injector.deps[vt.Type]; ok {
			vf.Set(value)
		}
	}

}

func canSet(v reflect.Value) bool {
	return v.CanSet()
}

func shouldSet(v reflect.StructField) bool {
	return v.Tag.Get("inject") == "yes"
}
