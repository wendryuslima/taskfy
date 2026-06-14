package handlers

import (
	"encoding/json"
	"net/http"
)

type TaskHandler struct {
	// taskUseCase *TaskUseCase
}

type TaskRequest struct {
	Title  string `json:"title"`
	UserID string `json:"user_id"`
}

type TaskResponse struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	UserID string `json:"user_id"`
}

func NewTaskHandler() *TaskHandler {
	return &TaskHandler{}
}

func (h *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	var requestBody TaskRequest

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "dados inválidos", http.StatusBadRequest)
		return
	}

	listTasks, err := h.taskUseCase.ListTasks(requestBody.Title, requestBody.UserID)
	if err != nil {
		http.Error(w, "erro ao listar tarefas", http.StatusInternalServerError)
		return
	}

	response := make([]TaskResponse, len(listTasks))
	for i, task := range listTasks {
		response[i] = TaskResponse{
			ID:     task.ID,
			Title:  task.Title,
			UserID: task.UserID,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
