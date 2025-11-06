package main

import "fmt"

func main() {
	var name string
	var age int
	var rating float64
	var isStudent bool

	name = "Alex"
	age = 30
	rating = 4.5
	isStudent = false

	println("Name:", name)
	println("Age:", age)
	println("Rating:", rating)
	println("Is Student:", isStudent)

	fmt.Printf("Name: %s, Age: %d, Rating: %.1f, Is Student: %t\n", name, age, rating, isStudent)
}
