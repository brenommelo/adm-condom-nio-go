package routes

import (
	"os"

	"github.com/brenommelo/adm-condominio-go/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() {

	router := gin.Default()

	apiV1 := router.Group("/api/v1")
	{
		apiV1.GET("/health", handlers.HealthCheckHandler)
		apiV1.POST("/login", func(ctx *gin.Context) {})
	}

	router.Run(os.Getenv("PORT"))
}
