package main

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		println("received message", msg)
	default:
		println("no message receives")
	}

	msg := "hi"
	select {
	case messages <- msg:
		println("sent message", msg)
	default:
		println("no message sent")
	}

	select {
	case msg := <-messages:
		println("received message", msg)
	case sig := <-signals:
		println("received signal", sig)
	default:
		println("no activity
		")
	}
}
