package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/somuthink/pics_journal/core/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Cookies("jwt") // Read JWT from the cookie

	if c.Path() == "/login" {
		return c.Next()
	}

	if token == "" {
		return c.Redirect("/login")
	}

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (any, error) {
		return []byte(config.Cfg.JWT_TOKEN), nil
	})
	if err != nil {
		return c.Redirect("/login")
	}

	c.Locals("user", claims)
	return c.Next()
}
