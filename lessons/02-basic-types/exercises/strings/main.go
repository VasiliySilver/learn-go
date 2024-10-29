package main

import (
	"fmt"
	"strings"
)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func countVowels(s string) int {
	vowels := "аеёиоуыэюяАЕЁИОУЫЭЮЯaeiouAEIOU"
	count := 0
	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}

func main() {
	fmt.Print("Введите строку: ")
	var input string
	fmt.Scanln(&input)

	// Длина строки
	fmt.Printf("\nДлина строки: %d символов\n", len([]rune(input)))

	// Перевернутая строка
	reversed := reverseString(input)
	fmt.Printf("Перевернутая строка: %s\n", reversed)

	// Подсчет гласных
	vowelsCount := countVowels(input)
	fmt.Printf("Количество гласных букв: %d\n", vowelsCount)
}
