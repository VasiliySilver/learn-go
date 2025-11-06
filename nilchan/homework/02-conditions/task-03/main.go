package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var litres int
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Введите количество литров которое вы выпили:")
		ok := scanner.Scan()
		if !ok {
			fmt.Println("Input error")
			continue
		}

		text := scanner.Text()
		_, err := fmt.Sscanf(text, "%d", &litres)
		if err != nil {
			fmt.Println("Неверный формат литров")
			continue
		}

		if litres < 1 || litres > 3 {
			fmt.Println("Ты какой-то странный")
		} else {
			fmt.Println("А ты знаешь золотую середину!")
		}

	}
}
