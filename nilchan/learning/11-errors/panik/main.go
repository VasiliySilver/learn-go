package main

import "fmt"

func main() {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("Was a panic error:", p)
		}
	}()

	a := 0
	b := 1 / a

	fmt.Println("B:", b)
}
