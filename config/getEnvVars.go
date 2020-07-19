package config

import (
	"os"
	"reflect"
)

// C stands for config where the env variables are going to b stored
type C struct {
	PORT            string `envName:"PORT" default:"8088"`
	MONGODBCXSTRING string `envName:"MONGODBCXSTRING" default:"mongodb://localhost:27017"`
	DBONE           string `envName:"DBONE" default:"twittor"`
	USRCOL          string `envName:"USRCOL" default:"users"`
	JWTSECRET       string `envName:"JWTSECRET" default:"not the actual encryption key"`
}

// Get environment variables or the default value
func Get(fld string) string {
	confg := reflect.TypeOf(C{})
	f, _ := confg.FieldByName(fld)
	v, _ := f.Tag.Lookup("envName")
	dv, _ := f.Tag.Lookup("default")
	envValue := os.Getenv(v)

	if envValue == "" {
		return dv
	}

	return envValue

}
