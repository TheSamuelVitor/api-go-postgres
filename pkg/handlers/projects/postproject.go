package projects

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/models"
	"github.com/gin-gonic/gin"
)

type AddProjectRequestBody struct {
	Nome_projeto string `json:"nome_projeto"`
	Descricao    string `json:"descricao"`
	Id_equipe    int    `json:"id_equipe"`
}

func (h handler) PostProjects(c *gin.Context) {

	body := AddProjectRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var newProject models.Projeto

	newProject.Nome_projeto = body.Nome_projeto
	newProject.Descricao = body.Descricao
	newProject.Id_equipe = body.Id_equipe

	if result := h.DB.Create(&newProject); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &newProject)
}
