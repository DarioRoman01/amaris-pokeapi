package repository

import (
	"context"

	"github.com/DarioRoman01/pokeapi/entities"
)

type PokemonRepository interface {
	Create(context.Context, *entities.Pokemon) (*entities.Pokemon, error)
}
