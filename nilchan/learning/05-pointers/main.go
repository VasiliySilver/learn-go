package main

import "fmt"

func main() {

	number := 42

	pointer := &number

	fmt.Println("Value of number:", number)
	fmt.Println("Address of number:", &number)
	fmt.Println("Value of pointer:", pointer)
	fmt.Println("Value at the address stored in pointer:", *pointer)

	boo(number)
	fmt.Println("Number after boo:", number)

	foo(pointer)
	fmt.Println("Number after foo:", number)
	fmt.Println("Pointer after foo:", pointer)
	fmt.Println("Value at the address stored in pointer after foo:", *pointer)

}

func foo(n *int) {
	fmt.Println("Pointer: ", n, "Value at pointer:", *n)
	*n = 100
}

func boo(n int) {
	n = 10
}
