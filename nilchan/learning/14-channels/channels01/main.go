package main

import "fmt"

func main() {
	ch := make(chan int)
	// Exmaple nil cahnel
	// var ch chan int

	// Close channel
	close(ch)

	// Reead value from closed channel
	v1, ok1 := <-ch
	v2, ok2 := <-ch
	v3, ok3 := <-ch
	fmt.Println("v", v1, v2, v3)
	fmt.Println("o", ok1, ok2, ok3)

	// Close closed channel - got panic error
	// close(ch)

}
