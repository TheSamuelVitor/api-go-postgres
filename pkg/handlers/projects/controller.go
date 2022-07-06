package projects

import (
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

	routes := r.Group("/projects")
	routes.GET("/", h.GetProjects)
	routes.GET("/:id", h.GetProjectbyId)
	routes.POST("/", h.PostProjects)
	routes.PUT("/:id", h.Updateproject)
	routes.DELETE("/:id", h.DeleteProjectbyId)
}