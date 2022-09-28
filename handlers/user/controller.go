package user

import (
	"github.com/TheSamuelVitor/api-go-postgres/pkg/middlewares"
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
	
	routes := r.Group("/user", middlewares.Auth())
	routes.POST("/", h.CreateUser)
}
