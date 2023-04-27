package validators

import (
	"errors"
	"regexp"
	"strings"
)

type ValidatorEmail struct {
}

func (v *ValidatorEmail) Validate(email string) error {

	if email == "" {
		return errors.New("email is required")
	}

	if len(email) > 256 {
		return errors.New("email is too long")
	}

	if !isValidFormat(email) {
		return errors.New("email is invalid")
	}
	return nil
}

func isValidFormat(email string) bool {
	if !strings.Contains(email, "@") {
		return false
	}

	parts := strings.Split(email, "@")
	if !isValidUsername(parts[0]) {
		return false
	}

	if !isValidDomain(parts[1]) {
		return false
	}

	return true
}

func isValidUsername(username string) bool {
	if len(username) < 1 || len(username) > 64 {
		return false
	}

	if match, _ := regexp.MatchString("^[a-zA-Z0-9-][a-zA-Z0-9\\-]*[a-zA-Z0-9]$", username); !match {
		return false
	}

	return true
}

func isValidDomain(domain string) bool {
	if len(domain) < 3 || len(domain) > 255 {
		return false
	}

	if strings.HasPrefix(domain, "-") || strings.HasSuffix(domain, "-") {
		return false
	}

	parts := strings.Split(domain, ".")
	if len(parts[len(parts)-1]) < 2 || len(parts[len(parts)-1]) > 6 {
		return false
	}

	for _, part := range parts {
		if strings.HasPrefix(part, "-") || strings.HasSuffix(part, "-") {
			return false
		}

		if match, _ := regexp.MatchString("^[a-zA-Z0-9-][a-zA-Z0-9\\-]*[a-zA-Z0-9]$", part); !match {
			return false
		}
	}

	return true
}
