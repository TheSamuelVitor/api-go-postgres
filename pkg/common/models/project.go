package models

type Projeto struct {
	Id_projeto   int    `gorm:"primary_key" json:"id_projeto"`
	Nome_projeto string `json:"nome_projeto"`
	Descricao    string `json:"descricao"`
	Id_equipe    int    `json:"id_equipe"`
	Nome_equipe  string `json:"nome_equipe"`
}
