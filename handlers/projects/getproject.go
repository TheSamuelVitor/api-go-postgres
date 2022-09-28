package projects

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetProjects(c *gin.Context) {
	var projects []models.ProjetocomEquipe
	sql := "select projetos.id_projeto, projetos.nome_projeto, projetos.descricao, projetos.id_equipe, equipes.nome_equipe from projetos join equipes on projetos.id_equipe = equipes.id_equipe order by id_projeto;"

	if result := h.DB.Raw(sql).Scan(&projects); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &projects)
}
