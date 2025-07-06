package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `json:"category" gorm:"foreignKey:CategoryID"`
	Tags        []Tag    `json:"tags" gorm:"many2many:task_tags"`
	UserID      uint
}

type Tag struct {
	gorm.Model
	Title string `json:"title" binding:"required"`
	Tasks []Task `json:"tasks" gorm:"many2many:task_tags"`
}

type Category struct {
	gorm.Model
	Title string `json:"title" gorm:"unique"`
	Tasks []Task
}
