package main

import (
	"fmt"
	"learn-go/nilchan/learning/07-modules-and-packages/study/greeting"
	"learn-go/nilchan/learning/07-modules-and-packages/study/user"
)

func main() {
	greeting.SayHello("John")
	greeting.SayBad("Bad")
	fmt.Println("Modules and Packages Study")
	fmt.Println("Random Number:", greeting.GiveMeInt())
	user := user.NewUser("Alice", 30)
	user.PrintUser()
	user.PrettyPrint()
}
