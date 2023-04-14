package cat_test

import (
	cats "kitten-backend-go-app/app/core/domain/cat"
	residence "kitten-backend-go-app/app/core/domain/residence"
	"testing"
	"time"
)

func TestCat(t *testing.T) {
	// create a new Cat
	cat := &cats.Cat{
		ID:         "1",
		Name:       "Whiskers",
		Birthday:   time.Date(2010, time.January, 1, 0, 0, 0, 0, time.UTC),
		Breed:      "Persian",
		Behaviors:  []*cats.Behaviors{},
		Health:     &cats.HealthStatus{},
		Diet:       &cats.Diet{},
		Activity:   &cats.Activity{},
		Sleep:      []*cats.CatSleep{},
		Residences: []*residence.CatResidence{},
	}

	// test that the Cat has the correct values
	if cat.ID != "1" {
		t.Errorf("Expected cat ID to be 1, but got %s", cat.ID)
	}

	if cat.Name != "Whiskers" {
		t.Errorf("Expected cat name to be Whiskers, but got %s", cat.Name)
	}

	expectedBirthday := time.Date(2010, time.January, 1, 0, 0, 0, 0, time.UTC)
	if !cat.Birthday.Equal(expectedBirthday) {
		t.Errorf("Expected cat birthday to be %v, but got %v", expectedBirthday, cat.Birthday)
	}

	if cat.Breed != "Persian" {
		t.Errorf("Expected cat breed to be Persian, but got %s", cat.Breed)
	}

	if len(cat.Behaviors) != 0 {
		t.Errorf("Expected cat behaviors to be empty, but got %v", cat.Behaviors)
	}

	if cat.Health == nil {
		t.Error("Expected cat health to not be nil")
	}

	if cat.Diet == nil {
		t.Error("Expected cat diet to not be nil")
	}

	if cat.Activity == nil {
		t.Error("Expected cat activity to not be nil")
	}

	if len(cat.Sleep) != 0 {
		t.Errorf("Expected cat sleep to be empty, but got %v", cat.Sleep)
	}
}
