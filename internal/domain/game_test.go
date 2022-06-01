package domain

import "testing"

func TestPlayMove(t *testing.T) {
	g := Game{Board: Board("X--------")}
	if err := g.PlayMove("X-------O"); err != nil {
		t.Errorf("Got an error while updating, new value should be X-------O, got error %v", err)
	}

	if string(g.Board) != "X-------O" {
		t.Errorf("Invalid Board value, should be X-------O, got %v", g.Board)
	}
}

func TestPlayMoveNoChangeError(t *testing.T) {
	g := Game{Board: Board("X--------")}
	if err := g.PlayMove("X--------"); err != ErrNoChange {
		t.Errorf("Should get an error %v, got %v", ErrNoChange, err)
	}
}

func TestPlayMoveMoreThanOneMoveError(t *testing.T) {
	g := Game{Board: Board("X--------")}
	if err := g.PlayMove("X------OO"); err != ErrMoreThanOneMove {
		t.Errorf("Should get an error %v, got %v", ErrMoreThanOneMove, err)
	}
}

func TestPlayMoveFieldAlreadyAssignedError(t *testing.T) {
	g := Game{Board: Board("X--------")}
	if err := g.PlayMove("O--------"); err != ErrFieldAlreadyAssigned {
		t.Errorf("Should get an error %v, got %v", ErrFieldAlreadyAssigned, err)
	}
}
