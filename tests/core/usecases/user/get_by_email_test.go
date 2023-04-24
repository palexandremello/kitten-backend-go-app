package user

import (
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
}
