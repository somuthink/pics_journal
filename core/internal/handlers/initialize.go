package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/somuthink/pics_journal/core/internal/handlers/api"
	"github.com/somuthink/pics_journal/core/internal/handlers/pages"
	"github.com/somuthink/pics_journal/core/internal/handlers/socket"
)

type Router interface {
	SetupRoutes(app fiber.Router)
}

func Initialize() {
	app := fiber.New()

	app.Static("/static", "../static")

	routes := []Router{
		&socket.SocketRoutes{},
		&api.APIRoutes{},
		&pages.PageRoutes{},
	}

	for _, route := range routes {
		route.SetupRoutes(app)
	}

	app.Listen(":3000")
}
