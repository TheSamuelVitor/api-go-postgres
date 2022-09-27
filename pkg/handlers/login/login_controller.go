package login

import (
	"github.com/TheSamuelVitor/api-go-postgres/pkg/models"
	"github.com/TheSamuelVitor/api-go-postgres/pkg/services"
	"github.com/gin-gonic/gin"
)

func (h handler) Login(c *gin.Context) {

	var p models.Login

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	var user models.User
	dbError := h.DB.Where("email = ?", p.Email).First(&user).Error
	if dbError != nil {
		c.JSON(400, gin.H{
			"Erro de autenticação": "usuario nao encontrado",
		})
		return
	}

	if user.Password != services.Sha256Encoder(p.Password) {
		c.JSON(400, gin.H{
			"Erro de autenticação": "senha digitada incorretamente",
		})
		return
	}

	token, err := services.NewJWTService().GenerateToken(user.Id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
