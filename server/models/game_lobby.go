package models

type GameLobby struct {
	NumberOfRounds int `json:"numberOfRounds,omitempty"`
	LengthOfRound  int `json:"lengthOfRound,omitempty"`
	PlayerMax      int `json:"playerMax,omitempty"`
}
