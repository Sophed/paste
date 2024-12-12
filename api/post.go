package api

import (
	"errors"
	"paste/alerts"
	"paste/util"

	"github.com/gofiber/fiber/v2"
)

func Post(c *fiber.Ctx) error {
	body := string(c.Body())
	if body == "" || !util.Filter(body) {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	content, err := format(body, c.GetReqHeaders())
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	id, err := util.Store([]byte(content))
	if err != nil {
		alerts.Error(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendString(id)
}

func format(body string, headers map[string][]string) (string, error) {
	format, contains := postType(headers)
	if !contains || format == "" {
		return "", errors.New("headers do not contain post type")
	}
	if format == "md" {
		return util.MDtoHTML([]byte(body)), nil
	} else {
		md := "```" + format + "\n" + body + "\n```"
		return util.MDtoHTML([]byte(md)), nil
	}
}

func postType(headers map[string][]string) (string, bool) {
	for field, ls := range headers {
		if field == "Post-Type" {
			return ls[0], true
		}
	}
	return "", false
}
