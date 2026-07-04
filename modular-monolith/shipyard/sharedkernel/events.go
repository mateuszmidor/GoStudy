package sharedkernel

type Event interface{}

type ProductCreated struct {
	Name     string
	Quantity uint
}
