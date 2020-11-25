package main

import (
	"errors"
	"fmt"
	"io"
)

// PathError is custom error type that implements "error" interface
type PathError struct {
	path string
}

// implement "error" interface
func (e PathError) Error() string {
	return e.path
}

func main() {
	// create an error chain:
	// platformError->fsError->pathError
	pathError := &PathError{path: `C:\user32`}
	fsError := fmt.Errorf("filesystem error: %w", pathError)    // %w for "wrap" - resulting fsError will have "Unwrap" method
	platformError := fmt.Errorf("windows32 error: %w", fsError) // %w for "wrap" - resulting platformError will have "Unwrap" method

	// entire error chain:
	// windows32 error: filesystem error: C:\user32
	fmt.Println(platformError)

	// take out outer layer:
	// filesystem error: C:\user32
	fmt.Println(errors.Unwrap(platformError)) // for errors.Unwrap(), platformError needs to implement "Unwrap". This is ensured by fmt.Errorf() with %w verb

	// take out 2 outer layers:
	// C:\user32
	fmt.Println(errors.Unwrap(errors.Unwrap(platformError)))

	// take out all 3 layers:
	// <nil>
	fmt.Println(errors.Unwrap(errors.Unwrap(errors.Unwrap(platformError))))

	// check platformError contains pathError
	// true
	fmt.Println("platformError contains pathError:", errors.Is(platformError, pathError))

	// check platformError contains pathError
	// false
	fmt.Println("platformError contains EOF:", errors.Is(platformError, io.EOF))

	// extract PathError from platformError
	pe := new(PathError)
	errors.As(platformError, &pe) // yes, provide pointer to pointer here (&pe)
	fmt.Println("platformError as PathError:", pe)
}
