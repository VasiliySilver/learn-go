package main

import (
	"fmt"
	"math"
)

// Вычисляет значение функции f(x, y) = x^2 + sqrt(y)
func CalculateFunction(x, y float64) float64 {
	// 1. Вычисляем x в квадрате (x^2)
	xSquared := math.Pow(x, 2)

	// 2. Проверяем что y < 0
	if y < 0 {
		fmt.Printf("Внимание - y < 0. Y - %.2f. Возвращается только X^2\n", y)
		return xSquared
	}

	// 3. Вычисляем квадратный корень из y
	sqrtY := math.Sqrt(y)

	return xSquared + sqrtY

}

func main() {
	var x float64
	var y float64

	fmt.Println("Введите поочередно два числа: ")
	fmt.Println("Введите x: ")
	fmt.Scan(&x)
	fmt.Println("")
	fmt.Println("Введите y: ")
	fmt.Scan(&y)
	fmt.Println("")

	result := CalculateFunction(x, y)

	fmt.Println("===============================================")
	fmt.Printf("Результат для f(x, y) = x^2 + sqrt(y) = %.2f\n", result)
	fmt.Println("===============================================")

}
