package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var beerDelicious bool
	var croutonsDelicious bool

	for {
		fmt.Println("Вкусное ли пиво?")
		getResponse(&beerDelicious)

		fmt.Println("Вкусные ли сухари")
		getResponse(&croutonsDelicious)

		if beerDelicious && croutonsDelicious {
			fmt.Println("Мы пойдем гулять")
		} else {
			fmt.Println("Мы пойдем гулять, только если пиво и сухарики будут вкусными")
		}

	}
}

func getResponse(linkBool *bool) {
	scanner := bufio.NewScanner(os.Stdin)
	ok := scanner.Scan()
	if !ok {
		fmt.Println("Ошиибка чтения из терминала...")
		return
	}

	text := scanner.Text()
	_, err := fmt.Sscanf(text, "%t", linkBool)
	if err != nil {
		fmt.Println("Ответ должен быть булевым значение true or false")
	}
}
