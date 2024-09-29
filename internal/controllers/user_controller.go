package controllers

import (
	"net/http"

	"github.com/brenommelo/adm-condominio-go/internal/config"
	"github.com/brenommelo/adm-condominio-go/internal/dto"
	"github.com/brenommelo/adm-condominio-go/internal/models"
	"github.com/brenommelo/adm-condominio-go/pkg/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {

	body := dto.SignupRequest{}

	if c.Bind(&body) != nil {
		utils.WriteErrorResponse(c, http.StatusBadRequest, "Failed to read body")
		return
	}

	if err := body.Validate(); err != nil {
		utils.WriteErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		utils.WriteErrorResponse(c, http.StatusBadRequest, "Failed to hash password")
		return
	}

	user := models.User{
		Email:     body.Email,
		Password:  string(hash),
		LastName:  body.LastName,
		FirstName: body.FirstName,
	}

	result := config.DB.Create(&user)
	if result.Error != nil {
		utils.WriteErrorResponse(c, http.StatusBadRequest, "Failed to create user")
		return
	}
	utils.WriteJSONResponse(c, http.StatusCreated, "user created successfully!", user)

}
