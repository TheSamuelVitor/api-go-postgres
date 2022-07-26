package models

type Membro struct {
	Id_membro   uint    `gorm:"primary_key" json:"id_membro"`
	Name_member string `json:"name_member"`
	Function    string `json:"function"`
	Id_equipe   uint    `json:"id_equipe"`
}
