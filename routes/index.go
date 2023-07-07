package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omarwaleed/goss/lib"
)

func RegisterIndexRouter(app *fiber.App, server *lib.Server) {
	app.Get("/health", HandleHealthRoute())
	app.Post("/join", lib.ValidateSecureKey(server), HandleJoinRoute(server))
}

func HandleHealthRoute() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.SendString("OK")
	}
}

func HandleJoinRoute(server *lib.Server) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// TODO: get the port from the request body and add it after attempting to send a health request to it
		return c.JSON(server)
	}
}
