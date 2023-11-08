package pokeapi

import "context"

type Client interface {
	GetPokemonData(ctx context.Context, id int) (*PokemonForm, error)
}
