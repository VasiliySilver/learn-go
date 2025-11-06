package main

import "fmt"

type User struct {
	Name   string  // ""
	Age    int     // 0
	Number string  // ""
	Closed bool    // false
	Rating float64 // 0.0
}

func main() {
	user := User{
		Name:   "Alice",        // ""
		Age:    30,             // 0
		Number: "123-456-7890", // ""
		Closed: false,          // false
		Rating: 4.5,            // 0.0
	}

	user2 := User{
		Name:   "Bob",
		Age:    25,
		Number: "123-444-9999",
	}
	fmt.Println("User1", user)
	fmt.Println("User2", user2)

	println("User Name:", user.Name)
	println("User Age:", user.Age)
	println("User Number:", user.Number)
	println("Is Account Closed?:", user.Closed)
	println("User Rating:", user.Rating)
}
