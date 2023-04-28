package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWage(t *testing.T) {

	t.Run("should return error if MonthlyWage is negative", func(t *testing.T) {

		monthlyWage := -100.0

		_, err := NewWage(monthlyWage)

		assert.Error(t, err)
	})

	t.Run("should return a validated MonthlyWage", func(t *testing.T) {

		monthlyWage := 20000.0

		result, err := NewWage(monthlyWage)

		assert.Equal(t, monthlyWage, result.Value())
		assert.Nil(t, err)
	})
}
