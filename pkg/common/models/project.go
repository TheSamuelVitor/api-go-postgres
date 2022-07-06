package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name_project string `json:"name_project"`
	Team         string `json:"id_team"`
}
