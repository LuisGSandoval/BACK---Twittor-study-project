package jwt

import (
	"time"

	"github.com/LuisGSandoval/twittor/config"
	"github.com/LuisGSandoval/twittor/models"
	"github.com/dgrijalva/jwt-go"
)

// GenerateJWT generates a jwt token with a payload assigned to it
func GenerateJWT(t models.User) (string, error) {

	JWTSECRET := config.Get("JWTSECRET")
	secret := []byte(JWTSECRET)

	payload := jwt.MapClaims{
		"_id":       t.ID.Hex(),
		"email":     t.Email,
		"name":      t.Name,
		"lastname":  t.Lastname,
		"bio":       t.Biography,
		"website":   t.Website,
		"birthDate": t.BirthDate,
		"location":  t.Location,

		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	// tokenStr = "Bearer " + tokenStr

	return tokenStr, nil
}
