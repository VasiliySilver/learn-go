# Урок 8: Обработка ошибок в Go

## Содержание
1. [Основы обработки ошибок](#основы-обработки-ошибок)
2. [Создание пользовательских ошибок](#создание-пользовательских-ошибок)
3. [Паники и восстановление](#паники-и-восстановление)
4. [Практические задания](#практические-задания)

## Основы обработки ошибок

### Интерфейс error
```go

// Встроенный интерфейс error
type error interface {
    Error() string
}

// Проверка ошибок
if err != nil {
    // обработка ошибки
    return err
}

// Множественные возвращаемые значения
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("деление на ноль")
    }
    return a / b, nil
}
```

### Работа с ошибками
```go

// Создание простой ошибки
err := errors.New("что-то пошло не так")

// Форматированная ошибка
err := fmt.Errorf("ошибка при обработке %s: %v", filename, err)

// Цепочка ошибок (Go 1.13+)
err = fmt.Errorf("дополнительный контекст: %w", err)

// Извлечение оригинальной ошибки
originalErr := errors.Unwrap(err)

// Проверка типа ошибки
if errors.Is(err, os.ErrNotExist) {
    // обработка конкретной ошибки
}
```

## Создание пользовательских ошибок

### Структура ошибки
```go

type ValidationError struct {
    Field string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("ошибка валидации поля %s: %s", e.Field, e.Message)
}

// Использование
func validateAge(age int) error {
    if age < 0 {
        return &ValidationError{
            Field: "age",
            Message: "возраст не может быть отрицательным",
        }
    }
    return nil
}
```

### Константные ошибки
```go

var (
    ErrNotFound = errors.New("элемент не найден")
    ErrInvalidInput = errors.New("некорректные входные данные")
)

func findItem(id string) error {
    return ErrNotFound
}
```

## Паники и восстановление

### Паника
```go

func doSomething() {
    panic("критическая ошибка")
}

// Восстановление после паники
func handlePanic() {
    if r := recover(); r != nil {
        fmt.Printf("Восстановление после паники: %v\n", r)
    }
}

func main() {
    defer handlePanic()
    doSomething()
}
```

## Практические задания

### Задание 1: Валидатор данных
Создайте систему валидации данных пользователя:
- Проверка различных полей (email, возраст, имя)
- Пользовательские типы ошибок
- Агрегация нескольких ошибок

### Задание 2: Безопасный парсер файлов
Реализуйте парсер конфигурационных файлов:
- Обработка различных форматов (JSON, YAML)
- Восстановление после ошибок
- Подробные сообщения об ошибках

### Задание 3: HTTP-сервер с обработкой ошибок
Создайте простой HTTP-сервер с правильной обработкой ошибок:
- Middleware для логирования ошибок
- Пользовательские HTTP-ошибки
- Graceful shutdown

## Решения

### Решение задания 1: Валидатор данных
```go

package main

import (
    "fmt"
    "regexp"
    "strings"
)

type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

type ValidationErrors []ValidationError

func (ve ValidationErrors) Error() string {
    var errors []string
    for _, err := range ve {
        errors = append(errors, err.Error())
    }
    return strings.Join(errors, "\n")
}

type User struct {
    Email string
    Age   int
    Name  string
}

func validateEmail(email string) error {
    emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
    if !emailRegex.MatchString(email) {
        return &ValidationError{
            Field:   "email",
            Message: "некорректный формат email",
        }
    }
    return nil
}

func validateAge(age int) error {
    if age < 0 || age > 150 {
        return &ValidationError{
            Field:   "age",
            Message: "возраст должен быть между 0 и 150",
        }
    }
    return nil
}

func validateName(name string) error {
    if len(name) < 2 {
        return &ValidationError{
            Field:   "name",
            Message: "имя должно содержать минимум 2 символа",
        }
    }
    return nil
}

func validateUser(user User) error {
    var errors ValidationErrors

    if err := validateEmail(user.Email); err != nil {
        errors = append(errors, *err.(*ValidationError))
    }
    if err := validateAge(user.Age); err != nil {
        errors = append(errors, *err.(*ValidationError))
    }
    if err := validateName(user.Name); err != nil {
        errors = append(errors, *err.(*ValidationError))
    }

    if len(errors) > 0 {
        return errors
    }
    return nil
}

func main() {
    users := []User{
        {Email: "invalid", Age: -1, Name: "A"},
        {Email: "valid@example.com", Age: 25, Name: "John"},
    }

    for _, user := range users {
        if err := validateUser(user); err != nil {
            fmt.Printf("Ошибка валидации для пользователя:\n%v\n\n", err)
        } else {
            fmt.Printf("Пользователь %s прошел валидацию\n\n", user.Name)
        }
    }
}
```

### Решение задания 2: Безопасный парсер файлов
```go

package main

import (
    "encoding/json"
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "path/filepath"
)

type ParseError struct {
    FileName string
    Format   string
    Err      error
}

func (e *ParseError) Error() string {
    return fmt.Sprintf("ошибка парсинга файла %s (%s): %v", e.FileName, e.Format, e.Err)
}

func (e *ParseError) Unwrap() error {
    return e.Err
}

type Config struct {
    Database struct {
        Host     string `json:"host" yaml:"host"`
        Port     int    `json:"port" yaml:"port"`
        Username string `json:"username" yaml:"username"`
        Password string `json:"password" yaml:"password"`
    } `json:"database" yaml:"database"`
}

func parseJSON(filename string) (*Config, error) {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, &ParseError{filename, "JSON", err}
    }

    var config Config
    if err := json.Unmarshal(data, &config); err != nil {
        return nil, &ParseError{filename, "JSON", err}
    }

    return &config, nil
}

func parseYAML(filename string) (*Config, error) {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, &ParseError{filename, "YAML", err}
    }

    var config Config
    if err := yaml.Unmarshal(data, &config); err != nil {
        return nil, &ParseError{filename, "YAML", err}
    }

    return &config, nil
}

func parseConfig(filename string) (*Config, error) {
    ext := filepath.Ext(filename)
    
    switch ext {
    case ".json":
        return parseJSON(filename)
    case ".yaml", ".yml":
        return parseYAML(filename)
    default:
        return nil, fmt.Errorf("неподдерживаемый формат файла: %s", ext)
    }
}

func main() {
    files := []string{
        "config.json",
        "config.yaml",
        "config.txt",
    }

    for _, file := range files {
        config, err := parseConfig(file)
        if err != nil {
            fmt.Printf("Ошибка: %v\n", err)
            continue
        }
        fmt.Printf("Успешно загружена конфигурация из %s\n", file)
        fmt.Printf("Database Host: %s\n", config.Database.Host)
    }
}
```

### Решение задания 3: HTTP-сервер с обработкой ошибок
```go

package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

type HTTPError struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

func (e *HTTPError) Error() string {
    return e.Message
}

func errorHandler(handler func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        err := handler(w, r)
        if err != nil {
            var httpError *HTTPError
            if e, ok := err.(*HTTPError); ok {
                httpError = e
            } else {
                httpError = &HTTPError{
                    Code:    http.StatusInternalServerError,
                    Message: "внутренняя ошибка сервера",
                }
                log.Printf("Неожиданная ошибка: %v", err)
            }

            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(httpError.Code)
            json.NewEncoder(w).Encode(httpError)
        }
    }
}

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
    })
}

func handleUser(w http.ResponseWriter, r *http.Request) error {
    if r.Method != http.MethodGet {
        return &HTTPError{
            Code:    http.StatusMethodNotAllowed,
            Message: "метод не поддерживается",
        }
    }

    userID := r.URL.Query().Get("id")
    if userID == "" {
        return &HTTPError{
            Code:    http.StatusBadRequest,
            Message: "отсутствует параметр id",
        }
    }

    // Имитация ошибки
    if userID == "error" {
        return fmt.Errorf("неожиданная ошибка при получении пользователя")
    }

    user := map[string]string{
        "id":   userID,
        "name": "John Doe",
    }

    return json.NewEncoder(w).Encode(user)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/user", errorHandler(handleUser))

    server := &http.Server{
        Addr:    ":8080",
        Handler: loggingMiddleware(mux),
    }

    // Graceful shutdown
    done := make(chan os.Signal, 1)
    signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Ошибка запуска сервера: %v", err)
        }
    }()

    log.Print("Сервер запущен")

    <-done
    log.Print("Сервер останавливается...")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := server.Shutdown(ctx); err != nil {
        log.Fatalf("Ошибка при остановке сервера: %v", err)
    }

    log.Print("Сервер остановлен")
}
```

## Дополнительные материалы
- [Error handling and Go](https://blog.golang.org/error-handling-and-go)
- [Working with Errors in Go 1.13](https://blog.golang.org/go1.13-errors)
- [Effective Error Handling in Go](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)

## Следующий урок
В следующем уроке мы изучим тестирование в Go, включая модульные тесты, бенчмарки и примеры.