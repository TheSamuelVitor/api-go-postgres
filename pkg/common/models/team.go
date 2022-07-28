package models

import "time"

type Equipe struct {
	Id_equipe       int    `gorm:"primary_key" json:"id_equipe"`
	Name_team       string `json:"name_team"`
	Data_de_criacao time.Time `json:"data_de_criacao"`
}
