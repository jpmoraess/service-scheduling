package entity

import (
	"github.com/jpmoraess/service-scheduling/internal/domain/vo"
	"time"
)

var (
	defaultStartTime = time.Date(0, 1, 1, 8, 30, 0, 0, time.UTC)
	defaultEndTime   = time.Date(0, 1, 1, 18, 30, 0, 0, time.UTC)
)

type WorkPlan struct {
	id             string
	professionalID string
	monday         *vo.Day
	tuesday        *vo.Day
	wednesday      *vo.Day
	thursday       *vo.Day
	friday         *vo.Day
	saturday       *vo.Day
	sunday         *vo.Day
}

func NewWorkPlan(professionalID string, monday, tuesday, wednesday, thursday, friday, saturday, sunday *vo.Day) (*WorkPlan, error) {
	return &WorkPlan{
		professionalID: professionalID,
		monday:         monday,
		tuesday:        tuesday,
		wednesday:      wednesday,
		thursday:       thursday,
		friday:         friday,
		saturday:       saturday,
		sunday:         sunday,
	}, nil
}

func RestoreWorkPlan(id, professionalID string, monday, tuesday, wednesday, thursday, friday, saturday, sunday *vo.Day) (*WorkPlan, error) {
	return &WorkPlan{
		id:             id,
		professionalID: professionalID,
		monday:         monday,
		tuesday:        tuesday,
		wednesday:      wednesday,
		thursday:       thursday,
		friday:         friday,
		saturday:       saturday,
		sunday:         sunday,
	}, nil
}

func DefaultWorkPlan() (*WorkPlan, error) {
	day, err := vo.NewDay(defaultStartTime, defaultEndTime, nil)
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

func (w *WorkPlan) GetDayFromWorkPlan(t time.Time) *vo.Day {
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

func (w *WorkPlan) SetID(id string) {
	w.id = id
}

func (w *WorkPlan) ID() string {
	return w.id
}

func (w *WorkPlan) ProfessionalID() string {
	return w.professionalID
}

func (w *WorkPlan) Monday() *vo.Day {
	return w.monday
}

func (w *WorkPlan) Tuesday() *vo.Day {
	return w.tuesday
}

func (w *WorkPlan) Wednesday() *vo.Day {
	return w.wednesday
}

func (w *WorkPlan) Thursday() *vo.Day {
	return w.thursday
}

func (w *WorkPlan) Friday() *vo.Day {
	return w.friday
}

func (w *WorkPlan) Saturday() *vo.Day {
	return w.saturday
}

func (w *WorkPlan) Sunday() *vo.Day {
	return w.sunday
}
