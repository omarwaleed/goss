package lib

import "github.com/gofiber/fiber/v2"

func ValidateSecureKey(server *Server) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if c.Get("X-Secure-Key") != server.SecureKey {
			return fiber.ErrUnauthorized
		}
		return c.Next()
	}
}
