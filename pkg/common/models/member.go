package models

type Membro struct {
	Id_membro   int    `json:"id_membro"`
	Name_member string `json:"name_member"`
	Function    string `json:"function"`
	Id_equipe   int    `json:"id_equipe"`
}
