package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Fluppy bird
	// 🐤🔵🔵

	// Птица Флуппи
	// 🐤
	// 🔵
	// 🔵
	fmt.Println("Fluppy bird")
	fmt.Println("--------------------")

	score := 0

	for {

		fmt.Println("Я подлетаю к трубе")
		fmt.Println("🔵 🔵 🐤")
		fmt.Println("")

		fmt.Println("Я пролетаю через трубу")
		fmt.Println("🔵 🐤 🔵")
		// 0 1 2 3 4
		// 1
		if rand.Intn(8) == 1 {
			fmt.Println("")
			fmt.Println("О нет! Я врезался в трубу!")
			fmt.Println("🔵 💥 🔵")
			break
		}

		fmt.Println("")
		fmt.Println("Я успешно пролетел через трубу!")
		fmt.Println("🐤 🔵 🔵")
		fmt.Println("")

		score++
		fmt.Println("Ваш счет:", score)
		fmt.Println("")

		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("Игра окончена! Ваш итоговый счет:", score)

}
