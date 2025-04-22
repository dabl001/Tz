package task

import (
	"context"
	"workmate-go/internal/model"
)

type Processor interface {
	ID() string
	Execute(ctx context.Context) (string, error)
}

type TaskManager interface {
	Create(ctx context.Context, p Processor) (*model.Task, error)
	Get(ctx context.Context, id string) (*model.Task, error)
}
