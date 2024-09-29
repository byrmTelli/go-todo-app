package handlers

import (
	"go-todo-app/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user models.User) (string, error) {
	claims := &jwt.MapClaims{
		"username":   user.Username,
		"email":      user.Email,
		"expiration": time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
