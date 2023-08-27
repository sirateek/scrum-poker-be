package config

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func Load() Config {
	var config Config
	ENV, ok := os.LookupEnv("ENV")
	if !ok {
		// Default value for ENV.
		ENV = "local"
	}

	path, err := os.Getwd()
	if err != nil {
		logrus.Panic("Load Config File Err: ", err)
	}
	configFile, err := os.Open(fmt.Sprintf("%s/config.%s.yaml", path, ENV))
	if err != nil {
		logrus.Panic("Load Config File Err: ", err)
	}

	v := viper.New()
	v.SetConfigType("yaml")
	err = v.ReadConfig(configFile)
	if err != nil {
		logrus.Panic("Load Config File Err: ", err)
	}

	err = v.Unmarshal(&config)
	if err != nil {
		logrus.Panic("Load Config File Err: ", err)
	}
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

	gin.SetMode(config.AppConfig.GinMode)
	config.AppConfig.CORSConfig = corsConfig

	return config
}
