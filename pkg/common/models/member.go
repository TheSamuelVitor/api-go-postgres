package models

type Membro struct {
	Id_membro   int    `gorm:"primary_key" json:"id_membro"`
	Name_member string `json:"name_member"`
	Funcao      string `json:"funcao"`
	Id_equipe   int    `json:"id_equipe"`
}
