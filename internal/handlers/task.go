package handlers

import (
	"encoding/json"
	"net/http"
	"taskfy/internal/helpers"
	"taskfy/internal/middleware"
	"taskfy/internal/usecase"
)

type TaskHandler struct {
	taskUseCase *usecase.TaskUseCase
}

type TaskRequest struct {
	Title  string `json:"title"`
	UserID string `json:"user_id"`
}

type TaskResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	UserID      string `json:"user_id"`
	IsCompleted bool   `json:"isCompleted"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func NewTaskHandler(taskUseCase usecase.TaskUseCase) *TaskHandler {
	return &TaskHandler{
		taskUseCase: &taskUseCase,
	}
}

func (h *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserIDFromContext(r.Context())
	if err != nil {
		helpers.ResponseError(w, http.StatusUnauthorized, err.Error())
		return
	}

	listTasks, err := h.taskUseCase.ListTasks(userID)
	if err != nil {
		helpers.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	response := make([]TaskResponse, len(listTasks))
	for i, task := range listTasks {
		response[i] = TaskResponse{
			ID:     task.Id,
			Title:  task.Title,
			UserID: task.UserId,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
