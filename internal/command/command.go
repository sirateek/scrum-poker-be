package command

import (
	"errors"
	"github.com/sirateek/poker-be/handler"
	"github.com/sirateek/poker-be/model"
	"github.com/sirupsen/logrus"
)

var (
	ErrCommandNotFound = errors.New("command not found")
)

type CommandHandler struct {
}

func (s *CommandHandler) Handle(wsConn *handler.WebSocketHandler, socketCommand model.SocketCommand) error {
	if socketCommand.Command == "" {
		logrus.Error("No Command")
		return nil
	}

	return nil
}
