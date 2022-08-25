package user

import (
	//"github.com/TheSamuelVitor/api-go-postgres/pkg/common/db"
	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/TheSamuelVitor/api-go-postgres/pkg/services"
	"github.com/gin-gonic/gin"
)

func (h handler) CreateUser(c *gin.Context) {
	//db := db.GetDatabase()

	var p models.User

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			" error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	p.Password = services.Sha256Encoder(p.Password)

	err = h.DB.Create(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create user: " + err.Error(),
		})
		return
	}

	c.Status(204)
}