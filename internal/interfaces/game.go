package interfaces

type NewGameRequest struct {
	Board string `json:"board"`
}

type NewGameResponse struct {
	URL string `json:"location"`
}
