package model

type CreateRoom struct {
	Name     string `json:"name"`
	Passcode string `json:"passcode"`
	DeckID   string `json:"deckId"`
}

type JoinRoom struct {
	ID       string `json:"id"`
	Passcode string `json:"passcode"`
}
