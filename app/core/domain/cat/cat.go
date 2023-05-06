package cats

import (
	"kitten-backend-go-app/app/core/domain/common"
	"kitten-backend-go-app/app/core/domain/health"
	"kitten-backend-go-app/app/core/domain/meal"
	"kitten-backend-go-app/app/core/domain/residence"
)

type Cat struct {
	ID         string
	Name       *common.Name
	Birthday   *Birthday
	Breed      string
	Weight     []*Weight
	Behaviors  []*Behaviors
	Health     *health.Health
	Meal       []*meal.Meal
	Activity   *Activity
	Sleep      []*CatSleep
	Residences []*residence.CatResidence
}
