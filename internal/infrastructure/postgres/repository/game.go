package repository

import (
	"context"

	"github.com/ivan-sabo/tic-tac-toe/internal/domain"
	"github.com/jmoiron/sqlx"
)

type GameDAO struct {
	ID       string `db:"game_id"`
	Board    string `db:"board"`
	Status   string `db:"status"`
	AIRole   string `db:"ai_role"`
	UserRole string `db:"user_role"`
}

type GameDAOs []GameDAO

func (g GameDAO) ToEntity() (domain.Game, error) {
	b, err := domain.FromString(g.Board)
	if err != nil {
		return domain.Game{}, err
	}

	return domain.Game{
		ID:       g.ID,
		Board:    b,
		Status:   domain.Status(g.Status),
		AIRole:   domain.Role(g.AIRole),
		UserRole: domain.Role(g.UserRole),
	}, nil
}

func (gs GameDAOs) ToEntities() (domain.Games, error) {
	games := make(domain.Games, len(gs))

	for _, g := range gs {
		ge, err := g.ToEntity()
		if err != nil {
			return domain.Games{}, err
		}

		games = append(games, ge)
	}

	return games, nil
}

type GamePostgre struct {
	DB *sqlx.DB
}

// List returns all games
func (g *GamePostgre) List(ctx context.Context) (domain.Games, error) {
	list := GameDAOs{}

	const q = `SELECT
	game_id, board, status, ai_role, user_role
	FROM games`

	if err := g.DB.SelectContext(ctx, &list, q); err != nil {
		return nil, err
	}

	games, err := list.ToEntities()
	if err != nil {
		return domain.Games{}, err
	}

	return games, nil
}
