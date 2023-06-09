package cats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeightValue(t *testing.T) {

	t.Run("should be able to return an error if weight is negative", func(t *testing.T) {
		weight := -10.

		weightValue, err := NewWeightValue(weight)

		assert.Nil(t, weightValue)
		assert.Error(t, err)
		assert.Equal(t, "weight value must be greater than 0", err.Error())
	})

	t.Run("should return an validated weight value", func(t *testing.T) {

		weight := 4.67

		weightValue, err := NewWeightValue(weight)

		assert.NotNil(t, weightValue)
		assert.NoError(t, err)
		assert.Equal(t, weight, weightValue.Value())
	})
}
