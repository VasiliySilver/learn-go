# Урок 10: Работу с базами данных в Go

## Содержание
1. [SQL базы данных](#sql-базы-данных)
2. [NoSQL базы данных](#nosql-базы-данных)
3. [Миграции](#миграции)
4. [Практические задания](#практические-задания)

## SQL базы данных

### Подключение к базе данных
```go

// Подключение к PostgreSQL
import (
    "database/sql"
    _ "github.com/lib/pq"
)

func connectDB() (*sql.DB, error) {
    connStr := "host=localhost port=5432 user=postgres password=secret dbname=myapp sslmode=disable"
    return sql.Open("postgres", connStr)
}

// Проверка подключения
func checkConnection(db *sql.DB) error {
    return db.Ping()
}
```

### Базовые операции
```go

// Создание таблицы
const createTableSQL = `
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)`

// CRUD операции
// Create
func createUser(db *sql.DB, name, email string) (int, error) {
    var id int
    err := db.QueryRow(`
        INSERT INTO users (name, email)
        VALUES ($1, $2)
        RETURNING id
    `, name, email).Scan(&id)
    return id, err
}

// Read
func getUser(db *sql.DB, id int) (*User, error) {
    user := &User{}
    err := db.QueryRow(`
        SELECT id, name, email, created_at
        FROM users
        WHERE id = $1
    `, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
    return user, err
}

// Update
func updateUser(db *sql.DB, user *User) error {
    _, err := db.Exec(`
        UPDATE users
        SET name = $1, email = $2
        WHERE id = $3
    `, user.Name, user.Email, user.ID)
    return err
}

// Delete
func deleteUser(db *sql.DB, id int) error {
    _, err := db.Exec(`
        DELETE FROM users
        WHERE id = $1
    `, id)
    return err
}
```

## NoSQL базы данных

### Работа с MongoDB
```go

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
)

type MongoUser struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    Name      string            `bson:"name"`
    Email     string            `bson:"email"`
    CreatedAt time.Time         `bson:"created_at"`
}

// Подключение
func connectMongo() (*mongo.Client, error) {
    ctx := context.Background()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        return nil, err
    }
    return client, client.Ping(ctx, nil)
}

// CRUD операции
func createMongoUser(collection *mongo.Collection, user *MongoUser) error {
    ctx := context.Background()
    _, err := collection.InsertOne(ctx, user)
    return err
}

func getMongoUser(collection *mongo.Collection, id primitive.ObjectID) (*MongoUser, error) {
    ctx := context.Background()
    var user MongoUser
    err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
    return &user, err
}
```

## Миграции

### Использование golang-migrate
```go

import (
    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

func runMigrations(dbURL string) error {
    m, err := migrate.New(
        "file://migrations",
        dbURL,
    )
    if err != nil {
        return err
    }
    
    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        return err
    }
    return nil
}
```

## Практические задания

### Задание 1: Библиотечная система
Создайте систему управления библиотекой:
- Таблицы для книг, авторов и читателей
- CRUD операции
- Поиск и фильтрация
- Транзакции для выдачи книг

### Задание 2: Система блогов
Реализуйте систему блогов с использованием MongoDB:
- Посты, комментарии и пользователи
- Поиск по тегам
- Пагинация
- Агрегации

### Задание 3: Микросервис с несколькими БД
Создайте микросервис, использующий разные базы данных:
- PostgreSQL для пользователей
- MongoDB для логов
- Redis для кеширования

## Решения

### Решение задания 1: Библиотечная система
```go

// library/models.go
package library

import "time"

type Book struct {
    ID        int
    Title     string
    AuthorID  int
    ISBN      string
    Available bool
    CreatedAt time.Time
}

type Author struct {
    ID        int
    Name      string
    CreatedAt time.Time
}

type Reader struct {
    ID        int
    Name      string
    Email     string
    CreatedAt time.Time
}

type Loan struct {
    ID        int
    BookID    int
    ReaderID  int
    LoanDate  time.Time
    ReturnDate *time.Time
}

// library/db.go
package library

import (
    "database/sql"
    "time"
)

type Library struct {
    db *sql.DB
}

func NewLibrary(db *sql.DB) *Library {
    return &Library{db: db}
}

func (l *Library) CreateBook(book *Book) error {
    return l.db.QueryRow(`
        INSERT INTO books (title, author_id, isbn, available)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at
    `, book.Title, book.AuthorID, book.ISBN, book.Available).
        Scan(&book.ID, &book.CreatedAt)
}

func (l *Library) LoanBook(bookID, readerID int) error {
    tx, err := l.db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    // Проверяем доступность книги
    var available bool
    err = tx.QueryRow("SELECT available FROM books WHERE id = $1", bookID).
        Scan(&available)
    if err != nil {
        return err
    }
    if !available {
        return ErrBookNotAvailable
    }

    // Создаем запись о выдаче
    _, err = tx.Exec(`
        INSERT INTO loans (book_id, reader_id, loan_date)
        VALUES ($1, $2, CURRENT_TIMESTAMP)
    `, bookID, readerID)
    if err != nil {
        return err
    }

    // Обновляем статус книги
    _, err = tx.Exec(`
        UPDATE books
        SET available = false
        WHERE id = $1
    `, bookID)
    if err != nil {
        return err
    }

    return tx.Commit()
}

func (l *Library) SearchBooks(query string) ([]Book, error) {
    rows, err := l.db.Query(`
        SELECT b.id, b.title, b.isbn, b.available, b.created_at, 
               a.id, a.name
        FROM books b
        JOIN authors a ON b.author_id = a.id
        WHERE b.title ILIKE $1 OR a.name ILIKE $1
    `, "%"+query+"%")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var books []Book
    for rows.Next() {
        var book Book
        var author Author
        err := rows.Scan(
            &book.ID, &book.Title, &book.ISBN, &book.Available,
            &book.CreatedAt, &author.ID, &author.Name,
        )
        if err != nil {
            return nil, err
        }
        books = append(books, book)
    }
    return books, nil
}
```

### Решение задания 2: Система блогов
```go

// blog/models.go
package blog

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type Post struct {
    ID        primitive.ObjectID   `bson:"_id,omitempty"`
    Title     string              `bson:"title"`
    Content   string              `bson:"content"`
    AuthorID  primitive.ObjectID   `bson:"author_id"`
    Tags      []string            `bson:"tags"`
    Comments  []Comment           `bson:"comments"`
    CreatedAt time.Time           `bson:"created_at"`
}

type Comment struct {
    ID        primitive.ObjectID   `bson:"_id,omitempty"`
    Content   string              `bson:"content"`
    AuthorID  primitive.ObjectID   `bson:"author_id"`
    CreatedAt time.Time           `bson:"created_at"`
}

// blog/repository.go
package blog

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type BlogRepository struct {
    posts *mongo.Collection
}

func NewBlogRepository(db *mongo.Database) *BlogRepository {
    return &BlogRepository{
        posts: db.Collection("posts"),
    }
}

func (r *BlogRepository) CreatePost(post *Post) error {
    ctx := context.Background()
    post.CreatedAt = time.Now()
    _, err := r.posts.InsertOne(ctx, post)
    return err
}

func (r *BlogRepository) SearchByTags(tags []string, page, pageSize int) ([]Post, error) {
    ctx := context.Background()
    
    opts := options.Find().
        SetSort(bson.D{{"created_at", -1}}).
        SetSkip(int64((page - 1) * pageSize)).
        SetLimit(int64(pageSize))

    filter := bson.M{"tags": bson.M{"$in": tags}}
    
    cursor, err := r.posts.Find(ctx, filter, opts)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var posts []Post
    if err = cursor.All(ctx, &posts); err != nil {
        return nil, err
    }
    return posts, nil
}

func (r *BlogRepository) AddComment(postID primitive.ObjectID, comment *Comment) error {
    ctx := context.Background()
    comment.CreatedAt = time.Now()
    
    update := bson.M{
        "$push": bson.M{"comments": comment},
    }
    
    _, err := r.posts.UpdateOne(ctx, bson.M{"_id": postID}, update)
    return err
}

func (r *BlogRepository) GetPopularTags() ([]string, error) {
    ctx := context.Background()
    
    pipeline := mongo.Pipeline{
        {{
            "$unwind": "$tags",
        }},
        {{
            "$group": bson.D{
                {"_id", "$tags"},
                {"count", bson.D{{"$sum", 1}}},
            },
        }},
        {{
            "$sort": bson.D{{"count", -1}},
        }},
        {{
            "$limit", 10,
        }},
    }
    
    cursor, err := r.posts.Aggregate(ctx, pipeline)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var results []struct {
        Tag   string `bson:"_id"`
        Count int    `bson:"count"`
    }
    if err = cursor.All(ctx, &results); err != nil {
        return nil, err
    }

    tags := make([]string, len(results))
    for i, result := range results {
        tags[i] = result.Tag
    }
    return tags, nil
}
```

### Решение задания 3: Микросервис с несколькими БД
```go

// service/config.go
package service

type Config struct {
    PostgresURL string
    MongoURL    string
    RedisAddr   string
}

// service/service.go
package service

import (
    "context"
    "database/sql"
    "encoding/json"
    "go.mongodb.org/mongo-driver/mongo"
    "github.com/go-redis/redis/v8"
)

type Service struct {
    db          *sql.DB
    mongo       *mongo.Client
    redis       *redis.Client
    userRepo    *UserRepository
    logRepo     *LogRepository
    cacheRepo   *CacheRepository
}

func NewService(config Config) (*Service, error) {
    // Подключение к PostgreSQL
    db, err := sql.Open("postgres", config.PostgresURL)
    if err != nil {
        return nil, err
    }

    // Подключение к MongoDB
    mongoClient, err := mongo.Connect(context.Background(), 
        options.Client().ApplyURI(config.MongoURL))
    if err != nil {
        return nil, err
    }

    // Подключение к Redis
    redisClient := redis.NewClient(&redis.Options{
        Addr: config.RedisAddr,
    })

    return &Service{
        db:        db,
        mongo:     mongoClient,
        redis:     redisClient,
        userRepo:  NewUserRepository(db),
        logRepo:   NewLogRepository(mongoClient.Database("logs")),
        cacheRepo: NewCacheRepository(redisClient),
    }, nil
}

// Работа с пользователями (PostgreSQL)
func (s *Service) CreateUser(user *User) error {
    // Создаем пользователя
    if err := s.userRepo.Create(user); err != nil {
        return err
    }

    // Логируем действие
    s.logRepo.Log("user_created", map[string]interface{}{
        "user_id": user.ID,
        "email":   user.Email,
    })

    // Инвалидируем кеш
    s.cacheRepo.Delete(context.Background(), "users_list")
    
    return nil
}

func (s *Service) GetUser(id int) (*User, error) {
    // Проверяем кеш
    cacheKey := fmt.Sprintf("user:%d", id)
    if cached, err := s.cacheRepo.Get(context.Background(), cacheKey); err == nil {
        var user User
        if err := json.Unmarshal([]byte(cached), &user); err == nil {
            return &user, nil
        }
    }

    // Получаем из БД
    user, err := s.userRepo.Get(id)
    if err != nil {
        return nil, err
    }

    // Сохраняем в кеш
    if userData, err := json.Marshal(user); err == nil {
        s.cacheRepo.Set(context.Background(), cacheKey, string(userData), time.Hour)
    }

    return user, nil
}

// Логирование (MongoDB)
func (s *Service) LogAction(action string, data map[string]interface{}) error {
    return s.logRepo.Log(action, data)
}

// Кеширование (Redis)
func (s *Service) CacheData(key string, value interface{}, expiration time.Duration) error {
    data, err := json.Marshal(value)
    if err != nil {
        return err
    }
    return s.cacheRepo.Set(context.Background(), key, string(data), expiration)
}
```

## Дополнительные материалы
- [database/sql package](https://golang.org/pkg/database/sql/)
- [MongoDB Go Driver](https://docs.mongodb.com/drivers/go/)
- [go-redis](https://github.com/go-redis/redis)
- [golang-migrate](https://github.com/golang-migrate/migrate)

## Следующий урок
В следующем уроке мы изучим создание REST API с использованием Go, включая маршрутизацию, middleware и документацию API.