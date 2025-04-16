package main

// Thread 1
func main() {
	ch := make(chan string) // channel is empty

	// Thread 2
	go func() {
		ch <- "Hello World!" // channel is filled
	}()

	msg := <-ch // Channel is emptied

	println(msg)
}
