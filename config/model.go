package config

import "github.com/gin-contrib/cors"

type Config struct {
	Env       string
	AppConfig AppConfig
	Decks     []Deck
}

type AppConfig struct {
	Port          string `mapstructure:"port" default:"3001"`
	GinMode       string `mapstructure:"ginMode" default:"release"`
	CORSConfig    cors.Config
	MaximumPlayer int `default:"100"`
}

type Deck struct {
	ID    string   `mapstructure:"id"`
	Name  string   `mapstructure:"name"`
	Cards []string `mapstructure:"cards"`
}
