package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (t *Task) Create(db *gorm.DB) error {
	return db.Create(t).Error
}
