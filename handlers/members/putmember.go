package members

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/models"
	"github.com/gin-gonic/gin"
)

type UpdateMemberRequestBody struct {
	Nome_membro string `json:"nome_membro"`
	Funcao      string `json:"funcao"`
	Id_equipe   int    `json:"id_equipe"`
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

	member.Nome_membro = body.Nome_membro
	member.Id_equipe = body.Id_equipe
	member.Funcao = body.Funcao

	h.DB.Save(&member)

	c.JSON(http.StatusOK, &member)
}
