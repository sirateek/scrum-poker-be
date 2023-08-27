package main

import (
	"fmt"
	gqlHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/sirateek/poker-be/config"
	"github.com/sirateek/poker-be/graph"
	"github.com/sirateek/poker-be/handler"
	"github.com/sirateek/poker-be/pkg/httpserver"
)

var appConfig config.Config

func init() {
	appConfig = config.Load()
}

func main() {
	server := httpserver.NewHttpServer(
		httpserver.WithListeningAddress(fmt.Sprint(":", appConfig.HttpServerPort)),
		httpserver.WithCORSConfig(appConfig.CORSConfig),
	)

	if appConfig.Env != "prod" {
		server.Engine.GET("/playground", gin.WrapH(playground.Handler("GraphQL playground", "/query")))
	}

	taskGqlHandler := gqlHandler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &handler.Resolver{}}))
	server.Engine.POST("/query", gin.WrapH(taskGqlHandler))

	// Run the http server
	server.Run()

	// Block the Process and do GracefullyShutdown
	server.GracefullyShutdown()
}
