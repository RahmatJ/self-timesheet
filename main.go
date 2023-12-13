package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"self-timesheet/internal/db/mongo"
	"self-timesheet/internal/timesheet"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("Starting Server ...")
	app := fiber.New()
	//TODO(Rahmat): Move this to env variable
	port := 3000
	address := fmt.Sprintf(":%d", port)

	_, err := mongo.InitClient()
	if err != nil {
		log.Fatal().Err(err).Msgf("Error occurred: %v", err)
		return
	}

	timesheet.NewRoute(app)

	err = app.Listen(address)
	if err != nil {
		log.Fatal().Err(err).Msgf("Error occurred: %v", err)
		return
	}
}
