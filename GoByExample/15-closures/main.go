package main

func intSeq() func() int {
	n := 0

	return func() int {
		n++
		return n
	}
}
func main() {
	gen1 := intSeq()
	println(gen1())
	println(gen1())
	println(gen1())

	gen2 := intSeq()
	println(gen2())
	println(gen2())
}
