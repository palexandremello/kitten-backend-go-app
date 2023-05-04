package vaccination

import "time"

// Card entity about Vaccination Pet
type Card struct {
	ID        string
	CatID     string
	Vaccines  []*Vaccine
	CreatedAt time.Time
	UpdatedAt time.Time
}
