package teams

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetMemberbyIdTeam(c* gin.Context) {
	id := c.Param("id")

	var members []models.Membro

	sql := "select * from membros where id_equipe = ?"

	if result := h.DB.Raw(sql, id).Scan(&members); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &members)
}