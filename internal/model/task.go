package model

import (
	"errors"
	"time"
)

type Task struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
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
