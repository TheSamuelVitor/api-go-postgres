package models

import "time"

type Equipe struct {
	Id_equipe   int       `gorm:"primary_key" json:"id_equipe"`
	Nome_equipe string    `json:"nome_equipe"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
