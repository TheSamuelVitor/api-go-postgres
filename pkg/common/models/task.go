package models

import "time"

type Tarefa struct {
	Id_tarefa       int       `gorm:"primary_key" json:"id_tarefa"`
	Name_task       string    `json:"name_task"`
	Description     string    `json:"description"`
	Id_projeto      int       `json:"id_project"`
	Id_membro       int       `json:"id_membro"`
	Data_de_criacao time.Time `json:"Data_de_criacao"`
}
