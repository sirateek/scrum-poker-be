package config

import "github.com/gin-contrib/cors"

type Config struct {
	Env            string
	HttpServerPort string `envconfig:"HTTP_SERVER_PORT" default:"3001"`
	GinMode        string `envconfig:"GIN_MODE" default:"release"`
	CORSConfig     cors.Config
}
