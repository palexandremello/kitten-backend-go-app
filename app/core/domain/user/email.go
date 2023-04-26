package users

type EmailValidator interface {
	Validate(email string) error
}

type Email struct {
	email string
}

func NewEmail(email string, validator EmailValidator) (*Email, error) {
	if err := validator.Validate(email); err != nil {
		return nil, err
	}

	return &Email{
		email: email,
	}, nil
}

func (e *Email) Value() string {
	return e.email
}
