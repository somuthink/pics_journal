package pages

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/somuthink/pics_journal/core/internal/crypto"
	"github.com/somuthink/pics_journal/core/internal/db/users"
)

func MePage(c *fiber.Ctx) error {
	user, err := users.SelectUser(crypto.GetUserID(c))
	if err != nil {
		return err
	}

	return adaptor.HTTPHandler(templ.Handler()(c))
}
