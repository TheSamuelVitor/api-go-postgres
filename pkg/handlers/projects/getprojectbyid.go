package projects

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetProjectbyId(c *gin.Context) {
	id := c.Param("id")

	var project models.ProjetocomEquipe
	sql := "select projetos.nome_projeto, projetos.descricao, equipes.nome_equipe from projetos join equipes on projetos.id_equipe = equipes.id_equipe where id_projeto = ?"

	if result := h.DB.Raw(sql, id).Scan(&project); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &project)
}
