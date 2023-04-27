package users

type EmailValidator interface {
	Validate(email string) error
}

type Email struct {
	value string
}

func NewEmail(email string, validator EmailValidator) (*Email, error) {
	if err := validator.Validate(email); err != nil {
		return nil, err
	}

	return &Email{
		value: email,
	}, nil
}

func (e *Email) Value() string {
	return e.value
}
