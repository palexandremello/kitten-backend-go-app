package cats

import (
	"errors"
	"time"
)

type Birthday struct {
	value time.Time
}

func NewBirthday(birthday time.Time) (*Birthday, error) {

	if birthday.IsZero() {
		return nil, errors.New("birthday date is required")
	}

	now := time.Now()
	if birthday.After(now) {
		return nil, errors.New("birthday date is invalid")
	}

	return &Birthday{value: birthday}, nil
}

func (b *Birthday) Value() time.Time {
	return b.value
}
