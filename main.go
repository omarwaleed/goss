package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omarwaleed/goss/lib"
	"github.com/omarwaleed/goss/routes"
)

func main() {

	server := &lib.Server{}

	app := fiber.New()

	routes.RegisterConfigRouter(app, server) // /config
}
