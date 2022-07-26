package projects

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type UpdateprojectRequestBody struct {
	Name_project string `json:"name_project"`
	Id_equipe    int    `json:"id_equipe"`
}

func (h handler) Updateproject(c *gin.Context) {
	id := c.Param("id")
	body := UpdateprojectRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var project models.Projeto
	if result := h.DB.First(&project, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	project.Name_project = body.Name_project
	project.Id_equipe = body.Id_equipe

	h.DB.Save(&project)

	c.JSON(http.StatusOK, &project)
}
