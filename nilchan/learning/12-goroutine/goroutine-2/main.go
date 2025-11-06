package main

import (
	"fmt"
	"time"
)

func foo() {
	for {
		fmt.Println("Hello")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go foo()
	go func() {
		for {
			fmt.Println("Anon")
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(1 * time.Second)

}
