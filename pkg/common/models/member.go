package models

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	Name_member string `json:"name_member"`
	Id_Team     int `json:"id_team"`
	Id_Task     int `json:"id_task"`
}
