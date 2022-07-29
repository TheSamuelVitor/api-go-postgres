package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h handler) CountTasks(c*gin.Context) {
	quantidade := 0
	sql := "select count(*) from tarefas"

	if result := h.DB.Raw(sql).Scan(&quantidade); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &quantidade)
}