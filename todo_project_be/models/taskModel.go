package models

import (
	"todo_project_be/config"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Priority    string `json:"priority"`
	Status      string `json:"status"`
}

func CreateTask(task *Task) error {
	err := config.DB.Create(task).Error
	if err != nil {
		return err
	}

	return nil
}
