package config

import "os"

// MONGODBCXSTRING brings the string that is stored in the env variables
func MONGODBCXSTRING() string {

	var MONGODBCXSTRING string = os.Getenv("MONGODBCXSTRING")
	if MONGODBCXSTRING == "" {
		MONGODBCXSTRING = "mongodb://localhost:27017"
	}
	return MONGODBCXSTRING
}

// PORT brings the PORT that is stored in the env variables
func PORT() string {
	var PORT string = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8088"
	}
	return PORT
}

// BDONE brings the db name selected from the environment variables
func BDONE() string {
	var BDONE string = os.Getenv("BDONE")
	if BDONE == "" {
		BDONE = "twittor"
	}
	return BDONE
}

// JWTSECRET brings the db name selected from the environment variables
func JWTSECRET() string {
	var JWTSECRET string = os.Getenv("JWTSECRET")
	if JWTSECRET == "" {
		JWTSECRET = "not the actual encryption key"
	}
	return JWTSECRET
}
