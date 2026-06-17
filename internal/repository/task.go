package repository

import (
	"taskfy/internal/domain"
)

type TaskRepositoryInMemory struct {
	tasks map[string]*domain.Task
}

func NewTaskRepository() *TaskRepositoryInMemory {
	return &TaskRepositoryInMemory{
		tasks: map[string]*domain.Task{},
	}

}

func (r *TaskRepositoryInMemory) GetAll(userID string) ([]*domain.Task, error) {
	var userTasks []*domain.Task

	for _, t := range r.tasks {
		if t.UserId != userID {
			continue
		}
		userTasks = append(userTasks, t)
	}

	return userTasks, nil

}
