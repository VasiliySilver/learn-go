# Урок 4: Функции в Go

## Содержание
1. [Объявление и вызов функций](#объявление-и-вызов-функций)
2. [Параметры и возвращаемые значения](#параметры-и-возвращаемые-значения)
3. [Анонимные функции и замыкания](#анонимные-функции-и-замыкания)
4. [Практические задания](#практические-задания)

## Объявление и вызов функций

### Базовый синтаксис
```go
// Простая функция
func sayHello() {
    fmt.Println("Hello!")
}

// Функция с параметрами
func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

// Функция с возвращаемым значением
func add(a, b int) int {
    return a + b
}
```

### Множественные возвращаемые значения
```go
// Функция с несколькими возвращаемыми значениями
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("деление на ноль")
    }
    return a / b, nil
}

// Именованные возвращаемые значения
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return // "голый" return
}
```

## Параметры и возвращаемые значения

### Варианты параметров
```go
// Variadic функция (переменное число аргументов)
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Функция с указателями
func increment(x *int) {
    *x++
}
```

### Функции как значения
```go
// Функция как тип
type Operation func(int, int) int

// Функция, принимающая функцию как параметр
func calculate(a, b int, op Operation) int {
    return op(a, b)
}
```

## Анонимные функции и замыкания

### Анонимные функции
```go
// Немедленно вызываемая анонимная функция
func main() {
    func() {
        fmt.Println("Анонимная функция")
    }()

    // Анонимная функция с параметрами
    func(x int) {
        fmt.Printf("Значение: %d\n", x)
    }(42)
}
```

### Замыкания
```go
// Функция, возвращающая функцию (замыкание)
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

## Практические задания

### Задание 1: Калькулятор с функциями
Создайте улучшенную версию калькулятора из предыдущего урока:
- Вынесите математические операции в отдельные функции
- Используйте функцию как тип для хранения операций
- Обработайте ошибки с помощью множественных возвращаемых значений

```go
package main

import (
    "errors"
    "fmt"
)

// Определяем тип для математических операций
type MathFunc func(float64, float64) (float64, error)

// Функции для операций
func add(a, b float64) (float64, error) {
    return a + b, nil
}

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("деление на ноль")
    }
    return a / b, nil
}

func main() {
    // Ваша реализация калькулятора
}
```

### Задание 2: Обработка данных
Создайте набор функций для обработки слайса чисел:
- Функция для поиска минимального и максимального значения (множественные возвращаемые значения)
- Функция для фильтрации чисел (принимает функцию-предикат)
- Функция для преобразования чисел (принимает функцию преобразования)

```go
package main

import "fmt"

// Найти минимальное и максимальное значение
func minMax(numbers []int) (min, max int) {
    // Ваш код
}

// Фильтрация чисел по предикату
func filter(numbers []int, predicate func(int) bool) []int {
    // Ваш код
}

// Преобразование чисел
func transform(numbers []int, transformer func(int) int) []int {
    // Ваш код
}
```

### Задание 3: Генератор последовательностей
Создайте набор функций для генерации различных последовательностей чисел:
- Функция-генератор для чисел Фибоначчи
- Функция-генератор для степеней двойки
- Функция высшего порядка для создания пользовательских последовательностей

```go
package main

import "fmt"

// Генератор чисел Фибоначчи
func fibonacciGenerator() func() int {
    // Ваш код
}

// Генератор степеней двойки
func powerOfTwoGenerator() func() int {
    // Ваш код
}

// Создание пользовательского генератора
func makeGenerator(start int, step func(int) int) func() int {
    // Ваш код
}
```

## Решения

### Решение задания 1: Калькулятор с функциями
```go
package main

import (
    "errors"
    "fmt"
)

type MathFunc func(float64, float64) (float64, error)

func add(a, b float64) (float64, error) {
    return a + b, nil
}

func subtract(a, b float64) (float64, error) {
    return a - b, nil
}

func multiply(a, b float64) (float64, error) {
    return a * b, nil
}

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("деление на ноль")
    }
    return a / b, nil
}

