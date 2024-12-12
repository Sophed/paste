package api

import (
	"os"
	"paste/app/components"
	"paste/app/views"
	"paste/util"

	"github.com/gofiber/fiber/v2"
)

func Get(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	entries, err := os.ReadDir(util.STORAGE_DIR)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.Name() != id {
			continue
		}
		data, err := os.ReadFile(util.STORAGE_DIR + "/" + entry.Name())
		if err != nil {
			return err
		}
		return c.SendString(components.ToString(views.Post(id, string(data))))
	}
	return c.SendStatus(fiber.StatusNotFound)
}
