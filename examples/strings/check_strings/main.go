 package main 

import (
	"fmt"
)

func main() {
	fmt.Print("Введите строку: ")
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}
