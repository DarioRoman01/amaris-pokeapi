package pokemon

import (
	"net/http"

	"github.com/DarioRoman01/pokeapi/entities"
	"github.com/DarioRoman01/pokeapi/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type PokemonDelivery struct {
	service   *service.Service
	validator *validator.Validate
}

func New(service *service.Service, validate *validator.Validate, router fiber.Router) {
	delivery := &PokemonDelivery{service: service, validator: validate}
	route := router.Group("/pokemon")
	delivery.CreatePokemon(route)
}

func (p *PokemonDelivery) CreatePokemon(route fiber.Router) {
	route.Post("/create", func(c *fiber.Ctx) error {
		var pokemon entities.Pokemon
		if err := c.BodyParser(&pokemon); err != nil {
			return fiber.NewError(http.StatusUnprocessableEntity, err.Error())
		}

		if err := p.validator.StructCtx(c.Context(), &pokemon); err != nil {
			return fiber.NewError(http.StatusBadRequest, err.Error())
		}

		createdPokemon, err := p.service.Pokemon.Create(c.Context(), &pokemon)
		if err != nil {
			return fiber.NewError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(fiber.Map{
			"data": createdPokemon,
		})
	})
}
