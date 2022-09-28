package tasks

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/models"
	"github.com/gin-gonic/gin"
)

func (h handler) DeleteTaskbyId(c *gin.Context) {
	id := c.Param("id")

	var task models.Tarefa

	if result := h.DB.First(&task, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&task)

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "deleted task sucessfully",
	})
}
