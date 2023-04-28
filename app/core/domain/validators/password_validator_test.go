package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordValidator(t *testing.T) {

	validator := PasswordValidator{}

	t.Run("should return error when password is too short", func(t *testing.T) {
		password := "Abc1"

		err := validator.Validate(password)

		assert.EqualError(t, err, ErrPasswordTooShort.Error())

	})

	t.Run("should return an error when missing upper case letter", func(t *testing.T) {
		password := "abc12345"

		err := validator.Validate(password)

		assert.EqualError(t, err, ErrPasswordMissingUpper.Error())
	})

	t.Run("should return an error when missing lower case letter", func(t *testing.T) {
		password := "ABC1235123"

		err := validator.Validate(password)

		assert.EqualError(t, err, ErrPasswordMissingLower.Error())
	})

	t.Run("should return an error when missing digit", func(t *testing.T) {
		password := "Abcdef#h"

		err := validator.Validate(password)

		assert.EqualError(t, err, ErrPasswordMissingDigit.Error())
	})
}
