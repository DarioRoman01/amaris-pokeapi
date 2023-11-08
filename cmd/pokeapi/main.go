package main

import (
	"log"

	"github.com/DarioRoman01/pokeapi/config"
	"github.com/DarioRoman01/pokeapi/internal/application"
	"github.com/DarioRoman01/pokeapi/internal/repository/postgres"
	"github.com/DarioRoman01/pokeapi/pkg/pokeapi"
	"github.com/go-playground/validator/v10"
)

func main() {
	conf, err := config.New()
	if err != nil {
		log.Fatalf("Main: config.New error: %v", err)
	}

	db, err := postgres.New(&conf.Postgres)
	if err != nil {
		log.Fatalf("Main: postgres.New error: %v", err)
	}

	validate := validator.New()
	pokeapiClient := pokeapi.NewClient()
	app := application.New(&application.Config{
		Name:          "pokeapi-wrapper",
		Validator:     validate,
		Config:        conf,
		PokeapiClient: pokeapiClient,
		Postgres:      db,
	})

	err = app.Start()
	if err != nil {
		log.Fatalf("Main: app.Start error: %v", err)
	}
}
