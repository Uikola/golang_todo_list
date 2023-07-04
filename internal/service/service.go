package service

import (
	"todolist/internal/model"
	"todolist/internal/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user model.User) (uint, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (uint, error)
}

type Task interface {
	GetAll(userID uint) []model.Task
	GetByID(userID, taskID uint) model.Task
	Create(userID uint, task model.Task) uint
	Delete(userID, taskID uint)
	Update(userID, taskID uint, input model.UpdateTaskInput) error
}

type Service struct {
	Authorization
	Task
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Task:          NewTaskService(repos.Task),
	}
}
