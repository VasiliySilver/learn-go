package main

import "fmt"

type Auto interface {
	StepOnGas()
	StepOnBrake()
}

type BMW struct{}

func (b BMW) StepOnGas() {
	fmt.Println("Я БМВ 550 лошадиных сил, жмем на газ!")
}

func (b BMW) StepOnBrake() {
	fmt.Println("Я БМВ, тормоз в пол, хорошее торможение!")
}

type Zhiga struct{}

func (b Zhiga) StepOnGas() {
	fmt.Println("Я Жига 2107 пробую не развалиться!")
}

func (b Zhiga) StepOnBrake() {
	fmt.Println("Я Жига тормоз в пол, надеюсь выживу!")
}

func rideBMW(bmw BMW) {
	fmt.Println("Я водитель! Сажусь в машину и нажимаю на газ...")
	bmw.StepOnGas()
}

func rideZhiga(zhiga Zhiga) {
	fmt.Println("Я водитель! Сажусь в машину и нажимаю на газ...")
	zhiga.StepOnGas()
}

func ride(auto Auto) {
	fmt.Println("Я водитель! Сажусь в машину и нажимаю на газ...")
	auto.StepOnGas()
}

func main() {
	bmw := BMW{}
	rideBMW(bmw)
	ride(bmw)

	zhiga := Zhiga{}
	zhiga.StepOnGas()
	rideZhiga(zhiga)
	ride(zhiga)
}
