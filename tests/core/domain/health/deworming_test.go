package health_test

import (
	"kitten-backend-go-app/app/core/domain/health"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDeworming(t *testing.T) {

	now := time.Now()

	deworming := health.Deworming{
		ID:          "any_id",
		CatID:       "any_id",
		DoseType:    health.FirstDose,
		WormingType: health.FleaControl,
		Date:        now,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	assert.Equal(t, "any_id", deworming.ID)
	assert.Equal(t, "any_id", deworming.CatID)
	assert.Equal(t, health.FirstDose, deworming.DoseType)
	assert.Equal(t, health.FleaControl, deworming.WormingType)
	assert.Equal(t, now, deworming.Date)
	assert.Equal(t, now, deworming.CreatedAt)
	assert.Equal(t, now, deworming.UpdatedAt)
}
