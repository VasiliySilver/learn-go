package main

func main() {
	name := "Original Name"
	println("Name before change:", name)

	changeName(&name)

	println("Name after change:", name)

}

func changeName(name *string) {
	*name = "Changed Name"
}
