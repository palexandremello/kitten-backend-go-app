package users

import (
	"errors"
	"time"
)

var ErrInvalidAge = errors.New("age is invalid")

type Age struct {
	birthdate time.Time
	age       int
}

func NewAge(birthdate time.Time) (*Age, error) {

	age := int(time.Since(birthdate).Hours() / 24 / 365.25)

	if age < 0 || age > 150 {
		return nil, ErrInvalidAge
	}

	return &Age{
		birthdate: birthdate,
		age:       age,
	}, nil
}

func (a *Age) Value() int {
	return a.age
}
