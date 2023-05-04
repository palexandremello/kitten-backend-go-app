package symptom

import "time"

// Symptom é a entidade responsável pelo Symptom do gato
type Symptom struct {
	ID             string
	CatID          string
	SymptomType    Type
	SymptomSubtype Subtype
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
