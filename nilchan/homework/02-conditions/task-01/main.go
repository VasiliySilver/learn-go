package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var age int

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter your age:")
		ok := scanner.Scan()

		if !ok {
			fmt.Println("Input error")
			return
		}

		text := scanner.Text()

		if text == "exit" {
			fmt.Println("Exiting the program")
			return
		}

		// Запись возраста в переменную age если введено корректное число
		_, err := fmt.Sscanf(text, "%d", &age)
		if err != nil {
			fmt.Println("Invalid age format: ", err)
			continue
		}

		if age < 0 {
			fmt.Println("Age cannot be negative")
			return
		} else if age < 18 {
			fmt.Println("You are a minor")
		} else if age <= 65 {
			fmt.Println("You are an adult")

		} else if age > 65 && age < 110 {
			fmt.Println("You are a senior")
		} else {
			fmt.Println("Invalid age")
		}
	}

}
