package users

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {

	t.Run("should create a new name", func(t *testing.T) {

		name, err := NewName("Paulo Mello")
		assert.Nil(t, err)
		assert.Equal(t, "Paulo Mello", name.Value())
	})

	t.Run("should return an error if name is empty", func(t *testing.T) {

		_, err := NewName("")
		assert.NotNil(t, err)
		assert.Equal(t, "name is required", err.Error())
	})

	t.Run("should return an error if name is too long", func(t *testing.T) {

		_, err := NewName("Paulo" + strings.Repeat("a", 100))

		assert.NotNil(t, err)

		assert.Equal(t, "name is too long", err.Error())
	})
}
