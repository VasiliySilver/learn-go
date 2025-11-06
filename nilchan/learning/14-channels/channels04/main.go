package main

import "fmt"

func main() {
	var ch chan string
	// var ch chan string = make(chan string)

	go func() {
		ch <- "Hello"
	}()

	v := <-ch
	fmt.Println("V:", v)

}
