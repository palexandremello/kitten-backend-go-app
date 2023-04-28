package users

import "errors"

var ErrInvalidWage = errors.New("wage is invalid")

type Wage struct {
	value float64
}

func NewWage(value float64) (*Wage, error) {

	if value < 0 {
		return nil, ErrInvalidWage
	}

	return &Wage{value: value}, nil
}

func (w *Wage) Value() float64 {
	return w.value
}
