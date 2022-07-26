package projects

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddProjectRequestBody struct {
	Name_project string `json:"name_project"`
	Description  string `json:"description"`
	Id_equipe    int    `json:"id_equipe"`
}

func (h handler) PostProjects(c *gin.Context) {

	body := AddProjectRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var newProject models.Projeto

	newProject.Name_project = body.Name_project
	newProject.Description = body.Description
	newProject.Id_equipe = body.Id_equipe

	if result := h.DB.Create(&newProject); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &newProject)
}
