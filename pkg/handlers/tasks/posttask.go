package tasks

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddTaskBodyResquest struct {
	Name_task  string `json:"name_task"`
	Id_project int    `json:"id_project"`
}

func (h handler) PostTask(c *gin.Context) {
	body := AddTaskBodyResquest{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var newTask models.Tarefa

	newTask.Name_task = body.Name_task
	newTask.Id_projeto = body.Id_project

	if result := h.DB.Create(&newTask); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &newTask)
}
