package teams

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddTeamRequestBody struct {
	Name_team string `json:"name_team"`
}

func (h handler) PostTeam(c *gin.Context) {
	body := AddTeamRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var team models.Team

	team.Name_team = body.Name_team

	if result := h.DB.Create(&team); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &team)
}
