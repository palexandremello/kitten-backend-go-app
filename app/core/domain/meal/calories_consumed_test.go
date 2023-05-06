package meal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaloriesConsumed(t *testing.T) {

	t.Run("should return an error when calories consumed is negative", func(t *testing.T) {

		_, err := NewCaloriesConsumed(-1.0)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "calories consumed cannot be negative")
	})

	t.Run("should validated calories consumed", func(t *testing.T) {
		caloriesConsumed := 230.50

		cc, err := NewCaloriesConsumed(caloriesConsumed)

		assert.Nil(t, err)
		assert.Equal(t, caloriesConsumed, cc.Value())

	})
}
