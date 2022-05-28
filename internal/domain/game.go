package domain

import (
	"github.com/google/uuid"
)

const (
	// RUNNING means that the game is still active
	RUNNING Status = "RUNNING"
	// X_WON means that the player with X won the game
	X_WON Status = "X_WON"
	// O_WON means that the player with O won the game
	O_WON Status = "O_WON"
	// DRAW means that the game ended undecided
	DRAW Status = "DRAW"
)

// Status holds the current state of the game
type Status string

// Game represents the main element in domain
type Game struct {
	ID     string
	Board  Board
	status Status
}

func NewGame() Game {
	return Game{
		ID:     uuid.NewString(),
		Board:  NewBoard(),
		status: RUNNING,
	}
}

// Staus returns the state of the game.
func (g *Game) Status() Status {
	return g.status
}

func (g *Game) PlayMove(nb Board) error {
	if err := validateMove(g.Board, nb); err != nil {
		return err
	}

	g.Board = nb

	return nil
}

func validateMove(b, nb Board) error {
	var moveDone bool = false

	for i := range nb {
		if nb[i] != b[i] {
			if moveDone {
				return ErrMoreThanOneMove
			}
			moveDone = true
			continue
		}
	}

	if !moveDone {
		return ErrNoChange
	}

	return nil
}
