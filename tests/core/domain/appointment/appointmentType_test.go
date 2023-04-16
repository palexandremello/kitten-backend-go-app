package appointments_test

import (
	appointments "kitten-backend-go-app/app/core/domain/appointment"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppointmentType(t *testing.T) {
	appointmentType := appointments.AppointmentType{
		ID:   "type-789",
		Name: "Checkup",
	}

	assert.Equal(t, "type-789", appointmentType.ID)
	assert.Equal(t, "Checkup", appointmentType.Name)
}
