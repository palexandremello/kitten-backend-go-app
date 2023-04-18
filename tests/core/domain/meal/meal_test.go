package meal

import (
	"kitten-backend-go-app/app/core/domain/foodtype"
	"kitten-backend-go-app/app/core/domain/meal"
	"kitten-backend-go-app/app/core/domain/mealtype"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMeal(t *testing.T) {
	mealType := mealtype.Type{
		ID:          "any_id",
		Name:        "Breakfast",
		Description: "The first meal of the day",
	}
	foodType := foodtype.Type{
		ID:          "any_id",
		Name:        "Ração Premier Pet Gatos Ambientes Internos Adultos",
		Description: "Sabor Frango",
	}
	now := time.Now()
	meal := meal.Meal{
		ID:        "any_id",
		MealType:  mealType,
		FoodType:  foodType,
		Amount:    100.0,
		CreatedAt: now,
		UpdatedAt: now,
	}

	assert.Equal(t, "any_id", meal.ID)
	assert.Equal(t, "", meal.DietID)
	assert.Equal(t, mealType, meal.MealType)
	assert.Equal(t, foodType, meal.FoodType)
	assert.Equal(t, float32(100.0), meal.Amount)
	assert.Equal(t, now, meal.CreatedAt)
	assert.Equal(t, now, meal.UpdatedAt)
}