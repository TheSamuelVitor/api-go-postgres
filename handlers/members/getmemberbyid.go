package members

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/models"
	"github.com/gin-gonic/gin"
)

type TarefacomProjeto struct {
	Id_tarefa    int    `gorm:"primary_key" json:"id_tarefa"`
	Nome_tarefa  string `json:"nome_tarefa"`
	Descricao    string `json:"descricao"`
	Id_projeto   int    `json:"id_projeto"`
	Nome_projeto string `json:"nome_projeto"`
}

type Info struct {
	Id_membro   int                `gorm:"primary_key" json:"id_membro"`
	Nome_membro string             `json:"nome_membro"`
	Funcao      string             `json:"funcao"`
	Id_equipe   int                `json:"id_equipe"`
	Nome_equipe string             `json:"nome_equipe"`
	Tarefas     []TarefacomProjeto `json:"tarefas,omitempty"`
}

// GetVideos godoc
// @Security    bearerToken
// @Summary     Shows member which has the id equals to given
// @Description Get all the existing members
// @Param  id path int  true "ID_membro"
// @Tags        members
// @Accept      json
// @Produce     json
// @Success     200 {object}  models.MembrocomEquipe
// @Router      /membros/{id} [get]
func (h handler) GetMemberbyId(c *gin.Context) {
	id := c.Param("id")

	var info Info
	var member models.MembrocomEquipe
	var tarefas []TarefacomProjeto

	comandoMembros := "select m.id_membro, m.nome_membro, m.funcao, m.id_equipe, e.nome_equipe from membros m join equipes e on e.id_equipe = m.id_equipe where id_membro = ?"
	result := h.DB.Raw(comandoMembros, id).Scan(&member)

	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	comandoTarefas := "select * from tarefas where id_membro = ?"
	result = h.DB.Raw(comandoTarefas, id).Scan(&tarefas)
	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	info.Id_membro = member.Id_membro
	info.Nome_membro = member.Nome_membro
	info.Funcao = member.Funcao
	info.Id_equipe = member.Id_equipe
	info.Nome_equipe = member.Nome_equipe
	info.Tarefas = tarefas

	c.JSON(http.StatusOK, &info)
}
