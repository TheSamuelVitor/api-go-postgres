package main

import (
	"github.com/TheSamuelVitor/api-go-postgres/webservice/equipe"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	equipes := server.Group("equipes")
	equipe.Router(equipes)

	server.Run()
}