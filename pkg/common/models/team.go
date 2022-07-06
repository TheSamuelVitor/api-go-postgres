package models

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name_team string `json:"name_team"`
}