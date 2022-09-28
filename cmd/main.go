package main

import (
	"os"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/db"
	"github.com/TheSamuelVitor/api-go-postgres/handlers/home"
	"github.com/TheSamuelVitor/api-go-postgres/handlers/login"
	"github.com/TheSamuelVitor/api-go-postgres/handlers/members"
	"github.com/TheSamuelVitor/api-go-postgres/handlers/projects"
	"github.com/TheSamuelVitor/api-go-postgres/handlers/tasks"
	"github.com/TheSamuelVitor/api-go-postgres/handlers/teams"
	"github.com/TheSamuelVitor/api-go-postgres/handlers/user"

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

	r.Use(CORSMiddleware())

	home.RegisterRoutes(r)
	members.RegisterRoutes(r, h)
	teams.RegisterRoutes(r, h)
	tasks.RegisterRoutes(r, h)
	projects.RegisterRoutes(r, h)
	user.RegisterRoutes(r, h)
	login.RegisterRoutes(r, h)

	r.Run(":"+port)
	// r.Run("localhost:3000")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }

}
