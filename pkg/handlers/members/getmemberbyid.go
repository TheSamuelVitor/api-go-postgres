package members

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetMemberbyId(c *gin.Context) {
	id := c.Param("id")
	
	var member models.MembroCompleto
	sql := "select m.id_membro, m.nome_membro, m.funcao, m.id_equipe, e.nome_equipe from membros m join equipes e on e.id_equipe = m.id_equipe where id_membro = ?;"
	
	if result := h.DB.Raw(sql, id).Scan(&member); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "member not found",
		})
		return
	}
	c.JSON(http.StatusOK, &member)
}