package teams

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB){
	h := &handler{
		DB: db,
	}

	routes := r.Group("/teams")
	routes.GET("/", h.GetTeams)
	routes.GET("/:id", h.GetTeambyId)
	routes.GET("/members/:id", h.GetMemberbyIdTeam)
	routes.POST("/", h.PostTeam)
	routes.PUT("/:id", h.PutTeam)
	routes.DELETE("/:id", h.DeleteTeams)
}