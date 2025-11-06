package main

import (
	"interface02/methods"
	"interface02/payments"

	"github.com/k0kubun/pp"
)

func main() {
	method := methods.NewBonus()
	paymentsModule := payments.NewPaymentModule(method)

	paymentsModule.Pay("Burger", 5)
	idPhone := paymentsModule.Pay("Phone", 500)
	idGame := paymentsModule.Pay("Game", 20)

	paymentsModule.Cancel(idPhone)

	allInfo := paymentsModule.AllInfo()
	pp.Println(allInfo)

	gameInfo := paymentsModule.Info(idGame)
	pp.Println(gameInfo)
}
