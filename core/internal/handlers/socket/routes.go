package socket

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/somuthink/pics_journal/core/internal/handlers/middleware"
)

type SocketRoutes struct{}

var HUB *SocketHub

func (s SocketRoutes) SetupRoutes(app fiber.Router) {
	HUB = NewSocketHub()
	go HUB.Run()

	socketGroup := app.Group("/sock")

	socketGroup.Use("/generate", middleware.WsMiddleware)
	socketGroup.Get("/generate", websocket.New(generate))
}
