package main

func main() {
	myMap := map[string]int{
		"Alice": 30,
		"Bob":   25,
		"Carol": 27,
	}

	for name, age := range myMap {
		println("Name:", name, "Age:", age)
	}

}
