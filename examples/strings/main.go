package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Print("Введите строку: ")
    var input string
    fmt.Scanln(&input)

    // Длина строки
    fmt.Printf("Длина строки: %d\n", len(input))

    // Переворачиваем строку
    runes := []rune(input)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    fmt.Printf("Перевернутая строка: %s\n", string(runes))

    // Подсчет гласных
    vowels := "aeiouAEIOU"
    count := 0
    for _, char := range input {
        if strings.ContainsRune(vowels, char) {
            count++
        }
    }
    fmt.Printf("Количество гласных: %d\n", count)
}

