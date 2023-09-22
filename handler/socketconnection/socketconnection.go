package socketconnection

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/sirateek/poker-be/model"
	"github.com/sirateek/poker-be/utils"
	"github.com/sirupsen/logrus"
)

type GoRoutineContext struct {
	Context context.Context
	Cancel  func()
}

type SocketConnection struct {
	ID                       string
	Conn                     *websocket.Conn
	Player                   *model.Player
	HandlerStrategy          Strategy
	SocketCommandHandler     *CommandHandler
	IncomingGoRoutineContext GoRoutineContext
	CommandErrorRate         int
	SpawnController          SpawnController
}

func NewSocketConnection(wsConn *websocket.Conn, socketCommandHandler *CommandHandler) *SocketConnection {
	connID := utils.RandStr(10)

	return &SocketConnection{
		ID:                   connID,
		Conn:                 wsConn,
		HandlerStrategy:      NewUnidentifyStrategy(),
		SocketCommandHandler: socketCommandHandler,
		CommandErrorRate:     0,
		SpawnController:      SpawnController{},
	}
}

func (s *SocketConnection) HandlePlayerController() {
	ctx := context.Background()
	cancelCtx, cancel := context.WithCancel(ctx)
	s.IncomingGoRoutineContext = GoRoutineContext{
		Context: cancelCtx,
		Cancel:  cancel,
	}
	s.SpawnController.SetValue(true)

	// HandleIncoming Message
	go func() {
		for {
			select {
			case <-cancelCtx.Done():
				logrus.Warnf("Client with id %s has been kicked", s.ID)
				err := s.Conn.Close()
				if err != nil {
					logrus.Error(err)
				}
				return
			default:
				if !s.SpawnController.GetShouldSpawn() {
					continue
				}
				logrus.Infof("Spawned the handler")
				s.SpawnController.SetValue(false)
				go s.HandleIncomingMessage()
			}
		}
	}()

	// Handle Strategy
	go func() {
		for {
			err := s.HandlerStrategy.Handle(s)
			if err != nil {
				logrus.Error("Strategy Error: ", err)
				// Stop the HandleIncoming Message GoRoutine
				s.IncomingGoRoutineContext.Cancel()
				return
			}

			// Set the next state
			nextState, err := s.HandlerStrategy.NextState(s)
			if err != nil {
				logrus.Error(err)
				return
			}

			if nextState == nil {
				logrus.Debug("No strategy to run for the player %s", s.Player.ID)
				return
			}

			s.HandlerStrategy = nextState
		}
	}()
}

func (s *SocketConnection) HandleIncomingMessage() {
	socketCommand := model.SocketCommand{}
	err := s.Conn.ReadJSON(&socketCommand)
	if err != nil {
		logrus.Error(err)
		return
	}

	//if socketCommand.Command == "" {
	//	return
	//}

	err = s.SocketCommandHandler.Handle(s, socketCommand)
	if err != nil {
		logrus.Error(err)
		s.CommandErrorRate++
	}

	// TODO: Implement Error Rate client kick.

	s.SpawnController.SetValue(true)
}

type Strategy interface {
	Handle(s *SocketConnection) error
	NextState(s *SocketConnection) (Strategy, error)
}
