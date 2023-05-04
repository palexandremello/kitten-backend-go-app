package symptom

// Subtype é o subtipo de sintomas de um SymptomType
type Subtype struct {
	ID          string
	Name        string
	Description string
	SymptomType Type
}
