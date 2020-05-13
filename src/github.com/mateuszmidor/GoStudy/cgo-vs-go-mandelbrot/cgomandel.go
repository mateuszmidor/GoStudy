package main

// #cgo CPPFLAGS: -DCGO=1
// #include <stdint.h>
// char* makeMandel(size_t width, size_t height, size_t numParallel);
// void freeMandel(char* data);
// void makeMandel2(size_t width, size_t height, char* data, size_t numParallel) ;
import "C"
import (
	"image"
	"unsafe"
)

// Variant 1: C side allocates memory
func cgoMandel(img *image.RGBA, numParallel int) {
	// get C memory allocated from C function
	crgba := C.makeMandel(width, height, C.size_t(numParallel))

	// convert to slice of bytes
	rgba := C.GoBytes(unsafe.Pointer(crgba), width*height*4)

	// copy bytes to image
	copy(img.Pix, rgba)

	// release C memory
	C.freeMandel(crgba)
}

// Variant 2: GO side allocates memory
func cgoMandel2(img *image.RGBA, numParallel int) {
	// cast GO slice of bytes to char*
	crgba := (*C.char)(unsafe.Pointer(&img.Pix[0]))

	// run C function to fill the image
	C.makeMandel2(width, height, crgba, C.size_t(numParallel))
}
