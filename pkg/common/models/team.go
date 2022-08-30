package models

type Equipe struct {
	Id_equipe   int    `gorm:"primary_key" json:"id_equipe"`
	Nome_equipe string `json:"nome_equipe"`
}

type EquipeSemId struct {
	Nome_equipe string `json:"nome_equipe"`
}