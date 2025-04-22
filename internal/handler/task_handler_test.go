package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"workmate-go/internal/model"
	"workmate-go/internal/task"
)

// Мок-менеджер
type mockManager struct {
	taskToReturn *model.Task
	errToReturn  error
}

func (m *mockManager) Create(ctx context.Context, p task.Processor) (*model.Task, error) {
	return m.taskToReturn, m.errToReturn
}

func (m *mockManager) Get(ctx context.Context, id string) (*model.Task, error) {
	return m.taskToReturn, m.errToReturn
}

func TestCreateTaskHandler(t *testing.T) {
	expectedTask := &model.Task{
		ID:        "test-id-123",
		Status:    model.StatusPending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	h := &TaskHandler{
		Manager: &mockManager{taskToReturn: expectedTask},
	}

	body := `{"input": "test input"}`
	req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	h.CreateTask(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", res.StatusCode)
	}

	var got model.Task
	err := json.NewDecoder(res.Body).Decode(&got)
	if err != nil {
		t.Fatalf("error decoding response: %v", err)
	}

	if got.ID != expectedTask.ID {
		t.Errorf("expected ID %s, got %s", expectedTask.ID, got.ID)
	}
}

func TestGetTaskHandler(t *testing.T) {
	expectedTask := &model.Task{
		ID:        "test-id-456",
		Status:    model.StatusCompleted,
		Result:    "done!",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	h := &TaskHandler{
		Manager: &mockManager{taskToReturn: expectedTask},
	}

	req := httptest.NewRequest(http.MethodGet, "/tasks?id=test-id-456", nil)
	w := httptest.NewRecorder()
	h.GetTask(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", res.StatusCode)
	}

	var got model.Task
	err := json.NewDecoder(res.Body).Decode(&got)
	if err != nil {
		t.Fatalf("error decoding response: %v", err)
	}

	if got.ID != expectedTask.ID {
		t.Errorf("expected ID %s, got %s", expectedTask.ID, got.ID)
	}
}
