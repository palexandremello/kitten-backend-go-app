package calories

import "errors"

type AgeFactor struct {
	value float64
}

func NewAgeFactor(age int) (*AgeFactor, error) {
	var factor float64

	if age < 0 {
		return nil, errors.New("age cannot be negative")
	}

	if age >= 15 {
		factor = 1.3
	} else if age >= 11 && age < 15 {
		factor = 1.1
	} else {
		factor = 1.0
	}

	return &AgeFactor{value: factor}, nil
}

func (af *AgeFactor) Value() float64 {
	return af.value
}
