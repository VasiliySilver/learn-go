package main

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr[0])
	arr[0] += 10

	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])
	fmt.Println(arr[3])
	fmt.Println(arr[4])
	arr2 := [6]int{}

	for i, j := range arr {
		fmt.Println(i, j)
		arr2[i] = j * 2
		fmt.Println(arr2)
	}

	arr2[5] = 77
	fmt.Println(arr2)

	for i := 0; i < len(arr2); i++ {
		fmt.Println(i, arr2[i])
	}

	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Println(i, arr[i])
	}
	fmt.Println(arr2)
}
