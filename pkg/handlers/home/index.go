package home

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

var Routes = []string{
	"------------------------------",
	"POSSIBLE ROUTES",
	"------------------------------",
	"ROTAS DOS MEMBROS",
	"GET    -  /                   ",
	"GET    -  /membros/           ",
	"POST   -  /membros/           ",
	"PUT    -  /membros/:id        ",
	"GET    -  /membros/:id        ",
	"DELETE -  /membros/:id        ",
	"------------------------------",
	"ROTAS DAS EQUIPES",
	"GET    -  /equipes/           ",
	"POST   -  /equipes/           ",
	"PUT    -  /equipes/:id        ",
	"GET    -  /equipes/:id        ",
	"DELETE -  /equipes/:id        ",
	"GET    -  /equipes/members/:id",
	"------------------------------",
	"ROTAS DOS TAREFAS",
	"GET    -  /tarefas/           ",
	"POST   -  /tarefas/           ",
	"PUT    -  /tarefas/:id        ",
	"GET    -  /tarefas/:id        ",
	"DELETE -  /tarefas/:id        ",
	"------------------------------",
	"ROTAS DOS PROJETOS",
	"GET    -  /projetos/          ",
	"POST   -  /projetos/          ",
	"PUT    -  /projetos/:id       ",
	"GET    -  /projetos/:id       ",
	"DELETE -  /projetos/:id       ",
	"GET    -  /projetos/tarefas/:id  ",
}

func telaInicial(c *gin.Context) {
	c.JSON(http.StatusOK, Routes)
}

func RegisterRoutes(g *gin.Engine){
	routes := g.Group("", middlewares.Auth())
	routes.GET("", telaInicial)
}