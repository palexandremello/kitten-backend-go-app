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

}
