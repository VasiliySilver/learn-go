# Урок 14: Безопасность в Go

## Содержание
1. [Криптография](#криптография)
2. [Защита от уязвимостей](#защита-от-уязвимостей)
3. [Аудит безопасности](#аудит-безопасности)
4. [Практические задания](#практические-задания)

## Криптография

### Хеширование паролей
```go


// security/password.go
package security

import (
    "golang.org/x/crypto/bcrypt"
    "crypto/rand"
    "encoding/base64"
)

// Генерация соли
func generateSalt(length int) (string, error) {
    bytes := make([]byte, length)
    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }
    return base64.URLEncoding.EncodeToString(bytes), nil
}

// Хеширование пароля
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}

// Проверка пароля
func CheckPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
```

### Шифрование данных
```go




// security/encryption.go
package security

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "errors"
    "io"
)

type Encryptor struct {
    key []byte
}

func NewEncryptor(key []byte) (*Encryptor, error) {
    if len(key) != 32 {
        return nil, errors.New("key must be 32 bytes")
    }
    return &Encryptor{key: key}, nil
}

func (e *Encryptor) Encrypt(data []byte) (string, error) {
    block, err := aes.NewCipher(e.key)
    if err != nil {
        return "", err
    }

    // Создаем nonce
    nonce := make([]byte, 12)
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return "", err
    }

    // Создаем GCM
    aesgcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }

    // Шифруем
    ciphertext := aesgcm.Seal(nil, nonce, data, nil)
    
    // Объединяем nonce и шифротекст
    result := make([]byte, len(nonce)+len(ciphertext))
    copy(result, nonce)
    copy(result[len(nonce):], ciphertext)

    return base64.StdEncoding.EncodeToString(result), nil
}

func (e *Encryptor) Decrypt(encrypted string) ([]byte, error) {
    data, err := base64.StdEncoding.DecodeString(encrypted)
    if err != nil {
        return nil, err
    }

    block, err := aes.NewCipher(e.key)
    if err != nil {
        return nil, err
    }

    aesgcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }

    if len(data) < 12 {
        return nil, errors.New("invalid ciphertext")
    }

    nonce := data[:12]
    ciphertext := data[12:]

    return aesgcm.Open(nil, nonce, ciphertext, nil)
}
```

## Защита от уязвимостей

### Защита от SQL-инъекций
```go




// security/sql.go
package security

import (
    "database/sql"
    "html"
    "strings"
)

// Безопасный запрос с параметрами
func SafeQuery(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
    // Используем подготовленные выражения
    stmt, err := db.Prepare(query)
    if err != nil {
        return nil, err
    }
    defer stmt.Close()
    
    return stmt.Query(args...)
}

// Санитизация входных данных
func SanitizeInput(input string) string {
    // Экранируем HTML
    escaped := html.EscapeString(input)
    // Удаляем потенциально опасные символы
    escaped = strings.Map(func(r rune) rune {
        if strings.ContainsRune("'\"<>();", r) {
            return -1
        }
        return r
    }, escaped)
    return escaped
}
```

### Защита от XSS и CSRF
```go




// security/web.go
package security

import (
    "crypto/rand"
    "encoding/base64"
    "net/http"
)

type CSRFToken struct {
    Secret []byte
}

func NewCSRFToken() (*CSRFToken, error) {
    secret := make([]byte, 32)
    if _, err := rand.Read(secret); err != nil {
        return nil, err
    }
    return &CSRFToken{Secret: secret}, nil
}

func (c *CSRFToken) Generate() string {
    token := make([]byte, 32)
    rand.Read(token)
    return base64.StdEncoding.EncodeToString(token)
}

func (c *CSRFToken) Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            token := r.Header.Get("X-CSRF-Token")
            if token == "" {
                http.Error(w, "CSRF token missing", http.StatusForbidden)
                return
            }
            // Проверка токена...
        }
        
        // Устанавливаем заголовки безопасности
        w.Header().Set("X-Frame-Options", "DENY")
        w.Header().Set("X-XSS-Protection", "1; mode=block")
        w.Header().Set("X-Content-Type-Options", "nosniff")
        w.Header().Set("Content-Security-Policy", 
            "default-src 'self'; script-src 'self'; style-src 'self';")
        
        next.ServeHTTP(w, r)
    })
}
```

## Аудит безопасности

### Логирование событий безопасности
```go




// security/audit.go
package security

import (
    "encoding/json"
    "time"
)

type SecurityEvent struct {
    Timestamp time.Time     `json:"timestamp"`
    Type      string       `json:"type"`
    UserID    string       `json:"user_id"`
    IP        string       `json:"ip"`
    Action    string       `json:"action"`
    Status    string       `json:"status"`
    Details   interface{}  `json:"details,omitempty"`
}

type SecurityAuditor struct {
    logger Logger
}

func NewSecurityAuditor(logger Logger) *SecurityAuditor {
    return &SecurityAuditor{logger: logger}
}

func (a *SecurityAuditor) LogEvent(event SecurityEvent) error {
    event.Timestamp = time.Now()
    
    data, err := json.Marshal(event)
    if err != nil {
        return err
    }
    
    return a.logger.Log(string(data))
}

// Middleware для аудита
func (a *SecurityAuditor) Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // Создаем wrapped response writer для получения статуса
        wrapped := NewResponseWriter(w)
        
        next.ServeHTTP(wrapped, r)
        
        // Логируем событие
        a.LogEvent(SecurityEvent{
            Type:    "http_request",
            UserID:  getUserID(r),
            IP:      r.RemoteAddr,
            Action:  r.Method + " " + r.URL.Path,
            Status:  http.StatusText(wrapped.Status()),
            Details: map[string]interface{}{
                "duration": time.Since(start),
                "status":   wrapped.Status(),
                "size":     wrapped.Size(),
            },
        })
    })
}
```

## Практические задания

### Задание 1: Безопасная аутентификация
Реализуйте систему аутентификации с:
- Безопасным хранением паролей
- Двухфакторной аутентификацией
- Защитой от брутфорса
- Аудитом действий

### Задание 2: Шифрование данных
Создайте систему безопасного хранения данных:
- Шифрование в покое
- Шифрование в движении
- Управление ключами
- Ротация ключей

### Задание 3: Защита API
Реализуйте комплексную защиту API:
- OAuth 2.0
- Rate limiting
- Input validation
- Security headers

## Решения

### Решение задания 1: Безопасная аутентификация
```go 




// auth/service.go
package auth

import (
    "time"
    "github.com/pquerna/otp/totp"
)

type AuthService struct {
    store          UserStore
    auditor        SecurityAuditor
    rateLimiter    RateLimiter
    tokenManager   TokenManager
}

type User struct {
    ID            string
    Email         string
    PasswordHash  string
    TOTPSecret    string
    FailedAttempts int
    LastFailedAt   time.Time
    Locked         bool
}

func (s *AuthService) Login(email, password, totpCode string) (*Token, error) {
    // Проверяем rate limit
    if err := s.rateLimiter.Check(email); err != nil {
        s.auditor.LogEvent(SecurityEvent{
            Type:    "login_attempt",
            UserID:  email,
            Action:  "rate_limit_exceeded",
            Status:  "failed",
        })
        return nil, err
    }
    
    // Получаем пользователя
    user, err := s.store.GetUserByEmail(email)
    if err != nil {
        return nil, err
    }
    
    // Проверяем блокировку
    if user.Locked {
        s.auditor.LogEvent(SecurityEvent{
            Type:    "login_attempt",
            UserID:  user.ID,
            Action:  "account_locked",
            Status:  "failed",
        })
        return nil, ErrAccountLocked
    }
    
    // Проверяем пароль
    if !CheckPassword(password, user.PasswordHash) {
        s.handleFailedAttempt(user)
        return nil, ErrInvalidCredentials
    }
    
    // Проверяем TOTP
    if !totp.Validate(totpCode, user.TOTPSecret) {
        s.handleFailedAttempt(user)
        return nil, ErrInvalidTOTP
    }
    
    // Сбрасываем счетчик неудачных попыток
    user.FailedAttempts = 0
    user.LastFailedAt = time.Time{}
    s.store.UpdateUser(user)
    
    // Создаем токен
    token, err := s.tokenManager.CreateToken(user.ID)
    if err != nil {
        return nil, err
    }
    
    s.auditor.LogEvent(SecurityEvent{
        Type:    "login",
        UserID:  user.ID,
        Action:  "login_success",
        Status:  "success",
    })
    
    return token, nil
}

func (s *AuthService) handleFailedAttempt(user *User) {
    user.FailedAttempts++
    user.LastFailedAt = time.Now()
    
    // Блокируем аккаунт после 5 неудачных попыток
    if user.FailedAttempts >= 5 {
        user.Locked = true
    }
    
    s.store.UpdateUser(user)
    
    s.auditor.LogEvent(SecurityEvent{
        Type:    "login_attempt",
        UserID:  user.ID,
        Action:  "invalid_credentials",
        Status:  "failed",
        Details: map[string]interface{}{
            "attempts": user.FailedAttempts,
            "locked":   user.Locked,
        },
    })
}
```

### Решение задания 2: Шифрование данных
```go 



// encryption/service.go
package encryption

import (
    "crypto/aes"
    "crypto/rand"
    "encoding/json"
    "time"
)

type EncryptionService struct {
    keyStore    KeyStore
    dataStore   DataStore
    auditor     SecurityAuditor
}

type Key struct {
    ID        string
    Value     []byte
    Version   int
    CreatedAt time.Time
    ExpiresAt time.Time
}

type EncryptedData struct {
    Data      []byte
    KeyID     string
    CreatedAt time.Time
}

func (s *EncryptionService) EncryptData(data []byte) (*EncryptedData, error) {
    // Получаем актуальный ключ
    key, err := s.keyStore.GetCurrentKey()
    if err != nil {
        return nil, err
    }
    
    // Создаем шифровальщик
    encryptor, err := NewEncryptor(key.Value)
    if err != nil {
        return nil, err
    }
    
    // Шифруем данные
    encrypted, err := encryptor.Encrypt(data)
    if err != nil {
        return nil, err
    }
    
    encData := &EncryptedData{
        Data:      []byte(encrypted),
        KeyID:     key.ID,
        CreatedAt: time.Now(),
    }
    
    // Сохраняем зашифрованные данные
    if err := s.dataStore.Save(encData); err != nil {
        return nil, err
    }
    
    s.auditor.LogEvent(SecurityEvent{
        Type:    "data_encryption",
        Action:  "encrypt",
        Status:  "success",
        Details: map[string]interface{}{
            "key_id":     key.ID,
            "key_version": key.Version,
        },
    })
    
    return encData, nil
}

func (s *EncryptionService) DecryptData(id string) ([]byte, error) {
    // Получаем зашифрованные данные
    encData, err := s.dataStore.Get(id)
    if err != nil {
        return nil, err
    }
    
    // Получаем ключ
    key, err := s.keyStore.GetKey(encData.KeyID)
    if err != nil {
        return nil, err
    }
    
    // Создаем шифровальщик
    encryptor, err := NewEncryptor(key.Value)
    if err != nil {
        return nil, err
    }
    
    // Расшифровываем данные
    decrypted, err := encryptor.Decrypt(string(encData.Data))
    if err != nil {
        return nil, err
    }
    
    s.auditor.LogEvent(SecurityEvent{
        Type:    "data_encryption",
        Action:  "decrypt",
        Status:  "success",
        Details: map[string]interface{}{
            "key_id":     key.ID,
            "key_version": key.Version,
        },
    })
    
    return decrypted, nil
}

func (s *EncryptionService) RotateKeys() error {
    // Генерируем новый ключ
    newKey := &Key{
        ID:        generateID(),
        Value:     make([]byte, 32),
        Version:   1,
        CreatedAt: time.Now(),
        ExpiresAt: time.Now().AddDate(0, 1, 0), // Срок действия 1 месяц
    }
    
    if _, err := rand.Read(newKey.Value); err != nil {
        return err
    }
    
    // Сохраняем новый ключ
    if err := s.keyStore.SaveKey(newKey); err != nil {
        return err
    }
    
    // Перешифровываем данные с новым ключом
    data, err := s.dataStore.GetAll()
    if err != nil {
        return err
    }
    
    for _, encData := range data {
        // Расшифровываем старым ключом
        decrypted, err := s.DecryptData(encData.ID)
        if err != nil {
            continue
        }
        
        // Шифруем новым ключом
        if _, err := s.EncryptData(decrypted); err != nil {
            continue
        }
    }
    
    s.auditor.LogEvent(SecurityEvent{
        Type:    "key_rotation",
        Action:  "rotate",
        Status:  "success",
        Details: map[string]interface{}{
            "new_key_id": newKey.ID,
            "version":    newKey.Version,
        },
    })
    
    return nil
}
```

### Решение задания 3: Защита API
```go




// api/security.go
package api

import (
    "net/http"
    "golang.org/x/time/rate"
    "github.com/go-oauth2/oauth2/v4/server"
)

type SecurityMiddleware struct {
    oauth      *server.Server
    limiter    *rate.Limiter
    validator  *InputValidator
    auditor    *SecurityAuditor
}

func (m *SecurityMiddleware) Secure(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Проверяем OAuth токен
        _, err := m.oauth.ValidationBearerToken(r)
        if err != nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            m.auditor.LogEvent(SecurityEvent{
                Type:    "api_auth",
                Action:  "unauthorized",
                Status:  "failed",
                Details: map[string]interface{}{
                    "error": err.Error(),
                },
            })
            return
        }
        
        // Проверяем rate limit
        if !m.limiter.Allow() {
            http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
            m.auditor.LogEvent(SecurityEvent{
                Type:    "api_rate_limit",
                Action:  "exceeded",
                Status:  "failed",
            })
            return
        }
        
        // Валидируем входные данные
        if err := m.validator.Validate(r); err != nil {
            http.Error(w, "Invalid Input", http.StatusBadRequest)
            m.auditor.LogEvent(SecurityEvent{
                Type:    "api_validation",
                Action:  "invalid_input",
                Status:  "failed",
                Details: map[string]interface{}{
                    "error": err.Error(),
                },
            })
            return
        }
        
        // Устанавливаем заголовки безопасности
        w.Header().Set("X-Content-Type-Options", "nosniff")
        w.Header().Set("X-Frame-Options", "DENY")
        w.Header().Set("Content-Security-Policy", 
            "default-src 'self'; frame-ancestors 'none';")
        w.Header().Set("Strict-Transport-Security", 
            "max-age=31536000; includeSubDomains")
        
        next.ServeHTTP(w, r)
    })
}

type InputValidator struct {
    rules map[string][]ValidationRule
}

func (v *InputValidator) Validate(r *http.Request) error {
    // Валидация параметров запроса
    for param, rules := range v.rules {
        value := r.FormValue(param)
        for _, rule := range rules {
            if err := rule.Validate(value); err != nil {
                return err
            }
        }
    }
    return nil
}
```

## Дополнительные материалы
- [Go Cryptography](https://golang.org/pkg/crypto/)
- [OWASP Go Security Cheat Sheet](https://github.com/OWASP/CheatSheetSeries/blob/master/cheatsheets/Go_Security_Cheat_Sheet.md)
- [OAuth 2.0 in Go](https://github.com/go-oauth2/oauth2)
- [Security Headers](https://securityheaders.com/)

## Следующий урок
В следующем уроке мы рассмотрим создание CLI-приложений в Go, включая работу с аргументами командной строки, интерактивный ввод и форматирование вывода.