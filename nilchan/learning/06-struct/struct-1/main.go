package main

import "fmt"

type User struct {
	Name   string  // ""
	Age    int     // 0
	Number string  // ""
	Closed bool    // false
	Rating float64 // 0.0
}

func NewUser(
	name string,
	age int,
	number string,
	closed bool,
	rating float64,
) User {
	if name == "" {
		fmt.Println("Name cannot be empty")
		return User{}
	}
	if age <= 0 || age >= 150 {
		fmt.Println("Age must be between 1 and 149")
		return User{}
	}
	if number == "" {
		fmt.Println("Number cannot be empty")
		return User{}
	}
	if rating < 0.0 || rating > 10.0 {
		fmt.Println("Rating must be between 0.0 and 10.0")
		return User{}
	}

	return User{
		Name:   name,
		Age:    age,
		Number: number,
		Closed: closed,
		Rating: rating,
	}
}

func (u User) greeting() {
	fmt.Println("Hello everyone!")
	fmt.Println("My name is", u.Name)
	fmt.Println("My rating is", u.Rating)
	u.RatingUp(3.0)
}

func (u User) Goodbye() {
	fmt.Println("Goodbye from", u.Name)
	fmt.Println("Final rating was", u.Rating)
}

func (u *User) RatingUp(rating float64) {
	if u.Rating+rating <= 10.0 {
		u.Rating += rating
		fmt.Println("New rating is", u.Rating)
	} else {
		fmt.Println("Rating cannot exceed 10.0")
	}
}

// func RatingUp(u *User, rating float64) {
// 	if u.Rating+rating <= 10.0 {
// 		u.Rating += rating
// 		fmt.Println("New rating is", u.Rating)
// 	} else {
// 		fmt.Println("Rating cannot exceed 10.0")
// 	}
// }

func main() {
	user := User{
		Name:   "Alice",
		Rating: 4.5,
	}

	fmt.Println("Before greeting, user rating is:", user.Rating)
	user.RatingUp(4.0)
	user.greeting()
	user.Goodbye()
	fmt.Println("After greeting, user rating is:", user.Rating)

}
