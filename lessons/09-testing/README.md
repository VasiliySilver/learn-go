# Урок 9: Тестирование в Go

## Содержание
1. [Модульное тестирование](#модульное-тестирование)
2. [Бенчмарки](#бенчмарки)
3. [Примеры в документации](#примеры-в-документации)
4. [Практические задания](#практические-задания)

## Модульное тестирование

### Основы тестирования
```go
// math/math_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}

// Таблица тестов
func TestMultiply(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive", 2, 3, 6},
        {"zero", 0, 5, 0},
        {"negative", -2, 3, -6},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Multiply(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Multiply(%d, %d) = %d; want %d",
                    tt.a, tt.b, result, tt.expected)
            }
        })
    }
}
```

### Вспомогательные функции
```go
func setupTestCase(t *testing.T) func() {
    t.Log("setup test case")
    return func() {
        t.Log("teardown test case")
    }
}

func TestWithSetup(t *testing.T) {
    teardown := setupTestCase(t)
    defer teardown()
    
    // тест
}
```

## Бенчмарки

### Написание бенчмарков
```go
func BenchmarkFibonacci(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Fibonacci(10)
    }
}

// Бенчмарк с разными размерами входных данных
func BenchmarkSort(b *testing.B) {
    sizes := []int{100, 1000, 10000}
    for _, size := range sizes {
        b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
            data := generateRandomSlice(size)
            b.ResetTimer()
            for i := 0; i < b.N; i++ {
                Sort(data)
            }
        })
    }
}
```

## Примеры в документации

### Написание примеров
```go
func ExampleHello() {
    fmt.Println(Hello("World"))
    // Output: Hello, World!
}

// Пример с несколькими выводами
func ExamplePrimes() {
    fmt.Println(Primes(10))
    fmt.Println(Primes(20))
    // Output:
    // [2 3 5 7]
    // [2 3 5 7 11 13 17 19]
}
```

## Практические задания

### Задание 1: Тестирование калькулятора
Создайте и протестируйте пакет калькулятора:
- Основные математические операции
- Обработка ошибок
- Таблицы тестов
- Бенчмарки для сравнения производительности

### Задание 2: Тестирование структур данных
Реализуйте и протестируйте структуру данных (например, стек или очередь):
- Тесты всех методов
- Проверка граничных случаев
- Примеры использования
- Бенчмарки операций

### Задание 3: Тестирование HTTP-сервера
Создайте тесты для HTTP-сервера:
- Тестирование обработчиков
- Моки для внешних зависимостей
- Тестирование ошибок
- Интеграционные тесты

## Решения

### Решение задания 1: Тестирование калькулятора
```go
// calculator/calculator.go
package calculator

import "errors"

var ErrDivideByZero = errors.New("деление на ноль")

type Calculator struct{}

func (c *Calculator) Add(a, b float64) float64 {
    return a + b
}

func (c *Calculator) Subtract(a, b float64) float64 {
    return a - b
}

func (c *Calculator) Multiply(a, b float64) float64 {
    return a * b
}

func (c *Calculator) Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, ErrDivideByZero
    }
    return a / b, nil
}

// calculator/calculator_test.go
package calculator

import (
    "testing"
)

func TestCalculator_Add(t *testing.T) {
    tests := []struct {
        name     string
        a, b     float64
        expected float64
    }{
        {"positive", 2, 3, 5},
        {"negative", -2, -3, -5},
        {"mixed", -2, 3, 1},
        {"zero", 0, 0, 0},
    }

    c := &Calculator{}
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := c.Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%f, %f) = %f; want %f",
                    tt.a, tt.b, result, tt.expected)
            }
        })
    }
}

func TestCalculator_Divide(t *testing.T) {
    tests := []struct {
        name        string
        a, b        float64
        expected    float64
        expectError bool
    }{
        {"normal", 6, 2, 3, false},
        {"zero", 1, 0, 0, true},
    }

    c := &Calculator{}
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := c.Divide(tt.a, tt.b)
            
            if tt.expectError {
                if err == nil {
                    t.Error("expected error, got nil")
                }
            } else {
                if err != nil {
                    t.Errorf("unexpected error: %v", err)
                }
                if result != tt.expected {
                    t.Errorf("Divide(%f, %f) = %f; want %f",
                        tt.a, tt.b, result, tt.expected)
                }
            }
        })
    }
}

func BenchmarkCalculator_Add(b *testing.B) {
    c := &Calculator{}
    for i := 0; i < b.N; i++ {
        c.Add(2, 3)
    }
}

func ExampleCalculator_Add() {
    c := &Calculator{}
    result := c.Add(2, 3)
    fmt.Printf("2 + 3 = %v\n", result)
    // Output: 2 + 3 = 5
}
```

### Решение задания 2: Тестирование структур данных
```go
// stack/stack.go
package stack

import "errors"

var ErrEmptyStack = errors.New("стек пуст")

type Stack struct {
    items []interface{}
}

func New() *Stack {
    return &Stack{
        items: make([]interface{}, 0),
    }
}

func (s *Stack) Push(item interface{}) {
    s.items = append(s.items, item)
}

func (s *Stack) Pop() (interface{}, error) {
    if len(s.items) == 0 {
        return nil, ErrEmptyStack
    }
    
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, nil
}

func (s *Stack) Peek() (interface{}, error) {
    if len(s.items) == 0 {
        return nil, ErrEmptyStack
    }
    return s.items[len(s.items)-1], nil
}

func (s *Stack) Size() int {
    return len(s.items)
}

// stack/stack_test.go
package stack

import (
    "testing"
)

func TestStack(t *testing.T) {
    s := New()
    
    // Test empty stack
    if s.Size() != 0 {
        t.Errorf("new stack size = %d; want 0", s.Size())
    }
    
    // Test Push
    s.Push(1)
    if s.Size() != 1 {
        t.Errorf("stack size after push = %d; want 1", s.Size())
    }
    
    // Test Peek
    value, err := s.Peek()
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }
    if value != 1 {
        t.Errorf("peek = %v; want 1", value)
    }
    
    // Test Pop
    value, err = s.Pop()
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }
    if value != 1 {
        t.Errorf("pop = %v; want 1", value)
    }
    if s.Size() != 0 {
        t.Errorf("size after pop = %d; want 0", s.Size())
    }
    
    // Test empty pop
    _, err = s.Pop()
    if err != ErrEmptyStack {
        t.Errorf("pop empty stack error = %v; want %v", err, ErrEmptyStack)
    }
}

func BenchmarkStack(b *testing.B) {
    s := New()
    b.Run("Push", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            s.Push(i)
        }
    })
    
    b.Run("Pop", func(b *testing.B) {
        for i := 0; i < b.N; i++ {
            s.Pop()
        }
    })
}

func ExampleStack() {
    s := New()
    s.Push(1)
    s.Push(2)
    value, _ := s.Pop()
    fmt.Println(value)
    // Output: 2
}
```

### Решение задания 3: Тестирование HTTP-сервера
```go
// server/server.go
package server

import (
    "encoding/json"
    "net/http"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

type UserHandler struct {
    users map[int]User
}

func NewUserHandler() *UserHandler {
    return &UserHandler{
        users: make(map[int]User),
    }
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "missing id parameter", http.StatusBadRequest)
        return
    }
    
    var userID int
    if _, err := fmt.Sscanf(id, "%d", &userID); err != nil {
        http.Error(w, "invalid id parameter", http.StatusBadRequest)
        return
    }
    
    user, ok := h.users[userID]
    if !ok {
        http.Error(w, "user not found", http.StatusNotFound)
        return
    }
    
    json.NewEncoder(w).Encode(user)
}

// server/server_test.go
package server

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestUserHandler_GetUser(t *testing.T) {
    handler := NewUserHandler()
    handler.users[1] = User{ID: 1, Name: "John"}
    
    tests := []struct {
        name       string
        userID     string
        wantStatus int
        wantUser   *User
    }{
        {
            name:       "existing user",
            userID:     "1",
            wantStatus: http.StatusOK,
            wantUser:   &User{ID: 1, Name: "John"},
        },
        {
            name:       "non-existing user",
            userID:     "2",
            wantStatus: http.StatusNotFound,
            wantUser:   nil,
        },
        {
            name:       "invalid id",
            userID:     "invalid",
            wantStatus: http.StatusBadRequest,
            wantUser:   nil,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            req := httptest.NewRequest("GET", "/user?id="+tt.userID, nil)
            w := httptest.NewRecorder()
            
            handler.GetUser(w, req)
            
            if w.Code != tt.wantStatus {
                t.Errorf("GetUser() status = %d; want %d", w.Code, tt.wantStatus)
            }
            
            if tt.wantUser != nil {
                var got User
                if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
                    t.Fatalf("Failed to decode response: %v", err)
                }
                if got != *tt.wantUser {
                    t.Errorf("GetUser() = %v; want %v", got, tt.wantUser)
                }
            }
        })
    }
}
```

## Дополнительные материалы
- [Testing package documentation](https://golang.org/pkg/testing/)
- [Go Blog: Using Subtests and Sub-benchmarks](https://blog.golang.org/subtests)
- [Go Blog: Examples in documentation](https://blog.golang.org/examples)
- [Go Blog: HTTP testing](https://blog.golang.org/http-testing)

## Следующий урок
В следующем уроке мы изучим работу с базами данных в Go, включая SQL и NoSQL решения.