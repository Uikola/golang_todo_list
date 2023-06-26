package repository

import (
	"gorm.io/gorm"
	"todolist/internal/model"
)

type TaskPostgres struct {
	db *gorm.DB
}

func NewTaskPostgres(db *gorm.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}

func (r *TaskPostgres) GetAll(userID uint) []model.Task {
	var tasks []model.Task
	r.db.Where("user_id = ?", userID).Find(&tasks)
	return tasks
}

func (r *TaskPostgres) GetByID(userID, taskID uint) model.Task {
	var task model.Task
	r.db.Where("user_id = ? AND id = ?", userID, taskID).First(&task)
	return task
}

func (r *TaskPostgres) Create(userID uint, task model.Task) uint {
	task.UserID = userID
	r.db.Create(&task)
	return task.ID
}

func (r *TaskPostgres) Delete(userID, taskID uint) {
	var task model.Task
	r.db.Where("user_id = ? AND id = ?", userID, taskID).Delete(&task)
	return
}

func (r *TaskPostgres) Update(userID, taskID uint, input model.UpdateTaskInput) {
	r.db.Model(&model.Task{}).Where("user_id = ? AND id = ?", userID, taskID).Updates(map[string]interface{}{"title": input.Title, "description": input.Description})
	return
}
