package main

func main() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	println(<-messages)
	println(<-messages)
}
