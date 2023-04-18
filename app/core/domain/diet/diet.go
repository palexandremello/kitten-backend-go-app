package diet

import "time"

// Diet é a entidade responsavel pela dieta
type Diet struct {
	ID             string
	CatID          string
	MaxMealsPerDay int
	MaxFoodAmount  float32
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
