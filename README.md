# WorkMate-Go

🎯 Простая HTTP API-система для управления долгими I/O-задачами (3–5 минут), написанная на Go с использованием Redis и PostgreSQL.

---

## 🚀 Возможности

- Создание задач, выполняемых асинхронно
- Получение статуса и результата задачи
- Redis для временного хранения
- PostgreSQL для истории выполненных задач
- Docker + `.env` + тесты

---

## 🧱 Технологии

- Go 1.22
- Redis 7
- PostgreSQL 15
- Docker / Docker Compose
- `net/http`, `database/sql`, `github.com/redis/go-redis/v9`, `github.com/lib/pq`

---

## ⚙️ Установка и запуск

1. Клонируй репозиторий:

```bash
git clone https://github.com/dabl001/WorkMate-Go.git
cd WorkMate-Go
```

2. Создай .env из шаблона:

```bash
cp .env.example .env
```

3. Собери и запусти проект:

```bash
docker-compose up --build
```

По умолчанию:

    API: http://localhost:8080

    Redis: localhost:6379

    PostgreSQL: localhost:5432

## 🔌 API
POST /tasks

### Создать задачу:

    POST http://localhost:8080/tasks

Content-Type: application/json

{

  "input": "your data"

}

GET /tasks?id=...

### Получить статус:

    GET http://localhost:8080/tasks?id=task-id

```markdown
### 🧪 Тесты

Запуск всех тестов:

```bash
go test ./...
```

Запуск отдельных пакетов:

```bash
go test ./internal/task
go test ./internal/storage
go test ./internal/handler
```
```

### 📁 Структура

cmd/server         → main.go (вход в приложение)

internal/model     → Task-модель

internal/task      → логика задач, менеджер, sample-задачи

internal/storage   → RedisStore, PostgresStore

internal/handler   → HTTP-эндпоинты

scripts/init.sql   → создание таблицы в PostgreSQL

---

## 🧩 Возможности для будущего развития

Проект спроектирован с расчётом на масштабирование. Возможные направления:

- ✅ Поддержка нескольких типов задач через интерфейс `Processor`
- ✅ Расширяемый `TaskManager` — легко внедрять новые хранилища
- 🧵 Очереди задач (Redis Streams / RabbitMQ / Kafka)
- 📊 Метрики и мониторинг (Prometheus / Grafana)
- 🛡️ Авторизация и аутентификация (JWT)
- 📂 Массовое создание задач через gRPC
- 🌍 Микросервисная архитектура с отдельными воркерами
- 🧠 Веб-интерфейс для отслеживания задач
- 🔄 Ретрай задач при сбоях


👤 Автор

Рабочее тестовое задание от Abylay Dauletkhan (dabl001)