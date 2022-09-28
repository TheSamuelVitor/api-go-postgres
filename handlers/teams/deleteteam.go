package teams

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/models"
	"github.com/gin-gonic/gin"
)

func (h handler) DeleteTeams(c *gin.Context) {
	id := c.Param("id")

	var team models.Equipe

	if result := h.DB.First(&team, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&team)

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "team deleted sucessfully",
	})
}
