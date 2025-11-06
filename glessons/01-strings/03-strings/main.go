package main

import (
	"fmt"
)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+i, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Print("Enter a string: ")
	var input string
	// fmt.Scanln(&input)
	input = "hello"

	// Length of the string
	fmt.Printf("\nLength of the string: %d characters\n", len([]rune(input)))

	// Reversed string
	reversed := reverseString(input)
	fmt.Printf("Reversed string: %s\n", reversed)

}
