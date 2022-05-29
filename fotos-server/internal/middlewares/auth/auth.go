package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthCheck(c *fiber.Ctx) error {
	var token string
	bearToken := c.Get("Authorization")

	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		token = onlyToken[1]
	} else {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Unauthorized"})
	}

	validToken, id := VerifyToken(token)

	if !validToken {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Unauthorized"})
	}
	c.Locals("id", id)

	return c.Next()
}
