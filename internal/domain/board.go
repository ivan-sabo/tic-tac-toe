package domain

import (
	"errors"
	"strings"
)

// Board represents the status of the board.
type Board string

// NewBoard creates a brand new Board.
func NewBoard() Board {
	return "---------"
}

// FromString recreates an existing Board out of string.
func FromString(s string) (Board, error) {
	if len(s) != 9 {
		return Board(""), ErrInvalidLength
	}
	return Board(s), nil
}

// String implements Stringer interface.
func (b *Board) String() string {
	return string(*b)
}

// Update does the next move.
func (b *Board) Update(n Board) error {
	oldState := strings.Split(b.String(), "")
	newState := strings.Split(n.String(), "")

	var moveDone bool = false

	for i, c := range newState {
		if err := validateField(newState[i]); err != nil {
			return err
		}
		if c != oldState[i] {
			if moveDone {
				return ErrMoreThanOneMove
			}
			newState[i] = c
			moveDone = true
			continue
		}
	}

	return nil
}

func validateField(f string) error {
	switch f {
	case "-", "X", "O":
		return nil
	}
	return ErrInvalidFieldValue
}

var (
	ErrInvalidLength     error = errors.New("board invalid length")
	ErrInvalidFieldValue error = errors.New("field value invalid")
	ErrMoreThanOneMove   error = errors.New("more than one move in new state")
)
