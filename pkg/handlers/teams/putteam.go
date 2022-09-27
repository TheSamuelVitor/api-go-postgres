package teams

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/models"
	"github.com/gin-gonic/gin"
)

type UpdateTeamRequestBody struct {
	Nome_equipe string `json:"nome_equipe"`
}

func (h handler) PutTeam(c *gin.Context) {
	id := c.Param("id")
	body := UpdateTeamRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var team models.Equipe

	if result := h.DB.First(&team, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	team.Nome_equipe = body.Nome_equipe

	h.DB.Save(&team)

	c.JSON(http.StatusOK, &team)

}
