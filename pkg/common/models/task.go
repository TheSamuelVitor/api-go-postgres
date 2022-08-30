package models

type Tarefa struct {
	Id_tarefa   int    `gorm:"primary_key" json:"id_tarefa"`
	Nome_tarefa string `json:"nome_tarefa"`
	Descricao   string `json:"descricao"`
	Id_projeto  int    `json:"id_projeto"`
	Id_membro   int    `json:"id_membro"`
}

type TarefacomProjetoeMembro struct {
	Nome_tarefa string `json:"nome_tarefa"`
	Descricao   string `json:"descricao"`
	Nome_projeto string `json:"nome_projeto"`
	Nome_membro string `json:"nome_membro"`
}
	