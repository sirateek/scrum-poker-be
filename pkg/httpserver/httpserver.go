package httpserver

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	Address string
	Engine  *gin.Engine
	server  *http.Server
}

func NewHttpServer(options ...Option) Server {
	engine := gin.Default()
	healthCheckGroup := engine.Group("healthcheck")
	healthCheckGroup.GET("/", healthCheckHandler)

	server := Server{
		Engine: engine,
	}

	// Apply Server Options.
	for _, option := range options {
		option(&server)
	}

	return server
}

// Run a server on Addr.
func (srv *Server) Run() {
	srv.server = &http.Server{
		Addr:    srv.Address,
		Handler: srv.Engine,
	}

	go func(server *http.Server) {
		logrus.Infof("Listen on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Errorf("Error while listen and serve: %v", err)
		}
	}(srv.server)
}

func (srv *Server) GracefullyShutdown() {
	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt, syscall.SIGTERM)
	<-wait

	logrus.Info("Shutting down http server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	if err := srv.server.Shutdown(ctx); err != nil {
		logrus.Errorf("Cannot shutdown server: %v", err)
	}

	cancel()
}
