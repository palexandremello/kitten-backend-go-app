package appointments

import "time"

type Appointment struct {
	ID              string
	CatID           string
	CreatorID       string
	AppointmentAt   time.Time
	AppointmentType AppointmentType
	Note            string
}
