package tasks

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type UpdateTaskRequestBody struct {
	Name_task string `json:"name_task"`
	Project    string `json:"id_project"`
}

func (h handler) PutTask(c *gin.Context) {
	id := c.Param("id")
	body := UpdateTaskRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var task models.Task

	if result := h.DB.First(&task, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	task.Name_task = body.Name_task
	task.Project = body.Project

	h.DB.Save(&task)

	c.JSON(http.StatusOK, &task)
}