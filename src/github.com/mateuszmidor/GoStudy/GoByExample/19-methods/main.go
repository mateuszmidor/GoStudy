package main

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return r.width*2 + r.height*2
}

func main() {
	r := rect{width: 10, height: 5}

	println("area:", r.area())
	println("perim:", r.perim())

	rp := &r
	println("area:", rp.area())
	println("perim:", rp.perim())
}
