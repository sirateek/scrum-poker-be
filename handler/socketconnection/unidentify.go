package socketconnection

import (
	"errors"
	"github.com/sirateek/poker-be/model"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	ErrClientNotIdentifyThemself = errors.New("client not identify themself")
)

type UnidentifyStrategy struct {
}

func NewUnidentifyStrategy() Strategy {
	return &UnidentifyStrategy{}
}

func (u UnidentifyStrategy) Handle(s *SocketConnection) error {
	isSendIdentifyYourSelf := false
	for i := 0; i < 50; i++ {
		if !isSendIdentifyYourSelf {
			s.Conn.WriteJSON(model.SocketCommand{
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
	return nil, nil
}
