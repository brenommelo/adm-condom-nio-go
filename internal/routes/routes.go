package routes

import (
	"net/http"
	"os"

	"github.com/brenommelo/adm-condominio-go/internal/controllers"
	"github.com/brenommelo/adm-condominio-go/internal/handlers"
	"github.com/brenommelo/adm-condominio-go/internal/middlewares"
	"github.com/brenommelo/adm-condominio-go/pkg/utils"
	"github.com/gin-gonic/gin"
)

func SetupRouter() {

	router := gin.Default()

	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/health", handlers.HealthCheckHandler)
		apiV1.POST("/login", handlers.Login)
		apiV1.POST("/signup", controllers.Signup)

		apiV1.GET("/ping", middlewares.RequireAuth, func(c *gin.Context) {
			user, exist := c.Get("user")
			if exist {
				utils.WriteJSONResponse(c, http.StatusAccepted, "", user)
			}

		})
	}

	router.Run(os.Getenv("PORT"))
}
