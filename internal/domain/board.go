package domain

import (
	"errors"
	"strings"
)

const (
	FieldX     rune = 'X'
	FieldO     rune = 'O'
	FieldEmpty rune = '-'
)

// Board represents the status of the board.
type Board [][]rune

// NewBoard creates a brand new Board.
func NewBoard() Board {
	return Board{{FieldEmpty, FieldEmpty, FieldEmpty}, {FieldEmpty, FieldEmpty, FieldEmpty}, {FieldEmpty, FieldEmpty, FieldEmpty}}
}

// IsEmpty checks if any moves ware played
func (b Board) IsEmpty() bool {
	for _, r := range b {
		for _, v := range r {
			if v != FieldEmpty {
				return false
			}
		}
	}
	return true
}

// FromString should be used to create a new Board structure from string.
func BoardFromString(s string) (Board, error) {
	if len(s) == 0 {
		return NewBoard(), nil
	}

	if err := validateLength(s); err != nil {
		return Board{}, err
	}

	for _, v := range s {
		if err := validateField(v); err != nil {
			return Board{}, err
		}
	}

	r := []rune(s)
	b := Board{r[:3], r[3:6], r[6:9]}

	return b, nil
}

// String implements Stringer interface.
func (b Board) String() string {
	s := []string{}
	for _, v := range b {
		s = append(s, string(v))
	}
	return strings.Join(s, "")
}

func validateField(f rune) error {
	switch f {
	case FieldEmpty, FieldX, FieldO:
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

func validateNewBoard(b Board) error {
	// There should be at most one "X" field and zero "O" fields in initial board
	countX := 0
	countO := 0
	for ri := range b {
		for ci := range b[ri] {
			validateField(b[ri][ci])
			switch b[ri][ci] {
			case FieldX:
				countX++
			case FieldO:
				countO++
			}
		}
	}

	if (countX > 1) || (countO > 0) {
		return ErrInvalidInitialBoard
	}

	return nil
}

var (
	ErrInvalidLength       error = errors.New("board invalid length")
	ErrInvalidFieldValue   error = errors.New("field value invalid")
	ErrInvalidInitialBoard error = errors.New("initial board value invalid")
)
