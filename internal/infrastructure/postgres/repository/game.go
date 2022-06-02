package repository

import (
	"context"
	"database/sql"
	"fmt"

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

func GameDAOFromEntity(g domain.Game) GameDAO {
	return GameDAO{
		ID:       g.ID,
		Board:    g.Board.String(),
		Status:   g.Status.String(),
		AIRole:   g.AIRole.String(),
		UserRole: g.UserRole.String(),
	}
}

type GameDAOs []GameDAO

func (g GameDAO) ToEntity() (domain.Game, error) {
	b, err := domain.BoardFromString(g.Board)
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
	games := make(domain.Games, 0, len(gs))

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

func (g *GamePostgre) Create(ctx context.Context, newGame domain.Game) (domain.Game, error) {
	gd := GameDAOFromEntity(newGame)

	const q = `INSERT INTO games
	(game_id, board, status, ai_role, user_role)
	VALUES($1, $2, $3, $4, $5)`

	if _, err := g.DB.ExecContext(ctx, q, gd.ID, gd.Board, gd.Status, gd.AIRole, gd.UserRole); err != nil {
		return domain.Game{}, fmt.Errorf("inserting product: %w", err)
	}

	ge, err := gd.ToEntity()
	if err != nil {
		return domain.Game{}, fmt.Errorf("converting to entity: %w", err)
	}

	return ge, nil
}

// Get tries to find a single game by uuid
func (g *GamePostgre) Get(ctx context.Context, uuid string) (domain.Game, error) {
	gd := GameDAO{}

	const q = `SELECT
	game_id, board, status, ai_role, user_role
	FROM games
	WHERE game_id = $1`

	if err := g.DB.GetContext(ctx, &gd, q, uuid); err != nil {
		if err == sql.ErrNoRows {
			return domain.Game{}, domain.ErrGameNotFound
		}
		return domain.Game{}, err
	}

	game, err := gd.ToEntity()
	if err != nil {
		return domain.Game{}, err
	}

	return game, nil
}
