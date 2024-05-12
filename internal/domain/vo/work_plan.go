package vo

import (
	"time"
)

type WorkPlan struct {
	Monday    *Day `bson:"monday" json:"monday"`
	Tuesday   *Day `bson:"tuesday" json:"tuesday"`
	Wednesday *Day `bson:"wednesday" json:"wednesday"`
	Thursday  *Day `bson:"thursday" json:"thursday"`
	Friday    *Day `bson:"friday" json:"friday"`
	Saturday  *Day `bson:"saturday" json:"saturday"`
	Sunday    *Day `bson:"sunday" json:"sunday"`
}

func NewWorkPlan(monday, tuesday, wednesday, thursday, friday, saturday, sunday *Day) (*WorkPlan, error) {
	return &WorkPlan{
		Monday:    monday,
		Tuesday:   tuesday,
		Wednesday: wednesday,
		Thursday:  thursday,
		Friday:    friday,
		Saturday:  saturday,
		Sunday:    sunday,
	}, nil
}

func DefaultWorkPlan() (*WorkPlan, error) {
	day, err := NewDay(time.Now(), time.Now(), nil)
	if err != nil {
		return nil, err
	}
	return &WorkPlan{
		Monday:    day,
		Tuesday:   day,
		Wednesday: day,
		Thursday:  day,
		Friday:    day,
		Saturday:  day,
	}, nil
}
