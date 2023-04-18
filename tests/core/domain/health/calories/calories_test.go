package calories_test

import (
	"kitten-backend-go-app/app/core/domain/health/calories"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCalories testing functions about Calories entity
func TestCalories(t *testing.T) {

	t.Run("should be able to create a Calories entity", func(t *testing.T) {
		cal := calories.Calories{
			ID:               "any_id",
			Weight:           4.5,
			IsSpyOrNeutral:   true,
			BMR:              260,
			ActivityLevel:    1.2,
			AgeFactor:        1,
			CastrationFactor: 1.2,
		}

		assert.Equal(t, "any_id", cal.ID)
		assert.Equal(t, float32(4.5), cal.Weight)
		assert.Equal(t, true, cal.IsSpyOrNeutral)
		assert.Equal(t, float32(260), cal.BMR)
		assert.Equal(t, float32(1.2), cal.ActivityLevel)
		assert.Equal(t, float32(1), cal.AgeFactor)
		assert.Equal(t, float32(1.2), cal.CastrationFactor)

	})
}
