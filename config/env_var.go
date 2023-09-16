package config

import "os"

type ConfigStruct struct {
	Port string
}

func LoadEnvVars() (ConfigStruct, error) {
	os.Setenv("PORT", ":3000")
	port := os.Getenv("PORT")

	env := ConfigStruct{
		Port: port,
	}

	return env, nil
}
