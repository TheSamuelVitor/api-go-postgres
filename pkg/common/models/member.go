package models

import "time"

type Membro struct {
	Id_membro   int       `gorm:"primary_key" json:"id_membro"`
	Nome_membro string    `json:"nome_membro"`
	Funcao      string    `json:"funcao"`
	Id_equipe   int       `json:"id_equipe"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type MembrocomEquipe struct {
	Id_membro   int    `gorm:"primary_key" json:"id_membro"`
	Nome_membro string `json:"nome_membro"`
	Funcao      string `json:"funcao"`
	Id_equipe   int    `json:"id_equipe"`
	Nome_equipe string `json:"nome_equipe"`
}
