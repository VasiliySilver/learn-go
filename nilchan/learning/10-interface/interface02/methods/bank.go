package methods

import (
	"fmt"
	"math/rand"
)

type Bank struct{}

func NewBank() Bank {
	return Bank{}
}

func (c Bank) Pay(usd int) int {
	fmt.Println("Payment through the bank !")
	fmt.Println("Value of payment: ", usd, "dollars")

	return rand.Int()
}

func (c Bank) Cancel(id int) {
	fmt.Println("Bank operation is cancelled! ID: ", id)
}
