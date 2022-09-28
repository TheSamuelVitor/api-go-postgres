package teams

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/models"
	"github.com/gin-gonic/gin"
)

type GetInfoOfTeam struct {
	ID_equipe   uint             `json:"id_equipe"`
	Nome_equipe string           `json:"nome_equipe"`
	Membros     []models.Membro  `json:"membros,omitempty"`
	Projetos    []models.Projeto `json:"projetos,omitempty"`
}

func (h handler) GetTeambyId(c *gin.Context) {
	id := c.Param("id")

	var team models.Equipe
	var info GetInfoOfTeam
	var membro []models.Membro
	var projeto []models.Projeto

	if result := h.DB.Raw("select * from equipes where id_equipe = ?", id).Scan(&team); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if result := h.DB.Raw("select * from membros where id_equipe = ?", id).Scan(&membro); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if result := h.DB.Raw("select * from projetos where id_equipe = ?", id).Scan(&projeto); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	info.ID_equipe = uint(team.Id_equipe)
	info.Nome_equipe = team.Nome_equipe
	info.Membros = membro
	info.Projetos = projeto

	c.JSON(http.StatusOK, &info)
}
