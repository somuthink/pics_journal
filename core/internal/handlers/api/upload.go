package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/somuthink/pics_journal/core/internal/crypto"
)

func uploadInput(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	// mu.Lock()
	// defer mu.Unlock()

	fileName := crypto.GenerateUniqueFileName(file.Filename)

	_ = crypto.GetUserID(c)

	destination := fmt.Sprintf("../static/images/uploads/%s", fileName)
	if err := c.SaveFile(file, destination); err != nil {
		return err
	}

	// input, err := inputs.InsertInput(fileName, userID)
	if err != nil {
		return err
	}

	return nil
}
