package main

import "fmt"

func main() {
	intSlice := make([]int, 0, 5)

	// intSlice = append(intSlice, 10, 20, 25, 30, 150)
	// intSlice = append(intSlice, 20)
	// intSlice = append(intSlice, 25)
	// intSlice = append(intSlice, 150)

	fmt.Println("Initial slice:", intSlice, "Length:", len(intSlice), "Capacity:", cap(intSlice))
}
