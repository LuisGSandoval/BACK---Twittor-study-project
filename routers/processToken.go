package routers

import (
	"errors"
	"strings"

	"github.com/LuisGSandoval/twittor/config"
	"github.com/LuisGSandoval/twittor/db"
	"github.com/LuisGSandoval/twittor/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// Email is the user's email returned from the request to the data base
var Email string

// UserID is the user id returned from the request to the data base
var UserID string

// ProcessToken extracts the values in the token
func ProcessToken(tk string) (*models.Claim, bool, string, error) {

	secret := config.Get("JWTSECRET")
	secretByte := []byte(secret)
	claims := &models.Claim{}

	// splitToken := strings.Split(tk, " ")

	// if len(splitToken) != 2 {
	// 	return claims, false, "", errors.New("invalid format token")
	// }

	tk = strings.TrimSpace(tk) // token without the bearer word

	jwtkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return secretByte, nil
	})

	if err != nil {
		return claims, false, "", err
	}

	if !jwtkn.Valid {
		return claims, false, "", errors.New("Invalid token")
	}

	_, userFound, _ := db.UserExistanceCheck(claims.Email)

	if userFound {
		Email = claims.Email
		UserID = claims.ID.Hex()
	}

	return claims, userFound, UserID, nil

}
