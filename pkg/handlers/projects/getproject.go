package projects

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetProjects(c *gin.Context) {
<<<<<<< HEAD
	var projects []models.ProjetoCompleto
	sql := "select projetos.id_projeto, projetos.nome_projeto, projetos.descricao, projetos.id_equipe, equipes.nome_equipe from projetos join equipes on projetos.id_equipe = equipes.id_equipe order by id_projeto;"
=======
	var projects []models.Projeto
	sql := "select * from projetos"
>>>>>>> parent of be0cd07 (update gets and models to show the respective caleed names)

	if result := h.DB.Raw(sql).Scan(&projects); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &projects)
}
