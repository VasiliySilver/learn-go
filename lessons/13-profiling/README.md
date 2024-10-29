# Урок 13: Профайлинг

## Содержание
1. [Профилирование](#профилирование)
2. [Оптимизация производительности](#оптимизация-производительности)
3. [Паттерны проектирования](#паттерны-проектирования)
4. [Практические задания](#практические-задания)

## Профилирование

### CPU профилирование
```go



// profiling/cpu.go
package main

import (
    "os"
    "runtime/pprof"
)

func startCPUProfile(filename string) func() {
    f, err := os.Create(filename)
    if err != nil {
        log.Fatal(err)
    }
    
    if err := pprof.StartCPUProfile(f); err != nil {
        log.Fatal(err)
    }
    
    return func() {
        pprof.StopCPUProfile()
        f.Close()
    }
}

func main() {
    // Запуск профилирования
    stop := startCPUProfile("cpu.prof")
    defer stop()
    
    // Ваш код здесь...
}
```

### Memory профилирование
```go



// profiling/memory.go
package main

import (
    "os"
    "runtime"
    "runtime/pprof"
)

func writeHeapProfile(filename string) {
    f, err := os.Create(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    
    runtime.GC() // Запуск GC перед записью профиля
    if err := pprof.WriteHeapProfile(f); err != nil {
        log.Fatal(err)
    }
}

// Пример использования с HTTP
func setupProfiler() {
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
}
```

## Оптимизация производительности

### Пул горутин
```go



// pool/worker.go
package pool

type Task func() error

type Pool struct {
    tasks    chan Task
    workers  int
    results  chan error
    done     chan struct{}
}

func NewPool(workers int) *Pool {
    return &Pool{
        tasks:   make(chan Task, workers),
        workers: workers,
        results: make(chan error, workers),
        done:    make(chan struct{}),
    }
}

func (p *Pool) Start() {
    for i := 0; i < p.workers; i++ {
        go p.worker()
    }
}

func (p *Pool) worker() {
    for task := range p.tasks {
        p.results <- task()
    }
}

func (p *Pool) Submit(task Task) {
    p.tasks <- task
}

func (p *Pool) Stop() {
    close(p.tasks)
    close(p.done)
}

// Пример использования
func Example() {
    pool := NewPool(5)
    pool.Start()
    defer pool.Stop()
    
    // Отправка задач
    for i := 0; i < 10; i++ {
        pool.Submit(func() error {
            // Выполнение задачи
            return nil
        })
    }
    
    // Обработка результатов
    for i := 0; i < 10; i++ {
        if err := <-pool.results; err != nil {
            log.Printf("Task error: %v", err)
        }
    }
}
```

### Оптимизация памяти
```go



// optimization/memory.go
package optimization

import (
    "sync"
)

// Пул объектов для переиспользования
var bufferPool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 0, 1024)
    },
}

func processData(data []byte) {
    // Получаем буфер из пула
    buf := bufferPool.Get().([]byte)
    buf = buf[:0] // Сброс длины, сохранение емкости
    
    // Используем буфер
    buf = append(buf, data...)
    
    // Возвращаем буфер в пул
    bufferPool.Put(buf)
}

// Оптимизация строк
func optimizeStrings(data []string) string {
    // Предварительное выделение памяти
    var builder strings.Builder
    builder.Grow(len(data) * 64) // Примерный размер
    
    for _, s := range data {
        builder.WriteString(s)
    }
    
    return builder.String()
}
```

## Паттерны проектирования

### Factory Method
```go



// patterns/factory.go
package patterns

type Storage interface {
    Save(data []byte) error
    Load(id string) ([]byte, error)
}

type FileStorage struct {
    path string
}

type S3Storage struct {
    bucket string
    region string
}

type StorageType string

const (
    FileStorageType StorageType = "file"
    S3StorageType   StorageType = "s3"
)

func NewStorage(storageType StorageType, config map[string]string) (Storage, error) {
    switch storageType {
    case FileStorageType:
        return &FileStorage{
            path: config["path"],
        }, nil
    case S3StorageType:
        return &S3Storage{
            bucket: config["bucket"],
            region: config["region"],
        }, nil
    default:
        return nil, fmt.Errorf("unknown storage type: %s", storageType)
    }
}
```

### Observer
```go



// patterns/observer.go
package patterns

type Event struct {
    Type    string
    Payload interface{}
}

type Observer interface {
    OnEvent(Event)
}

type Subject struct {
    observers map[Observer]bool
    mu        sync.RWMutex
}

func NewSubject() *Subject {
    return &Subject{
        observers: make(map[Observer]bool),
    }
}

func (s *Subject) Register(observer Observer) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.observers[observer] = true
}

func (s *Subject) Unregister(observer Observer) {
    s.mu.Lock()
    defer s.mu.Unlock()
    delete(s.observers, observer)
}

func (s *Subject) Notify(event Event) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    for observer := range s.observers {
        go observer.OnEvent(event)
    }
}
```

## Практические задания

### Задание 1: Оптимизация производительности
Оптимизируйте производительность сервиса обработки изображений:
- Профилирование узких мест
- Пул горутин для параллельной обработки
- Оптимизация памяти
- Кеширование результатов

### Задание 2: Реализация паттернов
Реализуйте систему обработки заказов с использованием паттернов:
- Factory Method для создания различных типов заказов
- Observer для уведомлений о статусе
- Chain of Responsibility для обработки
- Strategy для разных способов оплаты

### Задание 3: Профилирование и оптимизация API
Создайте инструменты для профилирования REST API:
- CPU и Memory профили
- Трейсинг запросов
- Метрики производительности
- Автоматическая генерация отчетов

## Решения

### Решение задания 1: Оптимизация производительности
```go



// image/processor.go
package image

import (
    "sync"
    "github.com/disintegration/imaging"
)

type Processor struct {
    workerPool *Pool
    cache      *Cache
}

func NewProcessor(workers int) *Processor {
    return &Processor{
        workerPool: NewPool(workers),
        cache:      NewCache(),
    }
}

type ProcessTask struct {
    ImageID   string
    Operation string
    Params    map[string]interface{}
}

func (p *Processor) Process(task ProcessTask) (string, error) {
    // Проверяем кеш
    if result := p.cache.Get(task.getCacheKey()); result != "" {
        return result, nil
    }
    
    // Создаем задачу для пула
    resultChan := make(chan string, 1)
    errChan := make(chan error, 1)
    
    p.workerPool.Submit(func() error {
        // Загрузка изображения
        img, err := imaging.Open(task.ImageID)
        if err != nil {
            errChan <- err
            return err
        }
        
        // Обработка изображения
        var processed *image.NRGBA
        switch task.Operation {
        case "resize":
            width := task.Params["width"].(int)
            height := task.Params["height"].(int)
            processed = imaging.Resize(img, width, height, imaging.Lanczos)
        case "crop":
            width := task.Params["width"].(int)
            height := task.Params["height"].(int)
            processed = imaging.CropCenter(img, width, height)
        // Другие операции...
        }
        
        // Сохранение результата
        outputPath := fmt.Sprintf("processed/%s_%s.jpg", 
            task.ImageID, task.Operation)
        err = imaging.Save(processed, outputPath)
        if err != nil {
            errChan <- err
            return err
        }
        
        // Сохраняем в кеш
        p.cache.Set(task.getCacheKey(), outputPath)
        
        resultChan <- outputPath
        return nil
    })
    
    // Ожидаем результат
    select {
    case result := <-resultChan:
        return result, nil
    case err := <-errChan:
        return "", err
    }
}

// Cache implementation
type Cache struct {
    data map[string]string
    mu   sync.RWMutex
}

func NewCache() *Cache {
    return &Cache{
        data: make(map[string]string),
    }
}

func (c *Cache) Get(key string) string {
    c.mu.RLock()
    defer c.mu.RUnlock()
    return c.data[key]
}

func (c *Cache) Set(key, value string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.data[key] = value
}
```

### Решение задания 2: Реализация паттернов
```go



// order/order.go
package order

type Order interface {
    Process() error
    GetTotal() float64
    GetStatus() string
}

type BaseOrder struct {
    ID     string
    Items  []Item
    Status string
    Total  float64
}

type RegularOrder struct {
    BaseOrder
}

type ExpressOrder struct {
    BaseOrder
    Priority int
}

// Factory Method
func NewOrder(orderType string, items []Item) (Order, error) {
    base := BaseOrder{
        ID:     generateID(),
        Items:  items,
        Status: "new",
    }
    
    switch orderType {
    case "regular":
        return &RegularOrder{BaseOrder: base}, nil
    case "express":
        return &ExpressOrder{
            BaseOrder: base,
            Priority:  1,
        }, nil
    default:
        return nil, fmt.Errorf("unknown order type: %s", orderType)
    }
}

// Observer
type OrderObserver interface {
    OnStatusChange(order Order)
}

type EmailNotifier struct {
    emailService EmailService
}

func (n *EmailNotifier) OnStatusChange(order Order) {
    n.emailService.SendNotification(
        "Order status changed",
        fmt.Sprintf("Order %s is now %s", order.ID, order.GetStatus()),
    )
}

// Chain of Responsibility
type OrderHandler interface {
    Handle(order Order) error
    SetNext(handler OrderHandler)
}

type ValidationHandler struct {
    next OrderHandler
}

func (h *ValidationHandler) Handle(order Order) error {
    // Validate order
    if err := validateOrder(order); err != nil {
        return err
    }
    
    if h.next != nil {
        return h.next.Handle(order)
    }
    return nil
}

// Strategy
type PaymentStrategy interface {
    Pay(amount float64) error
}

type CreditCardPayment struct {
    cardNumber string
}

func (p *CreditCardPayment) Pay(amount float64) error {
    // Process credit card payment
    return nil
}

type PayPalPayment struct {
    email string
}

func (p *PayPalPayment) Pay(amount float64) error {
    // Process PayPal payment
    return nil
}

// Order Service
type OrderService struct {
    observers  []OrderObserver
    handler   OrderHandler
    payment   PaymentStrategy
}

func (s *OrderService) ProcessOrder(order Order) error {
    // Validate and process using Chain of Responsibility
    if err := s.handler.Handle(order); err != nil {
        return err
    }
    
    // Process payment using Strategy
    if err := s.payment.Pay(order.GetTotal()); err != nil {
        return err
    }
    
    // Update status and notify observers
    order.Status = "completed"
    for _, observer := range s.observers {
        observer.OnStatusChange(order)
    }
    
    return nil
}
```

### Решение задания 3: Профилирование и оптимизация API
```go



// profiling/api.go
package profiling

import (
    "net/http"
    "runtime/pprof"
    "github.com/opentracing/opentracing-go"
)

type ProfilingMiddleware struct {
    next    http.Handler
    metrics *Metrics
    tracer  opentracing.Tracer
}

func NewProfilingMiddleware(next http.Handler) *ProfilingMiddleware {
    return &ProfilingMiddleware{
        next:    next,
        metrics: NewMetrics(),
        tracer:  opentracing.GlobalTracer(),
    }
}

func (m *ProfilingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // Start tracing
    span, ctx := opentracing.StartSpanFromContext(r.Context(), "http_request")
    defer span.Finish()
    
    // Add request info to span
    span.SetTag("http.method", r.Method)
    span.SetTag("http.url", r.URL.Path)
    
    // Start CPU profile if requested
    if r.Header.Get("X-Profile-CPU") == "true" {
        f, err := os.Create(fmt.Sprintf("cpu_%s.prof", time.Now().Format("20060102150405")))
        if err == nil {
            pprof.StartCPUProfile(f)
            defer func() {
                pprof.StopCPUProfile()
                f.Close()
            }()
        }
    }
    
    // Wrap response writer to capture status code
    wrapped := NewResponseWriter(w)
    
    // Process request
    start := time.Now()
    m.next.ServeHTTP(wrapped, r.WithContext(ctx))
    duration := time.Since(start)
    
    // Record metrics
    m.metrics.RecordRequest(r.Method, r.URL.Path, wrapped.Status(), duration)
    
    // Add response info to span
    span.SetTag("http.status_code", wrapped.Status())
    span.SetTag("http.duration", duration)
    
    // Generate memory profile if requested
    if r.Header.Get("X-Profile-Memory") == "true" {
        f, err := os.Create(fmt.Sprintf("memory_%s.prof", 
            time.Now().Format("20060102150405")))
        if err == nil {
            pprof.WriteHeapProfile(f)
            f.Close()
        }
    }
}

// Metrics collection
type Metrics struct {
    requestDuration *prometheus.HistogramVec
    requestTotal    *prometheus.CounterVec
}

func NewMetrics() *Metrics {
    return &Metrics{
        requestDuration: prometheus.NewHistogramVec(
            prometheus.HistogramOpts{
                Name:    "http_request_duration_seconds",
                Help:    "HTTP request duration in seconds",
                Buckets: prometheus.DefBuckets,
            },
            []string{"method", "path", "status"},
        ),
        requestTotal: prometheus.NewCounterVec(
            prometheus.CounterOpts{
                Name: "http_requests_total",
                Help: "Total number of HTTP requests",
            },
            []string{"method", "path", "status"},
        ),
    }
}

func (m *Metrics) RecordRequest(method, path string, status int, duration time.Duration) {
    labels := prometheus.Labels{
        "method": method,
        "path":   path,
        "status": fmt.Sprintf("%d", status),
    }
    m.requestDuration.With(labels).Observe(duration.Seconds())
    m.requestTotal.With(labels).Inc()
}
```

## Дополнительные материалы
- [Profiling Go Programs](https://blog.golang.org/pprof)
- [Go Design Patterns](https://github.com/tmrts/go-patterns)
- [High Performance Go Workshop](https://dave.cheney.net/high-performance-go-workshop/dotgo-paris.html)
- [Go Concurrency Patterns](https://talks.golang.org/2012/concurrency.slide)

## Следующий урок
В следующем уроке мы рассмотрим безопасность в Go-приложениях, включая криптографию, защиту от уязвимостей и аудит безопасности.