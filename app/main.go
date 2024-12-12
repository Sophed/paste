package main

import (
	"paste/api"
	"paste/app/views"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sophed/lg"
)

const PORT = 1337

func main() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	app.Get("/", views.PageIndex)
	app.Post("/post", api.Post)
	app.Get("/p/:id", api.Get)
	app.Static("/static", "static")
	lg.Info("server started at http://127.0.0.1:" + strconv.Itoa(PORT))
	lg.Fatl(app.Listen(":" + strconv.Itoa(PORT)))
}
