package repository

import (
	"context"

	"github.com/ivan-sabo/tic-tac-toe/internal/domain"
	"github.com/jmoiron/sqlx"
)

type GameDAO struct {
	ID       string `db:"id"`
	Board    string `db:"board"`
	Status   string `db:"status"`
	AIRole   string `db:"ai_role"`
	UserRole string `db:"user_role"`
}

func (g *GameDAO) ToEntity() (domain.Game, error) {
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

type GamePostgre struct {
	db *sqlx.DB
}

// List returns all games
func (g *GamePostgre) List(ctx context.Context) ([]domain.Game, error) {
	list := []domain.Game{}

	const q = `SELECT
	id, board, status, ai_role, user_role,
	FROM games`

	if err := g.db.SelectContext(ctx, &list, q); err != nil {
		return nil, err
	}

	return list, nil
}
