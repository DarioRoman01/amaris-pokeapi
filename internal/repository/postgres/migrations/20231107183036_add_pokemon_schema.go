package migrations

import (
	"context"
	"database/sql"

	"github.com/DarioRoman01/pokeapi/pkg/errors"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAddPokemonSchema, downAddPokemonSchema)
}

func upAddPokemonSchema(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS pokemons (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			generation TEXT NOT NULL,
			weight INTEGER NOT NULL,
			height INTEGER NOT NULL,
			version_group TEXT NOT NULL,
			is_default TEXT NOT NULL,
			types TEXT[] NOT NULL
		);
	`)

	if err != nil {
		return errors.Wrap(err, "migrations: upAddPokemonSchema tx.ExecContext error")
	}

	return nil
}

func downAddPokemonSchema(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.ExecContext(ctx, `
		DROP TABLE IF EXIXTS pokemons;	
	`)

	if err != nil {
		return errors.Wrap(err, "migrations: downAddPokemonSchema tx.ExecContext error")
	}

	return nil
}
