# Урок 6: Указатели, методы и интерфейсы в Go

## Содержание
1. [Указатели](#указатели)
2. [Методы](#методы)
3. [Интерфейсы](#интерфейсы)
4. [Практические задания](#практические-задания)

## Указатели

### Основы работы с указателями
```

// Объявление указателя
var p *int
number := 42
p = &number    // получение адреса переменной

// Разыменование указателя
value := *p    // получение значения по указателю
*p = 100       // изменение значения по указателю

// Указатель на структуру
type Person struct {
    Name string
    Age  int
}

person := &Person{Name: "John", Age: 25}
person.Name = "Bob"    // автоматическое разыменование
(*person).Age = 26     // явное разыменование
```

### Передача по указателю vs по значению
```

// Передача по значению (копия)
func increment(x int) {
    x++    // изменяется только локальная копия
}

// Передача по указателю
func incrementPtr(x *int) {
    *x++    // изменяется оригинальное значение
}
```

## Методы

### Объявление методов
```

type Rectangle struct {
    Width  float64
    Height float64
}

// Метод с receiver-значением
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Метод с receiver-указателем
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}
```

### Выбор типа receiver
```

// Value receiver - когда не нужно изменять состояние
func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Pointer receiver - когда нужно изменять состояние
func (r *Rectangle) SetWidth(width float64) {
    r.Width = width
}
```

## Интерфейсы

### Объявление и реализация
```

// Объявление интерфейса
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Неявная реализация интерфейса
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}
```

### Пустой интерфейс и утверждение типа
```

// Пустой интерфейс
var i interface{}
i = 42
i = "hello"

// Утверждение типа
str, ok := i.(string)
if ok {
    fmt.Printf("Это строка: %s\n", str)
}

// Switch по типу
switch v := i.(type) {
case int:
    fmt.Printf("Целое число: %d\n", v)
case string:
    fmt.Printf("Строка: %s\n", v)
default:
    fmt.Printf("Неизвестный тип\n")
}
```

## Практические задания

### Задание 1: Геометрические фигуры
Создайте систему для работы с геометрическими фигурами:
- Определите интерфейс Shape
- Реализуйте несколько типов фигур
- Создайте функции для работы с группами фигур

```

package main

import (
    "fmt"
    "math"
)

type Shape interface {
    Area() float64
    Perimeter() float64
}

// Ваша реализация различных фигур
```

### Задание 2: Связанный список
Реализуйте структуру данных "связанный список" с использованием указателей:
- Добавление и удаление элементов
- Поиск элементов
- Методы для работы со списком

### Задание 3: Стек и очередь
Реализуйте структуры данных стек и очередь:
- Используйте интерфейсы для определения поведения
- Реализуйте методы push, pop, peek
- Обработайте краевые случаи

## Решения

### Решение задания 1: Геометрические фигуры
```

package main

import (
    "fmt"
    "math"
)

type Shape interface {
    Area() float64
    Perimeter() float64
    String() string
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

func (c Circle) String() string {
    return fmt.Sprintf("Круг с радиусом %.2f", c.Radius)
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

func (r Rectangle) String() string {
    return fmt.Sprintf("Прямоугольник %.2fx%.2f", r.Width, r.Height)
}

type Triangle struct {
    A, B, C float64 // стороны треугольника
}

func (t Triangle) Area() float64 {
    // Формула Герона
    s := (t.A + t.B + t.C) / 2
    return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t Triangle) Perimeter() float64 {
    return t.A + t.B + t.C
}

func (t Triangle) String() string {
    return fmt.Sprintf("Треугольник со сторонами %.2f, %.2f, %.2f", t.A, t.B, t.C)
}

// Функция для работы с группой фигур
func PrintShapeInfo(shapes []Shape) {
    for _, shape := range shapes {
        fmt.Printf("%s:\n", shape.String())
        fmt.Printf("  Площадь: %.2f\n", shape.Area())
        fmt.Printf("  Периметр: %.2f\n", shape.Perimeter())
    }
}

func main() {
    shapes := []Shape{
        Circle{Radius: 5},
        Rectangle{Width: 4, Height: 6},
        Triangle{A: 3, B: 4, C: 5},
    }

    PrintShapeInfo(shapes)
}
```

### Решение задания 2: Связанный список
```

package main

import (
    "fmt"
    "errors"
)

type Node struct {
    Value int
    Next  *Node
}

type LinkedList struct {
    Head *Node
    Size int
}

func NewLinkedList() *LinkedList {
    return &LinkedList{nil, 0}
}

func (l *LinkedList) AddFront(value int) {
    node := &Node{Value: value, Next: l.Head}
    l.Head = node
    l.Size++
}

func (l *LinkedList) AddBack(value int) {
    node := &Node{Value: value}
    if l.Head == nil {
        l.Head = node
    } else {
        current := l.Head
        for current.Next != nil {
            current = current.Next
        }
        current.Next = node
    }
    l.Size++
}

func (l *LinkedList) RemoveFront() error {
    if l.Head == nil {
        return errors.New("список пуст")
    }
    l.Head = l.Head.Next
    l.Size--
    return nil
}

func (l *LinkedList) Remove(value int) bool {
    if l.Head == nil {
        return false
    }

    if l.Head.Value == value {
        l.Head = l.Head.Next
        l.Size--
        return true
    }

    current := l.Head
    for current.Next != nil {
        if current.Next.Value == value {
            current.Next = current.Next.Next
            l.Size--
            return true
        }
        current = current.Next
    }
    return false
}

func (l *LinkedList) Find(value int) bool {
    current := l.Head
    for current != nil {
        if current.Value == value {
            return true
        }
        current = current.Next
    }
    return false
}

func (l *LinkedList) Print() {
    current := l.Head
    for current != nil {
        fmt.Printf("%d -> ", current.Value)
        current = current.Next
    }
    fmt.Println("nil")
}

func main() {
    list := NewLinkedList()
    
    // Добавление элементов
    list.AddBack(1)
    list.AddBack(2)
    list.AddBack(3)
    list.AddFront(0)
    
    fmt.Println("Исходный список:")
    list.Print()
    
    // Удаление элемента
    list.Remove(2)
    fmt.Println("После удаления 2:")
    list.Print()
    
    // Поиск элемента
    fmt.Printf("Найден элемент 3: %v\n", list.Find(3))
    fmt.Printf("Найден элемент 2: %v\n", list.Find(2))
}
```

### Решение задания 3: Стек и очередь
```

package main

import (
    "fmt"
    "errors"
)

// Интерфейсы
type Stack interface {
    Push(value interface{})
    Pop() (interface{}, error)
    Peek() (interface{}, error)
    IsEmpty() bool
    Size() int
}

type Queue interface {
    Enqueue(value interface{})
    Dequeue() (interface{}, error)
    Peek() (interface{}, error)
    IsEmpty() bool
    Size() int
}

// Реализация стека
type ArrayStack struct {
    items []interface{}
}

func NewStack() *ArrayStack {
    return &ArrayStack{items: make([]interface{}, 0)}
}

func (s *ArrayStack) Push(value interface{}) {
    s.items = append(s.items, value)
}

func (s *ArrayStack) Pop() (interface{}, error) {
    if s.IsEmpty() {
        return nil, errors.New("стек пуст")
    }
    value := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return value, nil
}

func (s *ArrayStack) Peek() (interface{}, error) {
    if s.IsEmpty() {
        return nil, errors.New("стек пуст")
    }
    return s.items[len(s.items)-1], nil
}

func (s *ArrayStack) IsEmpty() bool {
    return len(s.items) == 0
}

func (s *ArrayStack) Size() int {
    return len(s.items)
}

// Реализация очереди
type ArrayQueue struct {
    items []interface{}
}

func NewQueue() *ArrayQueue {
    return &ArrayQueue{items: make([]interface{}, 0)}
}

func (q *ArrayQueue) Enqueue(value interface{}) {
    q.items = append(q.items, value)
}

func (q *ArrayQueue) Dequeue() (interface{}, error) {
    if q.IsEmpty() {
        return nil, errors.New("очередь пуста")
    }
    value := q.items[0]
    q.items = q.items[1:]
    return value, nil
}

func (q *ArrayQueue) Peek() (interface{}, error) {
    if q.IsEmpty() {
        return nil, errors.New("очередь пуста")
    }
    return q.items[0], nil
}

func (q *ArrayQueue) IsEmpty() bool {
    return len(q.items) == 0
}

func (q *ArrayQueue) Size() int {
    return len(q.items)
}

func main() {
    // Тестирование стека
    fmt.Println("Тестирование стека:")
    stack := NewStack()
    
    stack.Push(1)
    stack.Push(2)
    stack.Push(3)
    
    fmt.Printf("Размер стека: %d\n", stack.Size())
    
    if value, err := stack.Peek(); err == nil {
        fmt.Printf("Верхний элемент: %v\n", value)
    }
    
    for !stack.IsEmpty() {
        if value, err := stack.Pop(); err == nil {
            fmt.Printf("Извлечено: %v\n", value)
        }
    }
    
    // Тестирование очереди
    fmt.Println("\nТестирование очереди:")
    queue := NewQueue()
    
    queue.Enqueue("A")
    queue.Enqueue("B")
    queue.Enqueue("C")
    
    fmt.Printf("Размер очереди: %d\n", queue.Size())
    
    if value, err := queue.Peek(); err == nil {
        fmt.Printf("Первый элемент: %v\n", value)
    }
    
    for !queue.IsEmpty() {
        if value, err := queue.Dequeue(); err == nil {
            fmt.Printf("Извлечено: %v\n", value)
        }
    }
}
```

## Дополнительные материалы
- [Go Tour - Pointers](https://tour.golang.org/moretypes/1)
- [Go Tour - Methods](https://tour.golang.org/methods/1)
- [Go Tour - Interfaces](https://tour.golang.org/methods/9)
- [Effective Go - Interfaces and Methods](https://golang.org/doc/effective_go#interfaces_and_methods)

## Следующий урок
В следующем уроке мы изучим горутины и каналы в Go, что позволит нам писать конкурентные программы.