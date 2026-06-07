package handler

import (
	"encoding/json"
	"net/http"
	"taskfy/internal/usecase"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}

}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var requestBody CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	createdUser, err := h.userUseCase.CreateUser(requestBody.Email, requestBody.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	responseBody := CreateUserResponse{
		ID:    createdUser.Id,
		Email: createdUser.Email,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseBody)

}
