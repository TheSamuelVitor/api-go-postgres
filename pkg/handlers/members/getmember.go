package members

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetMembers(c *gin.Context) {
	var membro []models.MembroCompleto

	sql := "select m.id_membro, m.nome_membro, m.funcao, m.id_equipe, e.nome_equipe from membros m join equipes e on e.id_equipe = m.id_equipe order by id_membro;" 

	if result := h.DB.Raw(sql).Scan(&membro); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &membro)
}