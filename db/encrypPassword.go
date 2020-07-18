package db

import (
	"golang.org/x/crypto/bcrypt"
)

// EncrypPassword encrypts the passwords of new users
func EncrypPassword(pw string) (string, error) {

	cost := 12
	bytes, err := bcrypt.GenerateFromPassword([]byte(pw), cost)

	if err != nil {
		return "", err
	}

	return string(bytes), nil

}
