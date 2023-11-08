package repository

import "github.com/DarioRoman01/pokeapi/internal/repository/pokemon"

type Repository struct {
	Pokemon PokemonRepository
}

func New(pokemonRepository *pokemon.PokemonRepository) *Repository {
	return &Repository{
		Pokemon: pokemonRepository,
	}
}
