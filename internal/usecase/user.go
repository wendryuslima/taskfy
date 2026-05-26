package usecase

import (
	"taskfy/internal/domain"
	"taskfy/internal/pkg/errors"
)

type UserUseCase struct {
	userRepository domain.UserRepository
}

func NewUserUseCase(paramUserRepository domain.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: paramUserRepository,
	}
}

func (uc *UserUseCase) CreateUser(email, password string) (*domain.User, error) {
	if email == "" || password == "" {
		return nil, errors.ErrInvalidEmailOrPassword
	}

	user := domain.NewUser(email, password)
	err := uc.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	error := uc.userRepository.CreateUser(user)
	if error != nil {
		return nil, errors.ErrUserCreationFailed
	}

	return user, nil

}

func (uc *UserUseCase) LoginUser(email, password string) (*domain.User, error) {
	if email == "" || password == "" {
		return nil, errors.ErrEmailAndPasswordRequired
	}

	user, err := uc.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, errors.ErrUserNotFound
	}

	if user.Password != password {
		return nil, errors.ErrInvalidEmailOrPassword
	}

	return user, nil
}
