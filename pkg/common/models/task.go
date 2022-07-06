package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name_task  string `json:"name_task"`
	Id_project int    `json:"id_project"`
}
