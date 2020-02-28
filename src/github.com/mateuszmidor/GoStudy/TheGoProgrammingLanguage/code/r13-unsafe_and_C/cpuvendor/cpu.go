package main

// below comments are preprocessed by cgo becouse of import "C".
// WARN: there can be NO BLANK LINE between comment and import "C"

/*
const char* get_vendor();
*/
import "C"

func getCPUVendor() string {
	return C.GoString(C.get_vendor())
}
