package model

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Age      int    `json:"age" binding:"required"`
}
