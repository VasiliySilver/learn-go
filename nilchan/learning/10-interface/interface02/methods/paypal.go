package methods

import (
	"fmt"
	"math/rand"
)

type PayPal struct{}

func NewPayPal() PayPal {
	return PayPal{}
}

func (c PayPal) Pay(usd int) int {
	fmt.Println("Payment through the bank !")
	fmt.Println("Value of payment: ", usd, "dollars")

	return rand.Int()
}

func (c PayPal) Cancel(id int) {
	fmt.Println("PayPal operation is cancelled! ID: ", id)
}
