package timesheet

import "context"

type Repository interface {
	Create(ctx context.Context, payload CreateTimesheetRecordPayload) error
}
