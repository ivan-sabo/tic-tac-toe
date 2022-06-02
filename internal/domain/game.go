package domain

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

const (
	// running means that the game is still active
	running Status = "RUNNING"
	// xWin means that the player with X won the game
	xWin Status = "X_WON"
	// oWin means that the player with O won the game
	oWin Status = "O_WON"
	// draw means that the game ended undecided
	draw Status = "DRAW"

	roleX Role = "X"
	roleO Role = "O"
)

func init() {
	rand.Seed(time.Now().UnixMicro())
}

// Status holds the current state of the game
type Status string

func (s Status) String() string {
	return string(s)
}

// Role tells if a player uses X or O role
type Role string

func (r Role) String() string {
	return string(r)
}

func (r Role) field() rune {
	if r == roleX {
		return FieldX
	}

	return FieldO
}

// Game represents the main element in domain
type Game struct {
	ID       string
	Board    Board
	Status   Status
	AIRole   Role
	UserRole Role
}

// Games represents a collection of Game domain objects
type Games []Game

func StartGame(b Board) (Game, error) {
	g := Game{
		ID:     uuid.NewString(),
		Board:  b,
		Status: running,
	}

	g.AIRole = roleO
	g.UserRole = roleX
	if b.IsEmpty() {
		g.AIRole = roleX
		g.UserRole = roleO
	}

	if err := g.playAIMove(); err != nil {
		return Game{}, err
	}

	return g, nil
}

// UpdateStatus calculates the state of the game.
func (g *Game) updateStatus(field rune) {
	b := g.Board

	matchCount := 0

	// Check rows
	for ri := 0; ri < 3; ri++ {
		matchCount = 0
		for ci := 0; ci < 3; ci++ {
			if b[ri][ci] == field {
				matchCount++
			}
		}

		if matchCount == 3 {
			if field == FieldO {
				g.Status = oWin
				return
			}
			if field == FieldX {
				g.Status = xWin
				return
			}
		}
	}

	// Check columns
	for ci := 0; ci < 3; ci++ {
		matchCount = 0
		for ri := 0; ri < 3; ri++ {
			if b[ri][ci] == field {
				matchCount++
			}
		}

		if matchCount == 3 {
			if field == FieldO {
				g.Status = oWin
				return
			}
			if field == FieldX {
				g.Status = xWin
				return
			}
		}
	}

	// Check diagonal
	if (b[0][0] == field) && (b[1][1] == field) && (b[2][2] == field) {
		if field == FieldO {
			g.Status = oWin
			return
		}
		if field == FieldX {
			g.Status = xWin
			return
		}
	}

	// Check counter diagonal
	if (b[0][2] == field) && (b[1][1] == field) && (b[2][0] == field) {
		if field == FieldO {
			g.Status = oWin
			return
		}
		if field == FieldX {
			g.Status = xWin
			return
		}
	}

	// Check draw
	// If there is an empty field, the game is still not finished, otherwise, call it draw
	for ri := 0; ri < 3; ri++ {
		for ci := 0; ci < 3; ci++ {
			if b[ri][ci] == FieldEmpty {
				return
			}
		}
	}
	g.Status = draw
}

func (g *Game) PlayUserMove(nb Board) error {
	if err := validateMove(g.Board, nb); err != nil {
		return err
	}

	g.Board = nb
	g.updateStatus(g.UserRole.field())

	if g.Status == running {
		if err := g.playAIMove(); err != nil {
			return err
		}
	}

	return nil
}

func validateMove(b, nb Board) error {
	var moveDone bool = false

	for ri := 0; ri < 3; ri++ {
		for ci := 0; ci < 3; ci++ {
			if nb[ri][ci] != b[ri][ci] {
				if b[ri][ci] != FieldEmpty {
					return ErrFieldAlreadyAssigned
				}

				if moveDone {
					return ErrMoreThanOneMove
				}
				moveDone = true
				continue
			}
		}
	}

	if !moveDone {
		return ErrNoChange
	}

	return nil
}

func (g *Game) playAIMove() error {
	countEmpty := 0
	for ri := 0; ri < 3; ri++ {
		for ci := 0; ci < 3; ci++ {
			if g.Board[ri][ci] == FieldEmpty {
				countEmpty++
			}
		}
	}

	r := rand.Intn(countEmpty) + 1

	countEmpty = 0
playMove:
	for ri := 0; ri < 3; ri++ {
		for ci := 0; ci < 3; ci++ {
			if g.Board[ri][ci] == FieldEmpty {
				countEmpty++

				if r == countEmpty {
					g.Board[ri][ci] = g.AIRole.field()
					break playMove
				}
			}
		}
	}

	g.updateStatus(g.AIRole.field())

	return nil
}

var (
	ErrMoreThanOneMove      error = errors.New("more than one move in new state")
	ErrNoChange             error = errors.New("no new change")
	ErrFieldAlreadyAssigned error = errors.New("cannot change the field that has already been played")
	ErrGameNotFound         error = errors.New("game not found")
)

type GameRepository interface {
	List(context.Context) (Games, error)
	Create(context.Context, Game) (Game, error)
	Get(context.Context, string) (Game, error)
	Update(context.Context, Game) error
	Delete(context.Context, string) error
}
