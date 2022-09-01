package middlewares

import (
	"net/http"

	"github.com/TheSamuelVitor/api-go-postgres/pkg/services"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const Bearer_schema = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"message": "authorization is required",
			})
		}

		token := header[len(Bearer_schema):]
		if services.NewJWTService().ValidateToken(token) {
			c.AbortWithStatus(401)
		}
	}
}