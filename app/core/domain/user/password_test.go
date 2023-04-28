package users

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockPasswordValidator struct {
	mock.Mock
}

func (m *mockPasswordValidator) Validate(password string) error {
	args := m.Called(password)
	return args.Error(0)
}

func TestPassword(t *testing.T) {

	t.Run("should return a new password validated", func(t *testing.T) {

		password := "Abc1def#"
		validator := &mockPasswordValidator{}
		validator.On("Validate", password).Return(nil)

		result, err := NewPassword(password, validator)

		assert.Nil(t, err)
		assert.Equal(t, password, result.Value())
	})

	t.Run("should return an error if password is invalid", func(t *testing.T) {

		password := "invalid_password"

		validator := &mockPasswordValidator{}
		validator.On("Validate", password).Return(errors.New("password is invalid"))

		result, err := NewPassword(password, validator)

		assert.Error(t, err)
		assert.Nil(t, result)
	})

}
