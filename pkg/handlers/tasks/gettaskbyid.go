package tasks

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetTaskbyId(c *gin.Context) {
	id := c.Param("id")

	var task models.TarefacomProjetoeMembro
	sql := "select tarefas.nome_tarefa, tarefas.descricao, membros.nome_membro, projetos.nome_projeto from tarefas join membros on membros.id_membro = tarefas.id_membro join projetos on projetos.id_projeto = tarefas.id_projeto where id_tarefa = ?"

	if result := h.DB.Raw(sql, id).Scan(&task); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "task not found",
		})
		return
	}

	c.JSON(http.StatusOK, &task)
}
