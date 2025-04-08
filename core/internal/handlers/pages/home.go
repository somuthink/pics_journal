package pages

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/somuthink/pics_journal/core/internal/views/home"
)

func homePage(c *fiber.Ctx) error {
	return adaptor.HTTPHandler(templ.Handler(home.HomeIndex()))(c)
}
