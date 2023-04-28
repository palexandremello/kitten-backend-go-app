package validators

import (
	"errors"
	"regexp"
)

var (
	ErrPasswordTooShort      = errors.New("password is too short")
	ErrPasswordMissingUpper  = errors.New("password must contain at least one uppercase letter")
	ErrPasswordMissingLower  = errors.New("password must contain at least one lowercase letter")
	ErrPasswordMissingDigit  = errors.New("password must contain at least one digit")
	ErrPasswordInvalidFormat = errors.New("password format is invalid")
)

type PasswordValidator struct {
}

func (v *PasswordValidator) Validate(password string) error {

	if len(password) < 8 {
		return ErrPasswordTooShort
	}

	if !containsUpper(password) {
		return ErrPasswordMissingUpper
	}

	if !containsLower(password) {
		return ErrPasswordMissingLower
	}

	if !containsDigit(password) {
		return ErrPasswordMissingDigit
	}

	if !isValidPasswordFormat(password) {
		return ErrPasswordInvalidFormat
	}

	return nil

}

func containsUpper(s string) bool {
	r, _ := regexp.Compile("[A-Z]")
	return r.MatchString(s)
}

func containsLower(s string) bool {
	r, _ := regexp.Compile("[a-z]")
	return r.MatchString(s)
}

func containsDigit(s string) bool {
	r, _ := regexp.Compile("[0-9]")
	return r.MatchString(s)
}

func isValidPasswordFormat(s string) bool {
	r, _ := regexp.Compile(`^[a-zA-Z0-9!@#$%^&*()_+{}\[\]:;'"<>,.?/\\|-]{8,}$`)
	return r.MatchString(s)
}
