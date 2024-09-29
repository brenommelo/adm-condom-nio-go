package handlers

import (
	"net/http"

	"github.com/brenommelo/adm-condominio-go/internal/config"
	"github.com/brenommelo/adm-condominio-go/internal/dto"
	"github.com/brenommelo/adm-condominio-go/internal/models"
	"github.com/brenommelo/adm-condominio-go/pkg/auth"
	"github.com/brenommelo/adm-condominio-go/pkg/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HealthCheckHandler(c *gin.Context) {
	utils.WriteJSONResponse(c, http.StatusAccepted, "s", `{"status":"UP"}`)
}

func Login(c *gin.Context) {

	body := dto.SigninRequest{}

	if c.Bind(&body) != nil {
		utils.WriteErrorResponse(c, http.StatusBadRequest, "Failed to read body")
		return
	}

	if err := body.Validate(); err != nil {
		utils.WriteErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	config.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		utils.WriteErrorResponse(c, http.StatusBadRequest, "Invalid email or password")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		utils.WriteErrorResponse(c, http.StatusBadRequest, "Invalid email or password")
		return
	}

	jwtToken, err := auth.CreateJWT(int(user.ID))
	if err != nil {
		utils.WriteErrorResponse(c, http.StatusBadRequest, "Faild to create JWT token")
		return
	}

	// c.SetSameSite(http.SameSiteLaxMode)
	// c.SetCookie("Authorization", jwtToken, 3600, "", "", false, true)

	type response struct {
		Authorization string
	}

	resp := response{Authorization: jwtToken}
	utils.WriteJSONResponse(c, http.StatusAccepted, "Login realizado com sucesso", resp)
}
