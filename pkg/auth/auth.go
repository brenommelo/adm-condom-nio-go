package auth

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/brenommelo/adm-condominio-go/internal/config"
	"github.com/brenommelo/adm-condominio-go/internal/models"
	"github.com/golang-jwt/jwt/v4"
)

func CreateJWT(userID int) (string, error) {
	num, err := strconv.Atoi(os.Getenv("JWTExpirationInSeconds"))
	if err != nil {
		num = 3600
	}

	expiration := time.Second * time.Duration(num)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(int(userID)),
		"expiresAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func ValidateJWT(tokenString string) (models.User, error) {
	// Obtenha a chave secreta da variável de ambiente
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return models.User{}, fmt.Errorf("SECRET_KEY not set in environment")
	}

	// Parse e valida o token JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verifique se o método de assinatura é HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Retorna a chave secreta como um byte array para validar o token
		return []byte(secretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		if float64(time.Now().Unix()) > claims["expiresAt"].(float64) {
			return models.User{}, fmt.Errorf("token expired")
		}

		var user models.User
		config.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			return user, fmt.Errorf("unauthorized user")
		}
		return user, nil

	}

	return models.User{}, err
}
