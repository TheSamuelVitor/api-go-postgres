package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name_task string `json:"name_task"`
	Project   string `json:"id_project"`
}
