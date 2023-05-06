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
}
