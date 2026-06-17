package middleware

import (
	"context"
	"net/http"
	"strings"
	"taskfy/internal/domain"
	"taskfy/internal/helpers"
	"taskfy/internal/pkg/errors"
)

const (
	UserIDContextKey = "userID"
)

type AuthMiddleware struct {
	userRepository domain.UserRepository
}

func NewAuthMiddleware(userRepository domain.UserRepository) *AuthMiddleware {
	return &AuthMiddleware{
		userRepository: userRepository,
	}
}

func (m *AuthMiddleware) VerifyAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			helpers.ResponseError(w, http.StatusUnauthorized, errors.ErrInvalidAuthorizationHeader.Error())
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			helpers.ResponseError(w, http.StatusUnauthorized, errors.ErrInvalidAuthorizationHeader.Error())
			return
		}

		authenticatedUser, err := m.userRepository.GetUserByID(token)
		if err != nil {
			helpers.ResponseError(w, http.StatusUnauthorized, errors.ErrInvalidToken.Error())
			return

		}
		ctx := context.WithValue(r.Context(), UserIDContextKey, authenticatedUser.Id)
		next(w, r.WithContext(ctx))

	}

}

func GetUserIDFromContext(ctx context.Context) (string, error) {

	userID, ok := ctx.Value(UserIDContextKey).(string)
	if !ok {
		return "", errors.ErrUserCreationFailed
	}
	return userID, nil

}
