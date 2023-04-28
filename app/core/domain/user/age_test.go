package users

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAge(t *testing.T) {

	t.Run("should return a validated birthdate", func(t *testing.T) {

		birthdate := time.Now().AddDate(-20, 0, 0)

		age, err := NewAge(birthdate)

		assert.Equal(t, 20, age.Value())
		assert.Nil(t, err)
	})
}
