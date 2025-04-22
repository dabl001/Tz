package task

import (
	"context"
	"errors"
	"testing"
	"time"
	"workmate-go/internal/model"
)

// 🔧 Поддельный store
type mockStore struct {
	saved []*model.Task
}

func (m *mockStore) SaveTask(ctx context.Context, task *model.Task) error {
	m.saved = append(m.saved, task)
	return nil
}

func (m *mockStore) GetTask(ctx context.Context, id string) (*model.Task, error) {
	for _, t := range m.saved {
		if t.ID == id {
			return t, nil
		}
	}
	return nil, errors.New("not found")
}

// 🔧 Поддельный Processor
type mockProcessor struct {
	id     string
	result string
	err    error
}

func (p *mockProcessor) ID() string { return p.id }

func (p *mockProcessor) Execute(ctx context.Context) (string, error) {
	time.Sleep(100 * time.Millisecond) // имитируем работу
	return p.result, p.err
}

func TestTaskManager_CreateAndRun(t *testing.T) {
	store := &mockStore{}
	manager := NewManager(store, nil)

	task, err := manager.Create(context.Background(), &mockProcessor{
		id:     "input-abc",
		result: "done!",
		err:    nil,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if task.ID == "" {
		t.Fatal("expected task ID to be set")
	}

	time.Sleep(150 * time.Millisecond) // ждём выполнения

	if len(store.saved) < 2 {
		t.Fatalf("expected task to be saved twice (initial + after execution), got %d", len(store.saved))
	}

	final := store.saved[len(store.saved)-1]
	if final.Status != model.StatusCompleted {
		t.Fatalf("expected status to be completed, got %s", final.Status)
	}
	if final.Result != "done!" {
		t.Fatalf("unexpected result: %s", final.Result)
	}
}
