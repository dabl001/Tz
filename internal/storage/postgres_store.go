package storage

import (
	"context"
	"database/sql"
	"workmate-go/internal/model"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(db *sql.DB) *PostgresStore {
	return &PostgresStore{db: db}
}

func (s *PostgresStore) SaveCompletedTask(ctx context.Context, task *model.Task, input string) error {
	_, err := s.db.ExecContext(ctx, `
		INSERT INTO completed_tasks (id, input, result, error, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, task.ID, input, task.Result, task.Error, task.Status, task.CreatedAt, task.UpdatedAt)
	return err
}
