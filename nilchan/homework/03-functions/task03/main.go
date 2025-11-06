package main

import (
	"fmt"
	"math"
)

func CalculateDiscriminant(a, b, c float64) float64 {
	// КОРРЕКТНАЯ ФОРМУЛА: D = b^2 - 4ac
	return b*b - 4*a*c
}

// FindRoots теперь возвращает найденные корни и булево значение 'found',
// чтобы указать, сколько корней найдено (2, 1, или 0).
func FindRoots(a, b, discriminant float64) (x1, x2 float64, count int) {
	if discriminant > 0 {
		sqrtD := math.Sqrt(discriminant)
		x1 = (-b + sqrtD) / (2 * a)
		x2 = (-b - sqrtD) / (2 * a)
		count = 2
	} else if discriminant == 0 {
		x1 = -b / (2 * a)
		x2 = x1 // Второй корень совпадает с первым
		count = 1
	} else {
		// D < 0, корней нет. x1, x2 останутся равны 0 (zero value)
		count = 0
	}
	return
}

func main() {
	var a, b, c float64 // Сокращенное объявление

	fmt.Println("Поочередно введите три числа (a, b, c): ")
	fmt.Scan(&a, &b, &c)
	fmt.Printf("A: %.2f, B: %.2f, C: %.2f\n", a, b, c)

	if a == 0 {
		fmt.Println("Ошибка - коэффициент а не может быть равен нулю.")
		return
	}

	discriminant := CalculateDiscriminant(a, b, c)
	fmt.Printf("Дискриминант: %.2f\n", discriminant)

	// Присваиваем возвращаемые значения
	x1, x2, count := FindRoots(a, b, discriminant)

	// Логика вывода теперь находится в main
	fmt.Println("=========================================")
	if count == 2 {
		fmt.Println("Уравнение имеет два действительных корня:")
		fmt.Printf("x1 = %.2f\n", x1)
		fmt.Printf("x2 = %.2f\n", x2)
	} else if count == 1 {
		fmt.Println("Уравнение имеет один действительный корень:")
		fmt.Printf("x = %.2f\n", x1)
	} else { // count == 0
		fmt.Println("Дискриминант отрицателен. Действительных корней нет.")
	}
	fmt.Println("=========================================")
}
