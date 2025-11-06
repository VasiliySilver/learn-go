package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	transferPoint := make(chan int)

	// miner
	go func() {
		iterations := 3 + rand.Intn(4)

		fmt.Println("Iterations:", iterations)

		for i := 1; i <= iterations; i++ {
			time.Sleep(300 * time.Millisecond)
			transferPoint <- 10
		}

		close(transferPoint)

	}()

	coal := 0
	// for {
	// 	v, ok := <-transferPoint
	// 	if !ok {
	// 		fmt.Println("Шахтер отработал всю шахту")
	// 		break
	// 	}
	// 	coal += v
	// 	fmt.Println("Coal:", coal)
	// }

	// Обязательно использовать с close иначе будет deadlock
	for v := range transferPoint {
		coal += v
		fmt.Println("Coal:", coal)
	}

	fmt.Println("Суммарно добытый уголь -", coal)

}
