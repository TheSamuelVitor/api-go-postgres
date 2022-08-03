package members

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type UpdateMemberRequestBody struct {
	Name_member string `json:"name_member"`
	Funcao      string `json:"funcao"`
	Id_Team     int    `json:"id_team"`
}

func (h handler) PutMembers(c *gin.Context) {
	id := c.Param("id")
	body := UpdateMemberRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var member models.Membro
	sql := "select * from membros where id_membro = ?"

	if result := h.DB.Raw(sql, id).Scan(&member); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	member.Name_member = body.Name_member
	member.Id_equipe = body.Id_Team
	member.Funcao = body.Funcao

	h.DB.Save(&member)

	c.JSON(http.StatusOK, &member)
}
