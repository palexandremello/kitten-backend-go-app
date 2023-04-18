package mealtype

import (
	"kitten-backend-go-app/app/core/domain/mealtype"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMealType(t *testing.T) {
	// Cria um novo tipo de refeição

	mealType := mealtype.Type{
		ID:          "any_id",
		Name:        "Pequeno almoço",
		Description: "Refeição matinal",
	}

	// Testa se o ID está correto
	assert.Equal(t, "any_id", mealType.ID)

	// Testa se o nome está correto
	assert.Equal(t, "Pequeno almoço", mealType.Name)

	// Testa se a descrição está correta
	assert.Equal(t, "Refeição matinal", mealType.Description)

}
