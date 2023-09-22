package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirateek/poker-be/handler/socketconnection"
	"github.com/sirateek/poker-be/internal/player"
	"github.com/sirupsen/logrus"
	"net/http"
)

type WebSocketHandler struct {
	webSocketUpgrader    websocket.Upgrader
	playerService        player.Service
	socketCommandHandler *socketconnection.CommandHandler
	maximumPlayerSize    int
}

func NewWebSocketHandler(maxPlayerSize int, playerService player.Service) *WebSocketHandler {
	wsUpgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// Ignore the origin checking process
	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	return &WebSocketHandler{
		webSocketUpgrader:    wsUpgrader,
		maximumPlayerSize:    maxPlayerSize,
		playerService:        playerService,
		socketCommandHandler: &socketconnection.CommandHandler{},
	}
}

func (w *WebSocketHandler) Handle(c *gin.Context) {
	ws, err := w.webSocketUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	socketConnection := socketconnection.NewSocketConnection(ws, w.socketCommandHandler)

	logrus.Infof("A Client Connected with ID %s.", socketConnection.ID)
	go socketConnection.HandlePlayerController()
}
