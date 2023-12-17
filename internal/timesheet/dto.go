package timesheet

type (
	CreateTimesheetRequestPayload struct {
		Description string `json:"description"`
	}

	CreateTimesheetRecordPayload struct {
		Description string `bson:"description"`
	}
)
