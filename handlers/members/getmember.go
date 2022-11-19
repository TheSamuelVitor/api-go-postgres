package members

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/models"
	"github.com/gin-gonic/gin"
)

// GetVideos godoc
// @Security    bearerAuth
// @Summary     List existing members
// @Description Get all the existing members
// @Tags        members
// @Accept      json
// @Produce     json
// @Success     200 {array}  models.MembrocomEquipe
// @Router      /membros [get]
func (h handler) GetMembers(c *gin.Context) {
	var membro []models.MembrocomEquipe

	sql := "select membros.id_membro, membros.nome_membro, membros.funcao, membros.id_equipe, equipes.nome_equipe from membros join equipes on equipes.id_equipe = membros.id_equipe order by id_membro;"

	result := h.DB.Raw(sql).Scan(&membro)
	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &membro)
}
