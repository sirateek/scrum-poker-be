package socketconnection

import (
	"encoding/json"
	"errors"
	"github.com/sirateek/poker-be/internal/room"
	"github.com/sirateek/poker-be/model"
	"github.com/sirupsen/logrus"
)

var (
	ErrCommandNotFound = errors.New("command not found")
)

type RegisterCommandAttributes struct {
	Name string `json:"name"`
}

type JoinRoomCommandAttributes struct {
	RoomID   string `json:"roomID"`
	Passcode string `json:"passcode"`
}

type CommandHandler struct {
	RoomService room.Service
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

	case "PICK_CARD":
	case "HIDE_MY_CARD":
	case "SHOW_MY_CARD":
	case "FLIP_ALL_CARD_ON_TABLE":
	case "CLEAR_CARD_ON_TABLE":
	case "PULL_MY_CARD_OUT_OF_TABLE":
	}

	return nil
}
