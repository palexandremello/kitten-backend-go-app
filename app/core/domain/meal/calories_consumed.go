package meal

import "errors"

type CaloriesConsumed struct {
	value float64
}

func NewCaloriesConsumed(value float64) (*CaloriesConsumed, error) {

	if value < 0 {
		return nil, errors.New("calories consumed cannot be negative")
	}
	return &CaloriesConsumed{
		value: value,
	}, nil
}

func (cc *CaloriesConsumed) Value() float64 {
	return cc.value
}
