package main

import "fmt"

func main() {
	number := 15
	fmt.Println("Number:", number)

	var ptr *int

	if ptr != nil {
		fmt.Println("Pointer is not nil")
	} else {
		fmt.Println("Pointer is nil, allocating memory")
		ptr = &number
	}

	fmt.Println("Pointer address:", ptr)
	fmt.Println("Value at pointer:", *ptr)

}
