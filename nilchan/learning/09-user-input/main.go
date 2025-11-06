package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	// Введите комнадо: добавить лук
	// Вы хотите добавить лук

	// Введите команду: удалить морковь
	// Вы хотите кажется хотите удалить морковь

	// Неизвестная команда
	// |
	// help
	// Вот список команд!
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Введите команду:")

		ok := scanner.Scan()
		if !ok {
			fmt.Println("Ошибка ввода")
			return
		}

		// if ok := scanner.Scan(); !ok {
		// 	fmt.Println("Ошибка ввода")
		// 	return
		// }

		text := scanner.Text()
		fields := strings.Fields(text)
		pp.Println(fields)
		pp.Println(fields[0])

		// fmt.Println("Неизвестная команда")
		if len(fields) == 0 {
			fmt.Println("Вы ничего не ввели")
			return
		}

		fmt.Println("Вы ввели:", text)
		fmt.Println("Комманда:", fields[0])

		cmd := fields[0]

		if cmd == "exit" || cmd == "выход" {
			fmt.Println("Выход из программы")
			return
		} else if cmd == "добавить" {
			str := ""
			for i := 1; i < len(fields); i++ {
				str += fields[i]

				if i < len(fields)-1 {
					str += " "
				}
			}
			fmt.Println("Вы хотите добавить", str)
		} else if cmd == "удалить" {
			str := ""
			for i := 1; i < len(fields); i++ {
				str += fields[i]

				if i < len(fields)-1 {
					str += " "
				}
			}
			fmt.Println("Вы хотите кажется хотите удалить", str)
		} else if cmd == "помоги" || cmd == "help" {
			fmt.Println("Вот список команд:")
			fmt.Println("добавить <что-то>")
			fmt.Println("удалить <что-то>")
			fmt.Println("help - показать эту справку")
		} else {
			fmt.Println("Вы ввели неизвестную команду:", cmd)
		}
	}

}
