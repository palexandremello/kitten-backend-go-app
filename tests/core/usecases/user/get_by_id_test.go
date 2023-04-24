package user

import (
	"errors"
	users "kitten-backend-go-app/app/core/domain/user"
	"kitten-backend-go-app/app/core/usecases/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserByIDUseCase(t *testing.T) {

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

	t.Run("should return error if repository throws", func(t *testing.T) {

		userID := "any_id"
		expectedError := errors.New("something went wrong")

		mockRepo := new(RepositoryMock)
		mockRepo.On("GetByID", userID).Return(nil, expectedError)

		service := user.NewGetUserByIDService(mockRepo)
		user, err := service.GetUserByID(userID)

		assert.Nil(t, user)
		assert.EqualError(t, err, expectedError.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return nil and error if ID is empty", func(t *testing.T) {

		mockRepo := new(RepositoryMock)

		service := user.NewGetUserByIDService(mockRepo)
		user, err := service.GetUserByID("")

		assert.Nil(t, user)
		assert.EqualError(t, err, "ID is required")
		mockRepo.AssertExpectations(t)
	})
}
