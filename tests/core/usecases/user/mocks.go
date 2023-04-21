package user

import (
	users "kitten-backend-go-app/app/core/domain/user"

	"github.com/stretchr/testify/mock"
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

func (m *RepositoryMock) List() ([]*users.User, error) {
	args := m.Called()
	if users, ok := args.Get(0).([]*users.User); ok {
		return users, args.Error(1)
	}
	return nil, args.Error(1)
}
