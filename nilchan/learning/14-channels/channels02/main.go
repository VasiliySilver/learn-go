package main

func main() {
	// Create new channel
	ch := make(chan int)

	// Close channel
	close(ch)

	// Write value into closed channel - got panic error
	ch <- 10
}
