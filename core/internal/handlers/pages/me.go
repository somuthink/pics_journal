package pages

import (
	"time"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/somuthink/pics_journal/core/internal/crypto"
	"github.com/somuthink/pics_journal/core/internal/db/events"
	"github.com/somuthink/pics_journal/core/internal/db/users"
	"github.com/somuthink/pics_journal/core/internal/views/me"
)

func mePage(c *fiber.Ctx) error {
	userID := crypto.GetUserID(c)
	user, err := users.SelectUser(userID)
	if err != nil {
		return err
	}

	agenda, stats, err := events.SelectEventsAgenda(userID)
	if err != nil {
		return err
	}
	today := time.Now().Weekday().String()

	return adaptor.HTTPHandler(templ.Handler(me.MeIndex(user, agenda, stats, today)))(c)
}
