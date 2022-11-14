package models

type Track struct {
	CoverImage string `json:"coverImage"`
}

type SessionState struct {
	CurrentRound int    `json:"currentRound"`
	CurrentTrack Track  `json:"currentTrack"`
	CurrentLobby string `json:"currentLobby"`
	IsLobbyHost  bool   `json:"isLobbyHost"`
}
