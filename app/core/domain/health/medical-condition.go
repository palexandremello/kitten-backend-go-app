package health

import "time"

type MedicalCondition struct {
	ID            string
	CatID         string
	Condition     ConditionType
	DaignosisDate time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
