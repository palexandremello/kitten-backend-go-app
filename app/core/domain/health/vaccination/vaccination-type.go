package vaccination

type VaccinationType struct {
	Name        string
	Description string
}

type VaccinationTypes struct {
	V3 VaccinationType
	V4 VaccinationType
	V5 VaccinationType
}

// Type é uma estrutura que contém os tipos de vacina disponíveis
// e suas descrições.
var Type = VaccinationTypes{
	V3: VaccinationType{
		Name:        "V3",
		Description: "Protege contra panleucopenia, rinotraqueite e calicivirose",
	},
	V4: VaccinationType{
		Name:        "V4",
		Description: "Protege contra todas as vacinas da V3 e inclui proteção contra clamidiose",
	},
	V5: VaccinationType{
		Name:        "V5",
		Description: "Protege contra todas as vacinas da V3 e V4 e inclui proteção contra leucemia felina",
	},
}
