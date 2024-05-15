package vo

import (
	"time"
)

var (
	defaultStartTime = time.Date(0, 1, 1, 8, 30, 0, 0, time.UTC)
	defaultEndTime   = time.Date(0, 1, 1, 18, 30, 0, 0, time.UTC)
)

type WorkPlan struct {
	monday    *Day
	tuesday   *Day
	wednesday *Day
	thursday  *Day
	friday    *Day
	saturday  *Day
	sunday    *Day
}

func NewWorkPlan(monday, tuesday, wednesday, thursday, friday, saturday, sunday *Day) (*WorkPlan, error) {
	return &WorkPlan{
		monday:    monday,
		tuesday:   tuesday,
		wednesday: wednesday,
		thursday:  thursday,
		friday:    friday,
		saturday:  saturday,
		sunday:    sunday,
	}, nil
}

func DefaultWorkPlan() (*WorkPlan, error) {
	day, err := NewDay(defaultStartTime, defaultEndTime, nil)
	if err != nil {
		return nil, err
	}
	return &WorkPlan{
		monday:    day,
		tuesday:   day,
		wednesday: day,
		thursday:  day,
		friday:    day,
		saturday:  day,
		sunday:    day,
	}, nil
}

func (w *WorkPlan) GetDayFromWorkPlan(t time.Time) *Day {
	weekday := t.Weekday()
	switch weekday {
	case time.Monday:
		return w.monday
	case time.Tuesday:
		return w.tuesday
	case time.Wednesday:
		return w.wednesday
	case time.Thursday:
		return w.thursday
	case time.Friday:
		return w.friday
	case time.Saturday:
		return w.saturday
	case time.Sunday:
		return w.sunday
	default:
		return nil
	}
}

func (w *WorkPlan) Monday() *Day {
	return w.monday
}

func (w *WorkPlan) Tuesday() *Day {
	return w.tuesday
}

func (w *WorkPlan) Wednesday() *Day {
	return w.wednesday
}

func (w *WorkPlan) Thursday() *Day {
	return w.thursday
}

func (w *WorkPlan) Friday() *Day {
	return w.friday
}

func (w *WorkPlan) Saturday() *Day {
	return w.saturday
}

func (w *WorkPlan) Sunday() *Day {
	return w.sunday
}
