package main

import "fmt"

var GlobalCounter int = 10

func IncrementByOne() {
	GlobalCounter += 1
	fmt.Printf("Функция IncrementByOne: GlobalCounter = %d\n", GlobalCounter)
}

func DoubleValue() {
	GlobalCounter *= 2
	fmt.Printf("Функция DoubleValue: GlobalCounter = %d\n", GlobalCounter)
}

func ZeroValue() {
	GlobalCounter = 0
	fmt.Printf("Функция ZeroValue: GlobalCounter = %d\n", GlobalCounter)
}

func main() {
	IncrementByOne()
	DoubleValue()
	ZeroValue()
	fmt.Printf("GlobalCounter after all callings = %d\n", GlobalCounter)
}
