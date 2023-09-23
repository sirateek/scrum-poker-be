package model

type Card struct {
	Index        int    `json:"index"`
	DisplayValue string `json:"displayValue"`
}

type Deck struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Cards []*Card `json:"cards"`
}

type Player struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	HideCard    *bool  `json:"hideCard,omitempty"`
	PickedCard  *int   `json:"pickedCard,omitempty"`
	IsSpectator *bool  `json:"isSpectator,omitempty"`
}

type Room struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Deck     *Deck     `json:"deck"`
	Passcode *string   `json:"passcode,omitempty"`
	Players  []*Player `json:"players"`
}
