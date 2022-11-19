package main

import (
	// "os"

	"github.com/TheSamuelVitor/api-go-postgres/handlers/home"
	"github.com/TheSamuelVitor/api-go-postgres/handlers/login"
	"github.com/TheSamuelVitor/api-go-postgres/handlers/members"
	"github.com/TheSamuelVitor/api-go-postgres/handlers/projects"
	"github.com/TheSamuelVitor/api-go-postgres/handlers/tasks"
	"github.com/TheSamuelVitor/api-go-postgres/handlers/teams"
	"github.com/TheSamuelVitor/api-go-postgres/handlers/user"
	"github.com/TheSamuelVitor/api-go-postgres/pkg/common/db"
	"github.com/TheSamuelVitor/api-go-postgres/pkg/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/TheSamuelVitor/api-go-postgres/docs"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title          Swagger Example API
// @version        1.0
// @description    This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name  API Support
// @contact.url   http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	// port := os.Getenv("PORT")
	dbUrl := viper.Get("DB_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	docs.SwaggerInfo.Title = "Sistema de Gerenciamento de Projetos"
	docs.SwaggerInfo.Description = "Documentacao da API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3000"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	
	r.Use(middlewares.CORSMiddleware())

	home.RegisterRoutes(r)
	members.RegisterRoutes(r, h)
	teams.RegisterRoutes(r, h)
	tasks.RegisterRoutes(r, h)
	projects.RegisterRoutes(r, h)
	user.RegisterRoutes(r, h)
	login.RegisterRoutes(r, h)

	// r.Run(":"+port)
	r.Run("localhost:3000")
}
