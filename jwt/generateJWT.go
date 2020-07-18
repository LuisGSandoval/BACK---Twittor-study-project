package jwt

import (
	"time"

	"github.com/LuisGSandoval/twittor/config"
	"github.com/LuisGSandoval/twittor/models"
	"github.com/dgrijalva/jwt-go"
)

// GenerateJWT generates a jwt token with a payload assigned to it
func GenerateJWT(t models.User) (string, error) {

	secret := []byte(config.JWTSECRET())

	payload := jwt.MapClaims{
		"_id":      t.ID.Hex(),
		"email":    t.Email,
		"name":     t.Name,
		"lastname": t.Lastname,
		"bio":      t.Biography,
		"website":  t.Website,

		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	tokenStr = "Bearer " + tokenStr

	return tokenStr, nil
}
