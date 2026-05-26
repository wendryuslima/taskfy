package domain

import (
	"time"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
}

type User struct {
	Id        string
	Email     string
	Password  string
	CreatedAt time.Time
}

func (u *User) Validate() any {
	panic("unimplemented")
}

func NewUser(email, password string) *User {
	return &User{
		Id:        uuid.New().String(),
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	}
}
