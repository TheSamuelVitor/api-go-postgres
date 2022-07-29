package projects

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) CountProjects(c*gin.Context) {
	quantidade := 0
	sql := "select count(*) from projetos"

	if result := h.DB.Raw(sql).Scan(&quantidade); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &quantidade)
}