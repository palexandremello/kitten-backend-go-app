package health

import "time"

type WormingType int

const (
	Worming WormingType = iota + 1
	FleaControl
)

type DoseType int

const (
	FirstDose DoseType = iota + 1
	SecondDose
	ThirdDose
	FourthDose
	FifthDose
	Reinforcement
)

type Deworming struct {
	ID          string
	CatID       string
	DoseType    DoseType
	WormingType WormingType
	Date        time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
