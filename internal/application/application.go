package application

import (
	"fmt"

	"github.com/DarioRoman01/pokeapi/config"
	"github.com/DarioRoman01/pokeapi/internal/delivery"
	"github.com/DarioRoman01/pokeapi/internal/repository"
	pokemonRepository "github.com/DarioRoman01/pokeapi/internal/repository/pokemon"
	"github.com/DarioRoman01/pokeapi/internal/service"
	pokemonService "github.com/DarioRoman01/pokeapi/internal/service/pokemon"
	_ "github.com/DarioRoman01/pokeapi/pkg/migrations"
	"github.com/DarioRoman01/pokeapi/pkg/pokeapi"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type Application struct {
	name          string
	validator     *validator.Validate
	config        *config.Config
	pokeapiClient pokeapi.Client
	delivery      *fiber.App
}

type Config struct {
	Name          string
	Validator     *validator.Validate
	Config        *config.Config
	Postgres      *sqlx.DB
	PokeapiClient pokeapi.Client
}

func New(conf *Config) *Application {
	pokemonRepository := pokemonRepository.New(conf.Postgres)
	repository := repository.New(pokemonRepository)

	pokemonService := pokemonService.New(repository, conf.PokeapiClient)
	service := service.New(pokemonService)

	delivery := delivery.New(service, conf.Validator)
	return &Application{
		name:          conf.Name,
		validator:     conf.Validator,
		config:        conf.Config,
		pokeapiClient: conf.PokeapiClient,
		delivery:      delivery,
	}
}

func (app *Application) Start() error {
	return app.delivery.Listen(fmt.Sprintf(":%s", app.config.HTTPPort))
}
