package projects

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetProjectbyId(c *gin.Context) {
	id := c.Param("id")

	var project models.Projeto
	sql := "select * from projetos where id_projeto = ?"

	if result := h.DB.Raw(sql, id).Scan(&project); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "project not found",
		})
		return
	}

	c.JSON(http.StatusOK, &project)
}
