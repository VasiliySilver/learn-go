package main

import "github.com/k0kubun/pp"

type User struct {
	Name    string
	Rating  float64
	Premium bool
}

func main() {

	userArray := []User{
		{
			Name:    "Alice",
			Rating:  4.5,
			Premium: true,
		},
		{
			Name:    "Bob",
			Rating:  3.8,
			Premium: true,
		},
		{
			Name:    "Charlie",
			Rating:  4.2,
			Premium: false,
		},
	}

	// 1 - Append example
	userArray = append(
		userArray,
		User{
			Name:    "David",
			Rating:  4.0,
			Premium: true,
		},
	)

	// 2 - Update by index example (at index 1)
	userArray[1].Rating += 0.5

	// 3 - Delete by index example (at index 0)
	indexToDelete := 0
	userArray = append(
		userArray[:indexToDelete],
		userArray[indexToDelete+1:]...,
	)

	// 4 - Iterate and conditionally update example
	pp.Println("--------------------------")
	pp.Println("User Slice Length:", len(userArray))
	pp.Println("Before Update:", userArray)

	for i, u := range userArray {
		pp.Println("Index:", i, "Value:", u)
		if u.Premium {
			userArray[i].Rating += 1.0
		}
	}

	pp.Println("--------------------------")
	pp.Println("After Update:", userArray)
}
