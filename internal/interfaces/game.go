package interfaces

import (
	"fmt"

	"github.com/ivan-sabo/tic-tac-toe/internal/domain"
)

type NewGameRequest struct {
	Board string `json:"board"`
}

func (ngr NewGameRequest) ToEntity() (domain.Board, error) {
	b, err := domain.BoardFromString(ngr.Board)
	if err != nil {
		return domain.Board{}, fmt.Errorf("creating board domain entity: %w", err)
	}

	return b, nil
}

type NewGameResponse struct {
	URL string `json:"location"`
}

type GetGameReponse struct {
	ID     string `json:"id"`
	Board  string `json:"board"`
	Status string `json:"status"`
}

func NewGetGameReponse(g domain.Game) GetGameReponse {
	return GetGameReponse{
		ID:     g.ID,
		Board:  g.Board.String(),
		Status: g.Status.String(),
	}
}

type ListGameResponse []GetGameReponse

func NewListGameResponse(gs domain.Games) ListGameResponse {
	games := make(ListGameResponse, 0, len(gs))

	for _, g := range gs {
		games = append(games, NewGetGameReponse(g))
	}

	return games
}
