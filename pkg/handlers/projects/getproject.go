package projects

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetProjects(c *gin.Context) {
	var projects []models.Projeto
	sql := "select * from projetos"

	if result := h.DB.Raw(sql).Scan(&projects); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &projects)
}
