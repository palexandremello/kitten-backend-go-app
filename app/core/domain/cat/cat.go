package cats

import (
	"kitten-backend-go-app/app/core/domain/diet"
	"kitten-backend-go-app/app/core/domain/health"
	"kitten-backend-go-app/app/core/domain/residence"
	"time"
)

type Cat struct {
	ID         string
	Name       string
	Birthday   time.Time
	Breed      string
	Weight     []*Weight
	Behaviors  []*Behaviors
	Health     *health.Health
	Diet       *diet.Diet
	Activity   *Activity
	Sleep      []*CatSleep
	Residences []*residence.CatResidence
}
