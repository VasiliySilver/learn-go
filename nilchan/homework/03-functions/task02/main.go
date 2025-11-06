package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	guestsMap := map[string]int{
		"Ivan":    1,
		"Olga":    2,
		"Dmitriy": 3,
	}

	for name, roomNumber := range guestsMap {
		greetingGuest(name, roomNumber)
	}

	fmt.Println("\n✅ Все гости размещены!")
}

func greetingGuest(name string, room int) {
	var calledRoom int
	dayParts := [3]string{"morning", "afternoon", "evening"}

	// Случайное время суток
	timeOfDay := dayParts[rand.Intn(3)]
	fmt.Printf("Good %s, Mr(Mrs) %s!\n", timeOfDay, name)

	attempts := 0
	maxAttempts := 3

	for {
		attempts++

		// Ограничение попыток
		if attempts > maxAttempts {
			fmt.Printf("❌ %s не смог вспомнить номер комнаты!\n\n", name)
			return
		}

		fmt.Println("Please call your room number?")
		scanner := bufio.NewScanner(os.Stdin)

		ok := scanner.Scan()
		if !ok {
			fmt.Println("❌ Input error!")
			continue
		}

		text := scanner.Text()
		_, err := fmt.Sscanf(text, "%d", &calledRoom)
		if err != nil {
			fmt.Println("❌ Invalid input data, use digits")
			continue
		}

		// Проверка валидного номера комнаты (1-30)
		if calledRoom < 1 || calledRoom > 30 {
			fmt.Println("❌ We don't have such room number! Try again!")
			continue
		}

		// Проверка правильности номера
		if room != calledRoom {
			fmt.Printf("❌ Sorry, it's incorrect room! Your room: %d. Try again!\n", room)
			continue
		}

		fmt.Printf("✅ Your room number is: %d. Welcome!\n\n", room)
		break
	}
}
