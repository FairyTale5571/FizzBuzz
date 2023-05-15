package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func (app *App) Routes() {

	app.server.Use(logger.New())

	app.server.Static("", "./public")
	app.server.Get("/", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		return c.SendFile("index.html")
	})
	app.server.Get("/fizzbuzz/:number", app.fizz.IsFuzzBuzz())
}
