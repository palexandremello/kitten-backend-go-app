package cats

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBirthday(t *testing.T) {

	t.Run("should return an error if birthday date is empty", func(t *testing.T) {
		emptyValue := time.Time{}

		birthday, err := NewBirthday(emptyValue)

		assert.Error(t, err)
		assert.Nil(t, birthday)
		assert.Equal(t, "birthday date is required", err.Error())
	})

	t.Run("should return an error if birthday date is after now", func(t *testing.T) {
		futureValue := time.Now().AddDate(1, 0, 0)
		birthday, err := NewBirthday(futureValue)

		assert.Error(t, err)
		assert.Nil(t, birthday)
		assert.Equal(t, "birthday date is invalid", err.Error())
	})

	t.Run("should return a birthday date validated", func(t *testing.T) {
		birthdayValue := time.Date(2015, time.October, 8, 0, 0, 0, 0, time.UTC)
		birthday, err := NewBirthday(birthdayValue)

		assert.NoError(t, err)
		assert.NotNil(t, birthday)
		assert.Equal(t, birthdayValue, birthday.Value())
	})

}
