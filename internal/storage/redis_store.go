package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"workmate-go/internal/model"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(addr string) *RedisStore {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})
	return &RedisStore{client: rdb}
}

func (s *RedisStore) SaveTask(ctx context.Context, task *model.Task) error {
	data, err := json.Marshal(task)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("task:%s", task.ID)
	return s.client.Set(ctx, key, data, 24*time.Hour).Err()
}

func (s *RedisStore) GetTask(ctx context.Context, id string) (*model.Task, error) {
	key := fmt.Sprintf("task:%s", id)
	data, err := s.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var task model.Task
	if err := json.Unmarshal([]byte(data), &task); err != nil {
		return nil, err
	}
	return &task, nil
}
