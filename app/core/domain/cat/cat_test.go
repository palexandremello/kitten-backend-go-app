package cats

import (
	"kitten-backend-go-app/app/core/domain/common"
	"kitten-backend-go-app/app/core/domain/health"
	"kitten-backend-go-app/app/core/domain/meal"
	"kitten-backend-go-app/app/core/domain/residence"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCatEntity(t *testing.T) {
	t.Run("should have a cat", func(t *testing.T) {
		id := "any_id"
		name, _ := common.NewName("Felix")
		birthday, _ := NewBirthday(time.Now())
		breed := "Siamese"
		weight := []*Weight{}

		behaviors := []*Behaviors{}
		health := &health.Health{}
		meal := []*meal.Meal{}
		activity := &Activity{}
		sleep := []*CatSleep{}
		residences := []*residence.CatResidence{}

		cat := &Cat{
			ID:         id,
			Name:       name,
			Birthday:   birthday,
			Breed:      breed,
			Weight:     weight,
			Behaviors:  behaviors,
			Health:     health,
			Meal:       meal,
			Activity:   activity,
			Sleep:      sleep,
			Residences: residences,
		}

		assert.Equal(t, id, cat.ID)
	})

}
