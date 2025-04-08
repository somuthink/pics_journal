package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/somuthink/pics_journal/core/internal/handlers/middleware"
)

type APIRoutes struct{}

func (APIRoutes) SetupRoutes(app fiber.Router) {
	apiGroup := app.Group("/api")
	apiGroup.Post("/auth", login)

	app.Use(middleware.AuthMiddleware)
}
