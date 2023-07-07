package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omarwaleed/goss/lib"
)

func RegisterConfigRouter(app *fiber.App, server *lib.Server) {
	group := app.Group("/config")
	group.All("/*", lib.ValidateSecureKey(server))
	group.Get("/", HandleGetConfig(server))
}

func HandleGetConfig(server *lib.Server) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.JSON(server)
	}
}
