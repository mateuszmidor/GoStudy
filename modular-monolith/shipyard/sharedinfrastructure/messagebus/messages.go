package messagebus

type Message interface {
}

type ProductCreated struct {
	Name     string
	Quantity uint
}
