package user

import (
	users "kitten-backend-go-app/app/core/domain/user"
	"kitten-backend-go-app/app/core/interfaces/repositories/user"
)

// ListUserService interface
type ListUserService interface {
	ListUsers() ([]*users.User, error)
}

type listUserService struct {
	repo user.Repository
}

// NewListUserService service that list all users
func NewListUserService(repo user.Repository) ListUserService {
	return &listUserService{repo: repo}
}

func (s *listUserService) ListUsers() ([]*users.User, error) {
	usersChan := make(chan []*users.User)
	errChan := make(chan error)

	go func() {
		users, err := s.repo.List()
		if err != nil {
			errChan <- err
			close(usersChan)
			return
		}
		usersChan <- users
		close(usersChan)
	}()

	select {
	case users := <-usersChan:
		return users, nil

	case err := <-errChan:
		return nil, err

	}
}
