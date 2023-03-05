package config

import "github.com/kelseyhightower/envconfig"

type APIConfig struct {
	PORT string `envconfig:"PORT" default:"8080"`
}

func NewAPIConfig() (APIConfig, error) {
	var apiConfig APIConfig
	err := envconfig.Process("", &apiConfig)
	return apiConfig, err
}
