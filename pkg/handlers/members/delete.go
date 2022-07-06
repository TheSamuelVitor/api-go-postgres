package members

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) DeleteMembersbyId(c *gin.Context) {
	id := c.Param("id")

	var member models.Member

	if result := h.DB.First(&member, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&member)
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "member deleted sucessfully",
	})
}