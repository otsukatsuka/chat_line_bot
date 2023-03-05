package config

import "github.com/kelseyhightower/envconfig"

type LineMessageConfig struct {
	SECRET string `envconfig:"CHANNEL_SECRET" default:""`
	TOKEN  string `envconfig:"CHANNEL_TOKEN" default:""`
}

func NewLineMessageConfig() (LineMessageConfig, error) {
	var lineMessageConfig LineMessageConfig
	err := envconfig.Process("", &lineMessageConfig)
	return lineMessageConfig, err
}
