package members

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
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
	Id_membro   int             `gorm:"primary_key" json:"id_membro"`
	Nome_membro string          `json:"nome_membro"`
	Funcao      string          `json:"funcao"`
	Id_equipe   int             `json:"id_equipe"`
	Nome_equipe string          `json:"nome_equipe"`
	Tarefas     []TarefacomProjeto `json:"tarefas"`
}

func (h handler) GetMemberbyId(c *gin.Context) {
	id := c.Param("id")

	var info Info
	var member models.MembrocomEquipe
	var tarefas []TarefacomProjeto

	if result := h.DB.Raw("select m.id_membro, m.nome_membro, m.funcao, m.id_equipe, e.nome_equipe from membros m join equipes e on e.id_equipe = m.id_equipe where id_membro = ?", id).Scan(&member); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if result := h.DB.Raw("select t.id_tarefa, t.nome_tarefa, t.descricao, t.id_projeto, p.nome_projeto from tarefas t inner join projetos p on p.id_projeto = t.id_projeto inner join membros m on m.id_membro = t.id_membro where id_tarefa = ?", id).Scan(&tarefas); result.Error != nil {
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
