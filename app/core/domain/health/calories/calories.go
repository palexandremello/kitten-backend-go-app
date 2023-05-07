package calories

// Calories entity from a Cat with Metabolic Basal Rate
type Calories struct {
	ID               string
	Weight           float32
	IsSpyOrNeutral   bool
	BMR              float32
	ActivityLevel    float32
	AgeFactor        *AgeFactor
	CastrationFactor float32
}
