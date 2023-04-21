package user

import (
	users "kitten-backend-go-app/app/core/domain/user"
	list "kitten-backend-go-app/app/core/usecases/user"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListUsers(t *testing.T) {

	t.Run("should return a list of users", func(t *testing.T) {

		mockRepo := new(RepositoryMock)
		mockUsers := []*users.User{
			{
				Name:     "Paulo Mello",
				Email:    "palexandremello@gmail.com",
				Password: "password123",
			},
			{
				Name:     "Stereophonics",
				Email:    "stereophonics@email.com",
				Password: "maybe_tomorrow",
			},
		}

		mockRepo.On("List").Return(mockUsers, nil)

		service := list.NewListUserService(mockRepo)
		users, err := service.ListUsers()

		assert.NoError(t, err)
		assert.Equal(t, mockUsers, users)

		mockRepo.AssertExpectations(t)
	})

}
