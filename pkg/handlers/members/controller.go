package members

import (
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

	routes := r.Group("/members")
	routes.GET("/", h.GetMembers)
	routes.GET("/:id", h.GetMemberbyId)
	routes.POST("/", h.Postmembers)
	routes.PUT("/:id", h.PutMembers)
	routes.DELETE("/:id", h.DeleteMembersbyId)
}