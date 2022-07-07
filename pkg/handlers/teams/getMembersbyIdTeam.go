package teams

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetMemberbyIdTeam(c* gin.Context) {
	id := c.Param("id")

	var members []models.Member

	if result := h.DB.Find(&members, "id_team", id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &members)
}