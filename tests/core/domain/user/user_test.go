package user_test

import (
	users "kitten-backend-go-app/app/core/domain/user"
	"testing"
)

func TestUser(t *testing.T) {

	user := &users.User{
		ID:       "any_id",
		Name:     "Paulo Alexandre",
		Email:    "palexandremello@gmail.com",
		Password: "any_password",
	}

	if user.ID != "any_id" {
		t.Errorf("Error no user ID")
	}

	if user.Name != "Paulo Alexandre" {
		t.Errorf("Error ao setar Name")
	}
	if user.Email != "palexandremello@gmail.com" {
		t.Errorf("Error ao setar Email")
	}
	if user.Password != "any_password" {
		t.Errorf("Error ao setar password")
	}

}
