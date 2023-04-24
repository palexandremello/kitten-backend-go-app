package user

import (
	"errors"
	users "kitten-backend-go-app/app/core/domain/user"
	"kitten-backend-go-app/app/core/interfaces/repositories/user"
)

var ErrEmailRequired = errors.New("email is required")

type getUserByEmailUseCase struct {
	userRepo user.Repository
}

type GetUserByEmailUseCase interface {
	Execute(email string) (*users.User, error)
}

func NewGetUserByEmailUseCase(userRepo user.Repository) GetUserByEmailUseCase {
	return &getUserByEmailUseCase{userRepo: userRepo}
}

func (g *getUserByEmailUseCase) Execute(email string) (*users.User, error) {

	if email == "" {
		return nil, ErrEmailRequired
	}

	userChan := make(chan *users.User)
	errChan := make(chan error)

	go func() {
		user, err := g.userRepo.GetByEmail(email)
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
