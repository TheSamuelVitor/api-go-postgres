package projects

import (
	"github.com/TheSamuelVitor/api-go-postgres/pkg/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r* gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	r.Use(middlewares.CORSMiddleware())

	routes := r.Group("/projetos")
	routes.GET("/", h.GetProjects)
	routes.POST("/", h.PostProjects)
	routes.PUT("/:id", h.Updateproject)
	routes.GET("/:id", h.GetProjectbyId)
	routes.DELETE("/:id", h.DeleteProjectbyId)
	routes.GET("/tarefas/:id", h.GetTasksbyProjectId)
}