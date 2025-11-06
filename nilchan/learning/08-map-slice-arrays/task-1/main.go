package main

import "fmt"

func main() {
	// нам нужно хранить данные о людях
	// нужно по имени человека сразу понимать
	// был ли он судим

	criminal := map[string]bool{
		"Vasya":  true,
		"Petya":  false,
		"Antony": true,
		"Vova":   false,
		"Alex":   true,
	}

	c, ok := criminal["Dan"]

	if !ok {
		fmt.Println("This person isn't in our database")
		return
	}

	fmt.Println("This person in our database")
	if c {
		fmt.Println("This person is criminal")
	} else {
		fmt.Println("This person isn't criminal")
	}

	fmt.Println(c, ok)

}
