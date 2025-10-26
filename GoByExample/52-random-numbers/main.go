package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(rand.Intn(100)) // always 81
	fmt.Println(rand.Intn(100)) // always 87
	fmt.Println(rand.Float32()) // always 0.6645601

	// default random generator always gives same sequence
	// let's create generator based on time seed to get actual randominess
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println(r1.Intn(100)) // always something different
	fmt.Println(r1.Intn(100))

	// constant seed will give same sequence every time
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Println(r2.Intn(100)) // always 5
	fmt.Println(r2.Intn(100)) // always 87
}
