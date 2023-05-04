package vaccination

import "time"

// Vaccine entity
type Vaccine struct {
	Type        VaccinationType
	Date        time.Time
	Observation string
}
