package diet

import (
	"kitten-backend-go-app/app/core/domain/health/calories"
	"time"
)

// Diet é a entidade responsavel pela dieta
type Diet struct {
	ID             string
	CatID          string
	MaxMealsPerDay int
	MaxFoodAmount  *FoodAmount
	Calories       *calories.Calories
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
