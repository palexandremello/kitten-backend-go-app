package household_test

import (
	cats "kitten-backend-go-app/app/core/domain/cat"
	households "kitten-backend-go-app/app/core/domain/household"
	residence "kitten-backend-go-app/app/core/domain/residence"
	users "kitten-backend-go-app/app/core/domain/user"
	"testing"
	"time"
)

func TestNewHousehold(t *testing.T) {

	admin := &users.User{
		ID:       "any_id",
		Name:     "Paulo Alexandre",
		Email:    "palexandremello@gmail.com",
		Password: "any_password",
	}

	members := []*users.User{
		&users.User{
			ID:       "any_id",
			Name:     "Paulo Alexandre",
			Email:    "palexandremello@gmail.com",
			Password: "any_password",
		},
		&users.User{
			ID:       "any_id",
			Name:     "Tatsuro Yamashita",
			Email:    "yamashita@gmail.com",
			Password: "any_password",
		},
	}

	cats := []*cats.Cat{
		&cats.Cat{
			ID:         "any_id",
			Name:       "Lara",
			Birthday:   time.Date(2010, time.January, 1, 0, 0, 0, 0, time.UTC),
			Breed:      "SRD",
			Behaviors:  []*cats.Behaviors{},
			Health:     &cats.HealthStatus{},
			Diet:       &cats.Diet{},
			Activity:   &cats.Activity{},
			Sleep:      []*cats.CatSleep{},
			Residences: []*residence.CatResidence{},
		},
		&cats.Cat{
			ID:         "any_id",
			Name:       "Baco",
			Birthday:   time.Date(2010, time.January, 1, 0, 0, 0, 0, time.UTC),
			Breed:      "SRD",
			Behaviors:  []*cats.Behaviors{},
			Health:     &cats.HealthStatus{},
			Diet:       &cats.Diet{},
			Activity:   &cats.Activity{},
			Sleep:      []*cats.CatSleep{},
			Residences: []*residence.CatResidence{},
		},
	}

	latitude := float32(51.50)
	longitude := float32(-0.11)

	household := &households.Household{
		ID:        "any_id",
		Name:      "House 1",
		AdminID:   admin.ID,
		Members:   members,
		Cats:      cats,
		Latitude:  latitude,
		Longitude: longitude,
	}

	if household.ID != "any_id" {
		t.Errorf("Expected household ID to be 'any_id', but got '%s'", household.ID)
	}

	if household.Name != "House 1" {
		t.Errorf("Expected household Name to be 'House 1', but got '%s' ", household.Name)
	}
	if household.AdminID != admin.ID {
		t.Errorf("Expected household admin ID to be '%s', but got '%s'", admin.ID, household.AdminID)
	}

	if len(household.Members) != 2 {
		t.Errorf("Expected household to have 1 member, but got %d", len(household.Members))
	}

	if household.Members[0].ID != members[0].ID {
		t.Errorf("Expected household member ID to be '%s', but got '%s'", members[0].ID, household.Members[0].ID)
	}

	if len(household.Cats) != 2 {
		t.Errorf("Expected household to have 2 cat, but got %d", len(household.Cats))
	}

	if household.Cats[0].ID != cats[0].ID {
		t.Errorf("Expected household cat ID to be '%s', but got '%s'", cats[0].ID, household.Cats[0].ID)
	}

	if household.Latitude != latitude {
		t.Errorf("Expected household latitude to be '%f', but got '%f'", latitude, household.Latitude)
	}

	if household.Longitude != longitude {
		t.Errorf("Expected household longitude to be '%f', but got '%f'", longitude, household.Longitude)
	}

}
