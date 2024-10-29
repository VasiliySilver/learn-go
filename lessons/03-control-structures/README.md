
# Урок 3: Управляющие конструкции в Go

## Содержание
1. [Условные операторы](#условные-операторы)
2. [Циклы](#циклы)
3. [Switch/Case](#switchcase)
4. [Практические задания](#практические-задания)

## Условные операторы

### If/Else
```go
// Базовый синтаксис
if условие {
    // код
} else {
    // код
}

// If с инициализацией
if x := getValue(); x > 10 {
    // x доступен только внутри if/else блока
}

// If/else if/else
if условие1 {
    // код
} else if условие2 {
    // код
} else {
    // код
}
```

### Операторы сравнения
- `==` равно
- `!=` не равно
- `<` меньше
- `>` больше
- `<=` меньше или равно
- `>=` больше или равно

### Логические операторы
- `&&` И (AND)
- `||` ИЛИ (OR)
- `!` НЕ (NOT)

## Циклы

### For
```go
// Стандартный цикл
for i := 0; i < 10; i++ {
    // код
}

// Цикл while (только for)
for условие {
    // код
}

// Бесконечный цикл
for {
    // код
    if условие {
        break    // выход из цикла
    }
    if другое_условие {
        continue // переход к следующей итерации
    }
}

// Цикл по коллекции
for index, value := range collection {
    // код
}
```

## Switch/Case

### Простой switch
```go
switch значение {
case вариант1:
    // код
case вариант2, вариант3:
    // код
default:
    // код
}
```

### Switch без выражения
```go
switch {
case условие1:
    // код
case условие2:
    // код
default:
    // код
}
```

### Switch с инициализацией
```go
switch result := getValue(); result {
case вариант1:
    // код
case вариант2:
    // код
}
```

## Практические задания

### Задание 1: Угадай число
Создайте игру, где программа загадывает число от 1 до 100, а пользователь должен его угадать.
- Программа должна сообщать "больше" или "меньше" после каждой попытки
- Подсчитывать количество попыток
- Спрашивать, хочет ли пользователь сыграть еще раз

### Задание 2: FizzBuzz
Напишите программу, которая выводит числа от 1 до 100, но:
- Для чисел, кратных 3, выводит "Fizz"
- Для чисел, кратных 5, выводит "Buzz"
- Для чисел, кратных и 3, и 5, выводит "FizzBuzz"

### Задание 3: Калькулятор с меню
Создайте калькулятор с использованием switch/case:
- Показывать меню с доступными операциями
- Поддерживать основные математические операции
- Обрабатывать деление на ноль
- Позволять пользователю выполнять несколько операций подряд

## Решения

### Решение задания 1: Угадай число
```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func playGame() {
    rand.Seed(time.Now().UnixNano())
    target := rand.Intn(100) + 1
    attempts := 0

    fmt.Println("Я загадал число от 1 до 100. Попробуй угадать!")

    for {
        var guess int
        fmt.Print("Твой вариант: ")
        fmt.Scan(&guess)
        attempts++

        if guess < target {
            fmt.Println("Больше!")
        } else if guess > target {
            fmt.Println("Меньше!")
        } else {
            fmt.Printf("Поздравляю! Ты угадал за %d попыток!\n", attempts)
            break
        }
    }
}

func main() {
    for {
        playGame()
        
        var playAgain string
        fmt.Print("Хочешь сыграть еще раз? (да/нет): ")
        fmt.Scan(&playAgain)
        
        if playAgain != "да" {
            fmt.Println("Спасибо за игру!")
            break
        }
    }
}
```

### Решение задания 2: FizzBuzz
```go
package main

import "fmt"

func main() {
    for i := 1; i <= 100; i++ {
        switch {
        case i%3 == 0 && i%5 == 0:
            fmt.Println("FizzBuzz")
        case i%3 == 0:
            fmt.Println("Fizz")
        case i%5 == 0:
            fmt.Println("Buzz")
        default:
            fmt.Println(i)
        }
    }
}
```

### Решение задания 3: Калькулятор
```go
package main

import (
    "fmt"
    "math"
)

func showMenu() {
    fmt.Println("\nКалькулятор")
    fmt.Println("1. Сложение")
    fmt.Println("2. Вычитание")
    fmt.Println("3. Умножение")
    fmt.Println("4. Деление")
    fmt.Println("5. Возведение в степень")
    fmt.Println("6. Выход")
}

func getNumbers() (float64, float64) {
    var a, b float64
    fmt.Print("Введите первое число: ")
    fmt.Scan(&a)
    fmt.Print("Введите второе число: ")
    fmt.Scan(&b)
    return a, b
}

func main() {
    for {
        showMenu()
        
        var choice int
        fmt.Print("Выберите операцию (1-6): ")
        fmt.Scan(&choice)

        if choice == 6 {
            fmt.Println("До свидания!")
            break
        }

        if choice < 1 || choice > 6 {
            fmt.Println("Неверный выбор!")
            continue
        }

        a, b := getNumbers()

        switch choice {
        case 1:
            fmt.Printf("%.2f + %.2f = %.2f\n", a, b, a+b)
        case 2:
            fmt.Printf("%.2f - %.2f = %.2f\n", a, b, a-b)
        case 3:
            fmt.Printf("%.2f * %.2f = %.2f\n", a, b, a*b)
        case 4:
            if b == 0 {
                fmt.Println("Ошибка: деление на ноль!")
            } else {
                fmt.Printf("%.2f / %.2f = %.2f\n", a, b, a/b)
            }
        case 5:
            fmt.Printf("%.2f ^ %.2f = %.2f\n", a, b, math.Pow(a, b))
        }
    }
}
```

## Дополнительные материалы
- [Go Tour - Flow control statements](https://tour.golang.org/flowcontrol/1)
- [Go by Example - If/Else](https://gobyexample.com/if-else)
- [Go by Example - For](https://gobyexample.com/for)
- [Go by Example - Switch](https://gobyexample.com/switch)

## Следующий урок
В следующем уроке мы изучим функции в Go: как их объявлять, использовать и работать с возвращаемыми значениями.