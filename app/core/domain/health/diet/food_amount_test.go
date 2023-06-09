package diet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoodAmoun(t *testing.T) {

	t.Run("should reutrn an error if food amoun is negative", func(t *testing.T) {

		amount := float32(-10.0)

		foodAmount, err := NewFoodAmount(amount)

		assert.Nil(t, foodAmount)
		assert.Error(t, err)
		assert.Equal(t, "food amount cannot be negative", err.Error())
	})

	t.Run("should return a validated food amount", func(t *testing.T) {

		amount := float32(4.5)

		foodAmount, err := NewFoodAmount(amount)

		assert.NotNil(t, foodAmount)
		assert.NoError(t, err)
		assert.Equal(t, amount, foodAmount.Value())

	})
}
