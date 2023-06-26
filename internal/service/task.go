package service

import (
	"todolist/internal/model"
	"todolist/internal/repository"
)

type TaskService struct {
	repo repository.Task
}

func NewTaskService(repo repository.Task) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetAll(userID uint) []model.Task {
	return s.repo.GetAll(userID)
}

func (s *TaskService) GetByID(userID, taskID uint) model.Task {
	return s.repo.GetByID(userID, taskID)
}

func (s *TaskService) Create(userID uint, task model.Task) uint {
	return s.repo.Create(userID, task)
}

func (s *TaskService) Delete(userID, taskID uint) {
	s.repo.Delete(userID, taskID)
	return
}

func (s *TaskService) Update(userID, taskID uint, input model.UpdateTaskInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	s.repo.Update(userID, taskID, input)
	return nil
}
