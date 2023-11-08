package pokemon

import (
	"context"

	"github.com/DarioRoman01/pokeapi/entities"
	"github.com/DarioRoman01/pokeapi/internal/repository"
	"github.com/DarioRoman01/pokeapi/pkg/errors"
	"github.com/DarioRoman01/pokeapi/pkg/pokeapi"
)

type PokemonService struct {
	repository    *repository.Repository
	pokeapiClient pokeapi.Client
}

func New(repo *repository.Repository, pokeapiClient pokeapi.Client) *PokemonService {
	return &PokemonService{
		repository:    repo,
		pokeapiClient: pokeapiClient,
	}
}

func (p *PokemonService) Create(ctx context.Context, pokemon *entities.Pokemon) (*entities.Pokemon, error) {
	pokemonForm, err := p.pokeapiClient.GetPokemonData(ctx, pokemon.ID)
	if err != nil {
		return nil, errors.Wrap(err, "pokemon: PokemonService.Create p.pokeapiClient.GetPokemonData error")
	}

	pokemon.Name, err = p.parsePokemonName(pokemonForm)
	if err != nil {
		return nil, errors.Wrap(err, "pokemon: PokemonService.Create p.ParsePokemonName error")
	}

	pokemon.VersionGroup = pokemonForm.VersionGroup.Name
	pokemon.IsDefault = pokemonForm.IsDefault
	pokemon.Types = p.ParsePokemonTypes(pokemonForm.Types)

	createdPokemon, err := p.repository.Pokemon.Create(ctx, pokemon)
	if err != nil {
		return nil, errors.Wrap(err, "pokemon: PokemonService.Create p.repository.Pokemon.Create error")
	}

	return createdPokemon, nil

}

func (p *PokemonService) parsePokemonName(pokemonForm *pokeapi.PokemonForm) (string, error) {
	switch val := pokemonForm.Pokemon.(type) {
	case string:
		return val, nil

	case map[string]interface{}:
		return val["name"].(string), nil

	default:
		return "", errors.New("unable to parse pokemon name")
	}
}

func (p *PokemonService) ParsePokemonTypes(formTypes []pokeapi.Type) []string {
	types := make([]string, 0, len(formTypes))
	for _, t := range formTypes {
		types = append(types, t.Type.Name)
	}

	return types
}
