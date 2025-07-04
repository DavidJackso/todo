package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `json:"category" gorm:"foreignKey:CategoryID"`
	Tags        []Tag    `gorm:"many2many:user_tags"`
	UserID      uint
}

type Tag struct {
	gorm.Model
	Title string `json:"title"`
	Tasks []Task `gorm:"many2many:user_tags" `
}

type Category struct {
	gorm.Model
	Tasks []Task `gorm:"foreignKey:CategoryID"`
	Title string `json:"title"`
}
