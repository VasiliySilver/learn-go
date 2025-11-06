package main

import (
	"fmt"
	"time"
)

// This function describes miner walking to the mine
func mine(transferPoint chan int, n int) {
	fmt.Println("Поход рабочего в шахту", n, "начался...")
	time.Sleep(1 * time.Second)
	fmt.Println("Поход в шахту", n, "закончился!")

	transferPoint <- 10

	fmt.Println("Поход номер", n, "уголь передал!")

	// ...
}

func main() {
	coal := 0

	transferPoint := make(chan int)

	initTime := time.Now()

	go mine(transferPoint, 1)
	go mine(transferPoint, 2)
	go mine(transferPoint, 3)
	go mine(transferPoint, 4)

	coal += <-transferPoint
	coal += <-transferPoint
	coal += <-transferPoint

	fmt.Println("Добыли", coal, "угля!")
	fmt.Println("Прошло времени:", time.Since(initTime))

}
