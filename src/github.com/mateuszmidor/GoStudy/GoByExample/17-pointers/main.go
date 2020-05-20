package main

func zeroVal(n int) {
	n = 0
}

func zeroPtr(n *int) {
	*n = 0
}
func main() {
	n := 1
	println(n)

	zeroVal(n)
	println(n)

	zeroPtr(&n)
	println(n)
	println(&n)
}
