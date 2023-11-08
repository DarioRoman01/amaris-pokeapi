package service

type Service struct {
	Pokemon PokemonService
}

func New(pokemonService PokemonService) *Service {
	return &Service{
		Pokemon: pokemonService,
	}
}
