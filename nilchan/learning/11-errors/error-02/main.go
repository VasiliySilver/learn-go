package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type Car struct {
	Armor int
}

func (c *Car) Gas() (int, error) {
	if c.Armor-10 <= 0 {
		return 0, errors.New("we didn't initialize gas cause the car will be broken")
	}
	c.Armor -= 10
	kmch := rand.Intn(150)
	return kmch, nil
}

func main() {

	car := Car{
		Armor: 25,
	}

	for {
		fmt.Println("Car before", car)
		kmch, err := car.Gas()
		if err != nil {
			fmt.Println("Error push to gas:", err.Error())
			break
		}
		fmt.Println("kmch:", kmch)
		fmt.Println("Car after", car)
	}

}
