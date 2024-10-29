# Урок 11: Создание REST API в Go

## Содержание
1. [Основы REST API](#основы-rest-api)
2. [Маршрутизация](#маршрутизация)
3. [Middleware](#middleware)
4. [Документация API](#документация-api)
5. [Практические задания](#практические-задания)

## Основы REST API

### Базовая структура API
```go


// main.go
package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

type Response struct {
    Status  string      `json:"status"`
    Message string      `json:"message,omitempty"`
    Data    interface{} `json:"data,omitempty"`
}

func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(data)
}

func main() {
    r := mux.NewRouter()
    
    // Маршруты
    r.HandleFunc("/api/health", HealthCheck).Methods("GET")
    
    log.Fatal(http.ListenAndServe(":8080", r))
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
    JSONResponse(w, http.StatusOK, Response{
        Status: "success",
        Message: "API is running",
    })
}
```

### Обработка запросов
```go


type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        JSONResponse(w, http.StatusBadRequest, Response{
            Status:  "error",
            Message: "Invalid request payload",
        })
        return
    }
    
    // Логика создания пользователя...
    
    JSONResponse(w, http.StatusCreated, Response{
        Status: "success",
        Data:   user,
    })
}
```

## Маршрутизация

### Настройка маршрутов
```go


func setupRoutes(r *mux.Router) {
    // API версионирование
    api := r.PathPrefix("/api/v1").Subrouter()
    
    // Users routes
    users := api.PathPrefix("/users").Subrouter()
    users.HandleFunc("", GetUsers).Methods("GET")
    users.HandleFunc("", CreateUser).Methods("POST")
    users.HandleFunc("/{id:[0-9]+}", GetUser).Methods("GET")
    users.HandleFunc("/{id:[0-9]+}", UpdateUser).Methods("PUT")
    users.HandleFunc("/{id:[0-9]+}", DeleteUser).Methods("DELETE")
    
    // Другие группы маршрутов...
}
```

## Middleware

### Реализация middleware
```go


func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // Вызов следующего обработчика
        next.ServeHTTP(w, r)
        
        // Логирование после обработки
        log.Printf(
            "%s %s %s",
            r.Method,
            r.RequestURI,
            time.Since(start),
        )
    })
}

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            JSONResponse(w, http.StatusUnauthorized, Response{
                Status:  "error",
                Message: "No authorization token provided",
            })
            return
        }
        
        // Проверка токена...
        
        next.ServeHTTP(w, r)
    })
}
```

## Документация API

### Swagger документация
```go


// @title Example API
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @tag.name users
// @tag.description User management endpoints

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags users
// @Accept json
// @Produce json
// @Param user body User true "User object"
// @Success 201 {object} Response{data=User}
// @Failure 400 {object} Response
// @Router /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
    // Implementation...
}
```

## Практические задания

### Задание 1: API для блога
Создайте REST API для блога:
- CRUD операции для постов и комментариев
- Аутентификация и авторизация
- Валидация входных данных
- Swagger документация

### Задание 2: API для файлового хранилища
Реализуйте API для управления файлами:
- Загрузка и скачивание файлов
- Управление метаданными
- Контроль доступа
- Обработка больших файлов

### Задание 3: API для чата
Создайте API для чат-приложения:
- Управление пользователями и чатами
- WebSocket для реального времени
- Хранение истории сообщений
- Уведомления

## Решения

### Решение задания 1: API для блога
```go


// blog/models.go
package blog

type Post struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    AuthorID  int       `json:"author_id"`
    CreatedAt time.Time `json:"created_at"`
}

type Comment struct {
    ID        int       `json:"id"`
    PostID    int       `json:"post_id"`
    Content   string    `json:"content"`
    AuthorID  int       `json:"author_id"`
    CreatedAt time.Time `json:"created_at"`
}

// blog/handlers.go
package blog

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

type BlogHandler struct {
    store *Store
}

func NewBlogHandler(store *Store) *BlogHandler {
    return &BlogHandler{store: store}
}

// @Summary Create a new post
// @Tags posts
// @Accept json
// @Produce json
// @Param post body Post true "Post object"
// @Success 201 {object} Response{data=Post}
// @Failure 400 {object} Response
// @Security ApiKeyAuth
// @Router /posts [post]
func (h *BlogHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
    var post Post
    if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
        JSONResponse(w, http.StatusBadRequest, Response{
            Status:  "error",
            Message: "Invalid request payload",
        })
        return
    }
    
    // Получаем ID пользователя из контекста (установлен в AuthMiddleware)
    userID := r.Context().Value("user_id").(int)
    post.AuthorID = userID
    
    if err := h.store.CreatePost(&post); err != nil {
        JSONResponse(w, http.StatusInternalServerError, Response{
            Status:  "error",
            Message: "Failed to create post",
        })
        return
    }
    
    JSONResponse(w, http.StatusCreated, Response{
        Status: "success",
        Data:   post,
    })
}

// blog/middleware.go
package blog

import (
    "context"
    "net/http"
    "strings"
)

func (h *BlogHandler) AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            JSONResponse(w, http.StatusUnauthorized, Response{
                Status:  "error",
                Message: "No authorization token provided",
            })
            return
        }
        
        // Извлекаем токен
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            JSONResponse(w, http.StatusUnauthorized, Response{
                Status:  "error",
                Message: "Invalid authorization header",
            })
            return
        }
        
        // Проверяем токен и получаем ID пользователя
        userID, err := h.store.ValidateToken(parts[1])
        if err != nil {
            JSONResponse(w, http.StatusUnauthorized, Response{
                Status:  "error",
                Message: "Invalid token",
            })
            return
        }
        
        // Добавляем ID пользователя в контекст
        ctx := context.WithValue(r.Context(), "user_id", userID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// blog/main.go
package main

func main() {
    store := NewStore(/* db connection */)
    handler := NewBlogHandler(store)
    
    r := mux.NewRouter()
    api := r.PathPrefix("/api/v1").Subrouter()
    
    // Публичные маршруты
    api.HandleFunc("/login", handler.Login).Methods("POST")
    api.HandleFunc("/register", handler.Register).Methods("POST")
    
    // Защищенные маршруты
    protected := api.PathPrefix("").Subrouter()
    protected.Use(handler.AuthMiddleware)
    
    protected.HandleFunc("/posts", handler.CreatePost).Methods("POST")
    protected.HandleFunc("/posts", handler.GetPosts).Methods("GET")
    protected.HandleFunc("/posts/{id:[0-9]+}", handler.GetPost).Methods("GET")
    protected.HandleFunc("/posts/{id:[0-9]+}", handler.UpdatePost).Methods("PUT")
    protected.HandleFunc("/posts/{id:[0-9]+}", handler.DeletePost).Methods("DELETE")
    
    protected.HandleFunc("/posts/{id:[0-9]+}/comments", handler.CreateComment).Methods("POST")
    
    log.Fatal(http.ListenAndServe(":8080", r))
}
```

### Решение задания 2: API для файлового хранилища
```go


// storage/models.go
package storage

type File struct {
    ID          string    `json:"id"`
    Name        string    `json:"name"`
    Size        int64     `json:"size"`
    ContentType string    `json:"content_type"`
    OwnerID     int       `json:"owner_id"`
    Path        string    `json:"-"`
    CreatedAt   time.Time `json:"created_at"`
}

// storage/handlers.go
package storage

import (
    "io"
    "net/http"
    "path/filepath"
)

type StorageHandler struct {
    store      *Store
    uploadPath string
}

func (h *StorageHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
    // Максимальный размер файла - 32MB
    r.ParseMultipartForm(32 << 20)
    
    file, header, err := r.FormFile("file")
    if err != nil {
        JSONResponse(w, http.StatusBadRequest, Response{
            Status:  "error",
            Message: "No file provided",
        })
        return
    }
    defer file.Close()
    
    // Проверяем тип файла
    contentType := header.Header.Get("Content-Type")
    if !h.isAllowedType(contentType) {
        JSONResponse(w, http.StatusBadRequest, Response{
            Status:  "error",
            Message: "File type not allowed",
        })
        return
    }
    
    // Генерируем уникальное имя файла
    filename := generateUniqueFilename(header.Filename)
    path := filepath.Join(h.uploadPath, filename)
    
    // Создаем файл
    dst, err := os.Create(path)
    if err != nil {
        JSONResponse(w, http.StatusInternalServerError, Response{
            Status:  "error",
            Message: "Failed to create file",
        })
        return
    }
    defer dst.Close()
    
    // Копируем содержимое
    if _, err := io.Copy(dst, file); err != nil {
        JSONResponse(w, http.StatusInternalServerError, Response{
            Status:  "error",
            Message: "Failed to save file",
        })
        return
    }
    
    // Сохраняем метаданные
    fileInfo := &File{
        Name:        header.Filename,
        Size:        header.Size,
        ContentType: contentType,
        OwnerID:     r.Context().Value("user_id").(int),
        Path:        path,
    }
    
    if err := h.store.SaveFile(fileInfo); err != nil {
        os.Remove(path) // Удаляем файл при ошибке
        JSONResponse(w, http.StatusInternalServerError, Response{
            Status:  "error",
            Message: "Failed to save file metadata",
        })
        return
    }
    
    JSONResponse(w, http.StatusCreated, Response{
        Status: "success",
        Data:   fileInfo,
    })
}

func (h *StorageHandler) DownloadFile(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    fileID := vars["id"]
    
    file, err := h.store.GetFile(fileID)
    if err != nil {
        JSONResponse(w, http.StatusNotFound, Response{
            Status:  "error",
            Message: "File not found",
        })
        return
    }
    
    // Проверяем права доступа
    userID := r.Context().Value("user_id").(int)
    if !h.canAccessFile(userID, file) {
        JSONResponse(w, http.StatusForbidden, Response{
            Status:  "error",
            Message: "Access denied",
        })
        return
    }
    
    // Открываем файл
    f, err := os.Open(file.Path)
    if err != nil {
        JSONResponse(w, http.StatusInternalServerError, Response{
            Status:  "error",
            Message: "Failed to read file",
        })
        return
    }
    defer f.Close()
    
    // Устанавливаем заголовки
    w.Header().Set("Content-Type", file.ContentType)
    w.Header().Set("Content-Disposition", 
        fmt.Sprintf("attachment; filename=%s", file.Name))
    
    // Отправляем файл
    io.Copy(w, f)
}
```

### Решение задания 3: API для чата
```go


// chat/models.go
package chat

type Message struct {
    ID        int       `json:"id"`
    ChatID    int       `json:"chat_id"`
    UserID    int       `json:"user_id"`
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
}

type Chat struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Type      string    `json:"type"` // "private" или "group"
    CreatedAt time.Time `json:"created_at"`
}

// chat/websocket.go
package chat

import (
    "github.com/gorilla/websocket"
    "sync"
)

type Client struct {
    conn     *websocket.Conn
    send     chan []byte
    userID   int
    chatID   int
    hub      *Hub
}

type Hub struct {
    clients    map[*Client]bool
    broadcast  chan []byte
    register   chan *Client
    unregister chan *Client
    mutex      sync.Mutex
}

func NewHub() *Hub {
    return &Hub{
        clients:    make(map[*Client]bool),
        broadcast:  make(chan []byte),
        register:   make(chan *Client),
        unregister: make(chan *Client),
    }
}

func (h *Hub) Run() {
    for {
        select {
        case client := <-h.register:
            h.mutex.Lock()
            h.clients[client] = true
            h.mutex.Unlock()
            
        case client := <-h.unregister:
            h.mutex.Lock()
            if _, ok := h.clients[client]; ok {
                delete(h.clients, client)
                close(client.send)
            }
            h.mutex.Unlock()
            
        case message := <-h.broadcast:
            h.mutex.Lock()
            for client := range h.clients {
                select {
                case client.send <- message:
                default:
                    close(client.send)
                    delete(h.clients, client)
                }
            }
            h.mutex.Unlock()
        }
    }
}

// chat/handlers.go
package chat

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true // В продакшене нужна реальная проверка
    },
}

func (h *ChatHandler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    chatID, _ := strconv.Atoi(vars["id"])
    userID := r.Context().Value("user_id").(int)
    
    // Проверяем доступ к чату
    if !h.canAccessChat(userID, chatID) {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }
    
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Printf("Failed to upgrade connection: %v", err)
        return
    }
    
    client := &Client{
        conn:   conn,
        send:   make(chan []byte, 256),
        userID: userID,
        chatID: chatID,
        hub:    h.hub,
    }
    
    client.hub.register <- client
    
    // Запускаем горутины для чтения и записи
    go client.writePump()
    go client.readPump()
}

func (c *Client) readPump() {
    defer func() {
        c.hub.unregister <- c
        c.conn.Close()
    }()
    
    c.conn.SetReadLimit(maxMessageSize)
    c.conn.SetReadDeadline(time.Now().Add(pongWait))
    c.conn.SetPongHandler(func(string) error {
        c.conn.SetReadDeadline(time.Now().Add(pongWait))
        return nil
    })
    
    for {
        _, message, err := c.conn.ReadMessage()
        if err != nil {
            if websocket.IsUnexpectedCloseError(err,
                websocket.CloseGoingAway,
                websocket.CloseAbnormalClosure) {
                log.Printf("error: %v", err)
            }
            break
        }
        
        // Сохраняем сообщение в БД
        msg := &Message{
            ChatID:  c.chatID,
            UserID:  c.userID,
            Content: string(message),
        }
        
        if err := c.hub.store.SaveMessage(msg); err != nil {
            log.Printf("Failed to save message: %v", err)
            continue
        }
        
        // Отправляем всем в чате
        c.hub.broadcast <- message
    }
}
```

## Дополнительные материалы
- [RESTful Web Services](https://golang.org/doc/articles/wiki/#tmp_3)
- [Gorilla Mux](https://github.com/gorilla/mux)
- [Swagger with Go](https://github.com/swaggo/swag)
- [WebSocket](https://github.com/gorilla/websocket)

## Следующий урок
В следующем уроке мы изучим развертывание Go-приложений, включая контейнеризацию, CI/CD и мониторинг.