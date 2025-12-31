package main

import (
	"context"
	"fmt"
)

var counter = 1

func Greet(ctx context.Context, name string) (string, error) {
	// simulate a single error
	if counter > 0 {
		counter--
		return "", fmt.Errorf("some-error")
	}
	return fmt.Sprintf("Hello %s", name), nil
}
