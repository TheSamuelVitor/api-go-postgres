package models

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	Name_member string `json:"name_member"`
	Id_Team     string `json:"id_team"`
	Id_Task     string `json:"id_task"`
}
