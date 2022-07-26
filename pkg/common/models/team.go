package models

type Equipe struct {
	Id_equipe int `gorm:"primary_key" json:"id_equipe"`
	Name_team string `json:"name_team"`
}