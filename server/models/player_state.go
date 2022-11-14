package models

type PlayerState struct {
	Name         string `json:"name"`
	Points       int    `json:"points"`
	CurrentLobby string `json:"currentLobby"`
	IsLobbyHost  bool   `json:"isLobbyHost"`
}
