package model

import (
	"errors"
	"time"
)

type Task struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UserID      uint      `json:"user_id"`
}

type UpdateTaskInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateTaskInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update struct has no value")
	}

	return nil
}
