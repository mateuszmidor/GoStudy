// Project: 3 ways to call method on an object
// Usage: go run .
package main

import (
	"fmt"
)

// Num is number with distance method
type Num struct {
	val int
}

func (n Num) dist(i int) uint {
	return intAbs(n.val - i)
}

func intAbs(v int) uint {
	switch {
	case v >= 0:
		return uint(v)
	default:
		return uint(-v)
	}
}
func main() {
	case1()
	case2()
	case3()
}

func case1() {
	var n Num
	n.val = 5             // val is private, but accessible within the package main
	distance := n.dist(2) // case1: just call member function
	fmt.Println(distance)
}

func case2() {
	n := Num{5}
	distFuncHoldingObject := n.dist // case2: grab method together with the object
	distance := distFuncHoldingObject(1)
	fmt.Println(distance)
}

func case3() {
	n := Num{5}
	disFunc := Num.dist // case3: just grab the method, provide object when calling
	distance := disFunc(n, 0)
	fmt.Println(distance)
}
