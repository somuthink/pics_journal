package api

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/somuthink/pics_journal/core/internal/crypto"
	"github.com/somuthink/pics_journal/core/internal/db/users"
)

func login(c *fiber.Ctx) error {
	name := c.FormValue("username")
	password := c.FormValue("password")

	user, exists, err := users.InsertOrSelect(name, password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if exists {
		if !crypto.ComparePassword(user.Password, password) {
			return c.Status(400).JSON(fiber.Map{
				"message": "incorrect password",
			})
		}
	}

	token, err := crypto.GenerateToken(user.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:    "jwt",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 240),
	})

	c.Response().Header.Add("Hx-Redirect", "/")

	return c.Status(200).JSON(fiber.Map{
		"message": "user logged",
	})
}
