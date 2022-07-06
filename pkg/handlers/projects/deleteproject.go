package projects

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) DeleteProjectbyId(c *gin.Context) {
	id := c.Param("id")
	
	var project models.Project

	if result := h.DB.First(&project, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&project)
	
	c.Status(http.StatusOK)
}
