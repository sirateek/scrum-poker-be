package socketconnection

import (
	"encoding/json"
	"errors"
	"github.com/sirateek/poker-be/model"
	"github.com/sirupsen/logrus"
)

var (
	ErrCommandNotFound = errors.New("command not found")
)

type RegisterCommandAttributes struct {
	Name string `json:"name"`
}

type CommandHandler struct {
}

func (s *CommandHandler) Handle(wsConn *SocketConnection, socketCommand model.SocketCommand) error {
	if socketCommand.Command == "" {
		logrus.Error("No Command")
		return nil
	}
	switch socketCommand.Command {
	case "REGISTER":
		var registerAttributes RegisterCommandAttributes
		err := json.Unmarshal(socketCommand.Attributes, &registerAttributes)
		if err != nil {
			return err
		}
		wsConn.Player = &model.Player{
			ID:   wsConn.ID,
			Name: registerAttributes.Name,
		}
	}

	return nil
}
