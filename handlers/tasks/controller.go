package tasks

import (
	"github.com/TheSamuelVitor/api-go-postgres/pkg/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}


	routes := r.Group("/tarefas", middlewares.Auth())
	routes.GET("/", h.GetTask)
	routes.POST("/", h.PostTask)
	routes.PUT("/:id", h.PutTask)
	routes.GET("/:id", h.GetTaskbyId)
	routes.DELETE("/:id", h.DeleteTaskbyId)
}