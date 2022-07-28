package tasks

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetTask(c *gin.Context) {
	var tarefas []models.Tarefa
	sql := "select tarefas.id_tarefa, tarefas.name_task, tarefas.description, projetos.id_projeto, projetos.name_project, membros.id_membro, membros.name_member from tarefas join projetos on projetos.id_projeto = tarefas.id_projeto join membros on membros.id_membro = tarefas.id_membro"

	if result := h.DB.Raw(sql).Scan(&tarefas); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &tarefas)
}
