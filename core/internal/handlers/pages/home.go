package pages

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/somuthink/pics_journal/core/internal/crypto"
	"github.com/somuthink/pics_journal/core/internal/db/inputs"
	"github.com/somuthink/pics_journal/core/internal/views/home"
)

func homePage(c *fiber.Ctx) error {
	inps, err := inputs.SelectInputs(crypto.GetUserID(c), 7, true)
	if err != nil {
		return err
	}
	return adaptor.HTTPHandler(templ.Handler(home.HomeIndex(inps)))(c)
}
