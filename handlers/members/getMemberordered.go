package members

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MembrocomEquipe struct {
	Id_membro   int    `gorm:"primary_key" json:"id_membro"`
	Nome_membro string `json:"nome_membro"`
	Funcao      string `json:"funcao"`
	Id_equipe   int    `json:"id_equipe"`
	Nome_equipe string `json:"nome_equipe"`
}

func (h handler) GetMemberOrdered(c *gin.Context) {

	ordena := c.Param("ordena")
	var membros []MembrocomEquipe

	comandoSQL := "select m.id_membro, m.nome_membro, m.funcao, m.id_equipe, e.nome_equipe, m.created_at, m.updated_at from membros m join equipes e on e.id_equipe = m.id_equipe order by "+ordena
	result := h.DB.Raw(comandoSQL).Scan(&membros)
	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
 
	c.JSON(http.StatusOK, &membros)
}