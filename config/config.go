package config

import (
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"os"
)

func Load() Config {
	var config Config
	ENV, ok := os.LookupEnv("ENV")
	if !ok {
		// Default value for ENV.
		ENV = "local"
	}
	// Load the .env file only for dev env.
	EnvConfig, ok := os.LookupEnv("ENV_CONFIG")
	if !ok {
		EnvConfig = "./.env"
	}

	err := godotenv.Load(EnvConfig)
	if err != nil {
		logrus.Warn("Can't load env file")
	}

	envconfig.MustProcess("", &config)
	config.Env = ENV
	
	// Set CORS Config
	corsConfig := cors.Config{}
	if ENV != "prod" {
		corsConfig = cors.Config{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{"*"},
			AllowMethods: []string{"*"},
		}
	}
	config.CORSConfig = corsConfig

	return config
}
