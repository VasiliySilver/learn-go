package main

import "fmt"

// --- ФУНКЦИИ ИЗ ТРЕТЬЕГО ЗАДАНИЯ ПО УКАЗАТЕЛЯМ ---

// CheckStringPointer принимает указатель на строку, выводит его,
// проверяет на nil, и разыменовывает, если указатель не nil.
func CheckStringPointer(ptr *string) {
	fmt.Printf("Проверка *string. Адрес: %p\n", ptr) // %p выводит адрес указателя

	if ptr == nil {
		fmt.Println("Указатель nil! Значение отсутствует.")
	} else {
		// Разыменование и вывод значения
		fmt.Printf("Значение: %s\n", *ptr)
	}
}

// CheckIntPointer выполняет проверку для *int.
func CheckIntPointer(ptr *int) {
	fmt.Printf("Проверка *int. Адрес: %p\n", ptr)

	if ptr == nil {
		fmt.Println("Указатель nil! Значение отсутствует.")
	} else {
		fmt.Printf("Значение: %d\n", *ptr)
	}
}

// CheckFloatPointer выполняет проверку для *float64.
func CheckFloatPointer(ptr *float64) {
	fmt.Printf("Проверка *float64. Адрес: %p\n", ptr)

	if ptr == nil {
		fmt.Println("Указатель nil! Значение отсутствует.")
	} else {
		fmt.Printf("Значение: %.2f\n", *ptr)
	}
}

// CheckBoolPointer выполняет проверку для *bool.
func CheckBoolPointer(ptr *bool) {
	fmt.Printf("Проверка *bool. Адрес: %p\n", ptr)

	if ptr == nil {
		fmt.Println("Указатель nil! Значение отсутствует.")
	} else {
		fmt.Printf("Значение: %t\n", *ptr)
	}
}

// --- ОСНОВНАЯ ФУНКЦИЯ ---

func main() {
	// 1. Создание переменных всех типов
	name := "Sam"
	age := 17
	rating := 4.3
	isClosed := true

	// 2. Создание указателей на переменные
	ptrName := &name
	ptrAge := &age
	ptrRating := &rating
	ptrIsClosed := &isClosed

	// 3. Создание nil-указателей
	var nilPtrName *string
	var nilPtrAge *int
	var nilPtrRating *float64
	var nilPtrIsClosed *bool

	fmt.Println("\n===== ПРОВЕРКА НЕ-NIL УКАЗАТЕЛЕЙ =====")
	CheckStringPointer(ptrName)
	CheckIntPointer(ptrAge)
	CheckFloatPointer(ptrRating)
	CheckBoolPointer(ptrIsClosed)

	fmt.Println("\n===== ПРОВЕРКА NIL-УКАЗАТЕЛЕЙ =====")
	CheckStringPointer(nilPtrName)
	CheckIntPointer(nilPtrAge)
	CheckFloatPointer(nilPtrRating)
	CheckBoolPointer(nilPtrIsClosed)
}
