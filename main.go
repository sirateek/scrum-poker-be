package main

import (
	"fmt"
	gqlHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/sirateek/poker-be/config"
	"github.com/sirateek/poker-be/graph"
	"github.com/sirateek/poker-be/handler"
	"github.com/sirateek/poker-be/internal/deck"
	"github.com/sirateek/poker-be/internal/player"
	"github.com/sirateek/poker-be/internal/room"
	"github.com/sirateek/poker-be/pkg/httpserver"
	"github.com/sirateek/poker-be/utils"
)

var appConfig config.Config

func init() {
	appConfig = config.Load()
}

func main() {
	server := httpserver.NewHttpServer(
		httpserver.WithListeningAddress(fmt.Sprint(":", appConfig.AppConfig.Port)),
		httpserver.WithCORSConfig(appConfig.AppConfig.CORSConfig),
	)

	// Parse deck
	deckParser := utils.NewDeckConfigParser()
	decks := deckParser.Parse(appConfig.Decks)

	// Service
	playerService := player.NewService()
	deckService := deck.NewDeck(decks)
	roomService := room.NewService(deckService, playerService)

	// Utils
	contextManager := utils.ContextManager{}

	if appConfig.Env != "prod" {
		server.Engine.GET("/playground", gin.WrapH(playground.Handler("GraphQL playground", "/query")))
	}

	taskGqlHandler := gqlHandler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &handler.Resolver{
		DeckService:    deckService,
		RoomService:    roomService,
		ContextManager: contextManager,
	}}))
	server.Engine.POST("/query", handler.UseAuth(contextManager), gin.WrapH(taskGqlHandler))

	// Run the http server
	server.Run()

	// Block the Process and do GracefullyShutdown
	server.GracefullyShutdown()
}
