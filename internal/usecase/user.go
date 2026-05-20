package usecase

import (
	"taskfy/internal/domain"
	"taskfy/internal/pkg/errors"
)

type UserUseCase struct {
}

func (uc *UserUseCase) CreateUser(email, password string) (*domain.User, error) {
	if email == "" || password == "" {
		return nil, errors.ErrInvalidEmailOrPassword
	}

	user := domain.NewUser(email, password)

	return user, nil

}

func (lc *UserUseCase) LoginUser(email, password string) (*domain.User, error) {
	if email == "" || password == "" {
		return nil, errors.ErrEmailAndPasswordRequired
	}

	user := domain.NewUser(email, password)

	if user.Email != email || user.Password != password {
		return nil, errors.ErrInvalidEmailOrPassword
	}
	return user, nil

}
