// Дано age
// Если человек совершеннолетний - вывести я продам тебе пиво
// Иначе если возраст меньше 12 - вывести ты че малой совсем берега попутал пиво покупать?
// Иначе вывести то что алкоголь только от 18 лет
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

		for {
			ok := scanner.Scan()
			if !ok {
				fmt.Println("Input error")
			}

			text := scanner.Text()

			_, err := fmt.Scanf(text, "%d", &age)
			if err != nil {
				fmt.Println("Ivalid age format: ", err)
			}

			if text == "exit" {
				return
			}

			if age > 18 {
				fmt.Println("I will sell you the bear!")
				continue
			} else if age < 12 {
				fmt.Println("Are you serious, you are just kid")
			} else {
				fmt.Println("The correct age is 18!")
			}
		}
	}
}
