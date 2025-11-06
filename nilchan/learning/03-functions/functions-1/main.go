package main

import (
	"fmt"
)

func main() {
	fmt.Println("Before function calling")
	hello()
	fmt.Println("After function calling")
	aboba()
	num := 5
	result := square(num)
	fmt.Println("Square of", num, "is", result)

	restaurant()

}

func hello() {
	fmt.Println("Hello from function!")
}

func aboba() {
	a := 0
	b := 10
	c := 1

	fmt.Println("Summ of abc", a+b+c)
}

// Предположим мы работаем оффициантами в ресторане и у нас задача для каждого гостя:
// 1 - накрыть стол
// 2 - поприветствовать его по имени
// 3 - принять заказ
// 4 - принести блюдо

func restaurant() {
	restaurantGuest("Alice", "Pasta")
	restaurantGuest("Bob", "Steak")
	restaurantGuest("Charlie", "Salad")
}

func restaurantGuest(name string, order string) {
	fmt.Println("Накрываю на стол", name)
	fmt.Println("Приветствую", name)
	fmt.Println("Принимаю заказ у", name, "на", order)
	fmt.Println("Приношу", order, "для", name)
}

func square(n int) int {
	fmt.Println("Мы вошли в функцию square, переменная n:", n)
	fmt.Println("Вычисляем квадрат...")
	fmt.Println("Результат:", n*n)
	return n * n
}
