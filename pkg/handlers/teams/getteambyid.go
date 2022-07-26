package teams

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetTeambyId(c *gin.Context) {
	id := c.Param("id")

	var team models.Equipe
	sql := "select * from equipes where id_equipe = ?"

	if result := h.DB.Raw(sql, id).Scan(&team); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "team not found",
		})
		return
	}

	c.JSON(http.StatusOK, &team)
}
