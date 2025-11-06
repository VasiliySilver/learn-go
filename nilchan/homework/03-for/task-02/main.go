package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	// We have simple player
	// We have some object movin to the player
	// The player must simply jump or die
	// 🎮 Игрок (◆) прыгает
	// 🚧 Препятствия (█) появляются случайно
	// ⬆️ Гравитация и физика прыжка
	// 💥 Проверка столкновений
	// 📊 Счёт очков
	// obstaccle := rand.Intn(4)
	score := 0
	level := 1

	fmt.Println("\nДобро пожаловать в Geometry Dash!: 🎮")
	fmt.Println("=====================================")
	fmt.Println("Задача, прыгать выше препятствий!")
	fmt.Println("За каждый успешный прыжок +5 очков")
	fmt.Printf("=====================================\n")
	fmt.Println("Для старта нажми Enter:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	for {
		fmt.Printf("\n Уровень %d\n", level)
		fmt.Println("--------------")

		maxObstacle := 2 + (level / 2)
		if maxObstacle > 3 {
			maxObstacle = 5
		}

		fmt.Println("Игрок приближается к препятствию ◆")
		obstacle := rand.Intn(maxObstacle) + 1
		fmt.Println("Игрок прыгает через препятствие: ")
		fmt.Printf("Выбери цифру (1-%d): ", maxObstacle)

		jump := scanJump(maxObstacle)

		if jump > obstacle {
			fmt.Println("Успешный прыжок! ✅")
			score += 5
			level++
			fmt.Printf("Счёт: %d | Уровень: %d\n", score, level)
		} else {
			fmt.Println("\n\n💥 Ты попал в препятствие!")
			fmt.Println("====================================")
			fmt.Printf("🏁 Игра окончена!\n")
			fmt.Printf("📊 Финальный счёт: %d очков\n", score)
			fmt.Printf("🏆 Пройдено уровней: %d\n", level-1)
			fmt.Println("====================================")
			break
		}
		fmt.Printf("Высота препятствия: %d, высота прыжка: %d\n", obstacle, jump)
		fmt.Println("Игрок движется к следующему препятствию")

	}

	// Приближаюсь к препятствию ◆
	// Прыгаю через препятствие
	// Игрок попал в препятствие 💥
	// Двигается к следующему препятствию
}

func scanJump(maxHeight int) int {
	var res int
	for {
		scanner := bufio.NewScanner(os.Stdin)
		ok := scanner.Scan()

		if !ok {
			fmt.Println("Ошибка ввода")
			continue
		}

		text := scanner.Text()
		_, err := fmt.Sscanf(text, "%d", &res)

		if err != nil {
			fmt.Println("Невалидный формат введите от 1-ого до 3")
			continue
		}

		if res < 1 || res > maxHeight {
			fmt.Printf("Цифра должна быть от 1-ого до %d: ", maxHeight)
			continue
		}

		return res

	}

}
