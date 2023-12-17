package timesheet

import (
	"context"
	"github.com/rs/zerolog/log"
)

type useCaseImpl struct {
	repository Repository
}

func (u useCaseImpl) Create(c context.Context, payload CreateTimesheetRequestPayload) error {

	recordPayload := CreateTimesheetRecordPayload{
		Description: payload.Description,
	}

	err := u.repository.Create(c, recordPayload)
	if err != nil {
		log.Error().Msgf("Error when create record. Error: %v", err)
		return err
	}

	return nil
}

func NewUseCaseImpl(repository Repository) UseCase {
	return &useCaseImpl{
		repository: repository,
	}
}
