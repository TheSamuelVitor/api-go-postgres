package teams

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/models"
	"github.com/gin-gonic/gin"
)

type AddTeamRequestBody struct {
	Nome_equipe string `json:"nome_equipe"`
}

func (h handler) PostTeam(c *gin.Context) {
	body := AddTeamRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var team models.Equipe

	team.Nome_equipe = body.Nome_equipe

	if result := h.DB.Create(&team); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &team)
}
