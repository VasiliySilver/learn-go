# Урок 12: Развертывание Go-приложений

## Содержание
1. [Контейнеризация](#контейнеризация)
2. [CI/CD](#cicd)
3. [Мониторинг](#мониторинг)
4. [Практические задания](#практические-задания)

## Контейнеризация

### Dockerfile для Go-приложения
```dockerfile



# Многоэтапная сборка
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Копируем файлы с зависимостями
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/app

# Финальный образ
FROM alpine:3.18

WORKDIR /app

# Копируем бинарный файл из предыдущего этапа
COPY --from=builder /app/main .
COPY --from=builder /app/configs ./configs

# Определяем точку входа
ENTRYPOINT ["./main"]
```

### Docker Compose
```yaml



version: '3.8'

services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - DB_HOST=postgres
      - DB_USER=postgres
      - DB_PASSWORD=secret
      - DB_NAME=myapp
    depends_on:
      - postgres
      - redis

  postgres:
    image: postgres:15-alpine
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=secret
      - POSTGRES_DB=myapp
    volumes:
      - postgres_data:/var/lib/postgresql/data

  redis:
    image: redis:7-alpine
    volumes:
      - redis_data:/data

volumes:
  postgres_data:
  redis_data:
```

## CI/CD

### GitHub Actions
```yaml



name: CI/CD

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3

    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.21'

    - name: Run tests
      run: go test -v ./...

    - name: Run linter
      uses: golangci/golangci-lint-action@v3
      with:
        version: latest

  build:
    needs: test
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3

    - name: Build Docker image
      run: docker build -t myapp .

    - name: Login to Docker Hub
      uses: docker/login-action@v2
      with:
        username: ${{ secrets.DOCKER_HUB_USERNAME }}
        password: ${{ secrets.DOCKER_HUB_TOKEN }}

    - name: Push Docker image
      run: |
        docker tag myapp ${{ secrets.DOCKER_HUB_USERNAME }}/myapp:latest
        docker push ${{ secrets.DOCKER_HUB_USERNAME }}/myapp:latest

  deploy:
    needs: build
    runs-on: ubuntu-latest
    steps:
    - name: Deploy to production
      uses: appleboy/ssh-action@master
      with:
        host: ${{ secrets.SERVER_HOST }}
        username: ${{ secrets.SERVER_USER }}
        key: ${{ secrets.SSH_PRIVATE_KEY }}
        script: |
          docker pull ${{ secrets.DOCKER_HUB_USERNAME }}/myapp:latest
          docker-compose up -d
```

## Мониторинг

### Prometheus метрики
```go



package metrics

import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
)

var (
    RequestsTotal = promauto.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "endpoint", "status"},
    )

    RequestDuration = promauto.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "http_request_duration_seconds",
            Help:    "HTTP request duration in seconds",
            Buckets: prometheus.DefBuckets,
        },
        []string{"method", "endpoint"},
    )

    ActiveConnections = promauto.NewGauge(
        prometheus.GaugeOpts{
            Name: "active_connections",
            Help: "Number of currently active connections",
        },
    )
)

// Middleware для сбора метрик
func MetricsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // Оборачиваем ResponseWriter для получения статуса ответа
        wrapped := NewResponseWriter(w)
        
        next.ServeHTTP(wrapped, r)
        
        // Записываем метрики
        duration := time.Since(start).Seconds()
        RequestsTotal.WithLabelValues(
            r.Method,
            r.URL.Path,
            fmt.Sprintf("%d", wrapped.status),
        ).Inc()
        
        RequestDuration.WithLabelValues(
            r.Method,
            r.URL.Path,
        ).Observe(duration)
    })
}
```

### Grafana Dashboard
```json



{
  "annotations": {
    "list": []
  },
  "editable": true,
  "fiscalYearStartMonth": 0,
  "graphTooltip": 0,
  "links": [],
  "liveNow": false,
  "panels": [
    {
      "datasource": {
        "type": "prometheus",
        "uid": "prometheus"
      },
      "fieldConfig": {
        "defaults": {
          "color": {
            "mode": "palette-classic"
          },
          "custom": {
            "axisCenteredZero": false,
            "axisColorMode": "text",
            "axisLabel": "",
            "axisPlacement": "auto",
            "barAlignment": 0,
            "drawStyle": "line",
            "fillOpacity": 0,
            "gradientMode": "none",
            "hideFrom": {
              "legend": false,
              "tooltip": false,
              "viz": false
            },
            "lineInterpolation": "linear",
            "lineWidth": 1,
            "pointSize": 5,
            "scaleDistribution": {
              "type": "linear"
            },
            "showPoints": "auto",
            "spanNulls": false,
            "stacking": {
              "group": "A",
              "mode": "none"
            },
            "thresholdsStyle": {
              "mode": "off"
            }
          },
          "mappings": [],
          "thresholds": {
            "mode": "absolute",
            "steps": [
              {
                "color": "green",
                "value": null
              }
            ]
          }
        },
        "overrides": []
      },
      "gridPos": {
        "h": 8,
        "w": 12,
        "x": 0,
        "y": 0
      },
      "id": 1,
      "options": {
        "legend": {
          "calcs": [],
          "displayMode": "list",
          "placement": "bottom",
          "showLegend": true
        },
        "tooltip": {
          "mode": "single",
          "sort": "none"
        }
      },
      "targets": [
        {
          "datasource": {
            "type": "prometheus",
            "uid": "prometheus"
          },
          "expr": "rate(http_requests_total[5m])",
          "refId": "A"
        }
      ],
      "title": "Request Rate",
      "type": "timeseries"
    }
  ],
  "refresh": "5s",
  "schemaVersion": 38,
  "style": "dark",
  "tags": [],
  "templating": {
    "list": []
  },
  "time": {
    "from": "now-6h",
    "to": "now"
  },
  "timepicker": {},
  "timezone": "",
  "title": "API Monitoring",
  "version": 0,
  "weekStart": ""
}
```

## Практические задания

### Задание 1: Контейнеризация микросервисов
Создайте Docker-инфраструктуру для микросервисного приложения:
- Несколько сервисов
- Общая сеть
- Управление конфигурацией
- Persistent volumes

### Задание 2: Настройка CI/CD пайплайна
Реализуйте полный пайплайн CI/CD:
- Автоматические тесты
- Линтинг кода
- Сборка и публикация образов
- Автоматический деплой

### Задание 3: Система мониторинга
Настройте комплексный мониторинг:
- Prometheus метрики
- Grafana дашборды
- Алерты
- Логирование

## Решения

### Решение задания 1: Контейнеризация микросервисов
```yaml



# docker-compose.yml
version: '3.8'

services:
  api-gateway:
    build:
      context: ./api-gateway
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    environment:
      - AUTH_SERVICE_URL=http://auth-service:8081
      - USER_SERVICE_URL=http://user-service:8082
    depends_on:
      - auth-service
      - user-service
    networks:
      - microservices

  auth-service:
    build:
      context: ./auth-service
      dockerfile: Dockerfile
    environment:
      - DB_HOST=postgres
      - REDIS_HOST=redis
    depends_on:
      - postgres
      - redis
    networks:
      - microservices

  user-service:
    build:
      context: ./user-service
      dockerfile: Dockerfile
    environment:
      - DB_HOST=postgres
    depends_on:
      - postgres
    networks:
      - microservices

  postgres:
    image: postgres:15-alpine
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=secret
      - POSTGRES_DB=myapp
    volumes:
      - postgres_data:/var/lib/postgresql/data
    networks:
      - microservices

  redis:
    image: redis:7-alpine
    volumes:
      - redis_data:/data
    networks:
      - microservices

networks:
  microservices:
    driver: bridge

volumes:
  postgres_data:
  redis_data:
```

### Решение задания 2: Настройка CI/CD пайплайна
```yaml



# .github/workflows/ci-cd.yml
name: CI/CD Pipeline

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  lint:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3

    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.21'

    - name: Run linter
      uses: golangci/golangci-lint-action@v3
      with:
        version: latest

  test:
    runs-on: ubuntu-latest
    services:
      postgres:
        image: postgres:15
        env:
          POSTGRES_USER: postgres
          POSTGRES_PASSWORD: postgres
          POSTGRES_DB: test
        ports:
          - 5432:5432
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5

    steps:
    - uses: actions/checkout@v3

    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.21'

    - name: Run tests
      run: go test -v -race -coverprofile=coverage.txt ./...
      env:
        DB_HOST: localhost
        DB_USER: postgres
        DB_PASSWORD: postgres
        DB_NAME: test

    - name: Upload coverage
      uses: codecov/codecov-action@v3
      with:
        file: ./coverage.txt

  build-and-push:
    needs: [lint, test]
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3

    - name: Set up Docker Buildx
      uses: docker/setup-buildx-action@v2

    - name: Login to Docker Hub
      uses: docker/login-action@v2
      with:
        username: ${{ secrets.DOCKER_HUB_USERNAME }}
        password: ${{ secrets.DOCKER_HUB_TOKEN }}

    - name: Build and push images
      run: |
        for service in api-gateway auth-service user-service; do
          docker buildx build \
            --platform linux/amd64,linux/arm64 \
            --push \
            -t ${{ secrets.DOCKER_HUB_USERNAME }}/$service:${{ github.sha }} \
            -t ${{ secrets.DOCKER_HUB_USERNAME }}/$service:latest \
            ./$service
        done

  deploy:
    needs: build-and-push
    runs-on: ubuntu-latest
    steps:
    - name: Deploy to production
      uses: appleboy/ssh-action@master
      with:
        host: ${{ secrets.PROD_HOST }}
        username: ${{ secrets.PROD_USER }}
        key: ${{ secrets.SSH_PRIVATE_KEY }}
        script: |
          cd /opt/myapp
          docker-compose pull
          docker-compose up -d
          docker system prune -f
```

### Решение задания 3: Система мониторинга
```yaml



# prometheus/prometheus.yml
global:
  scrape_interval: 15s
  evaluation_interval: 15s

scrape_configs:
  - job_name: 'api-gateway'
    static_configs:
      - targets: ['api-gateway:8080']

  - job_name: 'auth-service'
    static_configs:
      - targets: ['auth-service:8081']

  - job_name: 'user-service'
    static_configs:
      - targets: ['user-service:8082']

# alertmanager/alertmanager.yml
global:
  resolve_timeout: 5m

route:
  group_by: ['alertname']
  group_wait: 10s
  group_interval: 10s
  repeat_interval: 1h
  receiver: 'slack'

receivers:
- name: 'slack'
  slack_configs:
  - api_url: 'https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX'
    channel: '#alerts'
    send_resolved: true

# docker-compose.monitoring.yml
version: '3.8'

services:
  prometheus:
    image: prom/prometheus:v2.45.0
    volumes:
      - ./prometheus:/etc/prometheus
      - prometheus_data:/prometheus
    command:
      - '--config.file=/etc/prometheus/prometheus.yml'
      - '--storage.tsdb.path=/prometheus'
      - '--web.console.libraries=/usr/share/prometheus/console_libraries'
      - '--web.console.templates=/usr/share/prometheus/consoles'
    ports:
      - "9090:9090"
    networks:
      - monitoring

  alertmanager:
    image: prom/alertmanager:v0.25.0
    volumes:
      - ./alertmanager:/etc/alertmanager
    command:
      - '--config.file=/etc/alertmanager/alertmanager.yml'
      - '--storage.path=/alertmanager'
    ports:
      - "9093:9093"
    networks:
      - monitoring

  grafana:
    image: grafana/grafana:10.0.3
    volumes:
      - grafana_data:/var/lib/grafana
      - ./grafana/provisioning:/etc/grafana/provisioning
    environment:
      - GF_SECURITY_ADMIN_USER=admin
      - GF_SECURITY_ADMIN_PASSWORD=secret
    ports:
      - "3000:3000"
    networks:
      - monitoring

  loki:
    image: grafana/loki:2.8.4
    ports:
      - "3100:3100"
    command: -config.file=/etc/loki/local-config.yaml
    networks:
      - monitoring

  promtail:
    image: grafana/promtail:2.8.4
    volumes:
      - /var/log:/var/log
      - ./promtail:/etc/promtail
    command: -config.file=/etc/promtail/config.yml
    networks:
      - monitoring

networks:
  monitoring:
    driver: bridge

volumes:
  prometheus_data:
  grafana_data:
```

## Дополнительные материалы
- [Docker Documentation](https://docs.docker.com/)
- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [Prometheus Documentation](https://prometheus.io/docs/introduction/overview/)
- [Grafana Documentation](https://grafana.com/docs/)

## Следующий урок
В следующем уроке мы изучим продвинутые темы Go, включая профилирование, оптимизацию производительности и паттерны проектирования.