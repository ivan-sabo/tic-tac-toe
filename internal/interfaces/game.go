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
