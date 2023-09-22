package model

import "encoding/json"

type SocketCommand struct {
	Command    string          `json:"command"`
	Attributes json.RawMessage `json:"attributes"`
}
