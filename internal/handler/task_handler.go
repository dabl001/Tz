package handler

import (
	"encoding/json"
	"net/http"
	"time"
	"workmate-go/internal/task"
)

type TaskHandler struct {
	Manager *task.Manager
}

type CreateTaskRequest struct {
	Input string `json:"input"`
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	newTask := task.NewSampleTask(req.Input, 5*time.Second)

	t, err := h.Manager.Create(r.Context(), newTask)
	if err != nil {
		http.Error(w, "failed to create task", http.StatusInternalServerError)
		return
	}

	writeJSON(w, t)
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	t, err := h.Manager.Get(r.Context(), id)
	if err != nil {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	writeJSON(w, t)
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
