package foodtype

import (
	"kitten-backend-go-app/app/core/domain/foodtype"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoodType(t *testing.T) {

	foodType := foodtype.Type{
		ID:          "any_id",
		Name:        "Ração seca",
		Description: "Alimento em grãos",
	}

	// Testa se o ID está correto
	assert.Equal(t, "any_id", foodType.ID)

	// Testa se o nome está correto
	assert.Equal(t, "Ração seca", foodType.Name)

	// Testa se a descrição está correta
	assert.Equal(t, "Alimento em grãos", foodType.Description)

}
