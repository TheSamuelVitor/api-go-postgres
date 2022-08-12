package tasks

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetTask(c *gin.Context) {
	var tarefas []models.Tarefa
	sql := "select tarefas.id_tarefa, tarefas.nome_tarefa, tarefas.descricao, membros.id_membro,  membros.nome_membro,projetos.id_projeto, projetos.nome_projeto from tarefas join membros on membros.id_membro = tarefas.id_membro join projetos on projetos.id_projeto = tarefas.id_projeto order by id_tarefa;"

	if result := h.DB.Raw(sql).Scan(&tarefas); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &tarefas)
}
