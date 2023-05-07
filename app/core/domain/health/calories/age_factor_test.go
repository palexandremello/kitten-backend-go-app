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
}
