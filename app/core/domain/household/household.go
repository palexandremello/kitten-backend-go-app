package households

import (
	cats "kitten-backend-go-app/app/core/domain/cats"
	users "kitten-backend-go-app/app/core/domain/user"
)

type Household struct {
	ID        string
	Name      string
	AdminID   string
	Members   []*users.User
	Cats      []*cats.Cat
	Latitude  float32
	Longitude float32
}
