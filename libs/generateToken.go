package libs

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id uint) (string, error) {
	JWT_SECRET := os.Getenv("JWT_SECRET")
	JWT_SECRET_BYTE := []byte(JWT_SECRET)

	claims := jwt.MapClaims{
		"id":        id,                                   // id of the signed in user
		"expiresAt": time.Now().Add(time.Hour * 1).Unix(), // expire in 1 hour
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(JWT_SECRET_BYTE)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
