package libs

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func ParseToken(token string) (jwt.MapClaims, bool) {
	JWT_SECRET := os.Getenv("JWT_SECRET")
	JWT_SECRET_BYTE := []byte(JWT_SECRET)

	parsed, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		// validate the signing method
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(JWT_SECRET_BYTE), nil
	})

	if err != nil {
		return nil, false
	}

	// check token validity
	claims, ok := parsed.Claims.(jwt.MapClaims)

	// change the expiresAt to int
	expiresAtInt := int(claims["expiresAt"].(float64))
	claims["expiresAt"] = expiresAtInt

	// check if the token is not valid and is expired
	if !ok || !parsed.Valid || claims["expiresAt"].(int) < int(time.Now().Unix()) {
		return nil, false
	}

	// token is safe, return the payload and true
	return claims, true
}
