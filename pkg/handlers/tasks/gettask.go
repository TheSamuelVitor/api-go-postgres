package tasks

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetTask(c *gin.Context) {
	var tarefas []models.Tarefa
	sql := "select * from tarefas"

	if result := h.DB.Raw(sql).Scan(&tarefas); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &tarefas)
}
