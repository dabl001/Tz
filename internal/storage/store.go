package storage

import (
	"context"
	"workmate-go/internal/model"
)

type TaskStore interface {
	SaveTask(ctx context.Context, task *model.Task) error
	GetTask(ctx context.Context, id string) (*model.Task, error)
}
