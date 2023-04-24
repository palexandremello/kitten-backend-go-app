package user

import (
	users "kitten-backend-go-app/app/core/domain/user"
	"kitten-backend-go-app/app/core/usecases/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserByIDService(t *testing.T) {

	t.Run("should return a user by ID", func(t *testing.T) {

		userID := "any_id"
		expectedUser := &users.User{ID: userID}

		mockRepo := new(RepositoryMock)
		mockRepo.On("GetByID", userID).Return(expectedUser, nil)

		service := user.NewGetUserByIDService(mockRepo)
		user, err := service.GetUserByID(userID)

		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
		mockRepo.AssertExpectations(t)
	})
}
