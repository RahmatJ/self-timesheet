package timesheet

type implementation struct {
}

func (i implementation) Create(payload CreateTimesheetRequestPayload) error {
	return nil
}

func NewUseCaseImpl() UseCase {
	return &implementation{}
}
