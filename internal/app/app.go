package app

import (
	"github.com/fairytale5571/fizzBuzz/internal/fizzBuzz"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	server *fiber.App
	fizz   fizzBuzz.Provider
}

func New() *App {
	return &App{
		server: fiber.New(),
		fizz:   fizzBuzz.New(),
	}
}

func (app *App) Run() error {
	app.Routes()
	return app.server.Listen(":3000")
}

func (app *App) Stop() error {
	return app.server.Shutdown()
}
