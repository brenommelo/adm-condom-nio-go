package middlewares

import (
	"net/http"

	"github.com/brenommelo/adm-condominio-go/pkg/auth"
	"github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context) {

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.Next()
		return
	}
	user, err := auth.ValidateJWT(tokenString)

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.Next()
		return
	}
	c.Set("user", user)
	c.Next()
}
