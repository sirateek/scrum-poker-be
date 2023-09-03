package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirateek/poker-be/internal/player"
	"github.com/sirupsen/logrus"
	"net/http"
)

type WebSocketHandler struct {
	webSocketUpgrader websocket.Upgrader
	playerService     player.Service
	maximumPlayerSize int
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
		webSocketUpgrader: wsUpgrader,
		maximumPlayerSize: maxPlayerSize,
		playerService:     playerService,
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

	_, message, err := ws.ReadMessage()
	if err != nil {
		logrus.Error("Client failed to identify him self.")
		ws.Close()
		return
	}
	messageStr := string(message)
	logrus.Info(messageStr)
	playerData, err := w.playerService.RegisterPlayer(messageStr)
	if err != nil {
		logrus.Error("Client failed to register as Player. ", err)
		ws.Close()
	}
	logrus.Info(fmt.Sprintf("A Client Connected, Identified as %s (%s)", playerData.Name, playerData.ID))
	go HandleWebSocketConnection(ws)
}

func HandleWebSocketConnection(ws *websocket.Conn) {
	for {
		ws.ReadMessage()
	}
}
