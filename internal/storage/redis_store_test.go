package storage

import (
	"context"
	"os"
	"testing"
	"time"
	"workmate-go/internal/model"

	"github.com/redis/go-redis/v9"
)

func setupTestRedis() *RedisStore {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}
	client := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   1, // отдельная Redis-база для тестов
	})
	return &RedisStore{client: client}
}

func TestSaveAndGetTask(t *testing.T) {
	store := setupTestRedis()
	ctx := context.Background()

	task := &model.Task{
		ID:        "test-task-123",
		Status:    model.StatusPending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := store.SaveTask(ctx, task)
	if err != nil {
		t.Errorf("SaveTask() returned error: %v", err)
	}

	got, err := store.GetTask(ctx, "test-task-123")
	if err != nil {
		t.Errorf("GetTask() returned error: %v", err)
	}

	if got.ID != task.ID {
		t.Errorf("Expected ID %s, got %s", task.ID, got.ID)
	}
	if got.Status != task.Status {
		t.Errorf("Expected Status %s, got %s", task.Status, got.Status)
	}
}
