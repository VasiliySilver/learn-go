package methods

import (
	"fmt"
	"math/rand"
)

type Crypto struct{}

func NewCrypto() Crypto {
	return Crypto{}
}

func (c Crypto) Pay(usd int) int {
	fmt.Println("Payment with cryptocurrency!")
	fmt.Println("Value of payment: ", usd, "USDT")

	return rand.Int()
}

func (c Crypto) Cancel(id int) {
	fmt.Println("Crypto operation is cancelled! ID: ", id)
}
