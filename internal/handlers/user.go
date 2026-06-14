package handlers

import (
	"encoding/json"
	"net/http"
	"taskfy/internal/helpers"
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
		helpers.ResponseError(w, http.StatusBadRequest, "dados inválidos")
		return
	}

	createdUser, err := h.userUseCase.CreateUser(requestBody.Email, requestBody.Password)
	if err != nil {
		helpers.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseBody := CreateUserResponse{
		ID:    createdUser.Id,
		Email: createdUser.Email,
	}
	helpers.ResponseJson(w, http.StatusCreated, responseBody)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var requestBody LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		helpers.ResponseError(w, http.StatusBadRequest, "dados inválidos")
		return
	}

	authenticatedUser, err := h.userUseCase.LoginUser(requestBody.Email, requestBody.Password)
	if err != nil {
		helpers.ResponseError(w, http.StatusUnauthorized, err.Error())
		return
	}
	responseBody := LoginResponse{
		ID:    authenticatedUser.Id,
		Email: authenticatedUser.Email,
		Token: authenticatedUser.Id,
	}
	helpers.ResponseJson(w, http.StatusOK, responseBody)

}
