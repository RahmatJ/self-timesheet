package timesheet

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type handler struct {
	uc UseCase
}

func (h *handler) createTimesheet(ctx *fiber.Ctx) error {
	payload := new(CreateTimesheetRequestPayload)

	err := ctx.BodyParser(payload)
	if err != nil {
		log.Error().Msgf("Error when parsing body: %v", err)
		return nil
	}
	log.Info().Msgf("Starting create record with payload: %v", *payload)
	//  TODO(Rahmat): validate body
	err = h.uc.Create(*payload)
	if err != nil {
		log.Error().Msgf("Error when creating record timesheet: %v", err)
		return err
	}
	log.Info().Msg("Success create record")
	return nil
}

func NewHandler(app *fiber.App, uc UseCase) {
	h := &handler{
		uc: uc,
	}

	baseGroup := app.Group("/timesheet")
	baseGroup.Post("", h.createTimesheet)
}
