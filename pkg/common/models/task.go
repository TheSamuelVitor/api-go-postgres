package models

type Tarefa struct {
	Id_tarefa   uint    `gorm:"primary_key"` `json:"id_tarefa"`
	Name_task   string `json:"name_task"`
	Description string `json:"description"`
	Id_projeto  int    `json:"id_project"`
	Id_membro   int    `json:"id_membro"`
}
