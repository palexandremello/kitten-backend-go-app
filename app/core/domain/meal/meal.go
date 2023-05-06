package meal

import (
	"kitten-backend-go-app/app/core/domain/foodtype"
	"kitten-backend-go-app/app/core/domain/mealtype"
	"time"
)

// Meal é a entidade responsável pelo registro de refeições para o gato
type Meal struct {
	ID               string
	DietID           string
	MealType         mealtype.Type
	FoodType         foodtype.Type
	Amount           float32
	CaloriesConsumed *CaloriesConsumed
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
