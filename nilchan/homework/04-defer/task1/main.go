package main

import "fmt"

func foo() {
	fmt.Println("FOO")
	defer fmt.Println("foo")
}

func boo() {
	fmt.Println("BOO")
	defer foo()
}

func main() {
	fmt.Println("MAIN")
	boo()

	defer func() {
		fmt.Println("Hello")
	}()

	defer fmt.Println("Hello2")

	defer foo()
}
