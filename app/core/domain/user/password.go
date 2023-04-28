package users

type Password struct {
	value string
}

type PasswordValidator interface {
	Validate(password string) error
}

func NewPassword(password string, validator PasswordValidator) (*Password, error) {

	if err := validator.Validate(password); err != nil {
		return nil, err
	}

	return &Password{
		value: password,
	}, nil
}

func (p *Password) Value() string {
	return p.value
}
