package views

import (
	"paste/app/components"

	"github.com/gofiber/fiber/v2"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Index() Node {
	return components.View("~/",
		Script(Src("/static/js/index.js")),
		H2(
			Text("Paste Service ~ "),
			A(
				Text("soph"),
				Href("https://soph.systems"),
				Target("_blank"),
			),
		),
		Input(
			ID("type"),
			Placeholder("type"),
			Value("md"),
		),
		Textarea(
			ID("input"),
			Placeholder("start your post..."),
			Attr("spellcheck", "false"),
		),
		Br(),
		Button(Attr("onclick", "post()"),
			Text("Post"),
		),
	)
}

func PageIndex(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendString(components.ToString(Index()))
}
