package main

func main() {
	var name string = "Alex"
	var age int = 30
	var rating float64 = 4.5
	var isStudent bool = false

	// base math operations
	println("Current age:", age)
	sum := age + 5
	println("Sum: age + 5: ", sum)
	diff := age - 5
	println("Difference: age - 5: ", diff)
	product := age * 2
	println("Product: age * 2: ", product)
	quotient := age / 2
	println("Quotient: age / 2: ", quotient)
	remainder := age % 7
	println("Remainder: age % 7: ", remainder)

	// compound assignment and increment/decrement
	age++ // age = age + 1
	println("Age after age++:", age)
	age-- // age = age - 1
	println("Age after age--:", age)
	age *= 2 // age = age * 2
	println("Age after age *= 2:", age)
	age -= 1 // age = age - 1
	println("Age after age -= 1:", age)
	age += 3 // age = age + 3
	println("Age after age += 3:", age)
	age /= 2 // age = age / 2  - что делает эта операция? - целочисленное деление
	println("Age after age /= 2:", age)
	age %= 5 // age = age % 5 - что делает эта операция? - остаток от деления
	println("Age after age %= 5:", age)

	println("Name:", name)
	println("Age:", age)
	println("Rating:", rating)
	println("Is Student:", isStudent)

}
