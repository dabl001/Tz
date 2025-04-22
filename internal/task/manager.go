package task

import (
	"context"
	"log"
	"time"
	"workmate-go/internal/model"
	"workmate-go/internal/storage"

	"github.com/google/uuid"
)

type Manager struct {
	store storage.TaskStore
}

func NewManager(store storage.TaskStore) *Manager {
	return &Manager{store: store}
}

func (m *Manager) Create(ctx context.Context, p Processor) (*model.Task, error) {
	id := uuid.NewString()

	task := &model.Task{
		ID:        id,
		Status:    model.StatusPending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	log.Printf("🟡 Creating new task: %s", id)

	// Сохраняем как pending
	if err := m.store.SaveTask(ctx, task); err != nil {
		log.Printf("❌ Failed to save task %s: %v", id, err)
		return nil, err
	}

	log.Printf("✅ Task saved to store: %s", id)

	// Асинхронно запускаем выполнение
	go m.executeTask(context.Background(), task, p)

	return task, nil
}

func (m *Manager) executeTask(ctx context.Context, task *model.Task, p Processor) {
	log.Printf("▶️ Executing task: %s", task.ID)

	task.Status = model.StatusRunning
	task.UpdatedAt = time.Now()
	_ = m.store.SaveTask(ctx, task)

	result, err := p.Execute(ctx)
	if err != nil {
		task.Status = model.StatusFailed
		task.Error = err.Error()
		log.Printf("❌ Task failed: %s, error: %v", task.ID, err)
	} else {
		task.Status = model.StatusCompleted
		task.Result = result
		log.Printf("✅ Task completed: %s, result: %s", task.ID, result)
	}
	task.UpdatedAt = time.Now()
	_ = m.store.SaveTask(ctx, task)
}

func (m *Manager) Get(ctx context.Context, id string) (*model.Task, error) {
	return m.store.GetTask(ctx, id)
}
