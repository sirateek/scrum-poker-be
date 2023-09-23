package main

import (
	"fmt"
	"github.com/sirateek/poker-be/config"
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

	socketHandler := handler.NewWebSocketHandler(appConfig.AppConfig.MaximumPlayer, playerService)

	handler.NewDeck(server.Engine.Group("/deck"), deckService)
	handler.NewPlayer(server.Engine.Group("/player"), playerService)
	handler.NewRoom(server.Engine.Group("/room"), roomService, contextManager)

	// Socket
	server.Engine.GET("/ws", socketHandler.Handle)

	// Run the http server
	server.Run()

	// Block the Process and do GracefullyShutdown
	server.GracefullyShutdown()
}
