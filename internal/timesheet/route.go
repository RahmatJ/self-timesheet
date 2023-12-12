package timesheet

import "github.com/gofiber/fiber/v2"

func NewRoute(app *fiber.App) {
	uc := NewUseCaseImpl()
	NewHandler(app, uc)
}
