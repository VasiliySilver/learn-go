package main

import "fmt"

func main() {
	// name := "Alex"
	// age := 30
	// rating := 4.5
	// isStudent := false

	var name string = "Alex"
	var age int = 30
	var rating float64 = 4.5
	var isStudent bool = false

	println("Name:", name)
	println("Age:", age)
	println("Rating:", rating)
	println("Is Student:", isStudent)

	fmt.Printf("Name: %s, Age: %d, Rating: %.1f, Is Student: %t\n", name, age, rating, isStudent)
}
