package teams

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) DeleteTeams(c *gin.Context) {
	id := c.Param("id")

	var team models.Team

	if result := h.DB.First(&team, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&team)
	c.Status(http.StatusOK)
}
