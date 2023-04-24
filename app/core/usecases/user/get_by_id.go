package user

import (
	users "kitten-backend-go-app/app/core/domain/user"
	"kitten-backend-go-app/app/core/interfaces/repositories/user"
)

type GetUserByIDService interface {
	GetUserByID(id string) (*users.User, error)
}

type getUserByIDService struct {
	repo user.Repository
}

func NewGetUserByIDService(repo user.Repository) GetUserByIDService {
	return &getUserByIDService{repo: repo}
}

func (s *getUserByIDService) GetUserByID(id string) (*users.User, error) {
	userChan := make(chan *users.User)
	errChan := make(chan error)

	go func() {
		user, err := s.repo.GetByID(id)
		if err != nil {
			errChan <- err
			close(userChan)
			return
		}

		userChan <- user
		close(userChan)
	}()

	select {
	case user := <-userChan:
		return user, nil

	case err := <-errChan:
		return nil, err

	}
}
