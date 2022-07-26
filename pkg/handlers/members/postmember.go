package members

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddMemberRequestBody struct {
	Name_member string `json:"name_member"`
	Function    string `json:"function"`
	Id_Team     int    `json:"id_team"`
}

func (h handler) Postmembers(c *gin.Context) {
	body := AddMemberRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var newMember models.Membro

	newMember.Name_member = body.Name_member
	newMember.Id_equipe = body.Id_Team
	newMember.Function = body.Function

	if result := h.DB.Create(&newMember); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &newMember)
}
