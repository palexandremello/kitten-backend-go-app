package user

import (
	"errors"
	users "kitten-backend-go-app/app/core/domain/user"
	"kitten-backend-go-app/app/core/interfaces/repositories/user"
)

// Service interface creates a User
type Service interface {
	CreateUser(name string, email string, password string) error
}

type userService struct {
	repo user.Repository
}

// NewUserService service create a new user
func NewUserService(repo user.Repository) Service {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(name string, email string, password string) error {
	user := &users.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	errChan := make(chan error, 1)

	go func() {
		existingUser, err := s.repo.GetByEmail(email)
		if err != nil {
			errChan <- err
			return
		}

		if existingUser != nil {
			errChan <- errors.New("user with this email already exists")
			return
		}

		if err := s.repo.Save(user); err != nil {
			errChan <- err
			return
		}

		errChan <- nil
	}()

	return <-errChan
}
