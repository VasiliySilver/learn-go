package methods

import (
	"fmt"
	"math/rand"
)

type Bonus struct{}

func NewBonus() Bonus {
	return Bonus{}
}

func (c Bonus) Pay(usd int) int {
	fmt.Println("Payment through the bonus !")
	fmt.Println("Value of payment: ", usd, "bonuses")

	return rand.Int()
}

func (c Bonus) Cancel(id int) {
	fmt.Println("Bonus operation is cancelled! ID: ", id)
}
