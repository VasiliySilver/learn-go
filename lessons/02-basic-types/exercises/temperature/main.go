package main

import "fmt"

const (
	FreezingC = 0.0
	BoilingC  = 100.0
	FreezingF = 32.0
	BoilingF  = 212.0
)

func celsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func main() {
	for {
		fmt.Println("\nКалькулятор температуры")
		fmt.Println("1. Цельсий -> Фаренгейт")
		fmt.Println("2. Фаренгейт -> Цельсий")
		fmt.Println("3. Выход")
		fmt.Print("Выберите действие (1-3): ")

		var choice int
		fmt.Scan(&choice)

		if choice == 3 {
			fmt.Println("До свидания!")
			break
		}

		var temp float64
		switch choice {
		case 1:
			fmt.Print("Введите температуру в градусах Цельсия: ")
			fmt.Scan(&temp)
			result := celsiusToFahrenheit(temp)
			fmt.Printf("%.2f°C = %.2f°F\n", temp, result)

			// Дополнительная информация
			if temp <= FreezingC {
				fmt.Println("Вода замерзает при этой температуре!")
			} else if temp >= BoilingC {
				fmt.Println("Вода кипит при этой температуре!")
			}

		case 2:
			fmt.Print("Введите температуру в градусах Фаренгейта: ")
			fmt.Scan(&temp)
			result := fahrenheitToCelsius(temp)
			fmt.Printf("%.2f°F = %.2f°C\n", temp, result)

			// Дополнительная информация
			if temp <= FreezingF {
				fmt.Println("Вода замерзает при этой температуре!")
			} else if temp >= BoilingF {
				fmt.Println("Вода кипит при этой температуре!")
			}

		default:
			fmt.Println("Неверный выбор. Пожалуйста, выберите 1, 2 или 3.")
		}
	}
}
