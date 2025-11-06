package main

import "fmt"

func calcValue(a int, b int) (int, int) {
	res := a + b
	res2 := a - b
	return res, res2
}

func main() {
	a, b := calcValue(5, 3)
	fmt.Println(a, b)
}
