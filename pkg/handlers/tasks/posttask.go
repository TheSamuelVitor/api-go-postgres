package tasks

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddTaskBodyResquest struct {
	Nome_tarefa string `json:"nome_tarefa"`
	Descricao   string `json:"descricao"`
	Id_projeto  int    `json:"id_projeto"`
	Id_membro   int    `json:"id_membro"`
}

func (h handler) PostTask(c *gin.Context) {
	body := AddTaskBodyResquest{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var newTask models.Tarefa

	newTask.Nome_tarefa = body.Nome_tarefa
	newTask.Descricao = body.Descricao
	newTask.Id_projeto = body.Id_projeto
	newTask.Id_membro = body.Id_membro

	if result := h.DB.Create(&newTask); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &newTask)
}
