package tasks

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/models"
	"github.com/gin-gonic/gin"
)

type UpdateTaskRequestBody struct {
	Nome_tarefa string `json:"nome_tarefa"`
	Descricao   string `json:"descricao"`
	Id_projeto  int    `json:"id_projeto"`
	Id_membro   int    `json:"id_membro"`
}

func (h handler) PutTask(c *gin.Context) {
	id := c.Param("id")
	body := UpdateTaskRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var task models.Tarefa

	if result := h.DB.First(&task, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	task.Nome_tarefa = body.Nome_tarefa
	task.Descricao = body.Descricao
	task.Id_projeto = body.Id_projeto
	task.Id_membro = body.Id_membro

	h.DB.Save(&task)

	c.JSON(http.StatusOK, &task)
}
