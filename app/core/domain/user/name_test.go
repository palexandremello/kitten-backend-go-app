package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {

	t.Run("should create a new name", func(t *testing.T) {

		name, err := NewName("Paulo Mello")
		assert.Nil(t, err)
		assert.Equal(t, "Paulo Mello", name.Value())
	})

}
