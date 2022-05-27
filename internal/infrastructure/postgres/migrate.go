package postgres

import (
	"github.com/GuiaBolso/darwin"
	"github.com/jmoiron/sqlx"
)

var migrations = []darwin.Migration{
	{
		Version:     1,
		Description: "Add games",
		Script: `
CREATE TABLE games (
	game_id	UUID,
	board	TEXT,
	status	TEXT,
	
	PRIMARY KEY (game_id)
);`,
	},
}

func Migrate(db *sqlx.DB) error {
	driver := darwin.NewGenericDriver(db.DB, darwin.PostgresDialect{})

	d := darwin.New(driver, migrations, nil)

	return d.Migrate()
}
