package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name_project string `json:"name_project"`
	Id_Team      int    `json:"id_team"`
}
