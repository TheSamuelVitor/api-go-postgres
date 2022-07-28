package members

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetMembers(c *gin.Context) {
	var membro []models.Membro

	sql := "select membros.id_membro, membros.name_member, membros.function, equipes.id_equipe, equipes.name_team from membros join equipes on equipes.id_equipe = membros.id_equipe"

	if result := h.DB.Raw(sql).Scan(&membro); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &membro)
}