func calculate(a, b float64, operation MathFunc) {
    result, err := operation(a, b)
    if err != nil {
        fmt.Printf("Ошибка: %v\n", err)
        return
    }
    fmt.Printf("Результат: %.2f\n", result)
}

func main() {
    operations := map[string]MathFunc{
        "+": add,
        "-": subtract,
        "*": multiply,
        "/": divide,
    }

    for {
        fmt.Print("\nВведите операцию (+, -, *, /) или 'q' для выхода: ")
        var op string
        fmt.Scan(&op)

        if op == "q" {
            break
        }

        operation, exists := operations[op]
        if !exists {
            fmt.Println("Неизвестная операция")
            continue
        }

        var a, b float64
        fmt.Print("Введите первое число: ")
        fmt.Scan(&a)
        fmt.Print("Введите второе число: ")
        fmt.Scan(&b)

        calculate(a, b, operation)
    }
}
```

### Решение задания 2: Обработка данных
```go
package main

import "fmt"

func minMax(numbers []int) (min, max int) {
    if len(numbers) == 0 {
        return 0, 0
    }
    
    min, max = numbers[0], numbers[0]
    for _, num := range numbers[1:] {
        if num < min {
            min = num
        }
        if num > max {
            max = num
        }
    }
    return
}

func filter(numbers []int, predicate func(int) bool) []int {
    result := make([]int, 0)
    for _, num := range numbers {
        if predicate(num) {
            result = append(result, num)
        }
    }
    return result
}

func transform(numbers []int, transformer func(int) int) []int {
    result := make([]int, len(numbers))
    for i, num := range numbers {
        result[i] = transformer(num)
    }
    return result
}

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    // Пример использования minMax
    min, max := minMax(numbers)
    fmt.Printf("Мин: %d, Макс: %d\n", min, max)

    // Пример использования filter (четные числа)
    evens := filter(numbers, func(x int) bool {
        return x%2 == 0
    })
    fmt.Printf("Четные числа: %v\n", evens)

    // Пример использования transform (квадраты чисел)
    squares := transform(numbers, func(x int) int {
        return x * x
    })
    fmt.Printf("Квадраты: %v\n", squares)
}
```

### Решение задания 3: Генератор последовательностей
```go
package main

import "fmt"

func fibonacciGenerator() func() int {
    a, b := 0, 1
    return func() int {
        result := a
        a, b = b, a+b
        return result
    }
}

func powerOfTwoGenerator() func() int {
    current := 1
    return func() int {
        result := current
        current *= 2
        return result
    }
}

func makeGenerator(start int, step func(int) int) func() int {
    current := start
    return func() int {
        result := current
        current = step(current)
        return result
    }
}

func main() {
    // Пример использования генератора Фибоначчи
    fib := fibonacciGenerator()
    fmt.Println("Числа Фибоначчи:")
    for i := 0; i < 10; i++ {
        fmt.Printf("%d ", fib())
    }
    fmt.Println()

    // Пример использования генератора степеней двойки
    pow2 := powerOfTwoGenerator()
    fmt.Println("Степени двойки:")
    for i := 0; i < 10; i++ {
        fmt.Printf("%d ", pow2())
    }
    fmt.Println()

    // Пример использования пользовательского генератора
    // Генератор нечетных чисел
    oddGen := makeGenerator(1, func(x int) int {
        return x + 2
    })
    fmt.Println("Нечетные числа:")
    for i := 0; i < 10; i++ {
        fmt.Printf("%d ", oddGen())
    }
    fmt.Println()
}
```

## Дополнительные материалы
- [Go Tour - Functions](https://tour.golang.org/basics/4)
- [Go by Example - Functions](https://gobyexample.com/functions)
- [Go by Example - Closures](https://gobyexample.com/closures)
- [Effective Go - Functions](https://golang.org/doc/effective_go#functions)

## Следующий урок
В следующем уроке мы изучим составные типы данных в Go: массивы, срезы, карты и структуры.