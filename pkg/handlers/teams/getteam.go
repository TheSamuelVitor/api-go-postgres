package teams

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetTeams(c *gin.Context) {
	var teams []models.Equipe
	sql := "select * from equipes"

	if result := h.DB.Raw(sql).Scan(&teams); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &teams)
}
