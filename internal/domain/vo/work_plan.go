package vo

import (
	"time"
)

var (
	defaultStartTime = time.Date(0, 1, 1, 8, 30, 0, 0, time.UTC)
	defaultEndTime   = time.Date(0, 1, 1, 8, 30, 0, 0, time.UTC)
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
	day, err := NewDay(defaultStartTime, defaultEndTime, nil)
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

func (w *WorkPlan) GetDayFromWorkPlan(t time.Time) *Day {
	weekday := t.Weekday()
	switch weekday {
	case time.Monday:
		return w.Monday
	case time.Tuesday:
		return w.Tuesday
	case time.Wednesday:
		return w.Wednesday
	case time.Thursday:
		return w.Thursday
	case time.Friday:
		return w.Friday
	case time.Saturday:
		return w.Saturday
	case time.Sunday:
		return w.Sunday
	default:
		return nil
	}
}
