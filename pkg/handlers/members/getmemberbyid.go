package members

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetMemberbyId(c *gin.Context) {
	id := c.Param("id")

	var member models.Member

	if result := h.DB.First(&member, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "member not found",
		})
		return
	}
	c.JSON(http.StatusOK, &member)
}