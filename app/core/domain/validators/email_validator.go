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
	println(parts[0])
	if len(parts[0]) < 1 || len(parts[0]) > 64 {
		return false
	}

	if len(parts[1]) < 3 || len(parts[1]) > 255 {
		return false
	}

	domain := parts[1]

	if strings.HasPrefix(domain, "-") || strings.HasSuffix(domain, "-") {
		return false
	}

	if !strings.Contains(domain, ".") {
		return false
	}

	subdomains := strings.Split(domain, ".")

	for _, subdomain := range subdomains {
		if len(subdomains[len(subdomains)-1]) < 2 || len(subdomains[len(subdomains)-1]) > 6 {
			return false
		}

		if strings.HasPrefix(subdomain, "-") || strings.HasSuffix(subdomain, "-") {
			return false
		}

		if match, _ := regexp.MatchString("^[a-zA-Z0-9-][a-zA-Z0-9\\-]*[a-zA-Z0-9]$", subdomain); !match {
			return false
		}
	}

	return true
}
