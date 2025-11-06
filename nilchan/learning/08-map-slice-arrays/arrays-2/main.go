package main

import (
	"github.com/k0kubun/pp"
)

type User struct {
	Name    string
	Rating  float64
	Premium bool
}

func main() {

	userArray := [3]User{
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
	pp.Println("--------------------------")
	pp.Println("User Array Length:", len(userArray))
	pp.Println("Before Update:", userArray)

	for i, u := range userArray {
		pp.Println("Index:", i, "Value:", u)
		if u.Premium {

			userArray[i].Rating += 1.0
		}
	}

	// for i := 0; i < len(userArray); i++ {
	// 	if userArray[i].Premium {
	// 		userArray[i].Rating += 1.0
	// 	}
	// }
	pp.Println("--------------------------")
	pp.Println("After Update:", userArray)
}
