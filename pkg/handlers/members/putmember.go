package members

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type UpdateMemberRequestBody struct {
	Name_member string `json:"name_member"`
	Function    string `json:"function"`
	Id_Team     int    `json:"id_team"`
	Id_Task     int    `json:"id_task"`
}

func (h handler) PutMembers(c *gin.Context) {
	id := c.Param("id")
	body := UpdateMemberRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var member models.Member

	if result := h.DB.First(&member, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	member.Name_member = body.Name_member
	member.Id_Team = body.Id_Team
	member.Id_Task = body.Id_Task
	member.Function = body.Function

	h.DB.Save(&member)

	c.JSON(http.StatusOK, &member)
}
