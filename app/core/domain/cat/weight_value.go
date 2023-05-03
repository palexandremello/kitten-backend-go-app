package cats

import "errors"

type WeightValue struct {
	value float64
}

func NewWeightValue(weight float64) (*WeightValue, error) {

	if weight <= 0 {
		return nil, errors.New("weight value must be greater than 0")
	}

	return &WeightValue{value: weight}, nil
}

func (w *WeightValue) Value() float64 {
	return w.value
}
