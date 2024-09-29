package handlers

import (
	"net/http"

	"github.com/brenommelo/adm-condominio-go/pkg/utils"
	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(c *gin.Context) {
	utils.WriteJSONResponse(c, http.StatusAccepted, "s", `{"status":"UP"}`)
}
