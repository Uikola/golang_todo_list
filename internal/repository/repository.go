package repository

import (
	"gorm.io/gorm"
	"todolist/internal/model"
)

type Authorization interface {
	CreateUser(user model.User) (uint, error)
	GetUser(username, password string) model.User
}

type Task interface {
	Create(userID uint, task model.Task) uint
	GetAll(userID uint) []model.Task
	GetByID(userID, taskID uint) model.Task
	Update(userID, taskID uint, input model.UpdateTaskInput)
	Delete(userID, taskID uint)
}

type Repository struct {
	Authorization
	Task
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Task:          NewTaskPostgres(db),
	}
}
