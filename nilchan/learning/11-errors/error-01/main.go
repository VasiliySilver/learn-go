package main

import (
	"errors"
	"fmt"

	"github.com/k0kubun/pp"
)

type User struct {
	Name    string
	Balance int
}

func Pay(user *User, usd int) error {
	if user.Balance < usd {
		return errors.New("not enough money")
	}

	user.Balance -= usd

	return nil

}

func main() {
	user := User{
		Name:    "Jhon",
		Balance: 10,
	}

	pp.Println(user)

	err := Pay(&user, 11)

	if err != nil {
		fmt.Println("The paymant was not succes, cause: ", err.Error())
	} else {
		pp.Println(user)
		fmt.Println("The payment was success")
	}

}
