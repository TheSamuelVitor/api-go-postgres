package members

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/models"
	"github.com/gin-gonic/gin"
)

type AddMemberRequestBody struct {
	Nome_membro string `json:"nome_membro"`
	Funcao      string `json:"funcao"`
	Id_equipe   int    `json:"id_equipe"`
}

func (h handler) Postmembers(c *gin.Context) {
	body := AddMemberRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var newMember models.Membro

	newMember.Nome_membro = body.Nome_membro
	newMember.Id_equipe = body.Id_equipe
	newMember.Funcao = body.Funcao

	if result := h.DB.Create(&newMember); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &newMember)
}
