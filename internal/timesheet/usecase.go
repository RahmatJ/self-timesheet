package timesheet

import "context"

type UseCase interface {
	Create(ctx context.Context, payload CreateTimesheetRequestPayload) error
}
