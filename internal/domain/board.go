package domain

import (
	"errors"
)

// Board represents the status of the board.
type Board string

// NewBoard creates a brand new Board.
func NewBoard() Board {
	return "---------"
}

// FromString should be used to create a new Board structure from string.
func FromString(s string) (Board, error) {
	if err := validateLength(s); err != nil {
		return Board(""), err
	}

	for _, v := range s {
		if err := validateField(string(v)); err != nil {
			return Board(""), err
		}

	}

	return Board(s), nil
}

// String implements Stringer interface.
func (b Board) String() string {
	return string(b)
}

func validateField(f string) error {
	switch f {
	case "-", "X", "O":
		return nil
	}
	return ErrInvalidFieldValue
}

func validateLength(s string) error {
	if len(s) != 9 {
		return ErrInvalidLength
	}

	return nil
}

var (
	ErrInvalidLength     error = errors.New("board invalid length")
	ErrInvalidFieldValue error = errors.New("field value invalid")
	ErrMoreThanOneMove   error = errors.New("more than one move in new state")
	ErrNoChange          error = errors.New("no new change")
)
