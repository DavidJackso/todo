package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Tasks    []Task
}
