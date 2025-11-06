package main

import "fmt"

func main() {
	fmt.Println("Hello!")

	n := 10
	m := "hello"
	foo(n, m)

	fmt.Println("N after foo:", n)
	fmt.Println("M after foo:", m)

	result := sum(5, 7)
	fmt.Println("Sum result:", result)

	fmt.Println("Goodbye!")

}

func sum(a int, b int) int {
	fmt.Println("Calculating sum of", a, "and", b)
	return a + b
}

func foo(n int, m string) {
	fmt.Println("Got n:", n)
	fmt.Println("Got m:", m)

	n = -100
	m = "changed"
	fmt.Println("Changed n to:", n)
	fmt.Println("Changed m to:", m)
}
