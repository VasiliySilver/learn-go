package main

import "fmt"


func main() {
	weather := map[int] int{
		11: +3,
		12: +6,
		13: +11,
		14: -4,
		15: +1,
	}

	c, ok := weather[30]
	fmt.Println(ok)
	fmt.Println(c)

	fmt.Println("Here is my map", weather)

	for i, v := range weather{
		fmt.Println("Index:", i, "Value:", v)
		fmt.Println(weather[i])
	}

	weather[30] = +1

	fmt.Println(weather)
}
