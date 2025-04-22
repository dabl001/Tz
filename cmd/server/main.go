package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"workmate-go/internal/handler"
	"workmate-go/internal/storage"
	"workmate-go/internal/task"

	_ "github.com/lib/pq"
)

func main() {
	// Загружаем переменные окружения
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	// Подключение к Redis
	store := storage.NewRedisStore(redisAddr)

	pgDSN := os.Getenv("POSTGRES_DSN")
	db, err := sql.Open("postgres", pgDSN)
	if err != nil {
		log.Fatal("❌ Failed to connect to PostgreSQL:", err)
	}

	pgStore := storage.NewPostgresStore(db)

	// Создаём менеджер задач
	manager := task.NewManager(store, pgStore)

	// Инициализируем HTTP хендлер
	h := &handler.TaskHandler{Manager: manager}

	// Регистрируем эндпоинты
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.CreateTask(w, r)
		} else if r.Method == http.MethodGet {
			h.GetTask(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("🚀 Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
