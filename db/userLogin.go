package db

import (
	"github.com/LuisGSandoval/twittor/models"
	"golang.org/x/crypto/bcrypt"
)

// ValidatePasswords functionalty
func ValidatePasswords(email, password string) (models.User, bool) {

	user, succes, _ := UserExistanceCheck(email)

	if !succes {
		return user, false
	}

	var (
		bdPwbytes  = []byte(user.Password)
		resPwBytes = []byte(password)
	)

	err := bcrypt.CompareHashAndPassword(bdPwbytes, resPwBytes)

	if err != nil {
		return user, false
	}

	return user, true

}
