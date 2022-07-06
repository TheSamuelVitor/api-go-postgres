package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Routes = []string{
	"-----------------------",
	"     WELCOME",
	"-----------------------",
	"POSSIBLE ROUTES",
	"GET    /members/",
	"GET    /members/:id",
	"POST   /members/",
	"PUT    /members/:id",
	"DELETE /members/:id",
	"--------------------",
	"    TEAMS ROUTES",
	"--------------------",
	"GET    /teams/",
	"GET    /teams/:id",
	"POST   /teams/",
	"PUT    /teams/:id",
	"DELETE /teams/:id",
	"--------------------",
	"    TASKS ROUTES",
	"--------------------",
	"GET    /tasks/",
	"GET    /tasks/:id",
	"POST   /tasks/",
	"PUT    /tasks/:id",
	"DELETE /tasks/:id",
	"--------------------",
	"    PROJECTS ROUTES",
	"--------------------",
	"GET    /projects/",
	"GET    /projects/:id",
	"POST   /projects/",
	"PUT    /projects/:id",
	"DELETE /projects/:id",
}

func telaInicial(g *gin.Context) {
	g.IndentedJSON(http.StatusOK, Routes)
}

func RegisterRoutes(g *gin.Engine){
	routes := g.Group("")
	routes.GET("", telaInicial)
}