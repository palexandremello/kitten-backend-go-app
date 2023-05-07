package cats

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWeight(t *testing.T) {

	t.Run("should be able to create a new weight", func(t *testing.T) {

		validatedWeight := 5.67
		weightValue, _ := NewWeightValue(validatedWeight)

		date := time.Now()

		weight := &Weight{
			Date:        date,
			WeightValue: weightValue,
		}

		assert.NotNil(t, weight)
		assert.Equal(t, date, weight.Date)
		assert.Equal(t, validatedWeight, weight.WeightValue.Value())
	})
}
