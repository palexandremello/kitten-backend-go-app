package users

import "errors"

type Name struct {
	value string
}

func NewName(name string) (*Name, error) {

	if name == "" {
		return nil, errors.New("name is required")
	}

	if len(name) > 100 {
		return nil, errors.New("name is too long")
	}

	return &Name{value: name}, nil
}

func (n *Name) Value() string {
	return n.value
}
