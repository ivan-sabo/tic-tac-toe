package interfaces

import (
	"fmt"

	"github.com/ivan-sabo/tic-tac-toe/internal/domain"
)

type PostGameRequest struct {
	Board string `json:"board"`
}

func (ngr PostGameRequest) ToEntity() (domain.Board, error) {
	b, err := domain.BoardFromString(ngr.Board)
	if err != nil {
		return domain.Board{}, err
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

type PutGameRequest struct {
	Board string `json:"board"`
}

func (ngr PutGameRequest) ToEntity() (domain.Board, error) {
	b, err := domain.BoardFromString(ngr.Board)
	if err != nil {
		return domain.Board{}, fmt.Errorf("creating board domain entity: %w", err)
	}

	return b, nil
}

type PutGameReponse struct {
	ID     string `json:"id"`
	Board  string `json:"board"`
	Status string `json:"status"`
}

func NewPutGameReponse(g domain.Game) PutGameReponse {
	return PutGameReponse{
		ID:     g.ID,
		Board:  g.Board.String(),
		Status: g.Status.String(),
	}
}
