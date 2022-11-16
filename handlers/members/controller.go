package members

import (
	// "github.com/TheSamuelVitor/api-go-postgres/pkg/middlewares"
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

	routes := r.Group("/membros")
	routes.GET("/", h.GetMembers)
	routes.POST("/", h.Postmembers)
	routes.PUT("/:id", h.PutMembers)
	routes.GET("/:id", h.GetMemberbyId)
	routes.DELETE("/:id", h.DeleteMembersbyId)
}