package main

import (
	"os"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/db"
	"github.com/TheSamuelVitor/api-go-postgres/pkg/handlers/home"
	"github.com/TheSamuelVitor/api-go-postgres/pkg/handlers/members"
	"github.com/TheSamuelVitor/api-go-postgres/pkg/handlers/projects"
	"github.com/TheSamuelVitor/api-go-postgres/pkg/handlers/tasks"
	"github.com/TheSamuelVitor/api-go-postgres/pkg/handlers/teams"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := os.Getenv("PORT") 
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	home.RegisterRoutes(r)
	members.RegisterRoutes(r,h)
	teams.RegisterRoutes(r,h)
	tasks.RegisterRoutes(r,h)
	projects.RegisterRoutes(r,h)

	r.Run(port)
}