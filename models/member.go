package models

import "time"

type Membro struct {
	Id_membro   int       `gorm:"primary_key" json:"id_membro" example:"1"`
	Nome_membro string    `json:"nome_membro" example:"Samuel"`
	Funcao      string    `json:"funcao" example:"Desenvolvedor Backend"`
	Id_equipe   int       `json:"id_equipe" example:"7"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type MembrocomEquipe struct {
	Id_membro   int    `gorm:"primary_key" json:"id_membro" example:"1"`
	Nome_membro string `json:"nome_membro" example:"Samuel"`
	Funcao      string `json:"funcao" example:"Desenvolvedor Backend"`
	Id_equipe   int    `json:"id_equipe" example:"7"`
	Nome_equipe string `json:"nome_equipe" example:"Komanda"`
}
