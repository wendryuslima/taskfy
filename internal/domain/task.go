package domain

import (
	"time"

	"github.com/google/uuid"
)

type TaskRepository interface {
	GetAll(userID string) ([]*Task, error)
}

type Task struct {
	Id     string
	Title  string
	UserId string

	IsCompleted bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTask(title string, userId string, isCompleted bool) *Task {
	now := time.Now()
	return &Task{
		Id:          uuid.New().String(),
		Title:       title,
		UserId:      userId,
		IsCompleted: isCompleted,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (t *Task) UpdateTask(newTitle string) {
	t.Title = newTitle
	t.UpdatedAt = time.Now()

}

func (t *Task) MarkAsCompleted() {
	t.IsCompleted = !t.IsCompleted
	t.UpdatedAt = time.Now()
}
