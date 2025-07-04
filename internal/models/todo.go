package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description"`
	Category    Category `json:"category"`
	Tags        []Tag    `gorm:"many2many:user_tags"`
	UserID      uint
}

type Tag struct {
	gorm.Model
	Title string
	Tasks []Task `gorm:"many2many:user_tags"`
}

type Category struct {
	gorm.Model
	TaskID uint
	Title  string
}
