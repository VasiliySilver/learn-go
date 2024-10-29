# Урок 5: Составные типы данных в Go

## Содержание
1. [Массивы и срезы](#массивы-и-срезы)
2. [Карты (map)](#карты-map)
3. [Структуры](#структуры)
4. [Практические задания](#практические-задания)

## Массивы и срезы

### Массивы
```go
// Объявление массива
var arr [5]int                    // Массив из 5 целых чисел
numbers := [3]int{1, 2, 3}        // Инициализация при объявлении
matrix := [2][3]int{{1,2,3}, {4,5,6}}  // Многомерный массив

// Длина массива
length := len(arr)

// Обращение к элементам
arr[0] = 1     // Изменение элемента
value := arr[1] // Получение элемента
```

### Срезы
```go
// Объявление среза
var slice []int                // Пустой срез
numbers := []int{1, 2, 3, 4}   // Срез с элементами
matrix := [][]int{{1,2}, {3,4}} // Многомерный срез

// Создание среза с помощью make
slice := make([]int, 5)    // Длина и емкость 5
slice := make([]int, 3, 5) // Длина 3, емкость 5

// Операции со срезами
slice = append(slice, 6)        // Добавление элемента
slice2 := slice[1:4]           // Получение подсреза
copy(dst, src)                 // Копирование срезов
```

## Карты (map)

### Основные операции
```go
// Объявление карты
var m map[string]int                  // Пустая карта
scores := map[string]int{             // Инициализация при объявлении
    "Alice": 98,
    "Bob":   87,
}
grades := make(map[string]int)        // Создание с помощью make

// Работа с элементами
m["key"] = 42                        // Добавление/изменение
value, exists := m["key"]            // Проверка существования
delete(m, "key")                     // Удаление элемента
```

## Структуры

### Определение и использование
```go
// Определение структуры
type Person struct {
    Name    string
    Age     int
    Address string
}

// Создание экземпляра
p1 := Person{"John", 25, "New York"}
p2 := Person{
    Name: "Alice",
    Age: 30,
    Address: "London",
}
var p3 Person // Все поля имеют нулевые значения

// Доступ к полям
p1.Name = "John Doe"
age := p1.Age
```

## Практические задания

### Задание 1: Управление списком задач
Создайте программу для управления списком задач (todo list):
- Используйте срезы для хранения задач
- Реализуйте добавление, удаление и отображение задач
- Добавьте возможность отмечать задачи как выполненные

```go
package main

type Task struct {
    ID          int
    Description string
    Done        bool
}

type TodoList struct {
    tasks []Task
}

// Ваша реализация методов
```

### Задание 2: Телефонная книга
Создайте телефонную книгу с использованием карт:
- Хранение контактов (имя -> номер телефона)
- Добавление, удаление и поиск контактов
- Возможность хранения нескольких номеров для одного контакта

### Задание 3: Библиотечный каталог
Создайте систему управления библиотечным каталогом:
- Используйте структуры для хранения информации о книгах
- Реализуйте поиск по различным критериям
- Добавьте учет доступности книг

## Решения

### Решение задания 1: Управление списком задач
```go
package main

import (
    "fmt"
    "strings"
)

type Task struct {
    ID          int
    Description string
    Done        bool
}

type TodoList struct {
    tasks []Task
    lastID int
}

func NewTodoList() *TodoList {
    return &TodoList{
        tasks: make([]Task, 0),
        lastID: 0,
    }
}

func (tl *TodoList) AddTask(description string) {
    tl.lastID++
    task := Task{
        ID:          tl.lastID,
        Description: description,
        Done:        false,
    }
    tl.tasks = append(tl.tasks, task)
}

func (tl *TodoList) MarkDone(id int) bool {
    for i := range tl.tasks {
        if tl.tasks[i].ID == id {
            tl.tasks[i].Done = true
            return true
        }
    }
    return false
}

func (tl *TodoList) RemoveTask(id int) bool {
    for i := range tl.tasks {
        if tl.tasks[i].ID == id {
            tl.tasks = append(tl.tasks[:i], tl.tasks[i+1:]...)
            return true
        }
    }
    return false
}

func (tl *TodoList) ListTasks() {
    if len(tl.tasks) == 0 {
        fmt.Println("Список задач пуст")
        return
    }
    
    for _, task := range tl.tasks {
        status := " "
        if task.Done {
            status = "✓"
        }
        fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Description)
    }
}

func main() {
    todoList := NewTodoList()
    
    for {
        var command string
        fmt.Print("\nВведите команду (add/done/remove/list/quit): ")
        fmt.Scan(&command)
        
        switch strings.ToLower(command) {
        case "add":
            var desc string
            fmt.Print("Введите описание задачи: ")
            fmt.Scan(&desc)
            todoList.AddTask(desc)
            fmt.Println("Задача добавлена")
            
        case "done":
            var id int
            fmt.Print("Введите ID задачи: ")
            fmt.Scan(&id)
            if todoList.MarkDone(id) {
                fmt.Println("Задача отмечена как выполненная")
            } else {
                fmt.Println("Задача не найдена")
            }
            
        case "remove":
            var id int
            fmt.Print("Введите ID задачи: ")
            fmt.Scan(&id)
            if todoList.RemoveTask(id) {
                fmt.Println("Задача удалена")
            } else {
                fmt.Println("Задача не найдена")
            }
            
        case "list":
            todoList.ListTasks()
            
        case "quit":
            fmt.Println("До свидания!")
            return
            
        default:
            fmt.Println("Неизвестная команда")
        }
    }
}
```

### Решение задания 2: Телефонная книга
```go
package main

import (
    "fmt"
    "strings"
)

type PhoneBook struct {
    contacts map[string][]string
}

func NewPhoneBook() *PhoneBook {
    return &PhoneBook{
        contacts: make(map[string][]string),
    }
}

func (pb *PhoneBook) AddContact(name, phone string) {
    name = strings.ToLower(name)
    if numbers, exists := pb.contacts[name]; exists {
        // Проверяем, нет ли уже такого номера
        for _, num := range numbers {
            if num == phone {
                return
            }
        }
        pb.contacts[name] = append(numbers, phone)
    } else {
        pb.contacts[name] = []string{phone}
    }
}

func (pb *PhoneBook) RemoveContact(name string) bool {
    name = strings.ToLower(name)
    if _, exists := pb.contacts[name]; exists {
        delete(pb.contacts, name)
        return true
    }
    return false
}

func (pb *PhoneBook) FindContact(name string) []string {
    name = strings.ToLower(name)
    return pb.contacts[name]
}

func (pb *PhoneBook) ListContacts() {
    if len(pb.contacts) == 0 {
        fmt.Println("Телефонная книга пуста")
        return
    }
    
    for name, phones := range pb.contacts {
        fmt.Printf("%s: %v\n", name, phones)
    }
}

func main() {
    phoneBook := NewPhoneBook()
    
    for {
        var command string
        fmt.Print("\nВведите команду (add/remove/find/list/quit): ")
        fmt.Scan(&command)
        
        switch strings.ToLower(command) {
        case "add":
            var name, phone string
            fmt.Print("Введите имя: ")
            fmt.Scan(&name)
            fmt.Print("Введите номер телефона: ")
            fmt.Scan(&phone)
            phoneBook.AddContact(name, phone)
            fmt.Println("Контакт добавлен")
            
        case "remove":
            var name string
            fmt.Print("Введите имя: ")
            fmt.Scan(&name)
            if phoneBook.RemoveContact(name) {
                fmt.Println("Контакт удален")
            } else {
                fmt.Println("Контакт не найден")
            }
            
        case "find":
            var name string
            fmt.Print("Введите имя: ")
            fmt.Scan(&name)
            if phones := phoneBook.FindContact(name); len(phones) > 0 {
                fmt.Printf("Номера для %s: %v\n", name, phones)
            } else {
                fmt.Println("Контакт не найден")
            }
            
        case "list":
            phoneBook.ListContacts()
            
        case "quit":
            fmt.Println("До свидания!")
            return
            
        default:
            fmt.Println("Неизвестная команда")
        }
    }
}
```

### Решение задания 3: Библиотечный каталог
```go
package main

import (
    "fmt"
    "strings"
)

type Book struct {
    ID       int
    Title    string
    Author   string
    Year     int
    Available bool
}

type Library struct {
    books    []Book
    lastID   int
}

func NewLibrary() *Library {
    return &Library{
        books:  make([]Book, 0),
        lastID: 0,
    }
}

func (l *Library) AddBook(title, author string, year int) {
    l.lastID++
    book := Book{
        ID:        l.lastID,
        Title:     title,
        Author:    author,
        Year:      year,
        Available: true,
    }
    l.books = append(l.books, book)
}

func (l *Library) FindBooks(searchTerm string) []Book {
    var results []Book
    searchTerm = strings.ToLower(searchTerm)
    
    for _, book := range l.books {
        if strings.Contains(strings.ToLower(book.Title), searchTerm) ||
           strings.Contains(strings.ToLower(book.Author), searchTerm) {
            results = append(results, book)
        }
    }
    return results
}

func (l *Library) BorrowBook(id int) bool {
    for i := range l.books {
        if l.books[i].ID == id && l.books[i].Available {
            l.books[i].Available = false
            return true
        }
    }
    return false
}

func (l *Library) ReturnBook(id int) bool {
    for i := range l.books {
        if l.books[i].ID == id && !l.books[i].Available {
            l.books[i].Available = true
            return true
        }
    }
    return false
}

func (l *Library) ListBooks() {
    if len(l.books) == 0 {
        fmt.Println("Библиотека пуста")
        return
    }
    
    for _, book := range l.books {
        status := "доступна"
        if !book.Available {
            status = "выдана"
        }
        fmt.Printf("%d: %s by %s (%d) - %s\n",
            book.ID, book.Title, book.Author, book.Year, status)
    }
}

func main() {
    library := NewLibrary()
    
    for {
        var command string
        fmt.Print("\nВведите команду (add/find/borrow/return/list/quit): ")
        fmt.Scan(&command)
        
        switch strings.ToLower(command) {
        case "add":
            var title, author string
            var year int
            fmt.Print("Введите название: ")
            fmt.Scan(&title)
            fmt.Print("Введите автора: ")
            fmt.Scan(&author)
            fmt.Print("Введите год: ")
            fmt.Scan(&year)
            library.AddBook(title, author, year)
            fmt.Println("Книга добавлена")
            
        case "find":
            var term string
            fmt.Print("Введите поисковый запрос: ")
            fmt.Scan(&term)
            books := library.FindBooks(term)
            if len(books) > 0 {
                for _, book := range books {
                    fmt.Printf("%d: %s by %s (%d)\n",
                        book.ID, book.Title, book.Author, book.Year)
                }
            } else {
                fmt.Println("Книги не найдены")
            }
            
        case "borrow":
            var id int
            fmt.Print("Введите ID книги: ")
            fmt.Scan(&id)
            if library.BorrowBook(id) {
                fmt.Println("Книга выдана")
            } else {
                fmt.Println("Книга недоступна или не найдена")
            }
            
        case "return":
            var id int
            fmt.Print("Введите ID книги: ")
            fmt.Scan(&id)
            if library.ReturnBook(id) {
                fmt.Println("Книга возвращена")
            } else {
                fmt.Println("Книга уже доступна или не найдена")
            }
            
        case "list":
            library.ListBooks()
            
        case "quit":
            fmt.Println("До свидания!")
            return
            
        default:
            fmt.Println("Неизвестная команда")
        }
    }
}
```

## Дополнительные материалы
- [Go Tour - Arrays](https://tour.golang.org/moretypes/6)
- [Go Tour - Slices](https://tour.golang.org/moretypes/7)
- [Go Tour - Maps](https://tour.golang.org/moretypes/19)
- [Go Tour - Structs](https://tour.golang.org/moretypes/2)
- [Effective Go - Arrays, Slices and Maps](https://golang.org/doc/effective_go#arrays)

## Следующий урок
В следующем уроке мы изучим указатели и методы в Go, а также разберем интерфейсы.