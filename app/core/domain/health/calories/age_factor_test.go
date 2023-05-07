package calories

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAgeFactor(t *testing.T) {

	t.Run("should return an error if age is negative", func(t *testing.T) {

		age := -10

		_, err := NewAgeFactor(age)

		assert.NotNil(t, err)
		assert.Equal(t, "age cannot be negative", err.Error())
	})

	t.Run("should return 1.3 if age is greater than or equal to 15", func(t *testing.T) {

		age := 15

		ageFactor, _ := NewAgeFactor(age)

		assert.NotNil(t, ageFactor)
		assert.Equal(t, 1.3, ageFactor.Value())
	})

	t.Run("should return 1.1 if age is an senior cat (more than 11 and minor )", func(t *testing.T) {

		age := 12

		ageFactor, _ := NewAgeFactor(age)

		assert.NotNil(t, ageFactor)
		assert.Equal(t, 1.1, ageFactor.Value())
	})

}
