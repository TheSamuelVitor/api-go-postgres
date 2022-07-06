package projects

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type AddProjectRequestBody struct {
	Name_project string `json:"name_project"`
	Id_Team      int    `json:"id_team"`
}

func (h handler) PostProjects(c *gin.Context) {

	body := AddProjectRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var newProject models.Project

	newProject.Name_project = body.Name_project
	newProject.Id_Team = body.Id_Team

	if result := h.DB.Create(&newProject); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &newProject)
}
