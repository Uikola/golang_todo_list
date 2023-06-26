package repository

import (
	"gorm.io/gorm"
	"todolist/internal/model"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (uint, error) {
	r.db.Create(&user)
	return user.ID, nil
}

func (r *AuthPostgres) GetUser(username, password string) model.User {
	var user model.User
	r.db.Where("username = ? AND password = ?", username, password).First(&user)
	return user
}
