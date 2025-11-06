package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Random times within the next 24 hours:")
	for i := 0; i < 5; i++ {
		now := time.Now()
		randomInt := rand.Intn(5)
		fmt.Println("Random Integer:", randomInt)
		// Генерируем случайное время в пределах следующих 24 часов
		randomDuration := time.Duration(rand.Intn(24*60*60)) * time.Second
		// Добавляем случайную длительность к текущему времени
		randomTime := now.Add(randomDuration)
		// Выводим сгенерированное случайное время
		fmt.Println(randomTime.Format(time.RFC1123))
		time.Sleep(500 * time.Millisecond)
	}
}
