package socketconnection

import (
	"errors"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	ErrClientNotIdentifyThemself = errors.New("client not identify themself")
)

type UnidentifyStrategy struct {
	socketCommandHandler *SocketCommandHandler
}

func NewUnidentifyStrategy(socketCommandHandler *SocketCommandHandler) Strategy {
	return &UnidentifyStrategy{
		socketCommandHandler: socketCommandHandler,
	}
}

func (u UnidentifyStrategy) Handle(s *SocketConnection) error {
	isSendIdentifyYourSelf := false
	for i := 0; i < 10; i++ {
		if !isSendIdentifyYourSelf {
			s.Conn.WriteJSON(SocketCommand{
				Command: "IDENTIFY_U_R_SELF",
			})
			isSendIdentifyYourSelf = true
		}

		if s.Player != nil {
			logrus.Info("Successfully identify")
			return nil
		}

		// Check message every 500ms.
		time.Sleep(500 * time.Millisecond)
	}

	return ErrClientNotIdentifyThemself
}

func (u UnidentifyStrategy) NextState(s *SocketConnection) (Strategy, error) {
	//TODO implement me
	panic("implement me")
}
