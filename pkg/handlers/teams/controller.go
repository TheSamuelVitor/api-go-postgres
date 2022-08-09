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

	routes := r.Group("/equipes")
	routes.GET("/", h.GetTeams)
	routes.POST("/", h.PostTeam)
	routes.PUT("/:id", h.PutTeam)
	routes.GET("/:id", h.GetTeambyId)
	routes.DELETE("/:id", h.DeleteTeams)
	routes.GET("/members/:id", h.GetMemberbyIdTeam)
}