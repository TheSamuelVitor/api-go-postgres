package projects

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetTasksbyProjectId(c *gin.Context)  {
	id := c.Param("id")

	var tasks []models.Tarefa

	if result := h.DB.Find(&tasks, "id_projeto", id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &tasks)
}