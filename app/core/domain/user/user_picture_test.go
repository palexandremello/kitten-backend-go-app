package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserPicture(t *testing.T) {

	t.Run("should return an error when data is empty", func(t *testing.T) {

		_, err := NewUserPicture([]byte{}, "image/jpeg", "any_path")

		assert.Error(t, err)
	})

	t.Run("should return an error when image size is too large", func(t *testing.T) {

		data := make([]byte, 11<<20) // 11MB

		_, err := NewUserPicture(data, "image/jpeg", "any_path")

		assert.Error(t, err)
	})
}
