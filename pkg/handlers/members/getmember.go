package members

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetMembers(c *gin.Context) {
	var membro []models.Membro

	sql := "select membros.id_membro, membros.nome_membro, membros.funcao, membros.id_equipe, equipes.nome_equipe from membros join equipes on equipes.id_equipe = membros.id_equipe order by id_membro;"

	if result := h.DB.Raw(sql).Scan(&membro); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &membro)
}