package timesheet

type UseCase interface {
	Create(payload CreateTimesheetRequestPayload) error
}
