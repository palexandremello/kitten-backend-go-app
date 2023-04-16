package cat_test

import (
	cats "kitten-backend-go-app/app/core/domain/cat"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWeight(t *testing.T) {
	date := time.Now()
	weight := float32(5.2)
	w := &cats.Weight{
		Date:   date,
		Weight: weight,
	}

	assert.Equal(t, date, w.Date)
	assert.Equal(t, weight, w.Weight)

}
