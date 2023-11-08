package service

import (
	"context"

	"github.com/DarioRoman01/pokeapi/entities"
)

type PokemonService interface {
	Create(context.Context, *entities.Pokemon) (*entities.Pokemon, error)
}
