package models

import "time"

type Projeto struct {
	Id_projeto      int       `gorm:"primary_key" json:"id_projeto"`
	Name_project    string    `json:"name_project"`
	Description     string    `json:"description"`
	Id_equipe       int       `json:"id_team"`
	Data_de_criacao time.Time `json:"data_de_criacao"`
}
