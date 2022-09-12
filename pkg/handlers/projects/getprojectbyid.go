package projects

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type Info struct {
	Id_projeto   int             `gorm:"primary_key" json:"id_projeto"`
	Nome_projeto string          `json:"nome_projeto"`
	Descricao    string          `json:"descricao"`
	Id_equipe    int             `json:"id_equipe"`
	Nome_equipe  string          `json:"nome_equipe"`
	Tarefas      []models.Tarefa `json:"tarefas"`
}

func (h handler) GetProjectbyId(c *gin.Context) {
	id := c.Param("id")

	var info Info
	var projeto models.ProjetocomEquipe
	var tarefas []models.Tarefa

	sql := "select projetos.id_projeto, projetos.nome_projeto, projetos.descricao, projetos.id_equipe, equipes.nome_equipe from projetos join equipes on projetos.id_equipe = equipes.id_equipe where id_projeto = ?"

	if result := h.DB.Raw(sql, id).Scan(&projeto); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	sql2 := "select * from tarefas where id_projeto = ?"
	if result := h.DB.Raw(sql2, id).Scan(&tarefas); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	info.Id_projeto = projeto.Id_projeto
	info.Nome_projeto = projeto.Nome_projeto
	info.Descricao = projeto.Descricao
	info.Id_equipe = projeto.Id_equipe
	info.Nome_equipe = projeto.Nome_equipe
	info.Tarefas = tarefas

	c.JSON(http.StatusOK, &info)
}
