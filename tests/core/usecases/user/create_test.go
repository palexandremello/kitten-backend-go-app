package user

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"

	users "kitten-backend-go-app/app/core/domain/user"
	create "kitten-backend-go-app/app/core/usecases/user"
)

// RepositoryMock mock interface
type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) Save(user *users.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *RepositoryMock) GetByEmail(email string) (*users.User, error) {
	args := m.Called(email)
	if user, ok := args.Get(0).(*users.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestCreateUser(t *testing.T) {
	t.Run("should create user successfully", func(t *testing.T) {
		mockRepo := new(RepositoryMock)
		mockUser := &users.User{
			Name:     "John",
			Email:    "john@example.com",
			Password: "password",
		}
		mockRepo.On("GetByEmail", mock.AnythingOfType("string")).Return(nil, nil)
		mockRepo.On("Save", mock.AnythingOfType("*users.User")).Return(nil)

		t.Log(mockUser)
		usecase := create.NewCreateUserService(mockRepo)
		err := usecase.CreateUser(mockUser.Name, mockUser.Email, mockUser.Password)

		mockRepo.AssertCalled(t, "Save", mock.AnythingOfType("*users.User"))

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Should be able to return a Error if user already exists", func(t *testing.T) {
		mockRepo := new(RepositoryMock)
		service := create.NewCreateUserService(mockRepo)

		existingUser := &users.User{
			Name:     "Test User",
			Email:    "palexandremello@gmail.com",
			Password: "password123",
		}

		mockRepo.On("GetByEmail", "palexandremello@gmail.com").Return(existingUser, nil)

		err := service.CreateUser("Test User", "palexandremello@gmail.com", "password123")
		assert.Error(t, err)
		assert.Equal(t, "user with this email already exists", err.Error())

	})

	t.Run("Should be able to throws if CreateUser throws", func(t *testing.T) {
		mockRepo := new(RepositoryMock)
		service := create.NewCreateUserService(mockRepo)

		mockRepo.On("GetByEmail", "palexandremello@gmail.com").Return(nil, errors.New("repository error"))

		err := service.CreateUser("Test User", "palexandremello@gmail.com", "password123")
		assert.Error(t, err)

	})

}
