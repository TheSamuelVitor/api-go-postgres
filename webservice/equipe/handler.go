package equipe

import (
	"github.com/TheSamuelVitor/api-go-postgres/domain/equipe"
	modelApresentacao "github.com/TheSamuelVitor/api-go-postgres/domain/equipe/model"
	"github.com/gin-gonic/gin"
)

func novaEquipe(c *gin.Context) {

	req := modelApresentacao.ReqEquipe{}
	
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "could not create. Parameters were not passed correctly" + err.Error(),
		})
		return
	}

	equipe.NovaEquipe(&req)
	c.JSON(200, gin.H {
		"message": "cadastro realizado com sucesso",
	})
}