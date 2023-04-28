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
}
