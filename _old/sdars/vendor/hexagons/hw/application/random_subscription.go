package application

var active bool = true

func RandomSubscription() bool {
	active = !active
	return active
}
