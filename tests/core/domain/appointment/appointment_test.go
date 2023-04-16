package appointments_test

import (
	appointments "kitten-backend-go-app/app/core/domain/appointment"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAppointment(t *testing.T) {
	catID := "cat-123"
	creatorID := "user-456"
	appointmentAt := time.Now()
	appointmentType := appointments.AppointmentType{
		ID:   "type-789",
		Name: "Checkup",
	}
	note := "Bring vaccination records"

	appointment := appointments.Appointment{
		ID:              "any_id",
		CatID:           catID,
		CreatorID:       creatorID,
		AppointmentAt:   appointmentAt,
		AppointmentType: appointmentType,
		Note:            note,
	}

	assert.Equal(t, "any_id", appointment.ID)
	assert.Equal(t, catID, appointment.CatID)
	assert.Equal(t, creatorID, appointment.CreatorID)
	assert.Equal(t, appointmentAt, appointment.AppointmentAt)
	assert.Equal(t, appointmentType, appointment.AppointmentType)
	assert.Equal(t, note, appointment.Note)
}
