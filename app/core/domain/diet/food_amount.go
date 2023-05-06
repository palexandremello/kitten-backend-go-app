package diet

import "errors"

type FoodAmount struct {
	value float32
}

func NewFoodAmount(amount float32) (*FoodAmount, error) {

	if amount < 0 {
		return nil, errors.New("food amount cannot be negative")
	}

	return &FoodAmount{value: amount}, nil
}

func (fa *FoodAmount) Value() float32 {
	return fa.value
}
