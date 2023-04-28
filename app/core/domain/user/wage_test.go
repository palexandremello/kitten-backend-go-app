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
}
