package httpserver

import (
	"github.com/gin-contrib/cors"
)

type Option func(server *Server)

// WithCORSConfig is used to add the CORs config to the gin middleware.
func WithCORSConfig(corsConfig cors.Config) Option {
	return func(server *Server) {
		// Allow CORs Policy
		server.Engine.Use(cors.New(corsConfig))
	}
}

// WithListeningAddress is a mandatory option to specify the listening address of the http server.
func WithListeningAddress(address string) Option {
	return func(server *Server) {
		server.Address = address
	}
}
