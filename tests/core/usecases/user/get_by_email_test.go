package user

import (
	"errors"
	users "kitten-backend-go-app/app/core/domain/user"
	"kitten-backend-go-app/app/core/usecases/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUserByEmailUseCase(t *testing.T) {

	t.Run("should return a user by email", func(t *testing.T) {

		mockUser := &users.User{
			ID:       "abc123",
			Name:     "Paulo Mello",
			Email:    "pmello@example.com",
			Password: "password123",
		}

		mockRepo := new(RepositoryMock)
		mockRepo.On("GetByEmail", mock.Anything).Return(mockUser, nil)

		useCase := user.NewGetUserByEmailUseCase(mockRepo)

		user, err := useCase.Execute("pmello@example.com")

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, mockUser, user)
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
