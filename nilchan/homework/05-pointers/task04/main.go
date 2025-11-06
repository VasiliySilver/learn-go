package main

import "fmt"

func changeSize(sizePtr *float64) {
	if sizePtr == nil {
		fmt.Println("Current sizePtr is pointing to nil")
		fmt.Println("Operation by changing is cancelled")
		return
	}

	*sizePtr += 1.0
}

func main() {
	size := 2.0

	var ptrSize *float64

	fmt.Println("Current size is ", size)
	changeSize(&size)
	fmt.Println("Size after changing is ", size)

	changeSize(ptrSize)
}
