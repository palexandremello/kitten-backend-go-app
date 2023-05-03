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

}
