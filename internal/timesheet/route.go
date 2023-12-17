package timesheet

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRoute(app *fiber.App, db *mongo.Database) {
	repository := NewRepository(db)
	uc := NewUseCaseImpl(repository)
	NewHandler(app, uc)
}
