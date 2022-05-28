package domain

import "testing"

func TestNewBoard(t *testing.T) {
	b := NewBoard()

	if string(b) != "---------" {
		t.Errorf("Invalid value for a new board, should be ---------, got %v", string(b))
	}
}

func TestFromString(t *testing.T) {
	b, err := FromString("X---O---X")

	if err != nil {
		t.Errorf("Got error %v", err)
	}

	if string(b) != "X---O---X" {
		t.Errorf("Invalid FromString() value, should be X---O---X, got %v", string(b))
	}
}

func TestFromStringFieldValueError(t *testing.T) {
	b, err := FromString("X---O---C")

	if err != ErrInvalidFieldValue {
		t.Errorf("Invalid returned value for err, should be %v, got %v", ErrInvalidFieldValue, err)
	}

	if b != "" {
		t.Errorf("Invalid returned value for Board, should be \"\", got %v", b)
	}
}

func TestFromStringLengthError(t *testing.T) {
	b, err := FromString("X---O---X-")

	if err != ErrInvalidLength {
		t.Errorf("Invalid returned value for err, should be %v, got %v", ErrInvalidLength, err)
	}

	if b != "" {
		t.Errorf("Invalid returned value for Board, should be \"\", got %v", b)
	}
}

func TestString(t *testing.T) {
	b := Board("X---O---X")

	if b.String() != "X---O---X" {
		t.Errorf("Invalid value from String(), should be X---O---X, got %v", b.String())
	}
}
