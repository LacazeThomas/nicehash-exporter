package config

import (
	"github.com/caarlos0/env"
)

type AppConfig struct {
	APIUrl          string `env:"APIUrl"`
	XOrganizationId string `env:"XOrganizationId"`
	APIKey          string `env:"APIKey"`
	APISecret       string `env:"APISecret"`
	Environment     string `default:"dev" env:"ENVIRONMENT"`
}

var appConfig AppConfig

//IsDev return true if application is on dev stack
func (a AppConfig) IsDev() bool {
	return IsDev(a.Environment)
}

//IsDev return true if application is on dev stack
func IsDev(env string) bool {
	return env == "dev" || env == "development"
}

//InitAppConfig struct with env variables
func initAppConfig() {
	env.Parse(&appConfig)
}

//GetAppConfig return initialize config structure with variable env
func GetAppConfig() AppConfig {
	if (appConfig == AppConfig{}) {
		initAppConfig()
	}
	return appConfig
}

func SetAppConfig(cfg AppConfig) {
	appConfig = cfg
}
