package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Routes = []string{
	"------------------------------",
	"POSSIBLE ROUTES",
	"------------------------------",
	"ROTAS DOS MEMBROS",
	"GET    - /                    ",
	"GET    - /membros/            ",
	"POST   - /membros/            ",
	"PUT    - /membros/:id         ",
	"GET    - /membros/:id         ",
	"GET    - /membros/count       ",
	"DELETE - /membros/:id         ",
	"------------------------------",
	"ROTAS DAS EQUIPES",
	"GET    - /equipes/            ",
	"POST   - /equipes/            ",
	"PUT    - /equipes/:id         ",
	"GET    - /equipes/:id         ",
	"GET    - /equipes/count       ",
	"DELETE - /equipes/:id         ",
	"GET    - /equipes/members/:id ",
	"------------------------------",
	"ROTAS DOS TAREFAS",
	"GET    - /tarefas/            ",
	"POST   - /tarefas/            ",
	"PUT    - /tarefas/:id         ",
	"GET    - /tarefas/:id         ",
	"GET    - /tarefas/count       ",
	"DELETE - /tarefas/:id         ",
	"------------------------------",
	"ROTAS DOS PROJETOS",
	"GET    - /projetos/           ",
	"POST   - /projetos/           ",
	"PUT    - /projetos/:id        ",
	"GET    - /projetos/:id        ",
	"GET    - /projetos/count      ",
	"DELETE - /projetos/:id        ",
	"GET    - /projetos/tarefas/:id",
}

func telaInicial(g *gin.Context) {
	g.IndentedJSON(http.StatusOK, Routes)
}

func RegisterRoutes(g *gin.Engine){
	routes := g.Group("")
	routes.GET("", telaInicial)
}