package cats

import (
	"kitten-backend-go-app/app/core/domain/residence"
	"time"
)

type Cat struct {
	ID         string
	Name       string
	Birthday   time.Time
	Breed      string
	Behaviors  []*Behaviors
	Health     *HealthStatus
	Diet       *Diet
	Activity   *Activity
	Sleep      []*CatSleep
	Residences []*residence.CatResidence
}
