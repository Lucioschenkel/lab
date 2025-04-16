package main

import "fmt"

// Indicate that the channel is only used to send data
func receive(name string, hello chan<- string) {
	hello <- name
}

// Indicate that the channel is only used to receive data
func read(data <-chan string) {
	fmt.Println(<-data)
}

func main() {
	hello := make(chan string)
	go receive("Hello", hello)
	read(hello)
}
