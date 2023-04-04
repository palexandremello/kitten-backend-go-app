package core

import (
	"testing"

	"kitten-backend-go-app/app/core"
)

func TestUserEntity(t *testing.T) {
	user := &core.User{
		Id:        "uuid",
		Name:      "Paulo",
		Email:     "palexandremello@gmail.com",
		Password:  "any_password",
		AddressId: "any_id",
	}
	if user.Id != "uuid" {
		t.Errorf("Expected ID to be '1', but got '%s'", user.Id)
	}

	if user.Name != "Paulo" {
		t.Errorf("Expected Name to be 'John Doe', but got '%s'", user.Name)
	}

	if user.Email != "palexandremello@gmail.com" {
		t.Errorf("Expected Email to be 'john.doe@example.com', but got '%s'", user.Email)
	}

	if user.Password != "any_password" {
		t.Errorf("Expected Password to be 'secret', but got '%s'", user.Password)
	}

	if user.AddressId != "any_id" {
		t.Errorf("Expected AddressID to be '123', but got '%s'", user.AddressId)
	}
}
