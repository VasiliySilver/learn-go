package main

import "fmt"

var number int = 5

func main() {
	// number *= 4

	greeting("John")

	fmt.Println("Number is:", number)

	result := square(number)

	fmt.Println("Square of number is:", result)

	fmt.Println("Number is: ", number)

}

func greeting(name string) {
	if name == "" {
		fmt.Println("You didn't provide a name")
		return
	}
	fmt.Println("Hello,", name)
}

func square(number int) int {
	return number * number
}
