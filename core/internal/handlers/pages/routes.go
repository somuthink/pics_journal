package pages

import (
	"github.com/gofiber/fiber/v2"
	"github.com/somuthink/pics_journal/core/internal/handlers/middleware"
)

type PageRoutes struct{}

func (PageRoutes) SetupRoutes(app fiber.Router) {
	app.Get("/login", loginPage)

	app.Use(middleware.AuthMiddleware)

	app.Get("/", homePage)

	app.Get("/me", homePage)
}
