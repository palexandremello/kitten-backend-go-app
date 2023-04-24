package user

import (
	"errors"
	users "kitten-backend-go-app/app/core/domain/user"
	"kitten-backend-go-app/app/core/usecases/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserByEmailUseCase(t *testing.T) {

	t.Run("should return a user by email", func(t *testing.T) {

		userEmail := "pmello@mail.com"
		expectedUser := &users.User{Email: userEmail}

		mockRepo := new(RepositoryMock)
		mockRepo.On("GetByEmail", userEmail).Return(expectedUser, nil)

		usecase := user.NewGetUserByEmailUseCase(mockRepo)

		user, err := usecase.Execute(userEmail)

		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error if repository throws", func(t *testing.T) {

		userEmail := "pmello@mail.com"

		expectedError := errors.New("something went wrong")

		mockRepo := new(RepositoryMock)
		mockRepo.On("GetByEmail", userEmail).Return(nil, expectedError)

		usecase := user.NewGetUserByEmailUseCase(mockRepo)
		user, err := usecase.Execute(userEmail)

		assert.Nil(t, user)
		assert.EqualError(t, err, expectedError.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return nil and error if email is empty", func(t *testing.T) {

		mockRepo := new(RepositoryMock)

		usecase := user.NewGetUserByEmailUseCase(mockRepo)
		user, err := usecase.Execute("")

		assert.Nil(t, user)
		assert.EqualError(t, err, "email is required")
		mockRepo.AssertExpectations(t)
	})
}
