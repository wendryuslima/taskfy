package usecase

import (
	"taskfy/internal/domain"
)

type TaskUseCase struct {
	taskRepository domain.TaskRepository
}

func NewTaskUseCase(taskRepository domain.TaskRepository) *TaskUseCase {

	return &TaskUseCase{
		taskRepository: taskRepository,
	}

}

func (uc *TaskUseCase) ListTasks(userID string) ([]*domain.Task, error) {
	return uc.taskRepository.GetAll(userID)
}
