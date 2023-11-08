package pokeapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DarioRoman01/pokeapi/pkg/errors"
)

const baseUrl = "https://pokeapi.co/api/v2"

type PokeAPIClient struct {
	client *http.Client
}

func NewClient() *PokeAPIClient {
	return &PokeAPIClient{
		client: http.DefaultClient,
	}
}

func (p *PokeAPIClient) GetPokemonData(ctx context.Context, id int) (*PokemonForm, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/pokemon-form/%d", baseUrl, id), http.NoBody)
	if err != nil {
		return nil, errors.Wrap(err, "pokeapi: PokeAPIClient.GetPokemonData http.NewRequest error")
	}

	resp, err := p.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, errors.Wrap(err, "pokeapi: PokeAPIClient.GetPokemonData p.client.Do error")
	}

	if resp.StatusCode > 299 {
		return nil, errors.New("pokeapi: PokeAPIClient.GetPokemonData error: recieved an unexpected status code")
	}

	defer resp.Body.Close()

	var pokemon PokemonForm
	err = json.NewDecoder(resp.Body).Decode(&pokemon)
	if err != nil {
		return nil, errors.Wrap(err, "pokeapi: PokeAPIClient.GetPokemonData json.NewDecoder.Decode error")
	}
	fmt.Println(pokemon.Pokemon)
	return &pokemon, nil
}
