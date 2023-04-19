package user

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"

	users "kitten-backend-go-app/app/core/domain/user"
	"kitten-backend-go-app/app/core/usecases/user"
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
	return args.Get(0).(*users.User), args.Error(1)
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
		usecase := user.NewUserService(mockRepo)
		err := usecase.CreateUser(mockUser.Name, mockUser.Email, mockUser.Password)

		mockRepo.AssertCalled(t, "Save", mock.AnythingOfType("*users.User"))

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

}
