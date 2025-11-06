package user

import "fmt"

func NewUser(name string, age int) User {
	if name == "" {
		fmt.Println("Name cannot be empty")
		return User{}
	}
	if age == 0 {
		fmt.Println("Age must be positive")
		return User{}
	}
	return User{
		name: name,
		age:  age,
	}
}
