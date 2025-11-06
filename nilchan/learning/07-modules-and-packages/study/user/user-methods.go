package user

import "fmt"

func (u User) PrintUser() {
	fmt.Println("User{name: " + u.name + ", age: " + fmt.Sprint(u.age) + "}")
}

func (u User) GetName() string {
	return u.name
}

func (u User) GetAge() int {
	return u.age
}

func (u *User) SetName(name string) {
	if name == "" {
		fmt.Println("Name cannot be empty")
		return
	}
	u.name = name
}

func (u *User) SetAge(age int) {
	if age <= 0 {
		fmt.Println("Age must be positive")
		return
	}
	u.age = age
}
