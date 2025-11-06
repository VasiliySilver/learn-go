package main

import "fmt"

func main() {
	fmt.Println("I am main function")
	defer func() {
		fmt.Println("I am main function and I finished my work")
	}()

	hello()

	database()
	foo()
}

func hello() {
	fmt.Println("I am hello function")
	defer func() {
		fmt.Println("I am hello function and I finished my work")
	}()
}

func foo() {

	defer func() {
		fmt.Println("Defer 1")
	}()
	defer func() {
		fmt.Println("Defer 2")
	}()
	defer func() {
		fmt.Println("Defer 3")
	}()

	fmt.Println("I am foo function")
}

func database() {
	fmt.Println("Connecting to database...")

	defer func() {

		fmt.Println("Closing database connection...")
	}()

	fmt.Println("Performing database operations...")
}
