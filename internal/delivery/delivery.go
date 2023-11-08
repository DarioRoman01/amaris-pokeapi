package delivery

import (
	"github.com/DarioRoman01/pokeapi/internal/delivery/pokemon"
	"github.com/DarioRoman01/pokeapi/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func New(service *service.Service, validate *validator.Validate) *fiber.App {
	fiberApp := fiber.New()
	fiberApp.Use(cors.New())

	api := fiberApp.Group("/api/v1")
	pokemon.New(service, validate, api)
	return fiberApp
}
