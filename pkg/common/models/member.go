package models

import (
	"time"
)

type Membro struct {
	Id_membro       int       `gorm:"primary_key" json:"id_membro"`
	Name_member     string    `json:"name_member"`
	Function        string    `json:"function"`
	Id_equipe       int       `json:"id_equipe"`
	Data_de_criacao time.Time `json:"data_de_criacao"`
}
