# Урок 7: Горутины и каналы в Go

## Содержание
1. [Горутины](#горутины)
2. [Каналы](#каналы)
3. [Паттерны конкурентности](#паттерны-конкурентности)
4. [Практические задания](#практические-задания)

## Горутины

### Основы горутин
```go
// Запуск горутины
go func() {
    // код, выполняемый в горутине
}()

// Пример с именованной функцией
func printNumbers() {
    for i := 0; i < 5; i++ {
        fmt.Printf("%d ", i)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    go printNumbers()
    time.Sleep(time.Second) // Ждем завершения горутины
}
```

### Синхронизация горутин
```go
var wg sync.WaitGroup

func worker(id int) {
    defer wg.Done()
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i)
    }
    wg.Wait()
}
```

## Каналы

### Базовые операции с каналами
```go
// Создание канала
ch := make(chan int)    // Небуферизованный канал
ch := make(chan int, 5) // Буферизованный канал

// Отправка и получение данных
ch <- 42        // Отправка
value := <-ch   // Получение

// Закрытие канала
close(ch)

// Проверка закрытия канала
value, ok := <-ch
if !ok {
    fmt.Println("Канал закрыт")
}
```

### Направление каналов
```go
func send(ch chan<- int) {    // Только отправка
    ch <- 42
}

func receive(ch <-chan int) { // Только получение
    value := <-ch
    fmt.Println(value)
}
```

## Паттерны конкурентности

### Select
```go
select {
case v1 := <-ch1:
    fmt.Println("Получено из ch1:", v1)
case v2 := <-ch2:
    fmt.Println("Получено из ch2:", v2)
case ch3 <- 42:
    fmt.Println("Отправлено в ch3")
default:
    fmt.Println("Нет доступных операций")
}
```

### Таймауты
```go
select {
case result := <-ch:
    fmt.Println("Получен результат:", result)
case <-time.After(time.Second):
    fmt.Println("Таймаут")
}
```

## Практические задания

### Задание 1: Параллельный обработчик данных
Создайте программу, которая параллельно обрабатывает данные:
- Несколько горутин-воркеров
- Канал для распределения задач
- Сбор результатов через отдельный канал

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

// Ваша реализация параллельного обработчика
```

### Задание 2: Пинг-понг
Реализуйте игру в пинг-понг между двумя горутинами:
- Обмен сообщениями через канал
- Подсчет количества обменов
- Завершение по таймауту

### Задание 3: Генератор последовательностей
Создайте конвейер обработки данных:
- Генератор чисел
- Фильтр четных/нечетных
- Умножитель
- Сборщик результатов

## Решения

### Решение задания 1: Параллельный обработчик данных
```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Job struct {
    ID   int
    Data int
}

type Result struct {
    JobID int
    Value int
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job.ID)
        time.Sleep(100 * time.Millisecond) // Имитация обработки
        
        // Обработка данных (например, умножение на 2)
        results <- Result{
            JobID: job.ID,
            Value: job.Data * 2,
        }
    }
}

func main() {
    numJobs := 10
    numWorkers := 3
    
    jobs := make(chan Job, numJobs)
    results := make(chan Result, numJobs)
    
    // Запуск воркеров
    var wg sync.WaitGroup
    for w := 1; w <= numWorkers; w++ {
        wg.Add(1)
        go worker(w, jobs, results, &wg)
    }
    
    // Отправка задач
    go func() {
        for j := 1; j <= numJobs; j++ {
            jobs <- Job{ID: j, Data: j * 10}
        }
        close(jobs)
    }()
    
    // Ожидание завершения всех воркеров
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Сбор результатов
    for result := range results {
        fmt.Printf("Result: Job %d = %d\n", result.JobID, result.Value)
    }
}
```

### Решение задания 2: Пинг-понг
```go
package main

import (
    "fmt"
    "time"
)

type Ball struct {
    hits int
}

func player(name string, table chan *Ball) {
    for {
        ball := <-table // Получаем мяч
        ball.hits++
        fmt.Printf("%s hit the ball. Hits: %d\n", name, ball.hits)
        time.Sleep(100 * time.Millisecond)
        table <- ball // Отбиваем мяч
    }
}

func main() {
    table := make(chan *Ball)
    done := make(chan bool)
    
    // Запускаем игроков
    go player("Ping", table)
    go player("Pong", table)
    
    // Начинаем игру
    table <- new(Ball)
    
    // Играем 1 секунду
    go func() {
        time.Sleep(time.Second)
        done <- true
    }()
    
    <-done
    fmt.Println("Game over")
}
```

### Решение задания 3: Генератор последовательностей
```go
package main

import (
    "fmt"
)

func generator(done chan bool) <-chan int {
    out := make(chan int)
    go func() {
        for i := 1; ; i++ {
            select {
            case out <- i:
            case <-done:
                close(out)
                return
            }
        }
    }()
    return out
}

func filter(done chan bool, in <-chan int, fn func(int) bool) <-chan int {
    out := make(chan int)
    go func() {
        for num := range in {
            if fn(num) {
                select {
                case out <- num:
                case <-done:
                    close(out)
                    return
                }
            }
        }
        close(out)
    }()
    return out
}

func multiply(done chan bool, in <-chan int, factor int) <-chan int {
    out := make(chan int)
    go func() {
        for num := range in {
            select {
            case out <- num * factor:
            case <-done:
                close(out)
                return
            }
        }
        close(out)
    }()
    return out
}

func main() {
    done := make(chan bool)
    defer close(done)
    
    // Создаем конвейер
    numbers := generator(done)
    
    // Фильтруем четные числа
    evenNumbers := filter(done, numbers, func(n int) bool {
        return n%2 == 0
    })
    
    // Умножаем на 10
    multiplied := multiply(done, evenNumbers, 10)
    
    // Получаем первые 10 результатов
    for i := 0; i < 10; i++ {
        fmt.Printf("%d ", <-multiplied)
    }
    fmt.Println()
}
```

## Дополнительные материалы
- [Go Tour - Goroutines](https://tour.golang.org/concurrency/1)
- [Go Tour - Channels](https://tour.golang.org/concurrency/2)
- [Go by Example - Goroutines](https://gobyexample.com/goroutines)
- [Go by Example - Channels](https://gobyexample.com/channels)
- [Effective Go - Concurrency](https://golang.org/doc/effective_go#concurrency)

## Следующий урок
В следующем уроке мы изучим работу с ошибками в Go, включая создание, обработку и распространение ошибок.