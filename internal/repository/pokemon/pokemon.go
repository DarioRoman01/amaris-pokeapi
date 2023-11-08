package pokemon

import (
	"context"

	"github.com/DarioRoman01/pokeapi/entities"
	"github.com/DarioRoman01/pokeapi/pkg/errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

const pokemonsSchema = "pokemons"

type PokemonRepository struct {
	postgres *sqlx.DB
}

func New(postgres *sqlx.DB) *PokemonRepository {
	return &PokemonRepository{
		postgres: postgres,
	}
}

func (p *PokemonRepository) Create(ctx context.Context, pokemon *entities.Pokemon) (*entities.Pokemon, error) {
	query, args, err := sq.
		Insert(pokemonsSchema).
		Columns("id", "name", "generation", "weight", "height", "version_group", "is_default", "types").
		Values(
			pokemon.ID,
			pokemon.Name,
			pokemon.Generation,
			pokemon.Weight,
			pokemon.Height,
			pokemon.VersionGroup,
			pokemon.IsDefault,
			pq.Array(pokemon.Types),
		).
		Suffix("RETURNING *").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "pokemons repository: Create ToSql error")
	}

	var createdPokemon entities.Pokemon
	err = p.postgres.QueryRowx(query, args...).StructScan(&createdPokemon)
	if err != nil {
		return nil, errors.Wrap(err, "pokemons repository: Create postgres.Exec error")
	}

	return &createdPokemon, nil
}
