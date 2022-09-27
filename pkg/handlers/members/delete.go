package members

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h handler) DeleteMembersbyId(c *gin.Context) {
	id := c.Param("id")

	var member models.Membro
	sql := "delete from membros where id_membro = ?"

	result := h.DB.Raw(sql, id).Scan(&member)
	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "member deleted sucessfully",
	})
}
