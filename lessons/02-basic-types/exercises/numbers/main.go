package main

import (
	"fmt"
	"math"
)

func main() {
	// Арифметические операции с разными типами
	var a int = 10
	var b float64 = 3.14
	var c int32 = 5

	// Преобразование типов для операций
	fmt.Println("Демонстрация арифметических операций:")
	fmt.Printf("a + float64(c) = %.2f\n", float64(a)+float64(c))
	fmt.Printf("b * float64(a) = %.2f\n", b*float64(a))
	fmt.Printf("float64(a) / b = %.2f\n", float64(a)/b)

	// Вычисление площади круга
	const Pi = math.Pi
	radius := 5.0
	area := Pi * radius * radius
	fmt.Printf("\nПлощадь круга с радиусом %.1f: %.2f\n", radius, area)

	// Демонстрация переполнения
	var x uint8 = 255
	fmt.Printf("\nДемонстрация переполнения uint8:\n")
	fmt.Printf("x = %d\n", x)
	x++
	fmt.Printf("После x++ = %d\n", x) // Будет 0 из-за переполнения
}
