package main

import (
	"fmt"
	"math/rand"
)

func main() {
	foo(rand.Intn(10) + 1)
	boo(rand.Intn(10) + 1)

}

func foo(count int) {
	for i := 0; i < count; i++ {
		fmt.Println("Fooooooo")
	}
}

func boo(count int) {
	for i := 0; i < count; i++ {
		fmt.Println("Booooooo")
	}
}
