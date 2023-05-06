package health

import (
	"kitten-backend-go-app/app/core/domain/health/diet"
	s "kitten-backend-go-app/app/core/domain/health/symptom"
	"kitten-backend-go-app/app/core/domain/health/vaccination"
	"time"
)

// Health é a entidade responsável pela saúde do gato
type Health struct {
	ID               string
	CatID            string
	Conditions       []*MedicalCondition
	DewormingSchemes []*Deworming
	VaccinationCard  []*vaccination.Card
	CommonSymptoms   []*s.Symptom
	Diet             *diet.Diet
	IsSpyOrNeutral   bool
	SpyOrNeutralDate *time.Time
	LastCheckup      *time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
