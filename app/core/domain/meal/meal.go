package meal

import (
	"kitten-backend-go-app/app/core/domain/foodtype"
	"kitten-backend-go-app/app/core/domain/mealtype"
	"time"
)

// Meal é a entidade responsavel pelo registro de refeições para o Gato
type Meal struct {
	ID        string
	DietID    string
	MealType  mealtype.Type
	FoodType  foodtype.Type
	Amount    float32
	CreatedAt time.Time
	UpdatedAt time.Time
}
