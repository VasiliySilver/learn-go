package main

import "fmt"

func main() {
	fmt.Println("Function before calling sum")
	b := sum(10, 20)

	fmt.Println("Function after calling sum")

	fmt.Println("Sum is:", b)
}

func sum(a int, b int) int {
	return a + b
}
