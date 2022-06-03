package domain

import "testing"

func TestNewBoard(t *testing.T) {
	b := NewBoard()

	for ri := range b {
		for ci := range b[ri] {
			if b[ri][ci] != FieldEmpty {
				t.Errorf("all fields should have empty value (%v) initialised, have: %v", FieldEmpty, b)
			}
		}
	}
}
