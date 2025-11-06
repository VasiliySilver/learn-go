package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	intChn := make(chan int)
	strChn := make(chan string)

	go func() {
		time.Sleep(200 * time.Millisecond)
		i := 1
		intChn <- i
		i++
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		i := 1
		strChn <- "hi " + strconv.Itoa(i)
		i++

	}()

	time.Sleep(50 * time.Millisecond)
	select {
	case number := <-intChn:
		fmt.Println("intChn:", number)
	case str := <-strChn:

		fmt.Println("strChn:", str)
	default:
		fmt.Println("Никакой канал не готов!")
	}
}
