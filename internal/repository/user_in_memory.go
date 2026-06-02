package repository

import (
	"taskfy/internal/domain"
	customErrors "taskfy/internal/pkg/errors"
)

type UserRepositoryInMemory struct {
	users map[string]*domain.User
}

func NewUserRepositoryInMemory() *UserRepositoryInMemory {

	return &UserRepositoryInMemory{
		users: map[string]*domain.User{},
	}
}

func (r *UserRepositoryInMemory) CreateUser(user *domain.User) error {

	for _, u := range r.users {
		if u.Email == user.Email {
			return customErrors.ErrEmailAlreadyExists
		}
	}

	r.users[user.Id] = user
	return nil

}

func (r *UserRepositoryInMemory) GetUserByEmail(email string) (*domain.User, error) {

	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, customErrors.ErrUserNotFound

}
