package users

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockEmailValidator struct {
	mock.Mock
}

func (m *mockEmailValidator) Validate(email string) error {
	args := m.Called(email)
	return args.Error(0)
}

func TestEmail(t *testing.T) {

	t.Run("should return a new email validated", func(t *testing.T) {

		email := "palexandremello@gmail.com"
		validator := &mockEmailValidator{}
		validator.On("Validate", email).Return(nil)

		result, err := NewEmail(email, validator)

		assert.Nil(t, err)
		assert.Equal(t, email, result.Value())
	})

	t.Run("should return an error if email is invalid", func(t *testing.T) {

		email := "invalid_email"

		validator := &mockEmailValidator{}
		validator.On("Validate", email).Return(errors.New("email is invalid"))

		result, err := NewEmail(email, validator)

		assert.Error(t, err)
		assert.Nil(t, result)
	})

}
